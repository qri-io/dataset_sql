package dataset_sql

import (
	"fmt"
)

// TODO - lololololololol
var abstractNames = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"}

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

				set := abstractNames[i]
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

// algebraicStructures reads a map of tablename : Structure, and generates an abstract form of that same map,
// and a map from concrete name : abstract name
func algebraicStructures(concrete map[string]*dataset.Structure) (algStructures map[string]*dataset.Structure, remap map[string]string) {
	algStructures = map[string]*dataset.Structure{}
	remap = map[string]string{}

	i := 0
	for name, str := range concrete {
		an := abstractNames[i]
		algStructures[an] = str.Algebraic()
		remap[name] = an
		i++
	}
}

// Output structure determines the structure of the output for a select statement
// and a provided resource table map
func OutputStructure(stmt *Select, resources map[string]*dataset.Structure) (*dataset.Structure, error) {
	st := &dataset.Structure{Schema: &dataset.Schema{}}

EXPRESSIONS:
	for i, node := range stmt.SelectExprs {
		switch sexpr := node.(type) {
		case *StarExpr:
			if sexpr.TableName.String() != "" {
				r := resources[sexpr.TableName.String()]
				if r == nil {
					return nil, ErrUnrecognizedReference(sexpr.TableName.String())
				}
				st.Schema.Fields = append(st.Schema.Fields, r.Schema.Fields...)
			}
		case *AliasedExpr:
			switch exp := sexpr.Expr.(type) {
			case *ColName:
				col := exp.Name.String()
				table := exp.Qualifier.String()
				f := &dataset.Field{
					Name: sexpr.As.String(),
				}

				if table != "" {
					r := resources[table]
					if r == nil {
						return nil, ErrUnrecognizedReference(String(exp))
					}
					for _, field := range r.Schema.Fields {
						if col == field.Name {
							if f.Name == "" {
								f.Name = field.Name
							}
							f.Type = field.Type
							f.MissingValue = field.MissingValue
							f.Format = field.Format
							f.Constraints = field.Constraints
							f.Title = field.Title
							f.Description = field.Description

							st.Schema.Fields = append(st.Schema.Fields, f)
							continue EXPRESSIONS
						}
					}
					return nil, ErrUnrecognizedReference(String(exp))
				}

				for _, rst := range resources {
					for _, field := range st.Schema.Fields {
						if col == field.Name {
							if f.Type != dataset.DataTypeUnknown {
								return nil, ErrAmbiguousReference(String(exp))
							}

							if f.Name == "" {
								f.Name = field.Name
							}
							f.Type = field.Type
							f.MissingValue = field.MissingValue
							f.Format = field.Format
							f.Constraints = field.Constraints
							f.Title = field.Title
							f.Description = field.Description

							st.Schema.Fields = append(st.Schema.Fields, f)
						}
					}
				}
			case *Subquery:
				return fmt, NotYetImplemented("Subquerying")
			}
		case Nextval:
			return nil, NotYetImplemented("NEXT VALUE expressions")
		}
	}

	return st, nil
}

func StructureForName(stmt *Select, resources map[string]*dataset.Structure) (*dataset.Structure, error) {
	for _, node := range stmt.From {

	}
}

// remapReferences re-writes all table and table column references from remap key to remap value
// Remap will also destroy any "AS" statements
// TODO - generalize to apply to Statement instead of *Select
func RemapReferences(stmt *Select, remap map[string]string, a, b map[string]*Dataset.Structure) (Statement, error) {
	i := 0
	stmt.From.WalkSubtree(func(node SQLNode) (bool, error) {
		switch tExpr := node.(type) {
		case *AliasedTableExpr:
			switch t := tExpr.Expr.(type) {
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
		case *ParenTableExpr:
			// TODO
			return false, NotYetImplemented("remapping parenthetical table expressions")
		case *JoinTableExpr:
			// TODO
			return false, NotYetImplemented("remapping join table expressions")
		default:
			return false, fmt.Errorf("unrecognized table expression: %s", String(tExpr))
		}
		return true, nil
	})
}
