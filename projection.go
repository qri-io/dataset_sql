package dataset_sql

import (
	"bytes"
	"github.com/qri-io/dataset"
	"github.com/qri-io/dataset/load"
)

func buildDatabase(from map[string]*StructureData, ds *dataset.Structure) (data [][][][]byte, lengths []int, err error) {
	for _, resourceData := range from {
		dsData, err := load.FormatRows(resourceData.Structure, bytes.NewReader(resourceData.Data))
		if err != nil {
			return nil, nil, err
		}

		lengths = append(lengths, len(dsData))
		data = append(data, dsData)
	}
	return
}

// generateResultSchema determines the schema of the query & adds it to result
func generateResultSchema(stmt *Select, from map[string]*StructureData, result *dataset.Structure) {
	if result.Schema == nil {
		result.Schema = &dataset.Schema{}
	}

	for _, node := range stmt.SelectExprs {
		if star, ok := node.(*StarExpr); ok && node != nil {
			name := star.TableName.String()
			for tableName, resourceData := range from {
				// we add fields if the names match, or if no name is specified
				if tableName == name || name == "" {
					result.Schema.Fields = append(result.Schema.Fields, resourceData.Structure.Schema.Fields...)
				}
			}
		} else if expr, ok := node.(*AliasedExpr); ok && node != nil {
			result.Schema.Fields = append(result.Schema.Fields, &dataset.Field{
				Name: expr.ResultName(),
				Type: expr.FieldType(from),
			})
		}
	}
}

// TODO - put this in a better place
func buildSelectorProjection(sqlNode SQLNode, proj *[]int, from map[string]*StructureData) error {
	switch node := sqlNode.(type) {
	case *ColName:
		idx := 0
		for _, resourceData := range from {
			for _, f := range resourceData.Structure.Schema.Fields {
				if f.Name == node.Name.String() {
					*proj = append(*proj, idx)
					return nil
				}
				idx++
			}
		}
	case *StarExpr:
		*proj = append(*proj, intSeries(0, fromFieldCount(from))...)
	case *AliasedExpr:
		if err := buildSelectorProjection(node.Expr, proj, from); err != nil {
			return err
		}
	case Nextval:
		return NotYetImplemented("building projections from nextVal")
	case SelectExprs:
		return node.buildSelectorProjection(proj, from)
	case *FuncExpr:
		return node.Exprs.buildSelectorProjection(proj, from)
	case *GroupConcatExpr:
		// TODO - wtf is a group concat expr?
		return node.Exprs.buildSelectorProjection(proj, from)
	case *MatchExpr:
		// TODO - wtf is a match expr?
		return node.Columns.buildSelectorProjection(proj, from)
	case *Select:
		return node.SelectExprs.buildSelectorProjection(proj, from)
	}
	return nil
}

func (selexprs SelectExprs) buildSelectorProjection(proj *[]int, from map[string]*StructureData) error {
	for _, se := range selexprs {
		if err := buildSelectorProjection(se, proj, from); err != nil {
			return err
		}
	}
	return nil
}

// buildProjection constructs the intermediate "projection" table that the sql query must
// generate in order to select form
func buildProjection(selectors SelectExprs, from map[string]*StructureData) (proj []int, err error) {
	proj = []int{}
	err = buildSelectorProjection(selectors, &proj, from)
	return
	// for _, node := range selectors {
	//  if isUnqualifiedStarExpr(node) {
	//    return intSeries(0, fromFieldCount(from)), nil
	//  } else if isQualifiedStarExpr(node) {
	//    r, e := findStarExprStructure(node, from)
	//    if e != nil {
	//      return proj, e
	//    }
	//    proj = append(proj, intSeries(len(proj), len(r.Schema.Fields))...)
	//  } else {
	//    i, e := nodeColIndex(node, from)
	//    if e != nil {
	//      return proj, e
	//    }
	//    proj = append(proj, i)
	//  }

	// }
	// return
}

// projectRow takes a master row & fits it to the desired result, evaluating any expressions along the way.
func projectRow(stmt SelectExprs, projection []int, source [][]byte) (row [][]byte, err error) {
	row = make([][]byte, len(projection))
	for i, _ := range projection {
		_, val, e := stmt[i].Eval(row)
		if e != nil {
			return row, e
		}
		row[i] = val
		// if j == -1 {
		//  _, val, e := stmt[i].Eval(row)
		//  if e != nil {
		//    return row, e
		//  }
		//  row[i] = val
		//  // switch node := stmt[i].(type) {
		//  // case *AliasedExpr:
		//  //  _, val, e := node.Expr.Eval(row)
		//  //  if e != nil {
		//  //    return row, e
		//  //  }
		//  //  row[i] = val
		//  // }
		// } else {
		//  row[i] = source[j]
		// }
	}
	return
}
