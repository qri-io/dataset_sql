package dataset_sql

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/namespace"
)

type ExecOpt struct {
	Format dataset.DataFormat
}

func opts(options ...func(*ExecOpt)) *ExecOpt {
	o := &ExecOpt{
		Format: dataset.CsvDataFormat,
	}
	for _, option := range options {
		option(o)
	}
	return o
}

func execSelect(stmt *Select, ns namespace.StorableNamespace, opt *ExecOpt) (result *dataset.Dataset, resultBytes []byte, err error) {
	if stmt.OrderBy != nil {
		return nil, nil, NotYetImplemented("ORDER BY statements")
	}

	// TODO - This is a total hack to support DISTINCT statements for now
	// in the future this needs to be rolled in as a "hasRow" method
	// on the resultWriter interface
	var writtenRows [][][]byte

	result, err = buildResultDataset(stmt, ns, opt)
	if err != nil {
		return
	}

	// TODO... Sort each table by select sort criteria here?
	// TODO - column ambiguity check

	proj, err := buildProjection(result, stmt.SelectExprs)
	if err != nil {
		return result, nil, err
	}

	// Populate any ColName nodes with their type information
	if err := populateColNames(stmt, result); err != nil {
		return result, nil, err
	}

	w := newResultWriter(result, opt)

	limit := int64(0)
	offset := int64(0)
	added := int64(0)
	skipped := int64(0)
	if stmt.LimitOffset != nil {
		limit = stmt.LimitOffset.GetRowCount()
		offset = stmt.LimitOffset.GetOffset()
	}

	data, lengths, err := buildDatabase(result)
	if err != nil {
		return result, nil, err
	}

	indicies := make([]int, len(result.Datasets))
	if len(indicies) > 0 {
		indicies[len(indicies)-1] = -1
	}
	rowLen := masterRowLength(result)

	for {
		if limit > 0 && added == limit && stmt.OrderBy == nil {
			break
		}

		// generate the next master row from source datasests, bailing if we have nothing left to examine
		// statements that don't reference any datasets need a chance to return their results
		// so we bail if added is above zero and the slice is empty
		row := nextRow(result, indicies, lengths, rowLen, data)
		if row == nil || (len(row) == 0 && added > 0) {
			break
		}

		// check dst against criteria, only continue if it passes
		// TODO - confirm that the result dataset is the proper one to be passing in here?
		// see if we can't remove dataset altogether by embedding all info in the ast?
		if pass, err := stmt.Where.EvalBool(result, row); err != nil {
			return result, nil, err
		} else if !pass {
			continue
		}

		// check offset
		if offset > 0 && skipped < offset {
			skipped++
			continue
		}

		// project result row
		row, err = projectRow(result, stmt.SelectExprs, proj, row)
		if err != nil {
			return
		}

		// check distinct
		if stmt.Distinct != "" {
			unique := true
			for _, r := range writtenRows {
				if rowsEqual(row, r) {
					unique = false
					break
				}
			}
			if unique {
				writtenRows = append(writtenRows, row)
			} else {
				continue
			}
		}

		w.WriteRow(row)
		added++

		// we can advance the leftmost row if we make it here and there's a filtering clause.
		// b/c at this point we have a match for the leftmost combination.
		// TODO - I'm nervous of this because I haven't thought through multiple matches.
		// 				so for the moment we're skipping it.
		// if stmt.Where != nil {
		// 	if done := jumpRow(indicies, lengths); done {
		// 		break
		// 	}
		// }
	}

	if err := w.Close(); err != nil {
		return result, nil, err
	}

	if stmt.OrderBy != nil {

	}

	resultBytes = w.Bytes()
	return
}

func execUnion(node *Union, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("union statements")
}

func execInsert(node *Insert, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("insert statements")
}

func execUpdate(node *Update, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("update statements")
}

func execDelete(node *Delete, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("delete statements")
}

func execSet(node *Set, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("set statements")
}

func execDDL(node *DDL, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	return nil, nil, NotYetImplemented("ddl statements")
}

func execOther(node *Other, ns namespace.StorableNamespace, opt *ExecOpt) (*dataset.Dataset, []byte, error) {
	// TODO - lolololol
	return nil, nil, NotYetImplemented("other statements")
}

