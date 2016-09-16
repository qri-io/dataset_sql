// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sqlparser

import (
	"errors"
	"strings"

	"github.com/qri-io/sqlparser/deps/cistring"
	"github.com/qri-io/sqlparser/deps/sqltypes"
)

// Instructions for creating new types: If a type
// needs to satisfy an interface, declare that function
// along with that interface. This will help users
// identify the list of types to which they can assert
// those interfaces.
// If the member of a type has a string with a predefined
// list of values, declare those values as const following
// the type.
// For interfaces that define dummy functions to consolidate
// a set of types, define the function as iTypeName.
// This will help avoid name collisions.

// Parse parses the sql and returns a Statement, which
// is the AST representation of the query.
func Parse(sql string) (Statement, error) {
	tokenizer := NewStringTokenizer(sql)
	if yyParse(tokenizer) != 0 {
		return nil, errors.New(tokenizer.LastError)
	}
	return tokenizer.ParseTree, nil
}

// SQLNode defines the interface for all nodes
// generated by the parser.
type SQLNode interface {
	Format(buf *TrackedBuffer)
	// WalkSubtree calls visit on all underlying nodes
	// of the subtree, but not the current one. Walking
	// must be interrupted if visit returns an error.
	WalkSubtree(visit Visit) error
}

// Visit defines the signature of a function that
// can be used to visit all nodes of a parse tree.
type Visit func(node SQLNode) (kontinue bool, err error)

// Walk calls visit on every node.
// If visit returns true, the underlying nodes
// are also visited. If it returns an error, walking
// is interrupted, and the error is returned.
func Walk(visit Visit, nodes ...SQLNode) error {
	for _, node := range nodes {
		if node == nil {
			continue
		}
		kontinue, err := visit(node)
		if err != nil {
			return err
		}
		if kontinue {
			err = node.WalkSubtree(visit)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// String returns a string representation of an SQLNode.
func String(node SQLNode) string {
	buf := NewTrackedBuffer(nil)
	buf.Myprintf("%v", node)
	return buf.String()
}

// GenerateParsedQuery returns a ParsedQuery of the ast.
func GenerateParsedQuery(node SQLNode) *ParsedQuery {
	buf := NewTrackedBuffer(nil)
	buf.Myprintf("%v", node)
	return buf.ParsedQuery()
}

// Statement represents a statement.
type Statement interface {
	iStatement()
	SQLNode
}

func (*Union) iStatement()  {}
func (*Select) iStatement() {}
func (*Insert) iStatement() {}
func (*Update) iStatement() {}
func (*Delete) iStatement() {}
func (*Set) iStatement()    {}
func (*DDL) iStatement()    {}
func (*Other) iStatement()  {}

// SelectStatement any SELECT statement.
type SelectStatement interface {
	iSelectStatement()
	iStatement()
	iInsertRows()
	SQLNode
}

func (*Select) iSelectStatement() {}
func (*Union) iSelectStatement()  {}

// Select represents a SELECT statement.
type Select struct {
	Comments    Comments
	Distinct    string
	Hints       string
	SelectExprs SelectExprs
	From        TableExprs
	Where       *Where
	GroupBy     GroupBy
	Having      *Where
	OrderBy     OrderBy
	Limit       *Limit
	Lock        string
}

// Select.Distinct
const (
	DistinctStr      = "distinct "
	StraightJoinHint = "straight_join "
)

// Select.Lock
const (
	ForUpdateStr = " for update"
	ShareModeStr = " lock in share mode"
)

// Format formats the node.
func (node *Select) Format(buf *TrackedBuffer) {
	buf.Myprintf("select %v%s%s%v from %v%v%v%v%v%v%s",
		node.Comments, node.Distinct, node.Hints, node.SelectExprs,
		node.From, node.Where,
		node.GroupBy, node.Having, node.OrderBy,
		node.Limit, node.Lock)
}

// WalkSubtree walks the nodes of the subtree
func (node *Select) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Comments,
		node.SelectExprs,
		node.From,
		node.Where,
		node.GroupBy,
		node.Having,
		node.OrderBy,
		node.Limit,
	)
}

// AddWhere adds the boolean expression to the
// WHERE clause as an AND condition. If the expression
// is an OR clause, it parenthesizes it. Currently,
// the OR operator is the only one that's lower precedence
// than AND.
func (node *Select) AddWhere(expr BoolExpr) {
	if _, ok := expr.(*OrExpr); ok {
		expr = &ParenBoolExpr{Expr: expr}
	}
	if node.Where == nil {
		node.Where = &Where{
			Type: WhereStr,
			Expr: expr,
		}
		return
	}
	node.Where.Expr = &AndExpr{
		Left:  node.Where.Expr,
		Right: expr,
	}
	return
}

// AddHaving adds the boolean expression to the
// HAVING clause as an AND condition. If the expression
// is an OR clause, it parenthesizes it. Currently,
// the OR operator is the only one that's lower precedence
// than AND.
func (node *Select) AddHaving(expr BoolExpr) {
	if _, ok := expr.(*OrExpr); ok {
		expr = &ParenBoolExpr{Expr: expr}
	}
	if node.Having == nil {
		node.Having = &Where{
			Type: HavingStr,
			Expr: expr,
		}
		return
	}
	node.Having.Expr = &AndExpr{
		Left:  node.Having.Expr,
		Right: expr,
	}
	return
}

// Union represents a UNION statement.
type Union struct {
	Type        string
	Left, Right SelectStatement
}

// Union.Type
const (
	UnionStr         = "union"
	UnionAllStr      = "union all"
	UnionDistinctStr = "union distinct"
)

// Format formats the node.
func (node *Union) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s %v", node.Left, node.Type, node.Right)
}

