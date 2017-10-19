package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset"
)

// RowGenerator makes rows from SourceRows
// calling eval on a set of select expressions from a given
// SourceRow
type RowGenerator struct {
	exprs SelectExprs
	st    *dataset.Structure
}

func NewRowGenerator(ast Statement, resources map[string]*dataset.Structure, result *dataset.Structure, cols *[]*ColName) (rg *RowGenerator, err error) {
	// funcs, err := AggregateFuncs(ast)
	// if err != nil {
	// 	return nil, err
	// }
	rg = &RowGenerator{}

	rg.exprs, err = generateResultSelectExprs(ast.(*Select).SelectExprs, resources, result, cols)
	rg.st = result
	return
}

// GenerateRow generates a row
func (rg *RowGenerator) GenerateRow() ([][]byte, error) {
	row := make([][]byte, len(rg.exprs))
	for i, expr := range rg.exprs {
		_, data, err := expr.Eval()
		if err != nil {
			return nil, err
		}
		row[i] = data
	}
	return row, nil
}

func (rg *RowGenerator) GenerateAggregateRow() ([][]byte, error) {
	return nil, fmt.Errorf("aggregate row not finished")
}

func (rg *RowGenerator) Structure() *dataset.Structure {
	return rg.st
}

// Gather all mentioned tables, attaching them to a *dataset.Structure
// TODO - refactor this out
// func buildResultStructure(stmt *Select, store cafs.Filestore, resources map[string]*dataset.Dataset, opts *ExecOpt) (from map[string]*dataset.Structure, result *dataset.Structure, err error) {
// 	structures := map[string]*dataset.Structure{}
// 	for name, ds := range resources {
// 		st := ds.Structure

// 		// file, e := store.Get(ds.Data)
// 		// if e != nil {
// 		//   err = fmt.Errorf("error getting dataset file: %s: %s", ds.Data, e.Error())
// 		//   return
// 		// }

// 		// TODO - shim until structured data refactor
// 		data, e := ioutil.ReadAll(file)
// 		if e != nil {
// 			err = fmt.Errorf("error loading dataset data: %s: %s", ds.Data, e.Error())
// 			return
// 		}

// 		from[name] = &StructureData{
// 			Structure: st,
// 			Data:      data,
// 		}

// 		structures[name] = st
// 	}

// 	result, err = ResultStructure(stmt, structures, opts)
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// generateResultSchema determines the schema of the query & adds it to result
// func generateResultSchema(ast Statement, from map[string]*dataset.Structure, result *dataset.Structure, cols []*ColName) (SelectExprs, error) {
// 	if result.Schema == nil {
// 		result.Schema = &dataset.Schema{}
// 	}

// 	switch stmt := ast.(type) {
// 	case *Select:
// 		for _, node := range stmt.SelectExprs {
// 			if star, ok := node.(*StarExpr); ok && node != nil {
// 				name := star.TableName.String()
// 				for tableName, resourceData := range from {
// 					// we add fields if the names match, or if no name is specified
// 					if tableName == name || name == "" {
// 						result.Schema.Fields = append(result.Schema.Fields, resourceData.Schema.Fields...)
// 					}
// 				}
// 			} else if expr, ok := node.(*AliasedExpr); ok && node != nil {
// 				result.Schema.Fields = append(result.Schema.Fields, &dataset.Field{
// 					Name: expr.ResultName(),
// 					Type: expr.FieldType(from),
// 				})
// 			}
// 		}
// 		return stmt.SelectExprs, nil
// 	}

// 	return nil, NotYetImplemented("result schemas for statements other than select")
// }

func generateResultSelectExprs(stmt SelectExprs, from map[string]*dataset.Structure, result *dataset.Structure, cols *[]*ColName) (se SelectExprs, err error) {
	if result.Schema == nil {
		result.Schema = &dataset.Schema{}
	}

	se = SelectExprs{}
	err = stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		if node == nil {
			return true, nil
		}

		switch n := node.(type) {
		case *StarExpr:
			if node != nil {
				name := n.TableName.String()
				for tableName, resourceData := range from {
					// we add fields if the names match, or if no name is specified
					if tableName == name || name == "" {
						result.Schema.Fields = append(result.Schema.Fields, resourceData.Schema.Fields...)

						for i, f := range resourceData.Schema.Fields {
							col := &ColName{
								Name:      NewColIdent(f.Name),
								Qualifier: TableName{Qualifier: NewTableIdent(tableName)},
								Metadata: StructureRef{
									TableName: tableName,
									Field:     f,
									ColIndex:  i,
								},
							}
							*cols = append(*cols, col)
							se = append(se, &AliasedExpr{As: NewColIdent(f.Name), Expr: col})
						}
					}
				}
			}
		case *AliasedExpr:
			if n != nil {
				se = append(se, n)
				result.Schema.Fields = append(result.Schema.Fields, &dataset.Field{
					Name: n.ResultName(),
					Type: n.FieldType(from),
				})
			}
			// case Nextval:
		}
		return true, nil
	})

	return
}
