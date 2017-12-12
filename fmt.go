package dataset_sql

import (
	"fmt"

	"github.com/ipfs/go-datastore"
	"github.com/qri-io/cafs"
	"github.com/qri-io/cafs/memfs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/datatypes"
	"github.com/qri-io/dataset/dsfs"
	"github.com/qri-io/dataset/validate"
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

// QueryRecordPath returns the hash of an abstracted query to be excectuted with a given set of resources.
// the returned key can be used to see if a a given query has been run before
func QueryRecordPath(store cafs.Filestore, q *dataset.Transform, opts ...func(o *ExecOpt)) (datastore.Key, error) {
	save := &dataset.Transform{}
	save.Assign(q)
	stmt, abst, err := Format(save, func(o *ExecOpt) {
		o.Format = dataset.CSVDataFormat
	})
	if err != nil {
		return datastore.NewKey(""), fmt.Errorf("formatting error: %s", err.Error())
	}

	save.Assign(abst)
	save.Kind = dataset.KindTransform
	save.Data = String(stmt)

	if save.Structure == nil {
		return datastore.NewKey(""), fmt.Errorf("structure required to save abstract transform")
	}

	save.Structure = save.Structure.Abstract()

	// ensure all dataset references are abstract
	for key, r := range save.Resources {
		data := r.Data
		rsc := dataset.Abstract(r)
		rsc.Data = data
		save.Resources[key] = rsc
	}

	data, err := save.MarshalJSON()
	if err != nil {
		return datastore.NewKey(""), fmt.Errorf("error marshaling dataset abstract transform to json: %s", err.Error())
	}
	return store.Put(memfs.NewMemfileBytes(dsfs.PackageFileTransform.String(), data), false)
}

// Format places an sql statement in it's standard form.
// This will be *heavily* refined, improved, and moved into a
// separate package
// TODO - ^^
// It's expected that the query to be exectuted will be a string in the
// given datset.Data value
// Format will modify the incoming
func Format(q *dataset.Transform, opts ...func(o *ExecOpt)) (stmt Statement, abst *dataset.Transform, err error) {
	opt := DefaultExecOpts()
	for _, o := range opts {
		o(opt)
	}

	stmt, err = Parse(q.Data)
	if err != nil {
		return
	}

	if q.Syntax != "sql" {
		return nil, nil, fmt.Errorf("Invalid syntax: '%s' dataset_sql only supports sql syntax. ", q.Syntax)
	}

	if err = validResources(q.Resources); err != nil {
		return
	}

	err = RemoveUnusedReferences(stmt, q)
	if err != nil {
		return
	}

	if err = containsAmbiguousReference(stmt, q.Resources); err != nil {
		return
	}

	q.Structure, err = ResultStructure(stmt, q.Resources, opt)
	if err != nil {
		return
	}

	abst = &dataset.Transform{
		Resources: map[string]*dataset.Dataset{},
	}

	// this is a basic table-name rewriter from concrete to abstract
	i := 0
	remap := map[string]string{}
	stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		if ate, ok := node.(*AliasedTableExpr); ok && ate != nil {
			switch t := ate.Expr.(type) {
			case TableName:
				current := t.TableName()
				for set, prev := range remap {
					if current == prev {
						ate.Expr = TableName{Name: TableIdent{set}}
						return true, nil
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

	// collect table references
	for mapped, ref := range remap {
		if q.Resources[ref] == nil {
			err = fmt.Errorf("couldn't find resource for table name: %s", ref)
			return
		}
		abst.Resources[mapped] = dataset.Abstract(q.Resources[ref])
		abst.Resources[mapped].Data = q.Resources[ref].Data
	}

	// This is a basic column-name rewriter from concrete to abstract
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
									Name:      NewColIdent(abst.Resources[abstName].Structure.Schema.Fields[i].Name),
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
		return
	}

	q.Syntax = "sql"
	abst.Syntax = "sql"
	abst.Data = String(stmt)
	abst.Structure = q.Structure.Abstract()

	err = PrepareStatement(stmt, abst.Resources)
	return
}

func validResources(resources map[string]*dataset.Dataset) error {
	for name, ds := range resources {
		if err := validate.ValidName(name); err != nil {
			return err
		}
		if ds == nil {
			return fmt.Errorf("invalid resource reference: %s", name)
		}
		if ds.Structure == nil {
			return fmt.Errorf("dataset structure is required for resource '%s'", name)
		}
	}
	return nil
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
func ResultStructure(stmt Statement, resources map[string]*dataset.Dataset, opts *ExecOpt) (*dataset.Structure, error) {
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
						st.Schema.Fields = append(st.Schema.Fields, r.Structure.Schema.Fields...)
					}
				}
			} else {
				if r := resources[name]; r != nil {
					st.Schema.Fields = append(st.Schema.Fields, r.Structure.Schema.Fields...)
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
					for _, field := range r.Structure.Schema.Fields {
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
					for _, field := range rsc.Structure.Schema.Fields {
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
func RemoveUnusedReferences(stmt Statement, q *dataset.Transform) error {
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