// WalkSubtree walks the nodes of the subtree
func (node *Union) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.Right,
	)
}

// Insert represents an INSERT statement.
type Insert struct {
	Comments Comments
	Ignore   string
	Table    *TableName
	Columns  Columns
	Rows     InsertRows
	OnDup    OnDup
}

// Format formats the node.
func (node *Insert) Format(buf *TrackedBuffer) {
	buf.Myprintf("insert %v%sinto %v%v %v%v",
		node.Comments, node.Ignore,
		node.Table, node.Columns, node.Rows, node.OnDup)
}

// WalkSubtree walks the nodes of the subtree
func (node *Insert) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Comments,
		node.Table,
		node.Columns,
		node.Rows,
		node.OnDup,
	)
}

// InsertRows represents the rows for an INSERT statement.
type InsertRows interface {
	iInsertRows()
	SQLNode
}

func (*Select) iInsertRows() {}
func (*Union) iInsertRows()  {}
func (Values) iInsertRows()  {}

// Update represents an UPDATE statement.
type Update struct {
	Comments Comments
	Table    *TableName
	Exprs    UpdateExprs
	Where    *Where
	OrderBy  OrderBy
	Limit    *Limit
}

// Format formats the node.
func (node *Update) Format(buf *TrackedBuffer) {
	buf.Myprintf("update %v%v set %v%v%v%v",
		node.Comments, node.Table,
		node.Exprs, node.Where, node.OrderBy, node.Limit)
}

// WalkSubtree walks the nodes of the subtree
func (node *Update) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Comments,
		node.Table,
		node.Exprs,
		node.Where,
		node.OrderBy,
		node.Limit,
	)
}

// Delete represents a DELETE statement.
type Delete struct {
	Comments Comments
	Table    *TableName
	Where    *Where
	OrderBy  OrderBy
	Limit    *Limit
}

// Format formats the node.
func (node *Delete) Format(buf *TrackedBuffer) {
	buf.Myprintf("delete %vfrom %v%v%v%v",
		node.Comments,
		node.Table, node.Where, node.OrderBy, node.Limit)
}

// WalkSubtree walks the nodes of the subtree
func (node *Delete) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Comments,
		node.Table,
		node.Where,
		node.OrderBy,
		node.Limit,
	)
}

// Set represents a SET statement.
type Set struct {
	Comments Comments
	Exprs    UpdateExprs
}

// Format formats the node.
func (node *Set) Format(buf *TrackedBuffer) {
	buf.Myprintf("set %v%v", node.Comments, node.Exprs)
}

// WalkSubtree walks the nodes of the subtree
func (node *Set) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Comments,
		node.Exprs,
	)
}

// DDL represents a CREATE, ALTER, DROP or RENAME statement.
// Table is set for AlterStr, DropStr, RenameStr.
// NewName is set for AlterStr, CreateStr, RenameStr.
// ColDefs is set for CreateStr
type DDL struct {
	Action   string
	Table    *TableName
	NewName  *TableName
	ColDefs  ColDefs
	IfExists bool
}

// DDL strings.
const (
	CreateStr = "create"
	AlterStr  = "alter"
	DropStr   = "drop"
	RenameStr = "rename"
)

// Format formats the node.
func (node *DDL) Format(buf *TrackedBuffer) {
	switch node.Action {
	case CreateStr:
		if len(node.ColDefs) > 0 {
			buf.Myprintf("%s table %v %v", node.Action, node.NewName, node.ColDefs)
		} else {
			buf.Myprintf("%s table %v", node.Action, node.NewName)
		}
	case DropStr:
		exists := ""
		if node.IfExists {
			exists = " if exists"
		}
		buf.Myprintf("%s table%s %v", node.Action, exists, node.Table)
	case RenameStr:
		buf.Myprintf("%s table %v %v", node.Action, node.Table, node.NewName)
	default:
		buf.Myprintf("%s table %v", node.Action, node.Table)
	}
}

