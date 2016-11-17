// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

%{
package dataset_sql

func setParseTree(yylex interface{}, stmt Statement) {
  yylex.(*Tokenizer).ParseTree = stmt
}

func setAllowComments(yylex interface{}, allow bool) {
  yylex.(*Tokenizer).AllowComments = allow
}

func incNesting(yylex interface{}) bool {
  yylex.(*Tokenizer).nesting++
  if yylex.(*Tokenizer).nesting == 200 {
    return true
  }
  return false
}

func decNesting(yylex interface{}) {
  yylex.(*Tokenizer).nesting--
}

func forceEOF(yylex interface{}) {
  yylex.(*Tokenizer).ForceEOF = true
}

%}

%union {
  empty       struct{}
  statement   Statement
  selStmt     SelectStatement
  byt         byte
  bytes       []byte
  bytes2      [][]byte
  str         string
  selectExprs SelectExprs
  selectExpr  SelectExpr
  columns     Columns
  colName     *ColName
  tableExprs  TableExprs
  tableExpr   TableExpr
  tableName   TableName
  indexHints  *IndexHints
  expr        Expr
  boolExpr    BoolExpr
  valExpr     ValExpr
  castValExpr *CastValExpr
  colTuple    ColTuple
  valExprs    ValExprs
  values      Values
  rowTuple    RowTuple
  subquery    *Subquery
  caseExpr    *CaseExpr
  whens       []*When
  when        *When
  orderBy     OrderBy
  order       *Order
  limitOffset *LimitOffset
  insRows     InsertRows
  updateExprs UpdateExprs
  updateExpr  *UpdateExpr
  colIdent    ColIdent
  colIdents   []ColIdent
  tableIdent  TableIdent

  dataType    *DataType
  colDef      *ColDef
  colDefs     ColDefs
  colConstr   *ColConstr
  colConstrs  ColConstrs
}

%token LEX_ERROR
%left <empty> UNION
%token <empty> SELECT INSERT UPDATE DELETE FROM CAST WHERE GROUP HAVING ORDER BY LIMIT FOR
%token <empty> OFFSET
%token <empty> ALL DISTINCT AS EXISTS ASC DESC INTO DUPLICATE KEY DEFAULT SET LOCK
%token <empty> VALUES LAST_INSERT_ID
%token <empty> NEXT VALUE
%left <empty> JOIN STRAIGHT_JOIN LEFT RIGHT INNER OUTER CROSS NATURAL USE FORCE
%left <empty> ON
%token <empty> '(' ',' ')'
%token <bytes> ID STRING NUMBER VALUE_ARG LIST_ARG COMMENT
%token <empty> NULL TRUE FALSE

// Precedence dictated by mysql. But the vitess grammar is simplified.
// Some of these operators don't conflict in our situation. Nevertheless,
// it's better to have these listed in the correct order. Also, we don't
// support all operators yet.
%left <empty> OR
%left <empty> AND
%right <empty> NOT
%left <empty> BETWEEN CASE WHEN THEN ELSE
%left <empty> '=' '<' '>' LE GE NE NULL_SAFE_EQUAL IS LIKE REGEXP IN
%left <empty> '|'
%left <empty> '&'
%left <empty> SHIFT_LEFT SHIFT_RIGHT
%left <empty> '+' '-'
%left <empty> '*' '/' '%'
%left <empty> '^'
%right <empty> '~' UNARY
%right <empty> INTERVAL
%nonassoc <empty> '.' RIGHT_ARROW
%left <empty> END

// DDL Tokens
%token <empty> CREATE ALTER DROP RENAME ANALYZE
%token <empty> TABLE INDEX VIEW TO IGNORE IF UNIQUE USING
%token <empty> SHOW DESCRIBE EXPLAIN
%token <empty> PRIMARY

// MySQL reserved words that are unused by this grammar will map to this token.
%token <empty> UNUSED

