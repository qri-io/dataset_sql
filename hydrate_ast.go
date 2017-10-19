package dataset_sql

import (
	"fmt"

	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
)

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
func PopulateTableInfo(tree SQLNode, resources map[string]*dataset.Structure) error {
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