// populateColNames adds type information to ColName nodes in the ast
func populateColNames(stmt *Select, ds *dataset.Dataset) error {
	return stmt.Where.WalkSubtree(func(node SQLNode) (bool, error) {
		if colName, ok := node.(*ColName); ok && node != nil {
			if colName.Qualifier != nil {
				idx := 0
				for _, d := range ds.Datasets {
					if d.Address.Equal(colName.Qualifier.TableAddress()) {
						for i, f := range d.Fields {
							if colName.Name.String() == f.Name {
								colName.Field = f
								colName.RowIndex = idx + i
								return true, nil
							}
						}
					}
					idx += len(d.Fields)
				}
				return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", colName.Name.String(), colName.Qualifier.TableAddress().String())
			} else {
				idx := 0
				for _, d := range ds.Datasets {
					for i, f := range d.Fields {
						if colName.Name.String() == f.Name {
							colName.Field = f
							colName.RowIndex = idx + i
							return true, nil
						}
					}
					idx += len(d.Fields)
				}
				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", colName.Name.String())
			}
		}

		return true, nil
	})
}

// projectRow takes a master row & fits it to the desired result, evaluating any expressions along the way.
func projectRow(ds *dataset.Dataset, stmt SelectExprs, projection []int, source [][]byte) (row [][]byte, err error) {
	row = make([][]byte, len(projection))
	for i, j := range projection {
		if j == -1 {
			if nsr, ok := stmt[i].(*NonStarExpr); ok {
				val, e := nsr.Expr.Eval(ds, row)
				if e != nil {
					return row, e
				}
				row[i] = val.Bytes()
			} else {
				return row, fmt.Errorf("select expression %d is invalid", i+1)
			}
		} else {
			row[i] = source[j]
		}
	}
	return
}

// Gather all mentioned tables, attaching them to a *dataset.Dataset
func buildResultDataset(stmt *Select, ns namespace.StorableNamespace, opt *ExecOpt) (result *dataset.Dataset, err error) {

	buf := NewTrackedBuffer(nil)
	stmt.Format(buf)

	result = &dataset.Dataset{
		Format: opt.Format,
		Query: &dataset.Query{
			Statement: buf.String(),
		},
	}

	for _, adr := range stmt.From.TableAddresses() {
		if ds, e := ns.Dataset(adr); e != nil {
			err = e
			return
		} else {
			store, err := ns.Store(ds.Address)
			if err != nil {
				return result, err
			}
			ds.Data, err = ds.FetchBytes(store)
			if err != nil {
				return result, err
			}
			result.Datasets = append(result.Datasets, ds)
		}
	}
	populateResultFields(stmt, result)
	return
}

func populateResultFields(stmt *Select, result *dataset.Dataset) {
	for _, node := range stmt.SelectExprs {
		if star, ok := node.(*StarExpr); ok && node != nil {
			name := string(star.TableName)
			for _, ds := range result.Datasets {
				// we add fields if the names match, or if no name is specified
				if ds.Name == name || name == "" {
					result.Fields = append(result.Fields, ds.Fields...)
				}
			}
		} else if expr, ok := node.(*NonStarExpr); ok && node != nil {
			result.Fields = append(result.Fields, &dataset.Field{
				Name: expr.ResultName(),
				Type: expr.FieldType(result),
			})
		}
	}
}

func buildProjection(ds *dataset.Dataset, selectors SelectExprs) (proj []int, err error) {
	for _, node := range selectors {
		if isUnqualifiedStarExpr(node) {
			return intSeries(0, subsetFieldCount(ds)), nil
		} else if isQualifiedStarExpr(node) {
			ds, e := findStarExprSubset(ds, node)
			if e != nil {
				return proj, e
			}
			proj = append(proj, intSeries(len(proj), len(ds.Fields))...)
		} else {
			i, e := nodeColIndex(ds, node)
			if e != nil {
				return proj, e
			}
			proj = append(proj, i)
		}

	}
	return
}

func buildDatabase(ds *dataset.Dataset) (data [][][][]byte, lengths []int, err error) {
	// data = make([][][][]byte, len(ds.Datasets))
	for _, d := range ds.Datasets {
		dsData, err := d.AllRows()
		if err != nil {
			return nil, nil, err
		}

		lengths = append(lengths, len(dsData))
		data = append(data, dsData)
	}
	return
}