// Accepted Data types will map to these tokens
%token <bytes> SMALLINT INTEGER BIGINT FLOAT DECIMAL NUMERIC REAL DOUBLE SMALLSERIAL SERIAL BIGSERIAL MONEY 
%token <bytes> CHAR_VARYING CHAR TEXT 
%token <bytes> BYTEA 
%token <bytes> TIMESTAMP DATE TIME // INTERVAL
%token <bytes> BOOLEAN
%token <bytes> ENUM
%token <bytes> POINT LINE LSEG BOX PATH POLYGON CIRCLE
%token <bytes> CIDR INET MACADDR
%token <bytes> BIT BIT_VARYING
%token <bytes> UUID
%token <bytes> XML JSON JSONB

%type <statement> command
%type <selStmt> select_statement
%type <statement> insert_statement update_statement delete_statement set_statement
%type <statement> create_statement alter_statement rename_statement drop_statement
%type <statement> analyze_statement other_statement
%type <bytes2> comment_opt comment_list
%type <str> union_op
%type <str> distinct_opt straight_join_opt
%type <selectExprs> select_expression_list
%type <selectExpr> select_expression
%type <expr> expression
%type <tableExprs> table_references
%type <tableExpr> table_reference table_factor join_table
%type <str> inner_join outer_join natural_join
%type <tableName> table_name
%type <indexHints> index_hint_list
%type <colIdents> index_list
%type <boolExpr> where_expression_opt
%type <boolExpr> boolean_expression condition
%type <str> compare
%type <insRows> row_list
%type <valExpr> value value_expression
%type <str> is_suffix
%type <colTuple> col_tuple
%type <valExprs> value_expression_list
%type <values> tuple_list
%type <rowTuple> row_tuple
%type <subquery> subquery
%type <colName> column_name
%type <caseExpr> case_expression
%type <whens> when_expression_list
%type <when> when_expression
%type <valExpr> value_expression_opt else_expression_opt
%type <valExprs> group_by_opt
%type <boolExpr> having_opt
%type <orderBy> order_by_opt order_list
%type <order> order
%type <str> asc_desc_opt
%type <limitOffset> limitOffset_opt
%type <str> lock_opt
%type <columns> column_list_opt column_list
%type <updateExprs> on_dup_opt
%type <updateExprs> update_list
%type <updateExpr> update_expression
%type <empty> for_from
%type <str> ignore_opt
%type <byt> exists_opt
%type <empty> not_exists_opt non_rename_operation to_opt constraint_opt using_opt
%type <colIdent> sql_id as_ci_opt
%type <tableIdent> table_id as_opt_id
%type <empty> as_opt
%type <empty> force_eof

%type <colDef> column_definition
%type <colDefs> column_definition_list_opt column_definition_list
%type <colConstr> column_constraint
%type <colConstrs> column_constraint_list
%type <dataType> data_type

%start any_command

%%

any_command:
  command
  {
    setParseTree(yylex, $1)
  }

command:
  select_statement
  {
    $$ = $1
  }
| insert_statement
| update_statement
| delete_statement
| set_statement
| create_statement
| alter_statement
| rename_statement
| drop_statement
| analyze_statement
| other_statement

select_statement:
  SELECT comment_opt distinct_opt straight_join_opt select_expression_list FROM table_references where_expression_opt group_by_opt having_opt order_by_opt limitOffset_opt lock_opt
  {
    $$ = &Select{Comments: Comments($2), Distinct: $3, Hints: $4, SelectExprs: $5, From: $7, Where: NewWhere(WhereStr, $8), GroupBy: GroupBy($9), Having: NewWhere(HavingStr, $10), OrderBy: $11, LimitOffset: $12, Lock: $13}
  }
| SELECT comment_opt NEXT sql_id for_from table_name
  {
    if $4.Lowered() != "value" {
      yylex.Error("expecting value after next")
      return 1
    }
    $$ = &Select{Comments: Comments($2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: $6}}}
  }
| select_statement union_op select_statement %prec UNION
  {
    $$ = &Union{Type: $2, Left: $1, Right: $3}
  }

insert_statement:
  INSERT comment_opt ignore_opt INTO table_name column_list_opt row_list on_dup_opt
  {
    $$ = &Insert{Comments: Comments($2), Ignore: $3, Table: $5, Columns: $6, Rows: $7, OnDup: OnDup($8)}
  }
