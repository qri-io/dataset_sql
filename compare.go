package sqlparser

import (
	"fmt"

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
		// TODO!!
		return BoolVal(false), ErrNotYetImplemented
	case NotLikeStr:
		// TODO!!
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
	return BoolVal(false), ErrNotYetImplemented
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
	return BoolVal(false), ErrNotYetImplemented
}

// Column Comparison
func (a *ColName) compare(op string, b ValExpr) (BoolVal, error) {
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
