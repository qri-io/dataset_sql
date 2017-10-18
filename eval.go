package dataset_sql

import (
	"bytes"
	"fmt"
	"github.com/qri-io/dataset/datatypes"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
)

// map bool to a unsigned 8 bit int
const QueryBoolType = q.Type_UINT8

func (node *AndExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	lt, lb, err := node.Left.Eval(row)
	if err != nil {
		return QueryBoolType, falseB, err
	}
	if lt != QueryBoolType {
		err = fmt.Errorf("non-boolean expression for left side of AND clause")
		return QueryBoolType, falseB, err
	}
	if !bytes.Equal(lb, trueB) {
		return QueryBoolType, falseB, nil
	}

	rt, rb, err := node.Right.Eval(row)
	if err != nil {
		return QueryBoolType, falseB, err
	}
	if rt != QueryBoolType {
		err = fmt.Errorf("non-boolean expression for right side of AND clause")
		return QueryBoolType, falseB, err
	}
	if !bytes.Equal(rb, trueB) {
		return QueryBoolType, falseB, nil
	}

	return QueryBoolType, trueB, nil
}

func (node *OrExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	lt, lb, err := node.Left.Eval(row)
	if err != nil {
		return QueryBoolType, falseB, err
	}
	if lt != QueryBoolType {
		err = fmt.Errorf("non-boolean expression for left side of AND clause: %s", String(node))
		return QueryBoolType, falseB, err
	}
	if bytes.Equal(lb, trueB) {
		return QueryBoolType, trueB, nil
	}

	rt, rb, err := node.Right.Eval(row)
	if err != nil {
		return QueryBoolType, falseB, err
	}
	if rt != QueryBoolType {
		err = fmt.Errorf("non-boolean expression for right side of AND clause: %s", String(node))
		return QueryBoolType, falseB, err
	}
	if bytes.Equal(rb, trueB) {
		return QueryBoolType, trueB, nil
	}

	return QueryBoolType, falseB, nil
}

func (node *NotExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	t, b, e := node.Expr.Eval(row)
	if t != QueryBoolType {
		e = fmt.Errorf("non-boolean expression for NOT expression: %s", String(node))
		return q.Type_NULL_TYPE, nil, e
	}
	if bytes.Equal(trueB, b) {
		return QueryBoolType, falseB, nil
	}
	// TODO - strange byte responses
	return QueryBoolType, trueB, nil
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
		t = q.Type_FLOAT32
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
	if sr, ok := node.Metadata.(StructureRef); ok {
		// TODO - this is a pretty decent indicator that we should switch
		// return types to our type system
		var t q.Type
		switch sr.Field.Type {
		case datatypes.Integer:
			t = q.Type_INT64
		case datatypes.Float:
			t = q.Type_FLOAT32
		case datatypes.String:
			t = q.Type_TEXT
		case datatypes.Boolean:
			t = QueryBoolType
		case datatypes.Date:
			t = q.Type_DATE
		default:
			return q.Type_NULL_TYPE, nil, fmt.Errorf("unsupported datatype for colname evaluation: %s", sr.Field.Type.String())
		}
		return t, row[sr.ColIndex], nil
	}
	return q.Type_NULL_TYPE, nil, fmt.Errorf("col ref %s hasn't been populated with structural information", node.Name.String())
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
	fmt.Println("eval func expr", node.Name.String())
	return node.fn.Eval(row)
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
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval finish")
}

func (node *Where) Eval(row [][]byte) (q.Type, []byte, error) {
	if node == nil {
		return QueryBoolType, trueB, nil
	}
	return node.Expr.Eval(row)
}

func (node *AliasedExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	return node.Expr.Eval(row)
}

func (nodes SelectExprs) Values(row [][]byte) (types []q.Type, vals [][]byte, err error) {
	for _, se := range nodes {
		switch node := se.(type) {
		case *StarExpr:
			ts, vs, e := node.Values(row)
			if e != nil {
				err = e
				return
			}
			types = append(types, ts...)
			vals = append(vals, vs...)
		case *AliasedExpr:
			t, v, e := node.Expr.Eval(row)
			if e != nil {
				err = e
				return
			}
			types = append(types, t)
			vals = append(vals, v)
		case Nextval:
			t, v, e := node.Value(row)
			if e != nil {
				err = e
				return
			}
			types = append(types, t)
			vals = append(vals, v)
		}
	}
	return
}

func (node *StarExpr) Values(row [][]byte) ([]q.Type, [][]byte, error) {
	return []q.Type{q.Type_NULL_TYPE}, nil, NotYetImplemented("star expession values")
}

func (node *Nextval) Value(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Nextval")
}

func (node TableName) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval TableName")
}

func (node OrderBy) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval OrderBy")
}

func (node GroupBy) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval GroupBy")
}

func (node TableExprs) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval TableExprs")
}

func (node *Set) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Set")
}

func (node Comments) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Comments")
}

func (node SelectExprs) Eval(row [][]byte) (q.Type, []byte, error) {
	if len(node) == 1 {
		return node[0].Eval(row)
	}
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval SelectExprs")
}

func (node *Limit) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Limit")
}

func (node Columns) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Columns")
}

func (node OnDup) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval OnDup")
}

func (node TableNames) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval TableNames")
}

func (node UpdateExprs) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval UpdateExprs")
}

func (node TableIdent) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval TableIdent")
}

func (node ColIdent) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ColIdent")
}

func (node IndexHints) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval IndexHints")
}

func (node Exprs) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Exprs")
}

func (node *ConvertType) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ConvertType")
}

func (node *When) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval When")
}

func (node *Order) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Order")
}

func (node *UpdateExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval UpdateExpr")
}

func (node *StarExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval StarExpr")
}

func (node Nextval) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Nextval")
}

func (node *Select) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Select")
}

func (node *AliasedTableExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval AliasedTableExpr")
}

func (node *ParenTableExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ParenTableExpr")
}

func (node *JoinTableExpr) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval JoinTableExpr")
}

func (node *Union) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Union")
}

func (node *ParenSelect) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval ParenSelect")
}

func (node *Insert) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Insert")
}

func (node *Delete) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Delete")
}

func (node *Update) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Update")
}

func (node *DDL) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval DDL")
}

func (node Values) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Values")
}

func (node *Show) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Show")
}

func (node *Use) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval Use")
}

func (node *OtherRead) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval OtherRead")
}

func (node *OtherAdmin) Eval(row [][]byte) (q.Type, []byte, error) {
	// TODO
	return q.Type_NULL_TYPE, nil, NotYetImplemented("eval OtherAdmin")
}
