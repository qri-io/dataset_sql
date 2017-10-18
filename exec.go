package dataset_sql

import (
	"fmt"
	"github.com/qri-io/cafs"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/dsio"
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
	if stmt.OrderBy != nil {
		return nil, nil, NotYetImplemented("ORDER BY statements")
	}

	result = ds.Query.Structure
	resources := map[string]*dataset.Structure{}
	for abst, _ := range remap {
		resources[abst] = ds.Query.Structures[abst]
	}

	if err := PopulateTableInfo(stmt, resources); err != nil {
		return result, nil, err
	}

	srg, err := NewSourceRowGenerator(store, ds.Resources)
	if err != nil {
		return result, nil, err
	}

	srf, err := NewSourceRowFilter(stmt)
	if err != nil {
		return result, nil, err
	}

	cols := CollectColNames(stmt)
	rg := NewRowGenerator(stmt, result)
	buf := dsio.NewBuffer(result)

	for srg.Next() && !srf.Done() {
		sr := srg.Row()

		if err := SetSourceRow(cols, sr); err != nil {
			return result, nil, err
		}

		if srf.Filter(sr) {

			row, err := rg.GenerateRow(sr)
			if err != nil {
				return result, nil, err
			}

			if err := buf.WriteRow(row); err != nil {
				return result, nil, err
			}

		}
	}

	// TODO - restore aggregate function writing
	// if agg {
	// 	row, err := aggFuncResults(funcs)
	// 	if err != nil {
	// 		return result, nil, err
	// 	}
	// 	// fmt.Println(row)
	// 	for _, r := range row {
	// 		fmt.Printf(string(r))
	// 	}
	// 	buf.WriteRow(row)
	// }

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

// populateColNames adds type information to ColName nodes in the ast
// func populateColNames(stmt *Select, from map[string]*StructureData) error {
// 	return stmt.WalkSubtree(func(sqlNode SQLNode) (bool, error) {
// 		switch node := sqlNode.(type) {
// 		case *ColName:
// 			if node.Qualifier.String() != "" {
// 				idx := 0
// 				for tableName, resourceData := range from {
// 					if node.Qualifier.String() == tableName {
// 						for i, f := range resourceData.Structure.Schema.Fields {
// 							if node.Name.String() == f.Name {
// 								node.Field = f
// 								node.RowIndex = idx + i
// 								return true, nil
// 							}
// 						}
// 					}
// 					idx += len(resourceData.Structure.Schema.Fields)
// 				}
// 				return false, fmt.Errorf("couldn't find field named '%s' in dataset '%s'", node.Name.String(), node.Qualifier.TableName())
// 			} else {
// 				idx := 0
// 				for _, resourceData := range from {
// 					for i, f := range resourceData.Structure.Schema.Fields {
// 						if node.Name.String() == f.Name {
// 							node.Field = f
// 							node.RowIndex = idx + i
// 							return true, nil
// 						}
// 					}
// 					idx += len(resourceData.Structure.Schema.Fields)
// 				}
// 				return false, fmt.Errorf("couldn't find field named '%s' in any of the specified datasets", node.Name.String())
// 			}
// 		}

// 		return true, nil
// 	})
// }

func aggFuncResults(funcs []AggFunc) (row [][]byte, err error) {
	row = make([][]byte, len(funcs))
	for i, fn := range funcs {
		row[i] = fn.Value()
	}
	return row, nil
}

// TODO - refactor StructureData to take a io.Reader instead of []byte
// type StructureData struct {
// 	Structure *dataset.Structure
// 	Data      []byte
// }

// fromFieldCount totals all fields in
// func fromFieldCount(from map[string]*StructureData) (count int) {
// 	for _, resourceData := range from {
// 		count += len(resourceData.Structure.Schema.Fields)
// 	}
// 	return
// }

// nodeColIndex finds the column index for a given node
// func nodeColIndex(node SelectExpr, from map[string]*StructureData) (idx int, err error) {
// 	if nse, ok := node.(*AliasedExpr); ok && node != nil {
// 		if colName, ok := nse.Expr.(*ColName); ok && node != nil {
// 			for _, resourceData := range from {
// 				for _, f := range resourceData.Structure.Schema.Fields {
// 					if f.Name == colName.Name.String() {
// 						return
// 					}
// 					idx++
// 				}
// 			}
// 		}
// 		return -1, nil
// 	}

// 	return 0, fmt.Errorf("node is not a non-star select expression")
// }

// intSeries returns a slice sized by length that counts from start upward
// func intSeries(start, length int) (series []int) {
// 	series = make([]int, length)
// 	for i := 0; i < length; i++ {
// 		series[i] = i + start
// 	}
// 	return
// }

// masterRowLength sums all fields of a dataset's children
// func masterRowLength(from map[string]*StructureData) (l int) {
// 	for _, resourceData := range from {
// 		l += len(resourceData.Structure.Schema.Fields)
// 	}
// 	return
// }

// nextRow generates the next master row for a dataset from the source datasets
// func nextRow(numStructures int, indicies, lengths []int, rowLen int, data [][][][]byte) (row [][]byte) {
// 	if incrIndicies(indicies, lengths) == nil {
// 		return nil
// 	} else {
// 		row = make([][]byte, rowLen)
// 		k := 0
// 		for i := 0; i < numStructures; i++ {
// 			// fmt.Println(i, indicies[i])
// 			for _, cell := range data[i][indicies[i]] {
// 				row[k] = cell
// 				k++
// 			}
// 		}
// 	}
// 	return
// }

// incrIndicies increments the index-counter, returning nil when
// counting is complete
// func incrIndicies(indicies, lengths []int) []int {
// 	for i := len(indicies) - 1; i >= 0; i-- {
// 		if indicies[i] < lengths[i]-1 {
// 			indicies[i]++
// 			break
// 		} else {
// 			if i-1 <= 0 && indicies[0] == lengths[0]-1 {
// 				return nil
// 			}
// 			indicies[i] = 0
// 			indicies[i-1]++
// 			break
// 		}
// 	}

// 	return indicies
// }

// jumpRow advances
// func jumpRow(indicies, lengths []int) bool {
// 	for i, idx := range indicies {
// 		if i == 0 {
// 			idx++
// 			if idx == lengths[i] {
// 				return true
// 			}
// 		} else if i == len(indicies)-1 {
// 			idx = -1
// 		} else {
// 			idx = 0
// 		}
// 	}
// 	return false
// }