// WalkSubtree walks the nodes of the subtree
func (node *DDL) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Table,
		node.NewName,
		node.ColDefs,
	)
}

// Other represents a SHOW, DESCRIBE, or EXPLAIN statement.
// It should be used only as an indicator. It does not contain
// the full AST for the statement.
type Other struct{}

// Format formats the node.
func (node *Other) Format(buf *TrackedBuffer) {
	buf.WriteString("other")
}

// WalkSubtree walks the nodes of the subtree
func (node *Other) WalkSubtree(visit Visit) error {
	return nil
}

// Comments represents a list of comments.
type Comments [][]byte

// Format formats the node.
func (node Comments) Format(buf *TrackedBuffer) {
	for _, c := range node {
		buf.Myprintf("%s ", c)
	}
}

// WalkSubtree walks the nodes of the subtree
func (node Comments) WalkSubtree(visit Visit) error {
	return nil
}

// SelectExprs represents SELECT expressions.
type SelectExprs []SelectExpr

// Format formats the node.
func (node SelectExprs) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node SelectExprs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// SelectExpr represents a SELECT expression.
type SelectExpr interface {
	iSelectExpr()
	SQLNode
}

func (*StarExpr) iSelectExpr()    {}
func (*NonStarExpr) iSelectExpr() {}
func (Nextval) iSelectExpr()      {}

// StarExpr defines a '*' or 'table.*' expression.
type StarExpr struct {
	TableName TableIdent
}

// Format formats the node.
func (node *StarExpr) Format(buf *TrackedBuffer) {
	if node.TableName != "" {
		buf.Myprintf("%v.", node.TableName)
	}
	buf.Myprintf("*")
}

// WalkSubtree walks the nodes of the subtree
func (node *StarExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.TableName,
	)
}

// NonStarExpr defines a non-'*' select expr.
type NonStarExpr struct {
	Expr Expr
	As   ColIdent
}

// Format formats the node.
func (node *NonStarExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v", node.Expr)
	if node.As.Original() != "" {
		buf.Myprintf(" as %v", node.As)
	}
}

// WalkSubtree walks the nodes of the subtree
func (node *NonStarExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
		node.As,
	)
}

// Nextval defines the NEXT VALUE expression.
type Nextval struct{}

// Format formats the node.
func (node Nextval) Format(buf *TrackedBuffer) {
	buf.Myprintf("next value")
}

// WalkSubtree walks the nodes of the subtree
func (node Nextval) WalkSubtree(visit Visit) error {
	return nil
}

// DataType represents a known data type
// Length is used in fields that take a defined length
// WithTZ is only used in Data types that support a time zone
type DataType struct {
	Type   string
	Length int
	WithTZ bool
}

func (node *DataType) Format(buf *TrackedBuffer) {
	buf.Myprintf("%s", node.Type)
}

func (node *DataType) WalkSubtree(visit Visit) error {
	return nil
}

// ColConstr represents a column constraint in a create table statement
type ColConstr struct {
	Constraint string
	Params     string
}

const (
	ColConstrNotNullStr    = "not null"
	ColConstrNullStr       = "null"
	ColConstrDefaultStr    = "default"
	ColConstrUniqueStr     = "unique"
	ColConstrPrimaryKeyStr = "primary key"
)

func (node *ColConstr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%s", node.Constraint)
}

func (node *ColConstr) WalkSubtree(visit Visit) error {
	return nil
}

type ColConstrs []*ColConstr

func (node ColConstrs) Format(buf *TrackedBuffer) {
	if len(node) == 0 {
		return
	}

	prefix := ""
	for i := 0; i < len(node); i++ {
		buf.Myprintf("%s%v", prefix, node[i])
		prefix = " "
	}
}

func (node ColConstrs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// ColDef represents a column definition for a create table statement
type ColDef struct {
	ColName     ColIdent
	ColType     *DataType
	Constraints ColConstrs
}

func (node *ColDef) Format(buf *TrackedBuffer) {
	if len(node.Constraints) == 0 {
		buf.Myprintf("%v %v", node.ColName, node.ColType)
	} else {
		buf.Myprintf("%v %v %v", node.ColName, node.ColType, node.Constraints)
	}
}

func (node *ColDef) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(visit,
		node.ColName,
		node.ColType)
}

// ColDefs is a list of ColDef used in a create table statement
type ColDefs []*ColDef

func (node ColDefs) Format(buf *TrackedBuffer) {
	if len(node) == 0 {
		return
	}

	prefix := ""
	buf.Myprintf("(")
	for i := 0; i < len(node); i++ {
		buf.Myprintf("%s%v", prefix, node[i])
		prefix = ", "
	}
	buf.Myprintf(")")
}

