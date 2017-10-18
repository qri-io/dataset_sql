package dataset_sql

import (
	"bytes"
	"fmt"
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dsio"
	"io/ioutil"
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

func Exec(store cafs.Filestore, ds *dataset.Dataset, options ...func(o *ExecOpt)) (result *dataset.Structure, resultBytes []byte, err error) {
	opts := &ExecOpt{
		Format: dataset.CsvDataFormat,
	}
	for _, option := range options {
		option(opts)
	}

	if ds.QuerySyntax != "sql" {
		return nil, nil, fmt.Errorf("Invalid syntax: '%s' sql_dataset only supports sql syntax. ", ds.QuerySyntax)
	}

	concreteStmt, err := Parse(ds.QueryString)
	if err != nil {
		return nil, nil, err
	}
	err = RemoveUnusedReferences(concreteStmt, ds)
	if err != nil {
		return nil, nil, err
	}

	strs := map[string]*dataset.Structure{}
	for name, ds := range ds.Resources {
		strs[name] = ds.Structure
	}

	ds.Structure, err = ResultStructure(concreteStmt, strs, opts)
	if err != nil {
		return nil, nil, err
	}

	_, stmt, remap, err := Format(ds.QueryString)
	if err != nil {
		return nil, nil, err
	}

	ds.Query = &dataset.Query{
		Structures: map[string]*dataset.Structure{},
	}

	// collect table references
	for mapped, ref := range remap {
		// for i, adr := range stmt.References() {
		if ds.Resources[ref] == nil {
			return nil, nil, fmt.Errorf("couldn't find resource for table name: %s", ref)
		}

		ds.Query.Structures[mapped] = ds.Resources[ref].Structure.Abstract()
	}

	// This is a basic-column name rewriter from concrete to abstract
	stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		// if ae, ok := node.(*AliasedExpr); ok && ae != nil {
		if cn, ok := node.(*ColName); ok && cn != nil {
			// TODO - check qualifier to avoid extra loopage
			// if cn.Qualifier.String() != "" {
			// 	for _, f := range ds.Query.Structures[cn.Qualifier.String()].Schema.Fields {
			// 		if cn.Name.String() ==
			// 	}
			// }
			for con, r := range ds.Resources {
				for i, f := range r.Structure.Schema.Fields {
					if f.Name == cn.Name.String() {
						for mapped, ref := range remap {
							if ref == con {
								// fmt.Println(ref, con, mapped)
								// fmt.Println("MATCH", ds.Query.Structures[mapped].Schema.Fields[i].Name)
								// fmt.Println(String(cn))
								// fmt.Println(String(&ColName{
								// 	Name:      NewColIdent(ds.Query.Structures[mapped].Schema.Fields[i].Name),
								// 	Qualifier: TableName{Name: NewTableIdent(mapped)},
								// }))

								*cn = ColName{
									Name:      NewColIdent(ds.Query.Structures[mapped].Schema.Fields[i].Name),
									Qualifier: TableName{Name: NewTableIdent(mapped)},
								}
							}
						}
						return false, nil
					}
				}
				// }

			}
		}
		return true, nil
	})

	ds.Query.Statement = String(stmt)

	ds.Query.Structure, err = ResultStructure(stmt, ds.Query.Structures, opts)
	if err != nil {
		return nil, nil, err
	}

	return stmt.exec(store, ds, remap, opts)
}

func (stmt *Select) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (result *dataset.Structure, resultBytes []byte, err error) {
	if stmt.OrderBy != nil {
		return nil, nil, NotYetImplemented("ORDER BY statements")
	}

	// TODO - This is a total hack to support DISTINCT statements for now
	// in the future this needs to be rolled in as a "hasRow" method
	// on the resultWriter interface
	// left like this it'll chew up memory
	var writtenRows [][][]byte

	// from, result, err := buildResultStructure(stmt, store, ds.Resources, opts)
	// if err != nil {
	// 	return
	// }

	result = ds.Query.Structure
	from := map[string]*StructureData{}
	for abst, con := range remap {
		file, e := store.Get(ds.Resources[con].Data)
		if e != nil {
			err = fmt.Errorf("error getting dataset file: %s: %s", ds.Data, e.Error())
			return
		}

		// TODO - this is a shim for now and should be removed asap
		data, e := ioutil.ReadAll(file)
		if e != nil {
			err = fmt.Errorf("error loading dataset data: %s: %s", ds.Data, e.Error())
			return
		}

		from[abst] = &StructureData{
			Structure: ds.Query.Structures[abst],
			Data:      data,
		}
	}

	// TODO... Sort each table by select sort criteria here?
	// TODO - column ambiguity check

	proj, err := buildProjection(stmt.SelectExprs, from)
	if err != nil {
		return result, nil, err
	}

	// Populate any ColName nodes with their type information
	// if err := populateColNames(stmt, from); err != nil {
	// 	return result, nil, err
	// }

	funcs, err := AggregateFuncs(stmt.SelectExprs, from)
	if err != nil {
		return result, nil, err
	}

	agg := len(funcs) > 0

	buf := dsio.NewBuffer(ds.Structure)

	limit, offset, err := stmt.Limit.Counts()
	if err != nil {
		return result, nil, err
	}

	added := int64(0)
	skipped := int64(0)

	data, lengths, err := buildDatabase(from, result)
	if err != nil {
		return result, nil, err
	}

	indicies := make([]int, len(from))
	if len(indicies) > 0 {
		indicies[len(indicies)-1] = -1
	}
	rowLen := masterRowLength(from)

	for {
		if limit > 0 && added == limit && stmt.OrderBy == nil {
			break
		}

		// generate the next master row from source datasests, bailing if we have nothing left to examine
		// statements that don't reference any datasets need a chance to return their results
		// so we bail if added is above zero and the slice is empty
		row := nextRow(len(from), indicies, lengths, rowLen, data)
		if row == nil || (len(row) == 0 && added > 0) {
			break
		}

		// check dst against criteria, only continue if it passes
		// TODO - confirm that the result dataset is the proper one to be passing in here?
		// see if we can't remove dataset altogether by embedding all info in the ast?
		if _, pass, err := stmt.Where.Eval(row); err != nil {
			return result, nil, err
		} else if bytes.Equal(pass, falseB) {
			continue
		}

		// check offset
		if offset > 0 && skipped < offset {
			skipped++
			continue
		}

		// project result row
		row, err = projectRow(stmt.SelectExprs, proj, row)
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

		if !agg {
			buf.WriteRow(row)
		}

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

	if agg {
		row, err := aggFuncResults(funcs, proj)
		if err != nil {
			return result, nil, err
		}
		// fmt.Println(row)
		for _, r := range row {
			fmt.Printf(string(r))
		}
		buf.WriteRow(row)
	}

	if err := buf.Close(); err != nil {
		return result, nil, err
	}

	if stmt.OrderBy != nil {

	}

	// TODO - rename / deref result var
	result = ds.Structure

	resultBytes = buf.Bytes()
	return
}

func (node *Union) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("union statements")
}
func (node *Insert) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("insert statements")
}
func (node *Update) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("update statements")
}
func (node *Delete) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("delete statements")
}
func (node *Set) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("set statements")
}
func (node *DDL) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("ddl statements")
}
func (node *ParenSelect) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("ParenSelect statements")
}
func (node *Show) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("Show statements")
}
func (node *Use) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("Use statements")
}
func (node *OtherRead) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("OtherRead statements")
}
func (node *OtherAdmin) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("OtherAdmin statements")
}

