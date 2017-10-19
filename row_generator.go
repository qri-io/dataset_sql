package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset"
)

// RowGenerator makes rows from SourceRows
// calling eval on a set of select expressions from a given
// SourceRow
type RowGenerator struct {
	ast Statement
	st  *dataset.Structure
}

func NewRowGenerator(ast Statement, st *dataset.Structure) *RowGenerator {
	// funcs, err := AggregateFuncs(ast)
	// if err != nil {
	// 	return nil, err
	// }

	return &RowGenerator{
		ast: ast,
		st:  st,
	}
}

// GenerateRow generates a row
func (rg *RowGenerator) GenerateRow(sr SourceRow) ([][]byte, error) {
	return nil, nil
}

func (rg *RowGenerator) GenerateAggregateRow() ([][]byte, error) {
	return nil, fmt.Errorf("aggregate row not finished")
}

func (rg *RowGenerator) Structure() *dataset.Structure {
	return rg.st
}

// Gather all mentioned tables, attaching them to a *dataset.Structure
// TODO - refactor this out
// func buildResultStructure(stmt *Select, store cafs.Filestore, resources map[string]*dataset.Dataset, opts *ExecOpt) (from map[string]*dataset.Structure, result *dataset.Structure, err error) {
// 	structures := map[string]*dataset.Structure{}
// 	for name, ds := range resources {
// 		st := ds.Structure

// 		// file, e := store.Get(ds.Data)
// 		// if e != nil {
// 		//   err = fmt.Errorf("error getting dataset file: %s: %s", ds.Data, e.Error())
// 		//   return
// 		// }

// 		// TODO - shim until structured data refactor
// 		data, e := ioutil.ReadAll(file)
// 		if e != nil {
// 			err = fmt.Errorf("error loading dataset data: %s: %s", ds.Data, e.Error())
// 			return
// 		}

// 		from[name] = &StructureData{
// 			Structure: st,
// 			Data:      data,
// 		}

// 		structures[name] = st
// 	}

// 	result, err = ResultStructure(stmt, structures, opts)
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// generateResultSchema determines the schema of the query & adds it to result
// func generateResultSchema(stmt *Select, from map[string]*StructureData, result *dataset.Structure) {
// 	if result.Schema == nil {
// 		result.Schema = &dataset.Schema{}
// 	}

// 	for _, node := range stmt.SelectExprs {
// 		if star, ok := node.(*StarExpr); ok && node != nil {
// 			name := star.TableName.String()
// 			for tableName, resourceData := range from {
// 				// we add fields if the names match, or if no name is specified
// 				if tableName == name || name == "" {
// 					result.Schema.Fields = append(result.Schema.Fields, resourceData.Structure.Schema.Fields...)
// 				}
// 			}
// 		} else if expr, ok := node.(*AliasedExpr); ok && node != nil {
// 			result.Schema.Fields = append(result.Schema.Fields, &dataset.Field{
// 				Name: expr.ResultName(),
// 				Type: expr.FieldType(from),
// 			})
// 		}
// 	}
// }
