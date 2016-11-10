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
		result = a == b.(StrVal)
	case LessThanStr:
		result = a < b.(StrVal)
	case GreaterThanStr:
		result = a > b.(StrVal)
	case LessEqualStr:
		result = a <= b.(StrVal)
	case GreaterEqualStr:
		result = a >= b.(StrVal)
	case NotEqualStr:
		result = a != b.(StrVal)
	case NullSafeEqualStr:
		result = a != b.(StrVal)
	case InStr:
		result = strings.Contains(a, b.(StrVal))
	case NotInStr:
		result = !strings.Contains(a, b.(StrVal))
	case LikeStr:
		// TODO!!
		return nil, ErrNotYetImplemented
	case NotLikeStr:
		// TODO!!
		return nil, ErrNotYetImplemented
	case RegexpStr:
		// TODO
		return nil, ErrNotYetImplemented
	case NotRegexpStr:
		// TODO
		return nil, ErrNotYetImplemented
	}

	return BoolVal(result), nil
}

// Numeric Comparison
func (a NumVal) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Value Comparison
func (a ValArg) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Null Comparison
func (a *NullVal) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Bool Comparison
func (a BoolVal) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Column Comparison
func (a *ColName) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Tuple Comparison
func (a ValTuple) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// Subquery Comparison
func (a *Subquery) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

// List Comparison
func (a ListArg) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}

func (a *BinaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}
func (a *UnaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}
func (a *IntervalExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}
func (a *FuncExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}
func (a *CaseExpr) compare(op string, b ValExpr) (BoolVal, error) {
	return nil, ErrNotYetImplemented
}