func findStarExprSubset(result *dataset.Dataset, node SelectExpr) (ds *dataset.Dataset, err error) {
	if star, ok := node.(*StarExpr); ok && node != nil {
		for _, d := range result.Datasets {
			if star.TableName.String() == d.Address.String() {
				return d, nil
			}
		}
	}
	return
}

func isUnqualifiedStarExpr(node SelectExpr) bool {
	if star, ok := node.(*StarExpr); ok && node != nil {
		if star.TableName.String() == "" {
			return true
		}
	}
	return false
}

func isQualifiedStarExpr(node SelectExpr) bool {
	if star, ok := node.(*StarExpr); ok && node != nil {
		if star.TableName.String() != "" {
			return true
		}
	}
	return false
}

func subsetFieldCount(result *dataset.Dataset) (count int) {
	for _, ds := range result.Datasets {
		count += len(ds.Fields)
	}
	return
}

// nodeColIndex
func nodeColIndex(result *dataset.Dataset, node SelectExpr) (idx int, err error) {
	if nse, ok := node.(*NonStarExpr); ok && node != nil {
		if colName, ok := nse.Expr.(*ColName); ok && node != nil {
			for _, ds := range result.Datasets {
				for _, f := range ds.Fields {
					if f.Name == colName.Name.String() {
						return
					}
					idx++
				}
			}
		}
		return -1, nil
	}

	return 0, fmt.Errorf("node is not a non-star select expression")
}

// intSeries returns a slice sized by length that counts from start upward
func intSeries(start, length int) (series []int) {
	series = make([]int, length)
	for i := 0; i < length; i++ {
		series[i] = i + start
	}
	return
}

// masterRowLength sums all fields of a dataset's children
func masterRowLength(ds *dataset.Dataset) (l int) {
	for _, d := range ds.Datasets {
		l += len(d.Fields)
	}
	return
}

// rowsEqual checks to see if two rows are identitical
func rowsEqual(a, b [][]byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, ai := range a {
		if !bytes.Equal(ai, b[i]) {
			return false
		}
	}
	return true
}

// nextRow generates the next master row for a dataset from the source datasets
func nextRow(ds *dataset.Dataset, indicies, lengths []int, rowLen int, data [][][][]byte) (row [][]byte) {
	if incrIndicies(indicies, lengths) == nil {
		return nil
	} else {
		row = make([][]byte, rowLen)
		k := 0
		for i := 0; i < len(ds.Datasets); i++ {
			// fmt.Println(i, indicies[i])
			for _, cell := range data[i][indicies[i]] {
				row[k] = cell
				k++
			}
		}
	}
	return
}

// incrIndicies increments the index-counter, returning nil when
// counting is complete
func incrIndicies(indicies, lengths []int) []int {
	for i := len(indicies) - 1; i >= 0; i-- {
		if indicies[i] < lengths[i]-1 {
			indicies[i]++
			break
		} else {
			if i-1 <= 0 && indicies[0] == lengths[0]-1 {
				return nil
			}
			indicies[i] = 0
			indicies[i-1]++
			break
		}
	}

	return indicies
}

// jumpRow advances
func jumpRow(indicies, lengths []int) bool {
	for i, idx := range indicies {
		if i == 0 {
			idx++
			if idx == lengths[i] {
				return true
			}
		} else if i == len(indicies)-1 {
			idx = -1
		} else {
			idx = 0
		}
	}
	return false
}

type resultWriter interface {
	WriteRow([][]byte) error
	Close() error
	Bytes() []byte
}

func newResultWriter(result *dataset.Dataset, o *ExecOpt) resultWriter {
	switch result.Format {
	case dataset.CsvDataFormat:
		buf := &bytes.Buffer{}
		return &csvResultWriter{
			buf:    buf,
			Writer: csv.NewWriter(buf),
		}
	case dataset.JsonDataFormat:
		return NewJsonWriter(result, true)
	case dataset.JsonArrayDataFormat:
		return NewJsonWriter(result, false)
	}
	return nil
}

type csvResultWriter struct {
	buf *bytes.Buffer
	*csv.Writer
}

func (cw *csvResultWriter) WriteRow(row [][]byte) error {
	strRow := make([]string, len(row))
	for i, col := range row {
		strRow[i] = string(col)
	}
	return cw.Write(strRow)
}

func (cw *csvResultWriter) Close() error {
	cw.Flush()
	return nil
}

func (cw *csvResultWriter) Bytes() []byte {
	return cw.buf.Bytes()
}
