package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset"
)

type StructureRef struct {
	TableName string
	ColIndex  int
	Field     *dataset.Field
}

// PopulateAST adds type information & data lookup locations to an AST
// for a given resource.
func StructureAST(tree SQLNode, resources map[string]*dataset.Structure) error {
	return tree.WalkSubtree(func(node SQLNode) (bool, error) {
		if col, ok := node.(*ColName); ok && node != nil {
			if col.Qualifier.TableName() != "" && resources[col.Qualifier.TableName()] != nil {
				for i, f := range resources[col.Qualifier.TableName()].Schema.Fields {
					if col.Name.String() == f.Name {
						col.Metadata = StructureRef{
							Field:     f,
							TableName: col.Qualifier.TableName(),
							ColIndex:  i,
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
							col.Metadata = StructureRef{
								Field:     f,
								TableName: tableName,
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
