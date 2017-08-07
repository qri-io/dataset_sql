package dataset_sql

import (
	"fmt"
)

// TODO - lololololololol
var names = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

// Format places an sql statement in it's standard form.
// This will be *heavily* refined, improved, and moved into a
// separate package
// TODO - milestone & break down this core piece of tech
func Format(sql string) (string, Statement, map[string]string, error) {
	remap := map[string]string{}
	stmt, err := Parse(sql)
	if err != nil {
		return "", nil, nil, err
	}

	sel, ok := stmt.(*Select)
	if !ok {
		return "", nil, nil, fmt.Errorf("dataset_sql: Format currently only supports 'select' statements")
	}

	i := 0
	sel.From.WalkSubtree(func(node SQLNode) (bool, error) {
		if ate, ok := node.(*AliasedTableExpr); ok && ate != nil {
			switch t := ate.Expr.(type) {
			case TableName:
				current := t.TableName()
				for set, prev := range remap {
					if current == prev {
						ate.Expr = TableName{Name: TableIdent{set}}
						return false, nil
					}
				}

				set := names[i]
				i++
				remap[set] = current
				ate.Expr = TableName{Name: TableIdent{set}}
				return false, nil
			}
		}
		return true, nil
	})

	buf := NewTrackedBuffer(nil)
	stmt.Format(buf)

	return buf.String(), stmt, remap, nil
}