| INSERT comment_opt ignore_opt INTO table_name SET update_list on_dup_opt
  {
    cols := make(Columns, 0, len($7))
    vals := make(ValTuple, 0, len($7))
    for _, updateList := range $7 {
      cols = append(cols, updateList.Name)
      vals = append(vals, updateList.Expr)
    }
    $$ = &Insert{Comments: Comments($2), Ignore: $3, Table: $5, Columns: cols, Rows: Values{vals}, OnDup: OnDup($8)}
  }

update_statement:
  UPDATE comment_opt table_name SET update_list where_expression_opt order_by_opt limitOffset_opt
  {
    $$ = &Update{Comments: Comments($2), Table: $3, Exprs: $5, Where: NewWhere(WhereStr, $6), OrderBy: $7, LimitOffset: $8}
  }

delete_statement:
  DELETE comment_opt FROM table_name where_expression_opt order_by_opt limitOffset_opt
  {
    $$ = &Delete{Comments: Comments($2), Table: $4, Where: NewWhere(WhereStr, $5), OrderBy: $6, LimitOffset: $7}
  }

set_statement:
  SET comment_opt update_list
  {
    $$ = &Set{Comments: Comments($2), Exprs: $3}
  }

create_statement:
  CREATE TABLE not_exists_opt table_name column_definition_list_opt force_eof
  {
    $$ = &DDL{Action: CreateStr, NewName: $4, ColDefs: $5 }
  }
| CREATE constraint_opt INDEX ID using_opt ON table_name force_eof
  {
    // Change this to an alter statement
    $$ = &DDL{Action: AlterStr, Table: $7, NewName: $7}
  }
| CREATE VIEW sql_id force_eof
  {
    $$ = &DDL{Action: CreateStr, NewName: TableName{TableIdent($3.Lowered())} } 
  }

alter_statement:
  ALTER ignore_opt TABLE table_name non_rename_operation force_eof
  {
    $$ = &DDL{Action: AlterStr, Table: $4, NewName: $4}
  }
| ALTER ignore_opt TABLE table_name RENAME to_opt table_name
  {
    // Change this to a rename statement
    $$ = &DDL{Action: RenameStr, Table: $4, NewName: $7}
  }
| ALTER VIEW sql_id force_eof
  {
    $$ = &DDL{Action: AlterStr, Table: TableName{ TableIdent($3.Lowered()) }, NewName: TableName{TableIdent($3.Lowered()) }}
  }

rename_statement:
  RENAME TABLE table_name TO table_name
  {
    $$ = &DDL{Action: RenameStr, Table: $3, NewName: $5}
  }

drop_statement:
  DROP TABLE exists_opt table_name
  {
    var exists bool
    if $3 != 0 {
      exists = true
    }
    $$ = &DDL{Action: DropStr, Table: $4, IfExists: exists}
  }
| DROP INDEX ID ON table_name
  {
    // Change this to an alter statement
    $$ = &DDL{Action: AlterStr, Table: $5, NewName: $5}
  }
| DROP VIEW exists_opt sql_id force_eof
  {
    var exists bool
        if $3 != 0 {
          exists = true
        }
    $$ = &DDL{Action: DropStr, Table: TableName{TableIdent($4.Lowered()) }, IfExists: exists}
  }

analyze_statement:
  ANALYZE TABLE table_name
  {
    $$ = &DDL{Action: AlterStr, Table: $3, NewName: $3}
  }

other_statement:
  SHOW force_eof
  {
    $$ = &Other{}
  }
| DESCRIBE force_eof
  {
    $$ = &Other{}
  }
| EXPLAIN force_eof
  {
    $$ = &Other{}
  }

comment_opt:
  {
    setAllowComments(yylex, true)
  }
  comment_list
  {
    $$ = $2
    setAllowComments(yylex, false)
  }

comment_list:
  {
    $$ = nil
  }
| comment_list COMMENT
  {
    $$ = append($1, $2)
  }

union_op:
  UNION
  {
    $$ = UnionStr
  }
| UNION ALL
  {
    $$ = UnionAllStr
  }
| UNION DISTINCT
  {
    $$ = UnionDistinctStr
  }

distinct_opt:
  {
    $$ = ""
  }
