package dataset_sql

import (
	"fmt"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
)

// map bool to a unsigned 8 bit int
const QueryBoolType = q.Type_UINT8

// TODO - finish
func (node *AndExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval AndExpr")
}

// TODO - finish
func (node OrExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval or expression")
}

// TODO - finish
func (node *NotExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval NotExpr")
}

// TODO - finish
func (node *ParenExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ParenExpr")
}

func (node *ComparisonExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	result, err := node.Compare(row)
	if err != nil {
		return q.Type_NULL_TYPE, nil, err
	}
	if result {
		return QueryBoolType, trueB, nil
	}
	return QueryBoolType, falseB, nil
}

// TODO - finish
func (node *RangeCond) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval RangeCond")
}

// TODO - finish
func (node *IsExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval IsExpr")
}

// TODO - finish
func (node *ExistsExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ExistsExpr")
}

func (node *SQLVal) Eval(row [][]byte) (q.Type, []byte, error) {
	var t q.Type
	switch node.Type {
	case StrVal:
		t = q.Type_TEXT
	case IntVal:
		t = q.Type_INT64
	case FloatVal:
		t = q.Type_FLOAT64
	case HexNum:
		t = q.Type_BINARY
	case HexVal:
		t = q.Type_BLOB
	case ValArg:
		// TODO - is this correct?
		t = q.Type_EXPRESSION
	}
	return t, node.Val, nil
}

func (node *NullVal) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, nil
}

func (node BoolVal) Eval(row [][]byte) (q.Type, []byte, error) {
	if bool(node) == true {
		return QueryBoolType, trueB, nil
	}
	return QueryBoolType, falseB, nil
}

// Eval evaluates the node against a row of data
func (node *ColName) Eval(row [][]byte) (q.Type, []byte, error) {
	// switch node.Field.Type {
	// case datatypes.Any:
	//  return row[node.RowIndex], nil
	// case datatypes.String:
	//  return row[node.RowIndex], nil
	// case datatypes.Float:
	//  return row[node.RowIndex], nil
	// case datatypes.Integer:
	//  return row[node.RowIndex], nil
	// case datatypes.Date:
	//  return row[node.RowIndex], nil
	// case datatypes.Boolean:
	//  val, err := datatypes.ParseBoolean(row[node.RowIndex])

	//  return BoolVal(val), err
	// case datatypes.Object:
	//  // TODO
	//  return NewStrVal(row[node.RowIndex]), nil
	// case datatypes.Array:
	//  // TODO
	//  return NewStrVal(row[node.RowIndex]), nil
	// }
	// return nil, fmt.Errorf("couldn't find a column named '%s'", node.Name)
	return q.Type_NULL_TYPE, row[node.RowIndex], nil
}

func (node ValTuple) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO - huh?
	return q.Type_NULL_TYPE, nil, NotYetImplemented("val tuple Eval")
}

// TODO - finish
func (node *Subquery) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Subquery")
}

func (node ListArg) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO - huh?
	return q.Type_NULL_TYPE, node, nil
}

// TODO - finish
func (node *BinaryExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval BinaryExpr")
}

func (node *UnaryExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval UnaryExpr")
}

// TODO - finish
func (node *IntervalExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval IntervalExpr")
}

// TODO - finish
func (node *CollateExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval CollateExpr")
}

// TODO - finish
func (node *FuncExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval FuncExpr")
}

// TODO - finish
func (node *GroupConcatExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval GroupConcatExpr")
}

func (node *ValuesFuncExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	if node.Resolved == nil {
		return q.Type_NULL_TYPE, nil, fmt.Errorf("invalid values expression: %s", String(node))
	}
	return node.Resolved.Eval(row)
}

// TODO - finish
func (node *ConvertExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ConvertExpr")
}

// TODO - finish
func (node *ConvertUsingExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ConvertUsingExpr")
}

// TODO - finish
func (node *MatchExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval MatchExpr")
}

// TODO - finish
func (node *CaseExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval CaseExpr")
}

func (node *Where) Eval(row [][]byte) (q.Type, []byte, error) {
	if node == nil {
		return QueryBoolType, trueB, nil
	}
	return node.Expr.Eval(row)
}
