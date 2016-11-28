package dataset_sql

import (
	"fmt"

	"github.com/qri-io/datatype"

	"strings"
)

var ErrInvalidComparison = fmt.Errorf("Cannot compare")

// String Comparison
func (a StrVal) compare(op string, b ValExpr) (BoolVal, error) {
	result := false
	switch op {
	case EqualStr:
		result = a.String() == b.(StrVal).String()
	case LessThanStr:
		result = a.String() < b.(StrVal).String()
	case GreaterThanStr:
		result = a.String() > b.(StrVal).String()
	case LessEqualStr:
		result = a.String() <= b.(StrVal).String()
	case GreaterEqualStr:
		result = a.String() >= b.(StrVal).String()
	case NotEqualStr:
		result = a.String() != b.(StrVal).String()
	case NullSafeEqualStr:
		result = a.String() != b.(StrVal).String()
	case InStr:
		result = strings.Contains(a.String(), b.(StrVal).String())
	case NotInStr:
		result = !strings.Contains(a.String(), b.(StrVal).String())
	case LikeStr:
		// TODO
		return BoolVal(false), ErrNotYetImplemented
	case NotLikeStr:
		// TODO
		return BoolVal(false), ErrNotYetImplemented
	case RegexpStr:
		// TODO
		return BoolVal(false), ErrNotYetImplemented
	case NotRegexpStr:
		// TODO
		return BoolVal(false), ErrNotYetImplemented
	}

	return BoolVal(result), nil
}

// Numeric Comparison
func (a NumVal) compare(op string, b ValExpr) (BoolVal, error) {
	ai := a.Int()
	bi := 0
	if i, ok := b.(NumVal); ok {
		bi = i.Int()
	} else {
		return BoolVal(false), ErrInvalidComparison
	}

	switch op {
	case EqualStr:
		return BoolVal(ai == bi), nil
	case LessThanStr:
		return BoolVal(ai < bi), nil
	case GreaterThanStr:
		return BoolVal(ai > bi), nil
	case LessEqualStr:
		return BoolVal(ai <= bi), nil
	case GreaterEqualStr:
		return BoolVal(ai >= bi), nil
	case NotEqualStr:
		return BoolVal(ai != bi), nil
	// case NotSafeEqualStr:
	// 	return BoolVal(ai == bi), nil
	// case InStr, NotInStr, LikeStr, NotLikeStr, RegexpStr, NotRegexpStr:
	default:
		return BoolVal(false), ErrInvalidComparison
	}

	return BoolVal(false), ErrInvalidComparison
}

// Value Comparison
func (a ValArg) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}

// Null Comparison
func (a *NullVal) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}

// Bool Comparison
func (a BoolVal) compare(op string, b ValExpr) (BoolVal, error) {
	if _, ok := b.(BoolVal); !ok {
		return BoolVal(false), ErrInvalidComparison
	}

	switch op {
	case EqualStr:
		return BoolVal(a == b), nil
	case NotEqualStr:
		return BoolVal(a != b), nil
	// case NotSafeEqualStr:
	// 	return BoolVal(a == b), nil
	default:
		return BoolVal(false), ErrInvalidComparison
	}
	return BoolVal(false), ErrInvalidComparison
}

// Column Comparison
func (a *ColName) compare(op string, b ValExpr) (BoolVal, error) {
	switch a.Type {
	case datatype.String.String():

	case datatype.Integer.String():
	case datatype.Float.String():
	case datatype.Date.String():
	default:
		return BoolVal(false), ErrInvalidComparison
	}
	return BoolVal(false), ErrNotYetImplemented
}

// Tuple Comparison
func (a ValTuple) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}

// Subquery Comparison
func (a *Subquery) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}

// List Comparison
func (a ListArg) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}

func (a *BinaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}
func (a *UnaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}
func (a *IntervalExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}
func (a *FuncExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}
func (a *CaseExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return BoolVal(false), ErrNotYetImplemented
}