| DISTINCT
  {
    $$ = DistinctStr
  }

straight_join_opt:
  {
    $$ = ""
  }
| STRAIGHT_JOIN
  {
    $$ = StraightJoinHint
  }

select_expression_list:
  select_expression
  {
    $$ = SelectExprs{$1}
  }
| select_expression_list ',' select_expression
  {
    $$ = append($$, $3)
  }

select_expression:
  '*'
  {
    $$ = &StarExpr{}
  }
| expression as_ci_opt
  {
    $$ = &NonStarExpr{Expr: $1, As: $2}
  }
| table_id RIGHT_ARROW '*'
  {
    $$ = &StarExpr{TableName: $1}
  }

expression:
  boolean_expression
  {
    $$ = $1
  }
| value_expression
  {
    $$ = $1
  }

as_ci_opt:
  {
    $$ = ColIdent{}
  }
| sql_id
  {
    $$ = $1
  }
| AS sql_id
  {
    $$ = $2
  }

table_references:
  table_reference
  {
    $$ = TableExprs{$1}
  }
| table_references ',' table_reference
  {
    $$ = append($$, $3)
  }

table_reference:
  table_factor
| join_table

table_factor:
  table_name as_opt_id index_hint_list
  {
    $$ = &AliasedTableExpr{Expr:$1, As: $2, Hints: $3}
  }
| subquery as_opt table_id
  {
    $$ = &AliasedTableExpr{Expr:$1, As: $3}
  }
| openb table_references closeb
  {
    $$ = &ParenTableExpr{Exprs: $2}
  }

// There is a grammar conflict here:
// 1: INSERT INTO a SELECT * FROM b JOIN c ON b.i = c.i
// 2: INSERT INTO a SELECT * FROM b JOIN c ON DUPLICATE KEY UPDATE a.i = 1
// When yacc encounters the ON clause, it cannot determine which way to
// resolve. The %prec override below makes the parser choose the
// first construct, which automatically makes the second construct a
// syntax error. This is the same behavior as MySQL.
join_table:
  table_reference inner_join table_factor %prec JOIN
  {
    $$ = &JoinTableExpr{LeftExpr: $1, Join: $2, RightExpr: $3}
  }
| table_reference inner_join table_factor ON boolean_expression
  {
    $$ = &JoinTableExpr{LeftExpr: $1, Join: $2, RightExpr: $3, On: $5}
  }
| table_reference outer_join table_reference ON boolean_expression
  {
    $$ = &JoinTableExpr{LeftExpr: $1, Join: $2, RightExpr: $3, On: $5}
  }
| table_reference natural_join table_factor
  {
    $$ = &JoinTableExpr{LeftExpr: $1, Join: $2, RightExpr: $3}
  }

as_opt:
  { $$ = struct{}{} }
| AS
  { $$ = struct{}{} }

as_opt_id:
  {
    $$ = ""
  }
| table_id
  {
    $$ = $1
  }
| AS table_id
  {
    $$ = $2
  }

inner_join:
  JOIN
  {
    $$ = JoinStr
  }
| INNER JOIN
  {
    $$ = JoinStr
  }
| CROSS JOIN
  {
    $$ = JoinStr
  }
| STRAIGHT_JOIN
  {
    $$ = StraightJoinStr
  }

outer_join:
  LEFT JOIN
  {
    $$ = LeftJoinStr
  }
| LEFT OUTER JOIN
  {
    $$ = LeftJoinStr
  }
| RIGHT JOIN
  {
    $$ = RightJoinStr
  }
| RIGHT OUTER JOIN
  {
    $$ = RightJoinStr
  }

natural_join:
 NATURAL JOIN
  {
    $$ = NaturalJoinStr
  }
| NATURAL outer_join
  {
    if $2 == LeftJoinStr {
      $$ = NaturalLeftJoinStr
    } else {
      $$ = NaturalRightJoinStr
    }
  }

table_name:
  table_id
  {
    $$ = TableName{$1}
  }
| table_name '.' table_id
  {
    $$ = append($1, $3)
  }

index_hint_list:
  {
    $$ = nil
  }