func (node ColDefs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// Columns represents an insert column list.
type Columns []ColIdent

// Format formats the node.
func (node Columns) Format(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	prefix := "("
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
	buf.WriteString(")")
}

// WalkSubtree walks the nodes of the subtree
func (node Columns) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// TableExprs represents a list of table expressions.
type TableExprs []TableExpr

// Format formats the node.
func (node TableExprs) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node TableExprs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// TableExpr represents a table expression.
type TableExpr interface {
	iTableExpr()
	SQLNode
}

func (*AliasedTableExpr) iTableExpr() {}
func (*ParenTableExpr) iTableExpr()   {}
func (*JoinTableExpr) iTableExpr()    {}

// AliasedTableExpr represents a table expression
// coupled with an optional alias or index hint.
// If As is empty, no alias was used.
type AliasedTableExpr struct {
	Expr  SimpleTableExpr
	As    TableIdent
	Hints *IndexHints
}

// Format formats the node.
func (node *AliasedTableExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v", node.Expr)
	if node.As != "" {
		buf.Myprintf(" as %v", node.As)
	}
	if node.Hints != nil {
		// Hint node provides the space padding.
		buf.Myprintf("%v", node.Hints)
	}
}

// WalkSubtree walks the nodes of the subtree
func (node *AliasedTableExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
		node.As,
		node.Hints,
	)
}

// SimpleTableExpr represents a simple table expression.
type SimpleTableExpr interface {
	iSimpleTableExpr()
	SQLNode
}

func (*TableName) iSimpleTableExpr() {}
func (*Subquery) iSimpleTableExpr()  {}

// TableName represents a table  name.
// Qualifier, if specified, represents a database.
// It's generally not supported because vitess has its own
// rules about which database to send a query to.
// User is added to deal with Qri requests
type TableName struct {
	User, Qualifier, Name TableIdent
}

// Format formats the node.
func (node *TableName) Format(buf *TrackedBuffer) {
	// node can be nil for unqualified column names.
	if node == nil {
		return
	}
	if node.User != "" && node.Qualifier != "" {
		buf.Myprintf("%v.%v.", node.User, node.Qualifier)
	} else if node.Qualifier != "" {
		buf.Myprintf("%v.", node.Qualifier)
	}
	buf.Myprintf("%v", node.Name)
}

// WalkSubtree walks the nodes of the subtree
func (node *TableName) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Name,
		node.Qualifier,
	)
}

// IsEmpty returns true if TableName is nil or empty.
func (node *TableName) IsEmpty() bool {
	return node == nil || (node.Qualifier == "" && node.Name == "")
}

// ParenTableExpr represents a parenthesized list of TableExpr.
type ParenTableExpr struct {
	Exprs TableExprs
}

// Format formats the node.
func (node *ParenTableExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("(%v)", node.Exprs)
}

// WalkSubtree walks the nodes of the subtree
func (node *ParenTableExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Exprs,
	)
}

// JoinTableExpr represents a TableExpr that's a JOIN operation.
type JoinTableExpr struct {
	LeftExpr  TableExpr
	Join      string
	RightExpr TableExpr
	On        BoolExpr
}

// JoinTableExpr.Join
const (
	JoinStr             = "join"
	StraightJoinStr     = "straight_join"
	LeftJoinStr         = "left join"
	RightJoinStr        = "right join"
	NaturalJoinStr      = "natural join"
	NaturalLeftJoinStr  = "natural left join"
	NaturalRightJoinStr = "natural right join"
)

// Format formats the node.
func (node *JoinTableExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s %v", node.LeftExpr, node.Join, node.RightExpr)
	if node.On != nil {
		buf.Myprintf(" on %v", node.On)
	}
}

// WalkSubtree walks the nodes of the subtree
func (node *JoinTableExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.LeftExpr,
		node.RightExpr,
		node.On,
	)
}

// IndexHints represents a list of index hints.
type IndexHints struct {
	Type    string
	Indexes []ColIdent
}

// Index hints.
const (
	UseStr    = "use "
	IgnoreStr = "ignore "
	ForceStr  = "force "
)

