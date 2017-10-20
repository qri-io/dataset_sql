package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset"
)

// remapStatement removes star expressions, replacing them with concrete colIdent
// pointers extracted from resources. It's important that no extraneous tables
// are in the resources map
func PrepareStatement(stmt Statement, resources map[string]*dataset.Structure) error {
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

	return stmt.WalkSubtree(visit(stmt))
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