| USE INDEX openb index_list closeb
  {
    $$ = &IndexHints{Type: UseStr, Indexes: $4}
  }
| IGNORE INDEX openb index_list closeb
  {
    $$ = &IndexHints{Type: IgnoreStr, Indexes: $4}
  }
| FORCE INDEX openb index_list closeb
  {
    $$ = &IndexHints{Type: ForceStr, Indexes: $4}
  }

index_list:
  sql_id
  {
    $$ = []ColIdent{$1}
  }
| index_list ',' sql_id
  {
    $$ = append($1, $3)
  }

where_expression_opt:
  {
    $$ = nil
  }
| WHERE boolean_expression
  {
    $$ = $2
  }

boolean_expression:
  condition
| boolean_expression AND boolean_expression
  {
    $$ = &AndExpr{Left: $1, Right: $3}
  }
| boolean_expression OR boolean_expression
  {
    $$ = &OrExpr{Left: $1, Right: $3}
  }
| NOT boolean_expression
  {
    $$ = &NotExpr{Expr: $2}
  }
| openb boolean_expression closeb
  {
    $$ = &ParenBoolExpr{Expr: $2}
  }
| boolean_expression IS is_suffix
  {
    $$ = &IsExpr{Operator: $3, Expr: $1}
  }

condition:
  TRUE
  {
    $$ = BoolVal(true)
  }
| FALSE
  {
    $$ = BoolVal(false)
  }
| value_expression compare value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: $2, Right: $3}
  }
| value_expression IN col_tuple
  {
    $$ = &ComparisonExpr{Left: $1, Operator: InStr, Right: $3}
  }
| value_expression NOT IN col_tuple
  {
    $$ = &ComparisonExpr{Left: $1, Operator: NotInStr, Right: $4}
  }
| value_expression LIKE value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: LikeStr, Right: $3}
  }
| value_expression NOT LIKE value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: NotLikeStr, Right: $4}
  }
| value_expression REGEXP value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: RegexpStr, Right: $3}
  }
| value_expression NOT REGEXP value_expression
  {
    $$ = &ComparisonExpr{Left: $1, Operator: NotRegexpStr, Right: $4}
  }
| value_expression BETWEEN value_expression AND value_expression
  {
    $$ = &RangeCond{Left: $1, Operator: BetweenStr, From: $3, To: $5}
  }
| value_expression NOT BETWEEN value_expression AND value_expression
  {
    $$ = &RangeCond{Left: $1, Operator: NotBetweenStr, From: $4, To: $6}
  }
| value_expression IS is_suffix
  {
    $$ = &IsExpr{Operator: $3, Expr: $1}
  }
| EXISTS subquery
  {
    $$ = &ExistsExpr{Subquery: $2}
  }

is_suffix:
  NULL
  {
    $$ = IsNullStr
  }
| NOT NULL
  {
    $$ = IsNotNullStr
  }
| TRUE
  {
    $$ = IsTrueStr
  }
| NOT TRUE
  {
    $$ = IsNotTrueStr
  }
| FALSE
  {
    $$ = IsFalseStr
  }
| NOT FALSE
  {
    $$ = IsNotFalseStr
  }

compare:
  '='
  {
    $$ = EqualStr
  }
| '<'
  {
    $$ = LessThanStr
  }
| '>'
  {
    $$ = GreaterThanStr
  }
| LE
  {
    $$ = LessEqualStr
  }
| GE
  {
    $$ = GreaterEqualStr
  }
| NE
  {
    $$ = NotEqualStr
  }
| NULL_SAFE_EQUAL
  {
    $$ = NullSafeEqualStr
  }

col_tuple:
  openb value_expression_list closeb
  {
    $$ = ValTuple($2)
  }
| subquery
  {
    $$ = $1
  }
| LIST_ARG
  {
    $$ = ListArg($1)
  }

subquery:
  openb select_statement closeb
  {
    $$ = &Subquery{$2}
  }

value_expression_list:
  value_expression
  {
    $$ = ValExprs{$1}
  }
| value_expression_list ',' value_expression
  {
    $$ = append($1, $3)
  }

value_expression:
  value
  {
    $$ = $1
  }
| column_name
  {
    $$ = $1
  }