// Format formats the node.
func (node *IndexHints) Format(buf *TrackedBuffer) {
	buf.Myprintf(" %sindex ", node.Type)
	prefix := "("
	for _, n := range node.Indexes {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
	buf.Myprintf(")")
}

// WalkSubtree walks the nodes of the subtree
func (node *IndexHints) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	for _, n := range node.Indexes {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// Where represents a WHERE or HAVING clause.
type Where struct {
	Type string
	Expr BoolExpr
}

// Where.Type
const (
	WhereStr  = "where"
	HavingStr = "having"
)

// NewWhere creates a WHERE or HAVING clause out
// of a BoolExpr. If the expression is nil, it returns nil.
func NewWhere(typ string, expr BoolExpr) *Where {
	if expr == nil {
		return nil
	}
	return &Where{Type: typ, Expr: expr}
}

// Format formats the node.
func (node *Where) Format(buf *TrackedBuffer) {
	if node == nil || node.Expr == nil {
		return
	}
	buf.Myprintf(" %s %v", node.Type, node.Expr)
}

// WalkSubtree walks the nodes of the subtree
func (node *Where) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// Expr represents an expression.
type Expr interface {
	iExpr()
	SQLNode
}

func (*AndExpr) iExpr()        {}
func (*OrExpr) iExpr()         {}
func (*NotExpr) iExpr()        {}
func (*ParenBoolExpr) iExpr()  {}
func (*ComparisonExpr) iExpr() {}
func (*RangeCond) iExpr()      {}
func (*IsExpr) iExpr()         {}
func (*ExistsExpr) iExpr()     {}
func (StrVal) iExpr()          {}
func (NumVal) iExpr()          {}
func (ValArg) iExpr()          {}
func (*NullVal) iExpr()        {}
func (BoolVal) iExpr()         {}
func (*ColName) iExpr()        {}
func (ValTuple) iExpr()        {}
func (*Subquery) iExpr()       {}
func (ListArg) iExpr()         {}
func (*BinaryExpr) iExpr()     {}
func (*UnaryExpr) iExpr()      {}
func (*IntervalExpr) iExpr()   {}
func (*FuncExpr) iExpr()       {}
func (*CaseExpr) iExpr()       {}

// BoolExpr represents a boolean expression.
type BoolExpr interface {
	iBoolExpr()
	Expr
}

func (BoolVal) iBoolExpr()         {}
func (*AndExpr) iBoolExpr()        {}
func (*OrExpr) iBoolExpr()         {}
func (*NotExpr) iBoolExpr()        {}
func (*ParenBoolExpr) iBoolExpr()  {}
func (*ComparisonExpr) iBoolExpr() {}
func (*RangeCond) iBoolExpr()      {}
func (*IsExpr) iBoolExpr()         {}
func (*ExistsExpr) iBoolExpr()     {}

// AndExpr represents an AND expression.
type AndExpr struct {
	Left, Right BoolExpr
}

// Format formats the node.
func (node *AndExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v and %v", node.Left, node.Right)
}

// WalkSubtree walks the nodes of the subtree
func (node *AndExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.Right,
	)
}

// OrExpr represents an OR expression.
type OrExpr struct {
	Left, Right BoolExpr
}

// Format formats the node.
func (node *OrExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v or %v", node.Left, node.Right)
}

// WalkSubtree walks the nodes of the subtree
func (node *OrExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.Right,
	)
}

// NotExpr represents a NOT expression.
type NotExpr struct {
	Expr BoolExpr
}

// Format formats the node.
func (node *NotExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("not %v", node.Expr)
}

// WalkSubtree walks the nodes of the subtree
func (node *NotExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// ParenBoolExpr represents a parenthesized boolean expression.
type ParenBoolExpr struct {
	Expr BoolExpr
}

// Format formats the node.
func (node *ParenBoolExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("(%v)", node.Expr)
}

// WalkSubtree walks the nodes of the subtree
func (node *ParenBoolExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// ComparisonExpr represents a two-value comparison expression.
type ComparisonExpr struct {
	Operator    string
	Left, Right ValExpr
}

// ComparisonExpr.Operator
const (
	EqualStr         = "="
	LessThanStr      = "<"
	GreaterThanStr   = ">"
	LessEqualStr     = "<="
	GreaterEqualStr  = ">="
	NotEqualStr      = "!="
	NullSafeEqualStr = "<=>"
	InStr            = "in"
	NotInStr         = "not in"
	LikeStr          = "like"
	NotLikeStr       = "not like"
	RegexpStr        = "regexp"
	NotRegexpStr     = "not regexp"
)

// Format formats the node.
func (node *ComparisonExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s %v", node.Left, node.Operator, node.Right)
}

// WalkSubtree walks the nodes of the subtree
func (node *ComparisonExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.Right,
	)
}

// RangeCond represents a BETWEEN or a NOT BETWEEN expression.
type RangeCond struct {
	Operator string
	Left     ValExpr
	From, To ValExpr
}

// RangeCond.Operator
const (
	BetweenStr    = "between"
	NotBetweenStr = "not between"
)

// Format formats the node.
func (node *RangeCond) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s %v and %v", node.Left, node.Operator, node.From, node.To)
}

// WalkSubtree walks the nodes of the subtree
func (node *RangeCond) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.From,
		node.To,
	)
}

// IsExpr represents an IS ... or an IS NOT ... expression.
type IsExpr struct {
	Operator string
	Expr     Expr
}

// IsExpr.Operator
const (
	IsNullStr     = "is null"
	IsNotNullStr  = "is not null"
	IsTrueStr     = "is true"
	IsNotTrueStr  = "is not true"
	IsFalseStr    = "is false"
	IsNotFalseStr = "is not false"
)

