package dataset_sql

// import (
// 	"fmt"

// 	"strings"
// )

// var ErrInvalidComparison = fmt.Errorf("Cannot compare")

// // String Comparison
// func (a StrVal) compare(op string, b ValExpr) (BoolVal, error) {
// 	result := false
// 	switch op {
// 	case EqualStr:
// 		result = a.String() == b.(StrVal).String()
// 	case LessThanStr:
// 		result = a.String() < b.(StrVal).String()
// 	case GreaterThanStr:
// 		result = a.String() > b.(StrVal).String()
// 	case LessEqualStr:
// 		result = a.String() <= b.(StrVal).String()
// 	case GreaterEqualStr:
// 		result = a.String() >= b.(StrVal).String()
// 	case NotEqualStr:
// 		result = a.String() != b.(StrVal).String()
// 	case NullSafeEqualStr:
// 		result = a.String() != b.(StrVal).String()
// 	case InStr:
// 		result = strings.Contains(a.String(), b.(StrVal).String())
// 	case NotInStr:
// 		result = !strings.Contains(a.String(), b.(StrVal).String())
// 	case LikeStr:
// 		// TODO
// 		return BoolVal(false), NotYetImplemented("comparing like strings")
// 	case NotLikeStr:
// 		// TODO
// 		return BoolVal(false), NotYetImplemented("comparing not like like strings")
// 	case RegexpStr:
// 		// TODO
// 		return BoolVal(false), NotYetImplemented("comparing regex strings")
// 	case NotRegexpStr:
// 		// TODO
// 		return BoolVal(false), NotYetImplemented("comparing not-regex strings")
// 	}

// 	return BoolVal(result), nil
// }

// // Numeric Comparison
// func (a NumVal) compare(op string, b ValExpr) (BoolVal, error) {
// 	ai := a.Num()
// 	bi := float64(0)
// 	if i, ok := b.(NumVal); ok {
// 		bi = i.Num()
// 	} else {
// 		return BoolVal(false), ErrInvalidComparison
// 	}

// 	switch op {
// 	case EqualStr:
// 		return BoolVal(ai == bi), nil
// 	case LessThanStr:
// 		return BoolVal(ai < bi), nil
// 	case GreaterThanStr:
// 		return BoolVal(ai > bi), nil
// 	case LessEqualStr:
// 		return BoolVal(ai <= bi), nil
// 	case GreaterEqualStr:
// 		return BoolVal(ai >= bi), nil
// 	case NotEqualStr:
// 		return BoolVal(ai != bi), nil
// 	// case NotSafeEqualStr:
// 	// 	return BoolVal(ai == bi), nil
// 	// case InStr, NotInStr, LikeStr, NotLikeStr, RegexpStr, NotRegexpStr:
// 	default:
// 		return BoolVal(false), ErrInvalidComparison
// 	}

// 	return BoolVal(false), ErrInvalidComparison
// }

// // Value Comparison
// func (a ValArg) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("comparing value arguments")
// }

// // Null Comparison
// func (a *NullVal) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("comparing null operations")
// }

// // Bool Comparison
// func (a BoolVal) compare(op string, b ValExpr) (BoolVal, error) {
// 	if _, ok := b.(BoolVal); !ok {
// 		return BoolVal(false), ErrInvalidComparison
// 	}

// 	switch op {
// 	case EqualStr:
// 		return BoolVal(a == b), nil
// 	case NotEqualStr:
// 		return BoolVal(a != b), nil
// 	// case NotSafeEqualStr:
// 	// 	return BoolVal(a == b), nil
// 	default:
// 		return BoolVal(false), ErrInvalidComparison
// 	}
// 	return BoolVal(false), ErrInvalidComparison
// }

// // Column Comparison should never happen, columns should be evaluated into concrete values
// func (a *ColName) compare(op string, b ValExpr) (BoolVal, error) {
// 	// switch a.Type {
// 	// case datatype.String.String():

// 	// case datatype.Integer.String():
// 	// case datatype.Float.String():
// 	// case datatype.Date.String():
// 	// default:
// 	// 	return BoolVal(false), ErrInvalidComparison
// 	// }
// 	return BoolVal(false), ErrInvalidComparison
// }

// // Tuple Comparison
// func (a ValTuple) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("tuple comparison operations")
// }

// // Subquery Comparison
// func (a *Subquery) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("subquery comparison operations")
// }

// // List Comparison
// func (a ListArg) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("list-argument comparison operations")
// }

// func (a *BinaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("binary-value comparison operations")
// }
// func (a *UnaryExpr) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("unary expression comparison operations")
// }
// func (a *IntervalExpr) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), NotYetImplemented("interval value comparison operations")
// }

// // functions need to be evaluated before comparison
// func (a *FuncExpr) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), ErrInvalidComparison
// }

// // case expressions need to be evaluated before comparison
// func (a *CaseExpr) compare(op string, b ValExpr) (BoolVal, error) {
// 	return BoolVal(false), ErrInvalidComparison
// }