//| table_name RIGHT_ARROW sql_id
//  {
//    $$ = &ColName{Qualifier: $1, Name: $3}
//  }
| row_tuple
  {
    $$ = $1
  }
| CAST openb value_expression AS data_type closeb
  {
    $$ = &CastValExpr{ Val : $3, Type : $5}
  }
| value_expression '&' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: BitAndStr, Right: $3}
  }
| value_expression '|' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: BitOrStr, Right: $3}
  }
| value_expression '^' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: BitXorStr, Right: $3}
  }
| value_expression '+' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: PlusStr, Right: $3}
  }
| value_expression '-' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: MinusStr, Right: $3}
  }
| value_expression '*' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: MultStr, Right: $3}
  }
| value_expression '/' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: DivStr, Right: $3}
  }
| value_expression '%' value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: ModStr, Right: $3}
  }
| value_expression SHIFT_LEFT value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: ShiftLeftStr, Right: $3}
  }
| value_expression SHIFT_RIGHT value_expression
  {
    $$ = &BinaryExpr{Left: $1, Operator: ShiftRightStr, Right: $3}
  }
| '+'  value_expression %prec UNARY
  {
    if num, ok := $2.(NumVal); ok {
      $$ = num
    } else {
      $$ = &UnaryExpr{Operator: UPlusStr, Expr: $2}
    }
  }
| '-'  value_expression %prec UNARY
  {
    if num, ok := $2.(NumVal); ok {
      // Handle double negative
      if num[0] == '-' {
        $$ = num[1:]
      } else {
        $$ = append(NumVal("-"), num...)
      }
    } else {
      $$ = &UnaryExpr{Operator: UMinusStr, Expr: $2}
    }
  }
| '~'  value_expression
  {
    $$ = &UnaryExpr{Operator: TildaStr, Expr: $2}
  }
| INTERVAL value_expression sql_id
  {
    // This rule prevents the usage of INTERVAL
    // as a function. If support is needed for that,
    // we'll need to revisit this. The solution
    // will be non-trivial because of grammar conflicts.
    $$ = &IntervalExpr{Expr: $2, Unit: $3}
  }
| table_id openb closeb
  {
    $$ = &FuncExpr{Name: string($1)}
  }
| table_id openb select_expression_list closeb
  {
    $$ = &FuncExpr{Name: string($1), Exprs: $3}
  }
| table_id openb DISTINCT select_expression_list closeb
  {
    $$ = &FuncExpr{Name: string($1), Distinct: true, Exprs: $4}
  }
| IF openb select_expression_list closeb
  {
    $$ = &FuncExpr{Name: "if", Exprs: $3}
  }
| case_expression
  {
    $$ = $1
  }

case_expression:
  CASE value_expression_opt when_expression_list else_expression_opt END
  {
    $$ = &CaseExpr{Expr: $2, Whens: $3, Else: $4}
  }

value_expression_opt:
  {
    $$ = nil
  }
| value_expression
  {
    $$ = $1
  }

when_expression_list:
  when_expression
  {
    $$ = []*When{$1}
  }
| when_expression_list when_expression
  {
    $$ = append($1, $2)
  }

when_expression:
  WHEN boolean_expression THEN value_expression
  {
    $$ = &When{Cond: $2, Val: $4}
  }

else_expression_opt:
  {
    $$ = nil
  }
| ELSE value_expression
  {
    $$ = $2
  }

column_name:
  sql_id
  {
    $$ = &ColName{Name: $1}
  }
| table_id RIGHT_ARROW sql_id
  {
    $$ = &ColName{Qualifier: TableName{$1}, Name: $3 }
  }
| table_name RIGHT_ARROW sql_id
  {
    $$ = &ColName{Qualifier: $1, Name: $3}
  }

value:
  STRING
  {
    $$ = StrVal($1)
  }
| NUMBER
  {
    $$ = NumVal($1)
  }
| VALUE_ARG
  {
    $$ = ValArg($1)
  }
| NULL
  {
    $$ = &NullVal{}
  }
| TRUE
  {
    $$ = BoolVal(true)
  }
| FALSE
  {
    $$ = BoolVal(false)
  }

