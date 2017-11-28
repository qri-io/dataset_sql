package dataset_sql

import (
	"fmt"
	"github.com/ipfs/go-datastore"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
)

// StatementTableNames extracts the names of all referenced tables
// from a given statement
func StatementTableNames(sql string) ([]string, error) {
	stmt, err := Parse(sql)
	if err != nil {
		return nil, err
	}

	if sel, ok := stmt.(*Select); ok {
		return sel.From.TableNames(), nil
	}

	return nil, fmt.Errorf("unsupported statement type: %s", String(stmt))
}

// Format places an sql statement in it's standard form.
// This will be *heavily* refined, improved, and moved into a
// separate package
// TODO - ^^
func Format(q *dataset.Query) (string, Statement, map[string]string, error) {
	remap := map[string]string{}
	stmt, err := Parse(q.Abstract.Statement)
	if err != nil {
		return "", nil, nil, err
	}

	q.Abstract = &dataset.AbstractQuery{
		Structures: map[string]*dataset.Structure{},
	}

	i := 0
	stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		if ate, ok := node.(*AliasedTableExpr); ok && ate != nil {
			switch t := ate.Expr.(type) {
			case TableName:
				current := t.TableName()
				for set, prev := range remap {
					if current == prev {
						ate.Expr = TableName{Name: TableIdent{set}}
						// return true, nil
					}
				}

				set := dataset.AbstractTableName(i)
				i++
				remap[set] = current
				ate.Expr = TableName{Name: TableIdent{set}}
				return true, nil
			}
		}
		return true, nil
	})

	paths := map[string]datastore.Key{}
	// collect table references
	for mapped, ref := range remap {
		// for i, adr := range stmt.References() {
		if q.Resources[ref] == nil {
			return "", nil, nil, fmt.Errorf("couldn't find resource for table name: %s", ref)
		}
		paths[mapped] = q.Resources[ref].Data
		q.Abstract.Structures[mapped] = q.Resources[ref].Structure.Abstract()
	}

	if err := containsAmbiguousReference(stmt, q.Resources); err != nil {
		return "", nil, nil, err
	}

	// This is a basic-column name rewriter from concrete to abstract
	err = stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		if cn, ok := node.(*ColName); ok && cn != nil {
			t := cn.Qualifier.String()
			for concreteName, resource := range q.Resources {
				if t != "" && concreteName != t {
					continue
				}
				for i, f := range resource.Structure.Schema.Fields {
					if f.Name == cn.Name.String() {
						for abstName, conName := range remap {
							if conName == concreteName {
								*cn = ColName{
									Name:      NewColIdent(q.Abstract.Structures[abstName].Schema.Fields[i].Name),
									Qualifier: TableName{Name: NewTableIdent(abstName)},
								}
							}
						}
						return true, nil
					}
				}
			}
		}
		return true, nil
	})
	if err != nil {
		return "", nil, nil, err
	}

	buf := NewTrackedBuffer(nil)
	stmt.Format(buf)

	return buf.String(), stmt, remap, nil
}

func containsAmbiguousReference(stmt Statement, resources map[string]*dataset.Dataset) error {
	return stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		if col, ok := node.(*ColName); ok && col != nil {
			qual := col.Qualifier.String()
			if qual != "" {
				return true, nil
			}
			ref := 0
			for _, ds := range resources {
				if ds.Structure.StringFieldIndex(col.Name.String()) >= 0 {
					ref++
				}
				if ref > 1 {
					return false, fmt.Errorf("column reference '%s' is ambiguous, please specify the dataset name for this table", String(col))
				}
			}
		}
		return true, nil
	})
}

// abstractStructures reads a map of tablename : Structure, and generates an abstract form of that same map,
// and a map from concrete name : abstract name
func abstractStructures(concrete map[string]*dataset.Structure) (algStructures map[string]*dataset.Structure, remap map[string]string) {
	algStructures = map[string]*dataset.Structure{}
	remap = map[string]string{}

	i := 0
	for name, str := range concrete {
		an := dataset.AbstractColumnName(i)
		algStructures[an] = str.Abstract()
		remap[name] = an
		i++
	}

	return
}

// ResultStructure determines the structure of the output for a select statement
// and a provided resource table map
func ResultStructure(stmt Statement, resources map[string]*dataset.Structure, opts *ExecOpt) (*dataset.Structure, error) {
	sel, ok := stmt.(*Select)
	if !ok {
		return nil, NotYetImplemented("statements other than select")
	}

	st := &dataset.Structure{Format: opts.Format, Schema: &dataset.Schema{}}

EXPRESSIONS:
	for _, node := range sel.SelectExprs {
		switch sexpr := node.(type) {
		case *StarExpr:
			name := sexpr.TableName.String()
			if name == "" {
				// unqualified star expressions should join in the order tables are referenced in
				// the from clause
				for _, name := range sel.From.TableNames() {
					if r := resources[name]; r != nil {
						st.Schema.Fields = append(st.Schema.Fields, r.Schema.Fields...)
					}
				}
			} else {
				if r := resources[name]; r != nil {
					st.Schema.Fields = append(st.Schema.Fields, r.Schema.Fields...)
				}
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

				for _, rsc := range resources {
					for _, field := range rsc.Schema.Fields {
						if col == field.Name {
							if f.Type != datatypes.Unknown {
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
			case *FuncExpr:
				st.Schema.Fields = append(st.Schema.Fields, &dataset.Field{
					Name: exp.Name.String(),
					Type: exp.Datatype(),
				})

			case *Subquery:
				return nil, NotYetImplemented("Subquerying")
			}
		case Nextval:
			return nil, NotYetImplemented("NEXT VALUE expressions")
		}
	}

	return st, nil
}

// RemoveUnusedReferences sets ds.Resources to a new map that that contains
// only datasets refrerenced in the provided select statement,
// it errors if it cannot find a named dataset from the provided ds.Resources map.
func RemoveUnusedReferences(stmt Statement, q *dataset.Query) error {
	sel, ok := stmt.(*Select)
	if !ok {
		return NotYetImplemented("statements other than select")
	}

	resources := map[string]*dataset.Dataset{}
	for _, name := range sel.From.TableNames() {
		datas := q.Resources[name]
		if datas == nil {
			return ErrUnrecognizedReference(name)
		}
		resources[name] = datas
	}
	q.Resources = resources
	return nil
}

// RemapReferences re-writes all table and table column references from remap key to remap value
// Remap will destroy any table-aliasing ("as" statements)
// TODO - generalize to apply to Statement instead of *Select
// TODO - need to finish support for remapping column refs
func RemapReferences(stmt *Select, remap map[string]string, a, b map[string]*dataset.Structure) (Statement, error) {
	// i := 0
	err := stmt.From.WalkSubtree(func(node SQLNode) (bool, error) {
		switch tExpr := node.(type) {
		case *AliasedTableExpr:
			switch t := tExpr.Expr.(type) {
			case TableName:
				current := t.TableName()
				if remap[current] == "" {
					return false, ErrUnrecognizedReference(current)
				}

				tExpr.Expr = TableName{Name: TableIdent{remap[current]}}
				return true, nil
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
	return stmt, err
}
