package dataset_sql

import (
	"bytes"
	q "github.com/qri-io/dataset_sql/vt/proto/query"
	"testing"
)

type evalTestCase struct {
	exp Expr
	t   q.Type
	res []byte
	err error
}

func runEvalCases(t *testing.T, row [][]byte, cases []evalTestCase) {
	for i, c := range cases {
		typ, res, err := c.exp.Eval(row)
		if c.err != err {
			t.Errorf("case %d error mistmatch. expected: %s, got: %s", i, c.err, err)
			continue
		}
		if typ != c.t {
			t.Errorf("case %d type mismatch. expected: %d, got: %d", i, c.t, typ)
		}
		if !bytes.Equal(c.res, res) {
			t.Errorf("case %d res mismatch. expected: %s, got: %s", i, string(c.res), string(res))
		}
	}
}

func TestEvalAndExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
		{&AndExpr{Left: BoolVal(true), Right: BoolVal(false)}, QueryBoolType, falseB, nil},
		{&AndExpr{Left: BoolVal(true), Right: BoolVal(true)}, QueryBoolType, trueB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalOrExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
		{&OrExpr{Left: BoolVal(true), Right: BoolVal(false)}, QueryBoolType, trueB, nil},
		{&OrExpr{Left: BoolVal(false), Right: BoolVal(true)}, QueryBoolType, trueB, nil},
		{&OrExpr{Left: BoolVal(true), Right: BoolVal(true)}, QueryBoolType, trueB, nil},
		{&OrExpr{Left: BoolVal(false), Right: BoolVal(false)}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalNotExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
		{&NotExpr{Expr: BoolVal(false)}, QueryBoolType, trueB, nil},
		{&NotExpr{Expr: BoolVal(true)}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalParenExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ParenExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalComparisonExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
		{&ComparisonExpr{Operator: EqualStr, Left: BoolVal(true), Right: BoolVal(true)}, QueryBoolType, trueB, nil},
		{&ComparisonExpr{Operator: EqualStr, Left: BoolVal(true), Right: BoolVal(false)}, QueryBoolType, falseB, nil},
		{&ComparisonExpr{Operator: LikeStr, Left: &SQLVal{Type: StrVal, Val: []byte("apples")}, Right: &SQLVal{Type: StrVal, Val: []byte("apples")}}, QueryBoolType, trueB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalRangeCond(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&RangeCond{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalIsExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&IsExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalExistsExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ExistsExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalSQLVal(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&SQLVal{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalNullVal(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&NullVal{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalBoolVal(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&BoolVal{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalColName(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ColName{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalValTuple(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ValTuple{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalSubquery(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&Subquery{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalListArg(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ListArg{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalBinaryExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&BinaryExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalUnaryExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&UnaryExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalIntervalExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&IntervalExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalCollateExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&CollateExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalFuncExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&FuncExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalGroupConcatExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&GroupConcatExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalValuesFuncExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ValuesFuncExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalConvertExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ConvertExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalConvertUsingExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&ConvertUsingExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalMatchExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&MatchExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalCaseExpr(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&CaseExpr{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
func TestEvalWhere(t *testing.T) {
	row := [][]byte{}
	cases := []evalTestCase{
	//{&Where{}, QueryBoolType, falseB, nil},
	}
	runEvalCases(t, row, cases)
}