// Format formats the node.
func (node *IsExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s", node.Expr, node.Operator)
}

// WalkSubtree walks the nodes of the subtree
func (node *IsExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// ExistsExpr represents an EXISTS expression.
type ExistsExpr struct {
	Subquery *Subquery
}

// Format formats the node.
func (node *ExistsExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("exists %v", node.Subquery)
}

// WalkSubtree walks the nodes of the subtree
func (node *ExistsExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Subquery,
	)
}

// ValExpr represents a value expression.
type ValExpr interface {
	iValExpr()
	Expr
}

func (StrVal) iValExpr()        {}
func (NumVal) iValExpr()        {}
func (ValArg) iValExpr()        {}
func (*NullVal) iValExpr()      {}
func (BoolVal) iValExpr()       {}
func (*ColName) iValExpr()      {}
func (ValTuple) iValExpr()      {}
func (*Subquery) iValExpr()     {}
func (ListArg) iValExpr()       {}
func (*BinaryExpr) iValExpr()   {}
func (*UnaryExpr) iValExpr()    {}
func (*IntervalExpr) iValExpr() {}
func (*FuncExpr) iValExpr()     {}
func (*CaseExpr) iValExpr()     {}

// StrVal represents a string value.
type StrVal []byte

// Format formats the node.
func (node StrVal) Format(buf *TrackedBuffer) {
	s := sqltypes.MakeString([]byte(node))
	s.EncodeSQL(buf)
}

// WalkSubtree walks the nodes of the subtree
func (node StrVal) WalkSubtree(visit Visit) error {
	return nil
}

// NumVal represents a number.
type NumVal []byte

// Format formats the node.
func (node NumVal) Format(buf *TrackedBuffer) {
	buf.Myprintf("%s", []byte(node))
}

// WalkSubtree walks the nodes of the subtree
func (node NumVal) WalkSubtree(visit Visit) error {
	return nil
}

// ValArg represents a named bind var argument.
type ValArg []byte

// Format formats the node.
func (node ValArg) Format(buf *TrackedBuffer) {
	buf.WriteArg(string(node))
}

// WalkSubtree walks the nodes of the subtree
func (node ValArg) WalkSubtree(visit Visit) error {
	return nil
}

// NullVal represents a NULL value.
type NullVal struct{}

// Format formats the node.
func (node *NullVal) Format(buf *TrackedBuffer) {
	buf.Myprintf("null")
}

// WalkSubtree walks the nodes of the subtree
func (node *NullVal) WalkSubtree(visit Visit) error {
	return nil
}

// BoolVal is true or false.
type BoolVal bool

// Format formats the node.
func (node BoolVal) Format(buf *TrackedBuffer) {
	if node {
		buf.Myprintf("true")
	} else {
		buf.Myprintf("false")
	}
}

// WalkSubtree walks the nodes of the subtree
func (node BoolVal) WalkSubtree(visit Visit) error {
	return nil
}

// ColName represents a column name.
type ColName struct {
	// Metadata is not populated by the parser.
	// It's a placeholder for analyzers to store
	// additional data, typically info about which
	// table or column this node references.
	Metadata  interface{}
	Name      ColIdent
	Qualifier *TableName
}

// Format formats the node.
func (node *ColName) Format(buf *TrackedBuffer) {
	if !node.Qualifier.IsEmpty() {
		buf.Myprintf("%v.", node.Qualifier)
	}
	buf.Myprintf("%v", node.Name)
}

// WalkSubtree walks the nodes of the subtree
func (node *ColName) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Name,
		node.Qualifier,
	)
}

// ColTuple represents a list of column values.
// It can be ValTuple, Subquery, ListArg.
type ColTuple interface {
	iColTuple()
	ValExpr
}

func (ValTuple) iColTuple()  {}
func (*Subquery) iColTuple() {}
func (ListArg) iColTuple()   {}

// ValTuple represents a tuple of actual values.
type ValTuple ValExprs

// Format formats the node.
func (node ValTuple) Format(buf *TrackedBuffer) {
	buf.Myprintf("(%v)", ValExprs(node))
}

// WalkSubtree walks the nodes of the subtree
func (node ValTuple) WalkSubtree(visit Visit) error {
	return Walk(visit, ValExprs(node))
}

// ValExprs represents a list of value expressions.
// It's not a valid expression because it's not parenthesized.
type ValExprs []ValExpr

// Format formats the node.
func (node ValExprs) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node ValExprs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// Subquery represents a subquery.
type Subquery struct {
	Select SelectStatement
}

// Format formats the node.
func (node *Subquery) Format(buf *TrackedBuffer) {
	buf.Myprintf("(%v)", node.Select)
}

// WalkSubtree walks the nodes of the subtree
func (node *Subquery) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Select,
	)
}

