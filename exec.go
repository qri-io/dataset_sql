package dataset_sql

import (
	"fmt"
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
)

type ExecOpt struct {
	Format dataset.DataFormat
}

func opts(options ...func(*ExecOpt)) *ExecOpt {
	o := &ExecOpt{
		Format: dataset.CsvDataFormat,
	}
	for _, option := range options {
		option(o)
	}
	return o
}

func Exec(store cafs.Filestore, ds *dataset.Dataset, options ...func(o *ExecOpt)) (result *dataset.Structure, resultBytes []byte, err error) {
	opts := &ExecOpt{
		Format: dataset.CsvDataFormat,
	}
	for _, option := range options {
		option(opts)
	}

	if ds.QuerySyntax != "sql" {
		return nil, nil, fmt.Errorf("Invalid syntax: '%s' sql_dataset only supports sql syntax. ", ds.QuerySyntax)
	}

	concreteStmt, err := Parse(ds.QueryString)
	if err != nil {
		return nil, nil, err
	}
	err = RemoveUnusedReferences(concreteStmt, ds)
	if err != nil {
		return nil, nil, err
	}

	strs := map[string]*dataset.Structure{}
	for name, ds := range ds.Resources {
		strs[name] = ds.Structure
	}

	ds.Structure, err = ResultStructure(concreteStmt, strs, opts)
	if err != nil {
		return nil, nil, err
	}

	_, stmt, remap, err := Format(ds.QueryString)
	if err != nil {
		return nil, nil, err
	}

	ds.Query = &dataset.Query{
		Structures: map[string]*dataset.Structure{},
	}

	// collect table references
	for mapped, ref := range remap {
		// for i, adr := range stmt.References() {
		if ds.Resources[ref] == nil {
			return nil, nil, fmt.Errorf("couldn't find resource for table name: %s", ref)
		}

		ds.Query.Structures[mapped] = ds.Resources[ref].Structure.Abstract()
	}

	// This is a basic-column name rewriter from concrete to abstract
	stmt.WalkSubtree(func(node SQLNode) (bool, error) {
		// if ae, ok := node.(*AliasedExpr); ok && ae != nil {
		if cn, ok := node.(*ColName); ok && cn != nil {
			// TODO - check qualifier to avoid extra loopage
			// if cn.Qualifier.String() != "" {
			// 	for _, f := range ds.Query.Structures[cn.Qualifier.String()].Schema.Fields {
			// 		if cn.Name.String() ==
			// 	}
			// }
			for con, r := range ds.Resources {
				for i, f := range r.Structure.Schema.Fields {
					if f.Name == cn.Name.String() {
						for mapped, ref := range remap {
							if ref == con {
								// fmt.Println(ref, con, mapped)
								// fmt.Println("MATCH", ds.Query.Structures[mapped].Schema.Fields[i].Name)
								// fmt.Println(String(cn))
								// fmt.Println(String(&ColName{
								// 	Name:      NewColIdent(ds.Query.Structures[mapped].Schema.Fields[i].Name),
								// 	Qualifier: TableName{Name: NewTableIdent(mapped)},
								// }))

								*cn = ColName{
									Name:      NewColIdent(ds.Query.Structures[mapped].Schema.Fields[i].Name),
									Qualifier: TableName{Name: NewTableIdent(mapped)},
								}
							}
						}
						return false, nil
					}
				}
				// }

			}
		}
		return true, nil
	})

	ds.Query.Statement = String(stmt)

	ds.Query.Structure, err = ResultStructure(stmt, ds.Query.Structures, opts)
	if err != nil {
		return nil, nil, err
	}

	return stmt.exec(store, ds, remap, opts)
}

func (stmt *Select) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (result *dataset.Structure, resultBytes []byte, err error) {
	result = ds.Query.Structure
	resources := map[string]*dataset.Structure{}
	ads := map[string]*dataset.Dataset{}
	for abst, con := range remap {
		resources[abst] = ds.Query.Structures[abst]
		ads[abst] = ds.Resources[con]
	}

	if err := PrepareStatement(stmt, resources); err != nil {
		return result, nil, err
	}
	cols := CollectColNames(stmt)
	buf := NewResultBuffer(stmt, ds.Query.Structure.Abstract())

	srg, err := NewSourceRowGenerator(store, ads)
	if err != nil {
		return result, nil, err
	}

	srf, err := NewSourceRowFilter(stmt, buf)
	if err != nil {
		return result, nil, err
	}
	rrg, err := NewResultRowGenerator(stmt, result)
	if err != nil {
		return result, nil, err
	}

	for srg.Next() && !srf.Done() {
		sr, err := srg.Row()
		if err != nil {
			return result, nil, err
		}

		if err := SetSourceRow(cols, sr); err != nil {
			return result, nil, err
		}

		if srf.Match() {
			row, err := rrg.GenerateRow()
			if err == ErrAggStmt {
				continue
			} else if err != nil {
				return result, nil, err
			}

			if srf.ShouldWriteRow(row) {
				if err := buf.WriteRow(row); err != nil {
					return result, nil, err
				}
			}

		}
	}

	if rrg.HasAggregates() {
		row, err := rrg.GenerateAggregateRow()
		if err != nil {
			return result, nil, err
		}
		buf.WriteRow(row)
	}

	if err := buf.Close(); err != nil {
		return result, nil, err
	}

	if stmt.OrderBy != nil {

	}

	// TODO - rename / deref result var
	result = ds.Structure

	resultBytes = buf.Bytes()
	return
}

func (node *Union) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("union statements")
}
func (node *Insert) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("insert statements")
}
func (node *Update) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("update statements")
}
func (node *Delete) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("delete statements")
}
func (node *Set) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("set statements")
}
func (node *DDL) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("ddl statements")
}
func (node *ParenSelect) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("ParenSelect statements")
}
func (node *Show) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("Show statements")
}
func (node *Use) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("Use statements")
}
func (node *OtherRead) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("OtherRead statements")
}
func (node *OtherAdmin) exec(store cafs.Filestore, ds *dataset.Dataset, remap map[string]string, opts *ExecOpt) (*dataset.Structure, []byte, error) {
	return nil, nil, NotYetImplemented("OtherAdmin statements")
}
