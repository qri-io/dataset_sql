package dataset_sql

import (
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
)

func PrepareStatement(stmt Statement, resources map[string]*dataset.Structure) (err error) {
	err = fitASTResources(stmt, resources)
	if err != nil {
		return
	}

	return populateTableInfo(stmt, resources)
}

// fitASTResources removes star expressions, replacing them with concrete colIdent
// pointers extracted from resources. It's important that no extraneous tables
// are in the resources map
func fitASTResources(ast SQLNode, resources map[string]*dataset.Structure) error {
	var visit func(node SQLNode) func(node SQLNode) (bool, error)
	visit = func(parent SQLNode) func(node SQLNode) (bool, error) {
		return func(child SQLNode) (bool, error) {
			if child == nil {
				return false, nil
			}

			switch node := child.(type) {
			case *StarExpr:
				if node != nil {
					se, err := starExprSelectExprs(node, resources)
					if err != nil {
						return false, err
					}
					return false, replaceSelectExprs(parent, node, se)
				}
			}

			return true, nil
		}
	}

	return ast.WalkSubtree(visit(ast))
}

func starExprSelectExprs(star *StarExpr, resources map[string]*dataset.Structure) (SelectExprs, error) {
	name := star.TableName.String()
	for tableName, resourceData := range resources {
		// we add fields if the names match, or if no name is specified
		if tableName == name || name == "" && len(resources) == 1 {
			se := make(SelectExprs, len(resourceData.Schema.Fields))
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
				se[i] = &AliasedExpr{As: NewColIdent(f.Name), Expr: col}
			}
			return se, nil
		}
	}
	return nil, fmt.Errorf("couldn't find table for star expression: '%s'", name)
}

func replaceSelectExprs(parent, prev SQLNode, se SelectExprs) error {
	switch node := parent.(type) {
	case *Select:
		for i, exp := range node.SelectExprs {
			if exp == prev {
				node.SelectExprs = spliceSelectExprs(node.SelectExprs, se, i)
				return nil
			}
		}
	}
	return fmt.Errorf("couldn't find selectExprs for parent")
}

func spliceSelectExprs(a, b SelectExprs, pos int) SelectExprs {
	return append(a[:pos], append(b, a[pos+1:]...)...)
}

// StructureRef is placed on ColName SQLNodes to
// connect typing & data lookup information
type StructureRef struct {
	TableName string
	ColIndex  int
	Field     *dataset.Field
	QueryType q.Type
}

// PopulateAST adds type information & data lookup locations to an AST
// for a given resource.
// TODO - column ambiguity check
func populateTableInfo(tree SQLNode, resources map[string]*dataset.Structure) error {
	return tree.WalkSubtree(func(node SQLNode) (bool, error) {
		if col, ok := node.(*ColName); ok && node != nil {
			if col.Qualifier.TableName() != "" && resources[col.Qualifier.TableName()] != nil {
				for i, f := range resources[col.Qualifier.TableName()].Schema.Fields {
					if col.Name.String() == f.Name {
						qt := QueryTypeForDataType(f.Type)
						if qt == q.Type_NULL_TYPE {
							return false, fmt.Errorf("unsupported datatype for colname evaluation: %s", f.Type.String())
						}
						col.Metadata = StructureRef{
							Field:     f,
							TableName: col.Qualifier.TableName(),
							ColIndex:  i,
							QueryType: qt,
						}
						return true, nil
					}
				}
				return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", col.Name.String(), col.Qualifier.TableName())
			} else {
				for tableName, st := range resources {
					for i, f := range st.Schema.Fields {
						if col.Name.String() == f.Name {
							col.Qualifier = TableName{Name: NewTableIdent(tableName)}

							qt := QueryTypeForDataType(f.Type)
							if qt == q.Type_NULL_TYPE {
								return false, fmt.Errorf("unsupported datatype for colname evaluation: %s", f.Type.String())
							}
							col.Metadata = StructureRef{
								Field:     f,
								TableName: tableName,
								QueryType: qt,
								ColIndex:  i,
							}
							return true, nil
						}
					}
				}
				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", col.Name.String())
			}
		}
		return true, nil
	})
}

func CollectColNames(tree SQLNode) (cols []*ColName) {
	tree.WalkSubtree(func(node SQLNode) (bool, error) {
		if col, ok := node.(*ColName); ok && node != nil {
			cols = append(cols, col)
		}
		return true, nil
	})
	return
}

func SetSourceRow(cols []*ColName, sr SourceRow) error {
	for _, col := range cols {
		if col.Metadata.TableName == "" {
			return fmt.Errorf("col missing metadata: %#v", col)
		}
		if col.Metadata.ColIndex > len(sr[col.Metadata.TableName])-1 {
			return fmt.Errorf("index out of range to set column value: %s.%d", col.Metadata.TableName, col.Metadata.ColIndex)
		}
		col.Value = sr[col.Metadata.TableName][col.Metadata.ColIndex]
	}
	return nil
}

func QueryTypeForDataType(t datatypes.Type) q.Type {
	switch t {
	case datatypes.Integer:
		return q.Type_INT64
	case datatypes.Float:
		return q.Type_FLOAT32
	case datatypes.String:
		return q.Type_TEXT
	case datatypes.Boolean:
		return QueryBoolType
	case datatypes.Date:
		return q.Type_DATE
	default:
		return q.Type_NULL_TYPE
	}
}