group_by_opt:
  {
    $$ = nil
  }
| GROUP BY value_expression_list
  {
    $$ = $3
  }

having_opt:
  {
    $$ = nil
  }
| HAVING boolean_expression
  {
    $$ = $2
  }

order_by_opt:
  {
    $$ = nil
  }
| ORDER BY order_list
  {
    $$ = $3
  }

order_list:
  order
  {
    $$ = OrderBy{$1}
  }
| order_list ',' order
  {
    $$ = append($1, $3)
  }

order:
  value_expression asc_desc_opt
  {
    $$ = &Order{Expr: $1, Direction: $2}
  }

asc_desc_opt:
  {
    $$ = AscScr
  }
| ASC
  {
    $$ = AscScr
  }
| DESC
  {
    $$ = DescScr
  }

limitOffset_opt:
  {
    $$ = nil
  }
| LIMIT value_expression
  {
    $$ = &LimitOffset{Rowcount: $2}
  }
| LIMIT value_expression ',' value_expression
  {
    $$ = &LimitOffset{Offset: $2, Rowcount: $4}
  }
| LIMIT value_expression OFFSET value_expression
  {
    $$ = &LimitOffset{ Rowcount: $2, Offset: $4 }
  }

lock_opt:
  {
    $$ = ""
  }
| FOR UPDATE
  {
    $$ = ForUpdateStr
  }
| LOCK IN sql_id sql_id
  {
    if $3.Lowered() != "share" {
      yylex.Error("expecting share")
      return 1
    }
    if $4.Lowered() != "mode" {
      yylex.Error("expecting mode")
      return 1
    }
    $$ = ShareModeStr
  }

data_type:
  BIGINT
  { $$ = &DataType{ Type: "bigint"} }
| BIGSERIAL
  { $$ = &DataType{ Type: "bigserial"} }
| BIT
  { $$ = &DataType{ Type: "bit"} }
| BIT_VARYING
  { $$ = &DataType{ Type: "bit_varying"} }
| BOOLEAN
  { $$ = &DataType{ Type: "boolean"} }
| BOX
  { $$ = &DataType{ Type: "box"} }
| BYTEA
  { $$ = &DataType{ Type: "bytea"} }
| CHAR
  { $$ = &DataType{ Type: "char"} }
| CHAR_VARYING
  { $$ = &DataType{ Type: "char_varying"} }
| CIDR
  { $$ = &DataType{ Type: "cidr"} }
| CIRCLE
  { $$ = &DataType{ Type: "circle"} }
| DATE
  { $$ = &DataType{ Type: "date"} }
| DECIMAL
  { $$ = &DataType{ Type: "decimal"} }
| DOUBLE
  { $$ = &DataType{ Type: "double"} }
| ENUM
  { $$ = &DataType{ Type: "enum"} }
| FLOAT
  { $$ = &DataType{ Type: "float"} }
| INET
  { $$ = &DataType{ Type: "inet"} }
| INTEGER
  { $$ = &DataType{ Type: "integer"} }
//| INTERVAL
//  { $$ = &DataType{ Type: "interval"} }
| JSON
  { $$ = &DataType{ Type: "json"} }
| JSONB
  { $$ = &DataType{ Type: "jsonb"} }
| LINE
  { $$ = &DataType{ Type: "line"} }
| LSEG
  { $$ = &DataType{ Type: "lseg"} }
| MACADDR
  { $$ = &DataType{ Type: "macaddr"} }
| MONEY
  { $$ = &DataType{ Type: "money"} }
| NUMERIC
  { $$ = &DataType{ Type: "numeric"} }
| PATH
  { $$ = &DataType{ Type: "path"} }
| POINT
  { $$ = &DataType{ Type: "point"} }
| POLYGON
  { $$ = &DataType{ Type: "polygon"} }
| REAL
  { $$ = &DataType{ Type: "real"} }
| SERIAL
  { $$ = &DataType{ Type: "serial"} }
| SMALLINT
  { $$ = &DataType{ Type: "smallint"} }
| SMALLSERIAL
  { $$ = &DataType{ Type: "smallserial"} }
| TEXT
  { $$ = &DataType{ Type: "text"} }