// ListArg represents a named list argument.
type ListArg []byte

// Format formats the node.
func (node ListArg) Format(buf *TrackedBuffer) {
	buf.WriteArg(string(node))
}

// WalkSubtree walks the nodes of the subtree
func (node ListArg) WalkSubtree(visit Visit) error {
	return nil
}

// BinaryExpr represents a binary value expression.
type BinaryExpr struct {
	Operator    string
	Left, Right Expr
}

// BinaryExpr.Operator
const (
	BitAndStr     = "&"
	BitOrStr      = "|"
	BitXorStr     = "^"
	PlusStr       = "+"
	MinusStr      = "-"
	MultStr       = "*"
	DivStr        = "/"
	ModStr        = "%"
	ShiftLeftStr  = "<<"
	ShiftRightStr = ">>"
)

// Format formats the node.
func (node *BinaryExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s %v", node.Left, node.Operator, node.Right)
}

// WalkSubtree walks the nodes of the subtree
func (node *BinaryExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Left,
		node.Right,
	)
}

// UnaryExpr represents a unary value expression.
type UnaryExpr struct {
	Operator byte
	Expr     Expr
}

// UnaryExpr.Operator
const (
	UPlusStr  = '+'
	UMinusStr = '-'
	TildaStr  = '~'
)

// Format formats the node.
func (node *UnaryExpr) Format(buf *TrackedBuffer) {
	if _, unary := node.Expr.(*UnaryExpr); unary {
		buf.Myprintf("%c %v", node.Operator, node.Expr)
		return
	}
	buf.Myprintf("%c%v", node.Operator, node.Expr)
}

// WalkSubtree walks the nodes of the subtree
func (node *UnaryExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// IntervalExpr represents a date-time INTERVAL expression.
type IntervalExpr struct {
	Expr Expr
	Unit ColIdent
}

// Format formats the node.
func (node *IntervalExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("interval %v %v", node.Expr, node.Unit)
}

// WalkSubtree walks the nodes of the subtree
func (node *IntervalExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
		node.Unit,
	)
}

// FuncExpr represents a function call.
type FuncExpr struct {
	Name     string
	Distinct bool
	Exprs    SelectExprs
}

// Format formats the node.
func (node *FuncExpr) Format(buf *TrackedBuffer) {
	var distinct string
	if node.Distinct {
		distinct = "distinct "
	}
	// Function names should not be back-quoted even
	// if they match a reserved word. So, print the
	// name as is.
	buf.Myprintf("%s(%s%v)", node.Name, distinct, node.Exprs)
}

// WalkSubtree walks the nodes of the subtree
func (node *FuncExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Exprs,
	)
}

// Aggregates is a map of all aggregate functions.
var Aggregates = map[string]bool{
	"avg":          true,
	"bit_and":      true,
	"bit_or":       true,
	"bit_xor":      true,
	"count":        true,
	"group_concat": true,
	"max":          true,
	"min":          true,
	"std":          true,
	"stddev_pop":   true,
	"stddev_samp":  true,
	"stddev":       true,
	"sum":          true,
	"var_pop":      true,
	"var_samp":     true,
	"variance":     true,
}

// IsAggregate returns true if the function is an aggregate.
func (node *FuncExpr) IsAggregate() bool {
	return Aggregates[strings.ToLower(node.Name)]
}

// CaseExpr represents a CASE expression.
type CaseExpr struct {
	Expr  ValExpr
	Whens []*When
	Else  ValExpr
}

// Format formats the node.
func (node *CaseExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("case ")
	if node.Expr != nil {
		buf.Myprintf("%v ", node.Expr)
	}
	for _, when := range node.Whens {
		buf.Myprintf("%v ", when)
	}
	if node.Else != nil {
		buf.Myprintf("else %v ", node.Else)
	}
	buf.Myprintf("end")
}

// WalkSubtree walks the nodes of the subtree
func (node *CaseExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	if err := Walk(visit, node.Expr); err != nil {
		return err
	}
	for _, n := range node.Whens {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	if err := Walk(visit, node.Else); err != nil {
		return err
	}
	return nil
}

// When represents a WHEN sub-expression.
type When struct {
	Cond BoolExpr
	Val  ValExpr
}

// Format formats the node.
func (node *When) Format(buf *TrackedBuffer) {
	buf.Myprintf("when %v then %v", node.Cond, node.Val)
}

// WalkSubtree walks the nodes of the subtree
func (node *When) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Cond,
		node.Val,
	)
}

// GroupBy represents a GROUP BY clause.
type GroupBy []ValExpr

// Format formats the node.
func (node GroupBy) Format(buf *TrackedBuffer) {
	prefix := " group by "
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node GroupBy) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// OrderBy represents an ORDER By clause.
type OrderBy []*Order