// populateColNames adds type information to ColName nodes in the ast
// func populateColNames(stmt *Select, from map[string]*StructureData) error {
// 	return stmt.WalkSubtree(func(sqlNode SQLNode) (bool, error) {
// 		switch node := sqlNode.(type) {
// 		case *ColName:
// 			if node.Qualifier.String() != "" {
// 				idx := 0
// 				for tableName, resourceData := range from {
// 					if node.Qualifier.String() == tableName {
// 						for i, f := range resourceData.Structure.Schema.Fields {
// 							if node.Name.String() == f.Name {
// 								node.Field = f
// 								node.RowIndex = idx + i
// 								return true, nil
// 							}
// 						}
// 					}
// 					idx += len(resourceData.Structure.Schema.Fields)
// 				}
// 				return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", node.Name.String(), node.Qualifier.TableName())
// 			} else {
// 				idx := 0
// 				for _, resourceData := range from {
// 					for i, f := range resourceData.Structure.Schema.Fields {
// 						if node.Name.String() == f.Name {
// 							node.Field = f
// 							node.RowIndex = idx + i
// 							return true, nil
// 						}
// 					}
// 					idx += len(resourceData.Structure.Schema.Fields)
// 				}
// 				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", node.Name.String())
// 			}
// 		}

// 		return true, nil
// 	})
// }

func aggFuncResults(funcs []AggFunc, projection []int) (row [][]byte, err error) {
	row = make([][]byte, len(funcs))
	for i, fn := range funcs {
		row[i] = fn.Value()
	}
	return row, nil
}

// TODO - refactor StructureData to take a io.Reader instead of []byte
type StructureData struct {
	Structure *dataset.Structure
	Data      []byte
}

// Gather all mentioned tables, attaching them to a *dataset.Structure
// TODO - refactor this out
func buildResultStructure(stmt *Select, store cafs.Filestore, resources map[string]*dataset.Dataset, opts *ExecOpt) (from map[string]*StructureData, result *dataset.Structure, err error) {
	from = map[string]*StructureData{}
	structures := map[string]*dataset.Structure{}
	for name, ds := range resources {
		st := ds.Structure

		file, e := store.Get(ds.Data)
		if e != nil {
			err = fmt.Errorf("error getting dataset file: %s: %s", ds.Data, e.Error())
			return
		}

		// TODO - shim until structured data refactor
		data, e := ioutil.ReadAll(file)
		if e != nil {
			err = fmt.Errorf("error loading dataset data: %s: %s", ds.Data, e.Error())
			return
		}

		from[name] = &StructureData{
			Structure: st,
			Data:      data,
		}

		structures[name] = st
	}

	result, err = ResultStructure(stmt, structures, opts)
	if err != nil {
		return
	}

	return
}

// fromFieldCount totals all fields in
func fromFieldCount(from map[string]*StructureData) (count int) {
	for _, resourceData := range from {
		count += len(resourceData.Structure.Schema.Fields)
	}
	return
}

// nodeColIndex finds the column index for a given node
func nodeColIndex(node SelectExpr, from map[string]*StructureData) (idx int, err error) {
	if nse, ok := node.(*AliasedExpr); ok && node != nil {
		if colName, ok := nse.Expr.(*ColName); ok && node != nil {
			for _, resourceData := range from {
				for _, f := range resourceData.Structure.Schema.Fields {
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
func masterRowLength(from map[string]*StructureData) (l int) {
	for _, resourceData := range from {
		l += len(resourceData.Structure.Schema.Fields)
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
func nextRow(numStructures int, indicies, lengths []int, rowLen int, data [][][][]byte) (row [][]byte) {
	if incrIndicies(indicies, lengths) == nil {
		return nil
	} else {
		row = make([][]byte, rowLen)
		k := 0
		for i := 0; i < numStructures; i++ {
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
