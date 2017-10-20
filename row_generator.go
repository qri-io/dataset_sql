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
	aggs  []AggFunc
	st    *dataset.Structure
}

func NewRowGenerator(sel *Select, result *dataset.Structure) (rg *RowGenerator, err error) {
	// sel, ok := ast.(*Select)
	// if !ok {
	// 	return nil, NotYetImplemented("row generation for non-SELECT statements")
	// }

	rg = &RowGenerator{
		exprs: sel.SelectExprs,
		st:    result,
	}
	// rg.exprs, err = generateResultSelectExprs(sel.SelectExprs, resources, result, cols)
	// if err != nil {
	// 	return
	// }
	// rg.exprs = sel.SelectExprs
	// rg.st = result

	rg.aggs, err = AggregateFuncs(sel)
	if err != nil {
		return nil, err
	}
	return
}

var (
	ErrAggStmt   = fmt.Errorf("this statement only generates an aggregate result row")
	ErrTableStmt = fmt.Errorf("this statement doesn't generate an aggregate result row")
)

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

	if !rg.HasAggregates() {
		return row, nil
	}
	return nil, ErrAggStmt
}

func (rg *RowGenerator) HasAggregates() bool {
	return len(rg.aggs) > 0
}

func (rg *RowGenerator) GenerateAggregateRow() ([][]byte, error) {
	if rg.HasAggregates() {

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
	return nil, ErrTableStmt
}

func (rg *RowGenerator) Structure() *dataset.Structure {
	return rg.st
}

// TODO - make this suck less by re-writing asts to remove star expressions much earlier in execution
// func generateResultSelectExprs(stmt SelectExprs, from map[string]*dataset.Structure, result *dataset.Structure, cols *[]*ColName) (se SelectExprs, err error) {
// 	if result.Schema == nil {
// 		result.Schema = &dataset.Schema{}
// 	}

// 	se = SelectExprs{}
// 	err = stmt.WalkSubtree(func(node SQLNode) (bool, error) {
// 		if node == nil {
// 			return true, nil
// 		}

// 		switch n := node.(type) {
// 		case *StarExpr:
// 			if node != nil {
// 				name := n.TableName.String()
// 				for tableName, resourceData := range from {
// 					// we add fields if the names match, or if no name is specified
// 					if tableName == name || name == "" {
// 						result.Schema.Fields = append(result.Schema.Fields, resourceData.Schema.Fields...)

// 						for i, f := range resourceData.Schema.Fields {
// 							col := &ColName{
// 								Name:      NewColIdent(f.Name),
// 								Qualifier: TableName{Qualifier: NewTableIdent(tableName)},
// 								Metadata: StructureRef{
// 									TableName: tableName,
// 									Field:     f,
// 									ColIndex:  i,
// 								},
// 							}
// 							*cols = append(*cols, col)
// 							se = append(se, &AliasedExpr{As: NewColIdent(f.Name), Expr: col})
// 						}
// 					}
// 				}
// 			}
// 		case *AliasedExpr:
// 			if n != nil {
// 				se = append(se, n)
// 				result.Schema.Fields = append(result.Schema.Fields, &dataset.Field{
// 					Name: n.ResultName(),
// 					Type: n.FieldType(from),
// 				})
// 			}
// 			// case Nextval:
// 		}
// 		return true, nil
// 	})

// 	return
// }
