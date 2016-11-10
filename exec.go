package sqlparser

import (
	"bytes"
	"encoding/csv"

	"github.com/qri-io/dataset"
	"github.com/qri-io/datatype"
)

type Domain interface {
	DatsetForAddress(dataset.Address) (*dataset.Dataset, error)
}

func execSelect(stmt *Select, domain Domain) (result *dataset.Dataset, err error) {
	result = &dataset.Dataset{}
	// 1. Gather all mentioned tables, attaching them to dataset.Dataset
	for _, name := range stmt.From.TableNames() {
		if ds, e := domain.DatasetForAddress(dataset.NewAddress(name)); e != nil {
			err = e
			return
		} else {
			// if err = ds.Read(r.ObjectsPath()); err != nil {
			// 	return
			// }
			// ds.Data, err = ds.FetchBytes(r.base)
			// if err != nil {
			// 	return
			// }
			result.Datasets = append(result.Datasets, &ds.Dataset)
		}
	}

	// TODO... Sort each table by select sort criteria here?

	// 2. Build dataset destination fields
	for _, node := range stmt.SelectExprs {
		if star, ok := node.(*StarExpr); ok && node != nil {
			name := string(star.TableName)
			for i, ds := range result.Datasets {
				// we add fields if the names match, or if no name is specified
				if ds.Name == name || name == "" {
					result.Fields = append(result.Fields, ds.Fields...)
				}
			}
		} else if expr, ok := node.(*NonStarExpr); ok && node != nil {
			name := expr.FieldName()
			typeStr := expr.FieldType()
			idx := len(result.Fields) - 1
			result.Fields = append(result.Fields, &dataset.Field{
				Name: name,
				Type: datatype.TypeFromString(typeStr),
			})
		}
	}

	results := bytes.NewBuffer(nil)
	w := csv.NewWriter(results)
	result.Format = dataset.CsvDataFormat

	// 3. Populate dataset data by iterating through each dataset.dataset, projecting the source dataset onto the result dataset.
	// 		Then evaluate if the projected row passes any where clauses
	for i, ds := range result.Datasets {
		err = ds.EachRow(func(rowNum int, src [][]byte, e error) (err error) {
			if e != nil {
				return e
			}

			// check dst against criteria, only continue if it passes
			if pass, err := stmt.Where.Check(ds, src); err != nil {
				return err
			} else if !pass {
				return nil
			}

			dst := make([][]byte, len(result.Fields))
			for _, sExp := range stmt.SelectExprs {
				if err = sExp.Map(results, ds, src, dst); err != nil {
					return
				}
			}

			// ewwwwwwwwwww
			row := make([]string, len(dst))
			for i, col := range dst {
				row[i] = string(col)
			}

			// fmt.Println(row)
			w.Write(row)

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

func execUnion(node *Union, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("union statements are not yet supported")
}

func execInsert(node *Insert, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("insert statements are not yet supported")
}

func execUpdate(node *Update, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("update statements are not yet supported")
}

func execDelete(node *Delete, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("delete statements are not yet supported")
}

func execSet(node *Set, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("set statements are not yet supported")
}

func execDDL(node *DDL, domain Domain) (*dataset.Dataset, error) {
	return nil, fmt.Errof("ddl statements are not yet supported")
}

func execOther(node *Other, domain Domain) (*dataset.Dataset, error) {
	// TODO - lolololol
	return nil, fmt.Errof("other statements are not yet supported")
}
