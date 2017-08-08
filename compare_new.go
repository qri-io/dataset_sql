package dataset_sql

import (
	"fmt"
	"github.com/qri-io/dataset_sql/sqltypes"
	// pb "github.com/qri-io/dataset_sql/vt/proto/query"
)

func (node *ComparisonExpr) Compare(row [][]byte) (bool, error) {
	_, left, err := node.Left.Eval(row)
	if err != nil {
		return false, err
	}
	_, right, err := node.Right.Eval(row)
	if err != nil {
		return false, err
	}

	l, err := sqltypes.BuildValue(left)
	if err != nil {
		return false, err
	}
	r, err := sqltypes.BuildValue(right)
	if err != nil {
		return false, err
	}

	result, err := sqltypes.NullsafeCompare(l, r)
	if err != nil {
		return false, err
	}

	switch node.Operator {
	case EqualStr:
		return result == 0, nil
	case LessThanStr:
		return result == -1, nil
	case GreaterThanStr:
		return result == 1, nil
	case LessEqualStr:
		return result == -1 || result == 0, nil
	case GreaterEqualStr:
		return result == 1 || result == 0, nil
	case NotEqualStr:
		return result == -1 || result == 1, nil
	case NullSafeEqualStr:
		// TODO - work through NSE case
		return result == -1 || result == 1, nil
	case InStr:
		return false, NotYetImplemented("InStr comparison")
	case NotInStr:
		return false, NotYetImplemented("NotInStr comparison")
	case LikeStr:
		return false, NotYetImplemented("LikeStr comparison")
	case NotLikeStr:
		return false, NotYetImplemented("NotLikeStr comparison")
	case RegexpStr:
		return false, NotYetImplemented("RegexpStr comparison")
	case NotRegexpStr:
		return false, NotYetImplemented("NotRegexpStr comparison")
	case JSONExtractOp:
		return false, NotYetImplemented("JSONExtractOp comparison")
	case JSONUnquoteExtractOp:
		return false, NotYetImplemented("JSONUnquoteExtractOp comparison")
	}

	return false, fmt.Errorf("unknown comparison operation: '%s'", node.Operator)
}