// Format formats the node.
func (node OrderBy) Format(buf *TrackedBuffer) {
	prefix := " order by "
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node OrderBy) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// Order represents an ordering expression.
type Order struct {
	Expr      ValExpr
	Direction string
}

// Order.Direction
const (
	AscScr  = "asc"
	DescScr = "desc"
)

// Format formats the node.
func (node *Order) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v %s", node.Expr, node.Direction)
}

// WalkSubtree walks the nodes of the subtree
func (node *Order) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Expr,
	)
}

// Limit represents a LIMIT clause.
type Limit struct {
	Offset, Rowcount ValExpr
}

// Format formats the node.
func (node *Limit) Format(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	buf.Myprintf(" limit ")
	if node.Offset != nil {
		buf.Myprintf("%v, ", node.Offset)
	}
	buf.Myprintf("%v", node.Rowcount)
}

// WalkSubtree walks the nodes of the subtree
func (node *Limit) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Offset,
		node.Rowcount,
	)
}

// Values represents a VALUES clause.
type Values []RowTuple

// Format formats the node.
func (node Values) Format(buf *TrackedBuffer) {
	prefix := "values "
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node Values) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// RowTuple represents a row of values. It can be ValTuple, Subquery.
type RowTuple interface {
	iRowTuple()
	ValExpr
}

func (ValTuple) iRowTuple()  {}
func (*Subquery) iRowTuple() {}

// UpdateExprs represents a list of update expressions.
type UpdateExprs []*UpdateExpr

// Format formats the node.
func (node UpdateExprs) Format(buf *TrackedBuffer) {
	var prefix string
	for _, n := range node {
		buf.Myprintf("%s%v", prefix, n)
		prefix = ", "
	}
}

// WalkSubtree walks the nodes of the subtree
func (node UpdateExprs) WalkSubtree(visit Visit) error {
	for _, n := range node {
		if err := Walk(visit, n); err != nil {
			return err
		}
	}
	return nil
}

// UpdateExpr represents an update expression.
type UpdateExpr struct {
	Name ColIdent
	Expr ValExpr
}

// Format formats the node.
func (node *UpdateExpr) Format(buf *TrackedBuffer) {
	buf.Myprintf("%v = %v", node.Name, node.Expr)
}

// WalkSubtree walks the nodes of the subtree
func (node *UpdateExpr) WalkSubtree(visit Visit) error {
	if node == nil {
		return nil
	}
	return Walk(
		visit,
		node.Name,
		node.Expr,
	)
}

// OnDup represents an ON DUPLICATE KEY clause.
type OnDup UpdateExprs

// Format formats the node.
func (node OnDup) Format(buf *TrackedBuffer) {
	if node == nil {
		return
	}
	buf.Myprintf(" on duplicate key update %v", UpdateExprs(node))
}

// WalkSubtree walks the nodes of the subtree
func (node OnDup) WalkSubtree(visit Visit) error {
	return Walk(visit, UpdateExprs(node))
}

// ColIdent is a case insensitive SQL identifier. It will be escaped with
// backquotes if it matches a keyword.
type ColIdent cistring.CIString

// NewColIdent makes a new ColIdent.
func NewColIdent(str string) ColIdent {
	return ColIdent(cistring.New(str))
}

// Format formats the node.
func (node ColIdent) Format(buf *TrackedBuffer) {
	if _, ok := keywords[node.Lowered()]; ok {
		buf.Myprintf("`%s`", node.Original())
		return
	}
	buf.Myprintf("%s", node.Original())
}

// WalkSubtree walks the nodes of the subtree
func (node ColIdent) WalkSubtree(visit Visit) error {
	return nil
}

// Original returns the case-preserved column name.
func (node ColIdent) Original() string {
	return cistring.CIString(node).Original()
}

func (node ColIdent) String() string {
	return cistring.CIString(node).String()
}

// Lowered returns a lower-cased column name.
// This function should generally be used only for optimizing
// comparisons.
func (node ColIdent) Lowered() string {
	return cistring.CIString(node).Lowered()
}

// Equal performs a case-insensitive compare.
func (node ColIdent) Equal(in ColIdent) bool {
	return cistring.CIString(node).Equal(cistring.CIString(in))
}

// EqualString performs a case-insensitive compare with str.
func (node ColIdent) EqualString(str string) bool {
	return cistring.CIString(node).EqualString(str)
}

// TableIdent is a case sensitive SQL identifier. It will be escaped with
// backquotes if it matches a keyword.
type TableIdent string

// Format formats the node.
func (node TableIdent) Format(buf *TrackedBuffer) {
	name := string(node)
	if _, ok := keywords[strings.ToLower(name)]; ok {
		buf.Myprintf("`%s`", name)
		return
	}
	buf.Myprintf("%s", name)
}

// WalkSubtree walks the nodes of the subtree
func (node TableIdent) WalkSubtree(visit Visit) error {
	return nil
}
