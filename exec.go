package dataset_sql

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/namespace"
)

func execSelect(stmt *Select, ns namespace.StorableNamespace) (result *dataset.Dataset, err error) {
	result = &dataset.Dataset{}
	// 1. Gather all mentioned tables, attaching them to dataset.Dataset
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

	// TODO... Sort each table by select sort criteria here?

	// Build dataset destination fields
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

	// TODO - column ambiguity check

	// Populate any ColName nodes with their type information
	stmt.Where.WalkSubtree(func(node SQLNode) (bool, error) {
		if colName, ok := node.(*ColName); ok && node != nil {
			if colName.Qualifier != nil {
				if ds, err := result.DatasetForAddress(colName.Qualifier.TableAddress()); err == nil {
					if field := ds.FieldForName(colName.Name.String()); field != nil {
						colName.Type = field.Type.String()
						return true, nil
					}
					return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", colName.Name.String(), colName.Qualifier.TableAddress().String())
				} else {
					return false, err
				}
			} else {
				for _, ds := range result.Datasets {
					if field := ds.FieldForName(colName.Name.String()); field != nil {
						colName.Type = field.Type.String()
						return true, nil
					} else {
						return false, err
					}
				}
				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", colName.Name.String())
			}
		}

		return true, nil
	})

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

	// 3. Populate dataset data by iterating through each dataset.dataset, projecting the source dataset onto the result dataset.
	// 		Then evaluate if the projected row passes any where clauses
	for _, ds := range result.Datasets {
		err = ds.EachRow(func(rowNum int, src [][]byte, e error) (err error) {
			if e != nil {
				return e
			}

			if limit > 0 && added == limit {
				return errors.New("EOF")
			}

			// check dst against criteria, only continue if it passes
			if pass, err := stmt.Where.EvalBool(ds, src); err != nil {
				return err
			} else if !pass {
				return nil
			}

			// check offset
			if offset > 0 && skipped < offset {
				skipped++
				return nil
			}

			dst := make([][]byte, len(result.Fields))
			col := 0
			for _, sExp := range stmt.SelectExprs {
				if colsWritten, err := sExp.Map(col, ds, result, src, dst); err != nil {
					return err
				} else {
					col += colsWritten
				}
			}

			// ewwwwwwwwwww
			row := make([]string, len(dst))
			for i, col := range dst {
				row[i] = string(col)
			}

			w.Write(row)
			added++

			return nil
		})
		if err != nil {
			return
		}
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
