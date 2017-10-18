package dataset_sql

import (
	"bytes"
	"fmt"
	"github.com/qri-io/dataset_sql/sqltypes"
	"regexp"
	// pb "github.com/qri-io/dataset_sql/vt/proto/query"
)

var (
	reUs     = regexp.MustCompile("_")
	rePct    = regexp.MustCompile("%")
	rePctPct = regexp.MustCompile("%%")
)

func (node *ComparisonExpr) Compare() (bool, error) {
	_, left, err := node.Left.Eval()
	if err != nil {
		return false, err
	}
	_, right, err := node.Right.Eval()
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
		return CompareLike(l, r)
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

func CompareLike(str, expr sqltypes.Value) (bool, error) {
	// TODO - lowercasing here may possibly break user-supplied regexes
	// need to do a more sophisticated regex detect & replace :/
	exp := bytes.ToLower(expr.Bytes())
	exp = reUs.ReplaceAll(exp, []byte("."))
	exp = rePct.ReplaceAll(exp, []byte("x*"))
	exp = rePct.ReplaceAll(exp, []byte("x*"))

	expre, err := regexp.Compile(string(exp))
	if err != nil {
		return false, fmt.Errorf("error parsing like expression: %s", err.Error())
	}

	comp := bytes.ToLower(str.Bytes())
	return expre.Match(comp), nil
}
