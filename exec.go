package dataset_sql

import (
	"bytes"
	"encoding/csv"
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/namespace"
)

func execSelect(stmt *Select, ns namespace.StorableNamespace) (result *dataset.Dataset, err error) {
	result, err = buildResultDataset(stmt, ns)
	if err != nil {
		return
	}

	// TODO... Sort each table by select sort criteria here?
	// TODO - column ambiguity check

	proj, err := buildProjection(result, stmt.SelectExprs)
	if err != nil {
		return result, err
	}

	// Populate any ColName nodes with their type information
	if err := populateColNames(stmt, result); err != nil {
		return result, err
	}

	results := bytes.NewBuffer(nil)
	w := csv.NewWriter(results)
	result.Format = dataset.CsvDataFormat

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
		return result, err
	}

	indicies := make([]int, len(result.Datasets))
	indicies[len(indicies)-1] = -1
	rowLen := masterRowLength(result)

	for {
		if limit > 0 && added == limit {
			break
		}

		row := nextRow(result, indicies, lengths, rowLen, data)
		if row == nil {
			break
		}

		// check dst against criteria, only continue if it passes
		// TODO - confirm that the result dataset is the proper one to be passing in here?
		// see if we can't remove dataset altogether by embedding all info in the ast?
		if pass, err := stmt.Where.EvalBool(result, row); err != nil {
			return result, err
		} else if !pass {
			continue
		}

		// check offset
		if offset > 0 && skipped < offset {
			skipped++
			continue
		}

		// project result row
		row = projectRow(proj, row)

		// ewwwwwwwwwww
		strRow := make([]string, len(row))
		for i, col := range row {
			strRow[i] = string(col)
		}
		w.Write(strRow)
		added++
	}

	w.Flush()
	result.Data = results.Bytes()

	return
}

func execUnion(node *Union, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("union statements are not yet supported")
}

func execInsert(node *Insert, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("insert statements are not yet supported")
}

func execUpdate(node *Update, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("update statements are not yet supported")
}

func execDelete(node *Delete, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("delete statements are not yet supported")
}

func execSet(node *Set, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("set statements are not yet supported")
}

func execDDL(node *DDL, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	return nil, fmt.Errorf("ddl statements are not yet supported")
}

func execOther(node *Other, ns namespace.StorableNamespace) (*dataset.Dataset, error) {
	// TODO - lolololol
	return nil, fmt.Errorf("other statements are not yet supported")
}

// populateColNames adds type information to ColName nodes in the ast
func populateColNames(stmt *Select, ds *dataset.Dataset) error {
	return stmt.Where.WalkSubtree(func(node SQLNode) (bool, error) {
		if colName, ok := node.(*ColName); ok && node != nil {
			if colName.Qualifier != nil {
				if d, err := ds.DatasetForAddress(colName.Qualifier.TableAddress()); err == nil {
					if field := d.FieldForName(colName.Name.String()); field != nil {
						colName.Type = field.Type.String()
						return true, nil
					}
					return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", colName.Name.String(), colName.Qualifier.TableAddress().String())
				} else {
					return false, err
				}
			} else {
				for _, d := range ds.Datasets {
					if field := d.FieldForName(colName.Name.String()); field != nil {
						colName.Type = field.Type.String()
						return true, nil
					} else {
						return false, nil
					}
				}
				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", colName.Name.String())
			}
		}

		return true, nil
	})
}

func projectRow(projection []int, source [][]byte) (row [][]byte) {
	row = make([][]byte, len(projection))
	for i, j := range projection {
		row[i] = source[j]
	}
	return
}

// Gather all mentioned tables, attaching them to a *dataset.Dataset
func buildResultDataset(stmt *Select, ns namespace.StorableNamespace) (result *dataset.Dataset, err error) {
	result = &dataset.Dataset{}
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

func nodeColIndex(result *dataset.Dataset, node SelectExpr) (idx int, err error) {
	if nse, ok := node.(*NonStarExpr); ok && node != nil {
		if colName, ok := nse.Expr.(*ColName); ok && node != nil {
			for _, ds := range result.Datasets {
				for _, f := range ds.Fields {
					idx++
					if f.Name == colName.Name.String() {
						return
					}
				}
			}
		}
		return -1, nil
	}

	return 0, fmt.Errorf("node is not a non-star select expression")
}

func intSeries(start, length int) (series []int) {
	series = make([]int, length)
	for i := 0; i < length; i++ {
		series[i] = i + start
	}
	return
}

func masterRowLength(ds *dataset.Dataset) (l int) {
	for _, d := range ds.Datasets {
		l += len(d.Fields)
	}
	return
}

// nextRow generates the next master row for a dataset
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