| TIME
  { $$ = &DataType{ Type: "time"} }
| TIMESTAMP
  { $$ = &DataType{ Type: "timestamp"} }
| UUID
  { $$ = &DataType{ Type: "uuid"} }
| XML
  { $$ = &DataType{ Type: "xml"} }


column_constraint:
  NOT NULL
  {
    $$ = &ColConstr{ Constraint: ColConstrNotNullStr }
  }
| NULL
  {
    $$ = &ColConstr{ Constraint: ColConstrNullStr }
  }
| PRIMARY KEY
  {
    $$ = &ColConstr{ Constraint: ColConstrPrimaryKeyStr }
  }
// | DEFAULT sql_id
//   {
//     $$ = &ColConstr{ Constraint: ColConstrDefaultStr, params : string($2) }
//   }

column_constraint_list:
  column_constraint
  {
    $$ = ColConstrs{$1}
  }
| column_constraint_list column_constraint
  {
    $$ = append($$, $2)
  }


column_definition:
  sql_id data_type column_constraint_list 
  {
    $$ = &ColDef{ColName: $1, ColType: $2, Constraints: $3 }
  }
| sql_id data_type
  {
    $$ = &ColDef{ColName: $1, ColType: $2 }
  }

column_definition_list:
  column_definition
  {
    $$ = ColDefs{$1}
  }
| column_definition_list ',' column_definition
  {
    $$ = append($$, $3)
  }
  
column_definition_list_opt:
  {
    $$ = nil
  }
| openb column_definition_list closeb
  {
    $$ = $2
  }

column_list_opt:
  {
    $$ = nil
  }
| openb column_list closeb
  {
    $$ = $2
  }

column_list:
  sql_id
  {
    $$ = Columns{$1}
  }
| column_list ',' sql_id
  {
    $$ = append($$, $3)
  }



on_dup_opt:
  {
    $$ = nil
  }
| ON DUPLICATE KEY UPDATE update_list
  {
    $$ = $5
  }

row_list:
  VALUES tuple_list
  {
    $$ = $2
  }
| select_statement
  {
    $$ = $1
  }

tuple_list:
  row_tuple
  {
    $$ = Values{$1}
  }
| tuple_list ',' row_tuple
  {
    $$ = append($1, $3)
  }

row_tuple:
  openb value_expression_list closeb
  {
    $$ = ValTuple($2)
  }
| subquery
  {
    $$ = $1
  }

update_list:
  update_expression
  {
    $$ = UpdateExprs{$1}
  }
| update_list ',' update_expression
  {
    $$ = append($1, $3)
  }

update_expression:
  sql_id '=' value_expression
  {
    $$ = &UpdateExpr{Name: $1, Expr: $3}
  }

for_from:
  FOR
| FROM

exists_opt:
  { $$ = 0 }
| IF EXISTS
  { $$ = 1 }

not_exists_opt:
  { $$ = struct{}{} }
| IF NOT EXISTS
  { $$ = struct{}{} }

ignore_opt:
  { $$ = "" }
| IGNORE
  { $$ = IgnoreStr }

non_rename_operation:
  ALTER
  { $$ = struct{}{} }
| DEFAULT
  { $$ = struct{}{} }
| DROP
  { $$ = struct{}{} }
| ORDER
  { $$ = struct{}{} }
| UNUSED
  { $$ = struct{}{} }
| ID
  { $$ = struct{}{} }

to_opt:
  { $$ = struct{}{} }
| TO
  { $$ = struct{}{} }

constraint_opt:
  { $$ = struct{}{} }
| UNIQUE
  { $$ = struct{}{} }

using_opt:
  { $$ = struct{}{} }
| USING sql_id
  { $$ = struct{}{} }

sql_id:
  ID
  {
    $$ = NewColIdent(string($1))
  }

table_id:
  ID
  {
    $$ = TableIdent($1)
  }

openb:
  '('
  {
    if incNesting(yylex) {
      yylex.Error("max nesting level reached")
      return 1
    }
  }

closeb:
  ')'
  {
    decNesting(yylex)
  }

force_eof:
{
  forceEOF(yylex)
}
