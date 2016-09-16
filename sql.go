//line sql.y:6
package sqlparser

import __yyfmt__ "fmt"

//line sql.y:6
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

//line sql.y:34
type yySymType struct {
	yys         int
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
	tableName   *TableName
	indexHints  *IndexHints
	expr        Expr
	boolExpr    BoolExpr
	valExpr     ValExpr
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
	limit       *Limit
	insRows     InsertRows
	updateExprs UpdateExprs
	updateExpr  *UpdateExpr
	colIdent    ColIdent
	colIdents   []ColIdent
	tableIdent  TableIdent

	dataType   *DataType
	colDef     *ColDef
	colDefs    ColDefs
	colConstr  *ColConstr
	colConstrs ColConstrs
}

const LEX_ERROR = 57346
const UNION = 57347
const SELECT = 57348
const INSERT = 57349
const UPDATE = 57350
const DELETE = 57351
const FROM = 57352
const WHERE = 57353
const GROUP = 57354
const HAVING = 57355
const ORDER = 57356
const BY = 57357
const LIMIT = 57358
const FOR = 57359
const ALL = 57360
const DISTINCT = 57361
const AS = 57362
const EXISTS = 57363
const ASC = 57364
const DESC = 57365
const INTO = 57366
const DUPLICATE = 57367
const KEY = 57368
const DEFAULT = 57369
const SET = 57370
const LOCK = 57371
const VALUES = 57372
const LAST_INSERT_ID = 57373
const NEXT = 57374
const VALUE = 57375
const JOIN = 57376
const STRAIGHT_JOIN = 57377
const LEFT = 57378
const RIGHT = 57379
const INNER = 57380
const OUTER = 57381
const CROSS = 57382
const NATURAL = 57383
const USE = 57384
const FORCE = 57385
const ON = 57386
const ID = 57387
const STRING = 57388
const NUMBER = 57389
const VALUE_ARG = 57390
const LIST_ARG = 57391
const COMMENT = 57392
const NULL = 57393
const TRUE = 57394
const FALSE = 57395
const OR = 57396
const AND = 57397
const NOT = 57398
const BETWEEN = 57399
const CASE = 57400
const WHEN = 57401
const THEN = 57402
const ELSE = 57403
const LE = 57404
const GE = 57405
const NE = 57406
const NULL_SAFE_EQUAL = 57407
const IS = 57408
const LIKE = 57409
const REGEXP = 57410
const IN = 57411
const SHIFT_LEFT = 57412
const SHIFT_RIGHT = 57413
const UNARY = 57414
const INTERVAL = 57415
const END = 57416
const CREATE = 57417
const ALTER = 57418
const DROP = 57419
const RENAME = 57420
const ANALYZE = 57421
const TABLE = 57422
const INDEX = 57423
const VIEW = 57424
const TO = 57425
const IGNORE = 57426
const IF = 57427
const UNIQUE = 57428
const USING = 57429
const SHOW = 57430
const DESCRIBE = 57431
const EXPLAIN = 57432
const PRIMARY = 57433
const UNUSED = 57434
const SMALLINT = 57435
const INTEGER = 57436
const BIGINT = 57437
const FLOAT = 57438
const DECIMAL = 57439
const NUMERIC = 57440
const REAL = 57441
const DOUBLE = 57442
const SMALLSERIAL = 57443
const SERIAL = 57444
const BIGSERIAL = 57445
const MONEY = 57446
const CHAR_VARYING = 57447
const CHAR = 57448
const TEXT = 57449
const BYTEA = 57450
const TIMESTAMP = 57451
const DATE = 57452
const TIME = 57453
const BOOLEAN = 57454
const ENUM = 57455
const POINT = 57456
const LINE = 57457
const LSEG = 57458
const BOX = 57459
const PATH = 57460
const POLYGON = 57461
const CIRCLE = 57462
const CIDR = 57463
const INET = 57464
const MACADDR = 57465
const BIT = 57466
const BIT_VARYING = 57467
const UUID = 57468
const XML = 57469
const JSON = 57470
const JSONB = 57471

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"LEX_ERROR",
	"UNION",
	"SELECT",
	"INSERT",
	"UPDATE",
	"DELETE",
	"FROM",
	"WHERE",
	"GROUP",
	"HAVING",
	"ORDER",
	"BY",
	"LIMIT",
	"FOR",
	"ALL",
	"DISTINCT",
	"AS",
	"EXISTS",
	"ASC",
	"DESC",
	"INTO",
	"DUPLICATE",
	"KEY",
	"DEFAULT",
	"SET",
	"LOCK",
	"VALUES",
	"LAST_INSERT_ID",
	"NEXT",
	"VALUE",
	"JOIN",
	"STRAIGHT_JOIN",
	"LEFT",
	"RIGHT",
	"INNER",
	"OUTER",
	"CROSS",
	"NATURAL",
	"USE",
	"FORCE",
	"ON",
	"'('",
	"','",
	"')'",
	"ID",
	"STRING",
	"NUMBER",
	"VALUE_ARG",
	"LIST_ARG",
	"COMMENT",
	"NULL",
	"TRUE",
	"FALSE",
	"OR",
	"AND",
	"NOT",
	"BETWEEN",
	"CASE",
	"WHEN",
	"THEN",
	"ELSE",
	"'='",
	"'<'",
	"'>'",
	"LE",
	"GE",
	"NE",
	"NULL_SAFE_EQUAL",
	"IS",
	"LIKE",
	"REGEXP",
	"IN",
	"'|'",
	"'&'",
	"SHIFT_LEFT",
	"SHIFT_RIGHT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'^'",
	"'~'",
	"UNARY",
	"INTERVAL",
	"'.'",
	"END",
	"CREATE",
	"ALTER",
	"DROP",
	"RENAME",
	"ANALYZE",
	"TABLE",
	"INDEX",
	"VIEW",
	"TO",
	"IGNORE",
	"IF",
	"UNIQUE",
	"USING",
	"SHOW",
	"DESCRIBE",
	"EXPLAIN",
	"PRIMARY",
	"UNUSED",
	"SMALLINT",
	"INTEGER",
	"BIGINT",
	"FLOAT",
	"DECIMAL",
	"NUMERIC",
	"REAL",
	"DOUBLE",
	"SMALLSERIAL",
	"SERIAL",
	"BIGSERIAL",
	"MONEY",
	"CHAR_VARYING",
	"CHAR",
	"TEXT",
	"BYTEA",
	"TIMESTAMP",
	"DATE",
	"TIME",
	"BOOLEAN",
	"ENUM",
	"POINT",
	"LINE",
	"LSEG",
	"BOX",
	"PATH",
	"POLYGON",
	"CIRCLE",
	"CIDR",
	"INET",
	"MACADDR",
	"BIT",
	"BIT_VARYING",
	"UUID",
	"XML",
	"JSON",
	"JSONB",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 106,
	45, 272,
	89, 272,
	-2, 271,
	-1, 119,
	59, 168,
	60, 168,
	65, 168,
	66, 168,
	67, 168,
	68, 168,
	69, 168,
	70, 168,
	71, 168,
	73, 168,
	74, 168,
	75, 168,
	76, 168,
	77, 168,
	78, 168,
	79, 168,
	80, 168,
	81, 168,
	82, 168,
	83, 168,
	84, 168,
	85, 168,
	-2, 99,
	-1, 120,
	59, 169,
	60, 169,
	65, 169,
	66, 169,
	67, 169,
	68, 169,
	69, 169,
	70, 169,
	71, 169,
	73, 169,
	74, 169,
	75, 169,
	76, 169,
	77, 169,
	78, 169,
	79, 169,
	80, 169,
	81, 169,
	82, 169,
	83, 169,
	84, 169,
	85, 169,
	-2, 100,
}

const yyNprod = 276
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 851

var yyAct = [...]int{

	127, 198, 241, 458, 215, 62, 195, 383, 387, 340,
	47, 292, 112, 230, 100, 222, 373, 282, 229, 253,
	247, 231, 136, 197, 3, 148, 389, 99, 407, 409,
	243, 388, 121, 74, 67, 64, 48, 49, 69, 41,
	228, 71, 14, 15, 16, 17, 35, 40, 37, 41,
	94, 50, 38, 441, 440, 80, 122, 439, 244, 106,
	123, 124, 125, 68, 18, 126, 119, 120, 418, 104,
	108, 70, 129, 46, 128, 42, 355, 98, 214, 390,
	91, 84, 93, 63, 64, 105, 408, 64, 122, 134,
	122, 113, 114, 101, 139, 179, 471, 115, 169, 116,
	233, 106, 87, 163, 43, 44, 45, 89, 283, 159,
	371, 260, 117, 430, 283, 81, 199, 358, 359, 360,
	200, 201, 202, 203, 258, 259, 257, 19, 20, 22,
	21, 23, 218, 58, 165, 238, 208, 167, 168, 167,
	24, 25, 26, 72, 420, 122, 223, 77, 86, 225,
	219, 169, 256, 169, 143, 220, 122, 213, 65, 60,
	226, 60, 85, 14, 106, 237, 239, 88, 105, 90,
	138, 92, 75, 105, 95, 351, 236, 252, 194, 196,
	261, 262, 263, 133, 265, 266, 267, 268, 269, 270,
	271, 272, 273, 274, 242, 157, 205, 264, 158, 275,
	276, 278, 122, 60, 279, 60, 277, 217, 164, 467,
	244, 105, 64, 289, 277, 244, 425, 287, 290, 239,
	182, 183, 184, 179, 296, 180, 181, 182, 183, 184,
	179, 209, 280, 109, 286, 234, 65, 245, 246, 160,
	168, 167, 374, 145, 86, 356, 255, 161, 244, 244,
	105, 28, 337, 422, 354, 169, 361, 73, 166, 168,
	167, 219, 137, 109, 109, 363, 364, 365, 248, 250,
	251, 357, 204, 249, 169, 161, 206, 362, 338, 244,
	379, 244, 137, 368, 367, 295, 244, 122, 212, 211,
	79, 378, 109, 244, 435, 166, 223, 338, 376, 370,
	381, 384, 380, 76, 377, 154, 122, 121, 374, 224,
	234, 385, 97, 438, 402, 235, 109, 86, 152, 403,
	14, 109, 109, 109, 400, 335, 254, 437, 336, 401,
	399, 122, 255, 398, 106, 123, 124, 125, 55, 156,
	126, 119, 120, 39, 285, 108, 391, 129, 427, 428,
	464, 54, 394, 372, 396, 83, 413, 412, 405, 109,
	414, 395, 465, 397, 415, 444, 113, 114, 101, 431,
	146, 421, 115, 417, 116, 57, 82, 96, 419, 404,
	423, 346, 347, 151, 153, 150, 132, 117, 353, 434,
	235, 51, 52, 131, 293, 429, 294, 216, 109, 155,
	433, 29, 178, 177, 185, 186, 180, 181, 182, 183,
	184, 179, 254, 234, 234, 234, 234, 31, 32, 33,
	34, 393, 137, 442, 416, 61, 470, 443, 456, 14,
	28, 446, 384, 447, 30, 445, 450, 448, 109, 1,
	140, 219, 178, 177, 185, 186, 180, 181, 182, 183,
	184, 179, 297, 386, 459, 459, 459, 64, 460, 461,
	457, 466, 462, 468, 469, 221, 144, 352, 472, 349,
	162, 147, 473, 59, 474, 185, 186, 180, 181, 182,
	183, 184, 179, 59, 36, 227, 149, 59, 66, 130,
	288, 210, 463, 235, 235, 235, 235, 426, 14, 382,
	432, 392, 59, 449, 369, 451, 452, 59, 207, 281,
	118, 59, 111, 121, 59, 375, 110, 284, 170, 103,
	107, 406, 341, 59, 424, 135, 177, 185, 186, 180,
	181, 182, 183, 184, 179, 59, 339, 122, 59, 232,
	106, 123, 124, 125, 102, 78, 126, 119, 120, 53,
	27, 108, 56, 129, 178, 177, 185, 186, 180, 181,
	182, 183, 184, 179, 13, 12, 11, 10, 9, 8,
	7, 59, 113, 114, 6, 5, 4, 2, 115, 0,
	116, 0, 0, 0, 0, 0, 0, 0, 109, 0,
	109, 109, 0, 117, 453, 454, 455, 0, 0, 0,
	0, 59, 103, 0, 0, 0, 240, 103, 328, 315,
	298, 313, 310, 322, 326, 311, 329, 327, 299, 321,
	306, 305, 330, 304, 332, 309, 331, 302, 312, 324,
	318, 319, 303, 323, 325, 308, 307, 314, 320, 300,
	301, 333, 334, 316, 317, 103, 0, 0, 0, 0,
	0, 0, 0, 121, 0, 291, 14, 0, 0, 240,
	342, 345, 346, 347, 343, 59, 344, 348, 59, 0,
	436, 0, 0, 0, 350, 0, 59, 122, 0, 0,
	106, 123, 124, 125, 103, 0, 126, 119, 120, 0,
	0, 108, 0, 129, 0, 122, 0, 0, 106, 123,
	124, 125, 0, 0, 126, 141, 142, 0, 0, 0,
	0, 129, 113, 114, 0, 0, 0, 0, 115, 0,
	116, 366, 342, 345, 346, 347, 343, 0, 344, 348,
	113, 114, 0, 117, 0, 0, 115, 0, 116, 178,
	177, 185, 186, 180, 181, 182, 183, 184, 179, 0,
	0, 117, 0, 122, 0, 0, 106, 123, 124, 125,
	0, 0, 126, 141, 142, 0, 0, 0, 0, 129,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 59,
	59, 59, 59, 0, 0, 0, 0, 0, 113, 114,
	0, 0, 410, 411, 115, 0, 116, 0, 0, 0,
	0, 0, 0, 65, 172, 175, 0, 0, 0, 117,
	187, 188, 189, 190, 191, 192, 193, 176, 173, 174,
	171, 178, 177, 185, 186, 180, 181, 182, 183, 184,
	179, 178, 177, 185, 186, 180, 181, 182, 183, 184,
	179, 178, 177, 185, 186, 180, 181, 182, 183, 184,
	179,
}
var yyPact = [...]int{

	36, -1000, -1000, 425, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -50,
	-51, -21, 8, -23, -1000, -1000, -1000, 423, 373, 319,
	-1000, -61, 113, 415, 110, -67, -34, 110, -1000, -25,
	110, -1000, 113, -68, 124, -68, 113, -1000, -1000, -1000,
	-1000, -1000, -1000, 255, 110, -1000, 62, 352, 327, -8,
	-1000, 113, 102, -1000, 37, -1000, 113, 48, 121, -1000,
	113, -1000, -49, 113, 356, 268, 110, -1000, 286, -1000,
	376, -1000, 113, 110, 113, 411, 110, 708, 242, 349,
	-78, -1000, 291, -1000, 113, -1000, -1000, 113, -1000, 229,
	-1000, -1000, 188, 45, 183, 745, -1000, -1000, 632, 492,
	-1000, -1000, -1000, 708, 708, 708, 708, 242, -1000, -1000,
	-1000, 242, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 708,
	113, -1000, -1000, 261, 271, -11, 383, 632, -1000, 765,
	43, -1000, -1000, 650, -1000, 110, -1000, 265, 110, -1000,
	-59, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	111, 286, -1000, -1000, 110, 53, 11, 632, 632, 214,
	708, 100, 51, 708, 708, 708, 214, 708, 708, 708,
	708, 708, 708, 708, 708, 708, 708, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 26, 745, 202, 246, 168, 745,
	-1000, -1000, -1000, 755, 286, -1000, 423, 52, 765, -1000,
	314, 110, 110, 383, 113, 378, 381, 183, 116, 765,
	-1000, 239, -1000, 499, 113, -1000, -1000, 113, -1000, 251,
	688, -1000, -1000, 155, 368, 157, -1000, -1000, -1000, -1000,
	-13, -1000, 201, 286, -1000, 26, 79, -1000, -1000, 63,
	-1000, -1000, 765, -1000, 650, -1000, -1000, 100, 708, 708,
	708, 765, 765, 663, -1000, 397, 449, -1000, 138, 138,
	10, 10, 10, 145, 145, -1000, -1000, 708, -1000, -1000,
	201, 46, -1000, 632, 264, 242, 425, 198, 234, -1000,
	378, -1000, -1000, 708, 708, 110, -1000, -28, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 409, 111, 111,
	111, 111, -1000, 299, 296, -1000, 290, 280, 345, -14,
	-1000, 113, 113, -1000, 232, 110, -1000, 201, -1000, -1000,
	-1000, 168, -1000, 765, 765, 366, 708, 765, -1000, -22,
	-1000, 708, 81, -1000, 346, 207, -1000, -1000, -1000, 110,
	-1000, 478, 170, -1000, 326, -1000, -28, -1000, 59, -1000,
	343, -1000, 387, 374, 688, 250, 626, -1000, -1000, -1000,
	-1000, 293, -1000, 279, -1000, -1000, -1000, -40, -43, -44,
	-1000, -1000, -1000, -1000, -1000, -1000, 708, 765, -1000, 765,
	708, 339, 242, -1000, 708, 708, -1000, -1000, -1000, -1000,
	-1000, -1000, 383, 632, 708, 632, 632, -1000, -1000, 242,
	242, 242, 765, 765, 420, -1000, 765, -1000, 378, 183,
	160, 183, 183, 110, 110, 110, 110, 333, 163, -1000,
	163, 163, 102, -1000, 418, 21, -1000, 110, -1000, -1000,
	-1000, 110, -1000, 110, -1000,
}
var yyPgo = [...]int{

	0, 577, 23, 576, 575, 574, 570, 569, 568, 567,
	566, 565, 564, 401, 552, 550, 549, 545, 27, 14,
	544, 18, 13, 21, 539, 536, 9, 522, 100, 521,
	3, 22, 69, 520, 518, 517, 516, 6, 20, 19,
	1, 515, 12, 74, 512, 510, 509, 17, 508, 504,
	501, 500, 4, 499, 7, 497, 11, 492, 491, 490,
	16, 5, 83, 489, 343, 257, 488, 486, 485, 484,
	471, 0, 470, 440, 469, 467, 10, 15, 466, 465,
	8, 453, 452, 439, 434, 154, 2,
}
var yyR1 = [...]int{

	0, 83, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 3, 3, 4, 5,
	6, 7, 7, 7, 8, 8, 8, 9, 10, 10,
	10, 11, 12, 12, 12, 84, 13, 14, 14, 15,
	15, 15, 16, 16, 17, 17, 18, 18, 19, 19,
	19, 20, 20, 72, 72, 72, 21, 21, 22, 22,
	23, 23, 23, 24, 24, 24, 24, 75, 75, 74,
	74, 74, 25, 25, 25, 25, 26, 26, 26, 26,
	27, 27, 28, 28, 28, 29, 29, 29, 29, 30,
	30, 31, 31, 32, 32, 32, 32, 32, 32, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 38, 38, 38, 38, 38, 38, 34, 34,
	34, 34, 34, 34, 34, 39, 39, 39, 43, 40,
	40, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 45, 48, 48, 46, 46, 47, 49,
	49, 44, 44, 44, 36, 36, 36, 36, 36, 36,
	50, 50, 51, 51, 52, 52, 53, 53, 54, 55,
	55, 55, 56, 56, 56, 57, 57, 57, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 80, 80, 80, 81, 81,
	77, 77, 79, 79, 78, 78, 58, 58, 59, 59,
	60, 60, 35, 35, 41, 41, 42, 42, 61, 61,
	62, 63, 63, 65, 65, 66, 66, 64, 64, 67,
	67, 67, 67, 67, 67, 68, 68, 69, 69, 70,
	70, 71, 73, 85, 86, 76,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 13, 6, 3, 8, 8, 8, 7,
	3, 6, 8, 4, 6, 7, 4, 5, 4, 5,
	5, 3, 2, 2, 2, 0, 2, 0, 2, 1,
	2, 2, 0, 1, 0, 1, 1, 3, 1, 2,
	3, 1, 1, 0, 1, 2, 1, 3, 1, 1,
	3, 3, 3, 3, 5, 5, 3, 0, 1, 0,
	1, 2, 1, 2, 2, 1, 2, 3, 2, 3,
	2, 2, 1, 3, 5, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 1,
	1, 3, 3, 4, 3, 4, 3, 4, 5, 6,
	3, 2, 1, 2, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 2, 2, 3, 3, 4,
	5, 4, 1, 5, 0, 1, 1, 2, 4, 0,
	2, 1, 3, 5, 1, 1, 1, 1, 1, 1,
	0, 3, 0, 2, 0, 3, 1, 3, 2, 0,
	1, 1, 0, 2, 4, 0, 2, 4, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
	3, 2, 1, 3, 0, 3, 0, 3, 1, 3,
	0, 5, 2, 1, 1, 3, 3, 1, 1, 3,
	3, 1, 1, 0, 2, 0, 3, 0, 1, 1,
	1, 1, 1, 1, 1, 0, 1, 0, 1, 0,
	2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -83, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 28, 91,
	92, 94, 93, 95, 104, 105, 106, -15, 5, -13,
	-84, -13, -13, -13, -13, 96, -69, 98, 102, -64,
	98, 100, 96, 96, 97, 98, 96, -76, -76, -76,
	-2, 18, 19, -16, 32, 19, -14, -64, -28, -73,
	48, 10, -61, -62, -71, 48, -66, 101, 97, -71,
	96, -71, -28, -65, 101, 48, -65, -28, -17, 35,
	-71, 53, 24, 28, 89, -28, 46, 65, -28, 59,
	48, -76, -28, -76, 99, -28, 21, 44, -71, -18,
	-19, 82, -20, -73, -32, -37, 48, -33, 59, -85,
	-36, -44, -42, 80, 81, 86, 88, 101, -45, 55,
	56, 21, 45, 49, 50, 51, 54, -71, -43, 61,
	-63, 17, 10, -28, -61, -73, -31, 11, -62, -37,
	-73, 55, 56, -85, -78, -85, 21, -70, 103, -67,
	94, 92, 27, 93, 14, 108, 48, -28, -28, -76,
	10, 46, -72, -71, 20, 89, -85, 58, 57, 72,
	-34, 75, 59, 73, 74, 60, 72, 77, 76, 85,
	80, 81, 82, 83, 84, 78, 79, 65, 66, 67,
	68, 69, 70, 71, -32, -37, -32, -2, -40, -37,
	-37, -37, -37, -37, -85, -43, -85, -48, -37, -28,
	-58, 28, -85, -31, 89, -52, 14, -32, 89, -37,
	-76, -79, -77, -71, 44, -71, -76, -68, 99, -21,
	-22, -23, -24, -28, -43, -85, -19, -71, 82, -71,
	-73, -86, -18, 19, 47, -32, -32, -38, 54, 59,
	55, 56, -37, -39, -85, -43, 52, 75, 73, 74,
	60, -37, -37, -37, -38, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -86, -86, 46, -86, -71,
	-18, -46, -47, 62, -35, 30, -2, -61, -59, -71,
	-52, -73, -56, 16, 15, 46, -86, -82, 111, 119,
	140, 141, 128, 133, 124, 122, 121, 137, 136, 126,
	113, 116, 129, 112, 138, 110, 144, 145, 131, 132,
	139, 120, 114, 134, 130, 135, 115, 118, 109, 117,
	123, 127, 125, 142, 143, -28, -28, -31, 46, -25,
	-26, -27, 34, 38, 40, 35, 36, 37, 41, -74,
	-73, 20, -75, 20, -21, 89, -86, -18, 54, 55,
	56, -40, -39, -37, -37, -37, 58, -37, -86, -49,
	-47, 64, -32, -60, 44, -41, -42, -60, -86, 46,
	-56, -37, -53, -54, -37, -77, -81, -80, 59, 54,
	107, -76, -50, 12, -22, -23, -22, -23, 34, 34,
	34, 39, 34, 39, 34, -26, -29, 42, 100, 43,
	-73, -73, -86, -71, -86, -86, 58, -37, 90, -37,
	63, 25, 46, -71, 46, 46, -55, 22, 23, -80,
	54, 26, -51, 13, 15, 44, 44, 34, 34, 97,
	97, 97, -37, -37, 26, -42, -37, -54, -52, -32,
	-40, -32, -32, -85, -85, -85, 8, -56, -30, -71,
	-30, -30, -61, -57, 17, 29, -86, 46, -86, -86,
	8, 75, -71, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 35, 35, 35, 35, 35, 267,
	257, 0, 0, 0, 275, 275, 275, 0, 39, 42,
	37, 257, 0, 0, 0, 255, 0, 0, 268, 0,
	0, 258, 0, 253, 0, 253, 0, 32, 33, 34,
	15, 40, 41, 44, 0, 43, 36, 0, 0, 82,
	272, 0, 20, 248, 0, 271, 0, 0, 0, 275,
	0, 275, 0, 0, 0, 0, 0, 31, 0, 45,
	0, 38, 0, 0, 0, 91, 0, 0, 234, 0,
	269, 23, 0, 26, 0, 28, 254, 0, 275, 0,
	46, 48, 53, 0, 51, 52, -2, 93, 0, 0,
	131, 132, 133, 0, 0, 0, 0, 0, 152, -2,
	-2, 0, 273, 164, 165, 166, 167, 161, 247, 154,
	0, 251, 252, 236, 91, 83, 174, 0, 249, 250,
	0, 168, 169, 0, 275, 0, 256, 0, 0, 275,
	265, 259, 260, 261, 262, 263, 264, 27, 29, 30,
	0, 0, 49, 54, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 118, 119, 120,
	121, 122, 123, 124, 96, 0, 0, 0, 0, 129,
	144, 145, 146, 0, 0, 111, 0, 0, 155, 14,
	0, 0, 0, 174, 0, 182, 0, 92, 0, 129,
	21, 0, 232, 0, 0, 270, 24, 0, 266, 91,
	56, 58, 59, 69, 67, 0, 47, 55, 50, 162,
	0, 148, 0, 0, 274, 94, 95, 98, 112, 0,
	114, 116, 101, 102, 0, 126, 127, 0, 0, 0,
	0, 104, 106, 0, 110, 134, 135, 136, 137, 138,
	139, 140, 141, 142, 143, 97, 128, 0, 246, 147,
	0, 159, 156, 0, 240, 0, 243, 240, 0, 238,
	182, 84, 19, 0, 0, 0, 235, 231, 188, 189,
	190, 191, 192, 193, 194, 195, 196, 197, 198, 199,
	200, 201, 202, 203, 204, 205, 206, 207, 208, 209,
	210, 211, 212, 213, 214, 215, 216, 217, 218, 219,
	220, 221, 222, 223, 224, 275, 25, 170, 0, 0,
	0, 0, 72, 0, 0, 75, 0, 0, 0, 85,
	70, 0, 0, 68, 0, 0, 149, 0, 113, 115,
	117, 0, 103, 105, 107, 0, 0, 130, 151, 0,
	157, 0, 0, 16, 0, 242, 244, 17, 237, 0,
	18, 183, 175, 176, 179, 233, 230, 228, 0, 226,
	0, 22, 172, 0, 57, 63, 0, 66, 73, 74,
	76, 0, 78, 0, 80, 81, 60, 0, 0, 0,
	71, 61, 62, 163, 150, 125, 0, 108, 153, 160,
	0, 0, 0, 239, 0, 0, 178, 180, 181, 229,
	225, 227, 174, 0, 0, 0, 0, 77, 79, 0,
	0, 0, 109, 158, 0, 245, 184, 177, 182, 173,
	171, 64, 65, 0, 0, 0, 0, 185, 0, 89,
	0, 0, 241, 13, 0, 0, 86, 0, 87, 88,
	186, 0, 90, 0, 187,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 84, 77, 3,
	45, 47, 82, 80, 46, 81, 89, 83, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	66, 65, 67, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 85, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 76, 3, 86,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 48, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	68, 69, 70, 71, 72, 73, 74, 75, 78, 79,
	87, 88, 90, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:197
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:203
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:219
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, Limit: yyDollar[12].limit, Lock: yyDollar[13].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:223
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:231
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:237
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:241
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:253
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, Limit: yyDollar[8].limit}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, Limit: yyDollar[7].limit}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableName, ColDefs: yyDollar[5].colDefs}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:275
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:280
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:286
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:290
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:295
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}, NewName: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:301
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:307
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:315
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:320
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: &TableName{Name: TableIdent(yyDollar[4].colIdent.Lowered())}, IfExists: exists}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:330
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:340
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:344
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:349
		{
			setAllowComments(yylex, true)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:353
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:359
		{
			yyVAL.bytes2 = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:363
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:369
		{
			yyVAL.str = UnionStr
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:373
		{
			yyVAL.str = UnionAllStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:377
		{
			yyVAL.str = UnionDistinctStr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:382
		{
			yyVAL.str = ""
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:386
		{
			yyVAL.str = DistinctStr
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:391
		{
			yyVAL.str = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:395
		{
			yyVAL.str = StraightJoinHint
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:401
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:405
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:411
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:415
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:419
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:425
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:429
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:434
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:438
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:442
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:448
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:452
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:462
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:466
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:470
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:483
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:487
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:491
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:495
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:500
		{
			yyVAL.empty = struct{}{}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:502
		{
			yyVAL.empty = struct{}{}
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:505
		{
			yyVAL.tableIdent = ""
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:509
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:513
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:519
		{
			yyVAL.str = JoinStr
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:523
		{
			yyVAL.str = JoinStr
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:527
		{
			yyVAL.str = JoinStr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:531
		{
			yyVAL.str = StraightJoinStr
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:537
		{
			yyVAL.str = LeftJoinStr
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:541
		{
			yyVAL.str = LeftJoinStr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:545
		{
			yyVAL.str = RightJoinStr
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:549
		{
			yyVAL.str = RightJoinStr
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:555
		{
			yyVAL.str = NaturalJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:559
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:569
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:573
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:577
		{
			yyVAL.tableName = &TableName{User: yyDollar[1].tableIdent, Qualifier: yyDollar[3].tableIdent, Name: yyDollar[5].tableIdent}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:582
		{
			yyVAL.indexHints = nil
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:586
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:590
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:594
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:600
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:604
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:609
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:613
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:620
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:624
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:628
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:632
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:636
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:642
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:646
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:650
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:654
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:658
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:662
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:666
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:670
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:674
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:678
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:682
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:686
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:690
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:696
		{
			yyVAL.str = IsNullStr
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:700
		{
			yyVAL.str = IsNotNullStr
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:704
		{
			yyVAL.str = IsTrueStr
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:708
		{
			yyVAL.str = IsNotTrueStr
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:712
		{
			yyVAL.str = IsFalseStr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:716
		{
			yyVAL.str = IsNotFalseStr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:722
		{
			yyVAL.str = EqualStr
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:726
		{
			yyVAL.str = LessThanStr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:730
		{
			yyVAL.str = GreaterThanStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:734
		{
			yyVAL.str = LessEqualStr
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:738
		{
			yyVAL.str = GreaterEqualStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:742
		{
			yyVAL.str = NotEqualStr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:746
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:752
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:756
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:760
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:766
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:772
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:776
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:782
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:786
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:790
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:794
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:798
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:802
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:806
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:810
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:814
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:818
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:822
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:826
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:830
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:834
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:842
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				// Handle double negative
				if num[0] == '-' {
					yyVAL.valExpr = num[1:]
				} else {
					yyVAL.valExpr = append(NumVal("-"), num...)
				}
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UMinusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:855
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:859
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:867
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:871
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:875
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:879
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:883
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:889
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:894
		{
			yyVAL.valExpr = nil
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:898
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:904
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:908
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:914
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:919
		{
			yyVAL.valExpr = nil
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:923
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:929
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:933
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:937
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 164:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:943
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:947
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:951
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:955
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:959
		{
			yyVAL.valExpr = BoolVal(true)
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:963
		{
			yyVAL.valExpr = BoolVal(false)
		}
	case 170:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:968
		{
			yyVAL.valExprs = nil
		}
	case 171:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:972
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 172:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:977
		{
			yyVAL.boolExpr = nil
		}
	case 173:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:981
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 174:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:986
		{
			yyVAL.orderBy = nil
		}
	case 175:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:990
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 176:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:996
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 177:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 178:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.str = AscScr
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1015
		{
			yyVAL.str = AscScr
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1019
		{
			yyVAL.str = DescScr
		}
	case 182:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.limit = nil
		}
	case 183:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1028
		{
			yyVAL.limit = &Limit{Rowcount: yyDollar[2].valExpr}
		}
	case 184:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1032
		{
			yyVAL.limit = &Limit{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.str = ""
		}
	case 186:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.str = ForUpdateStr
		}
	case 187:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1045
		{
			if yyDollar[3].colIdent.Lowered() != "share" {
				yylex.Error("expecting share")
				return 1
			}
			if yyDollar[4].colIdent.Lowered() != "mode" {
				yylex.Error("expecting mode")
				return 1
			}
			yyVAL.str = ShareModeStr
		}
	case 188:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1059
		{
			yyVAL.dataType = &DataType{Type: "bigint"}
		}
	case 189:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1061
		{
			yyVAL.dataType = &DataType{Type: "bigserial"}
		}
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1063
		{
			yyVAL.dataType = &DataType{Type: "bit"}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1065
		{
			yyVAL.dataType = &DataType{Type: "bit_varying"}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1067
		{
			yyVAL.dataType = &DataType{Type: "boolean"}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1069
		{
			yyVAL.dataType = &DataType{Type: "box"}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1071
		{
			yyVAL.dataType = &DataType{Type: "bytea"}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1073
		{
			yyVAL.dataType = &DataType{Type: "char"}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.dataType = &DataType{Type: "char_varying"}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.dataType = &DataType{Type: "cidr"}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.dataType = &DataType{Type: "circle"}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.dataType = &DataType{Type: "date"}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.dataType = &DataType{Type: "decimal"}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.dataType = &DataType{Type: "double"}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.dataType = &DataType{Type: "enum"}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.dataType = &DataType{Type: "float"}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.dataType = &DataType{Type: "inet"}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.dataType = &DataType{Type: "integer"}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.dataType = &DataType{Type: "json"}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.dataType = &DataType{Type: "jsonb"}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.dataType = &DataType{Type: "line"}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.dataType = &DataType{Type: "lseg"}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.dataType = &DataType{Type: "macaddr"}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.dataType = &DataType{Type: "money"}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1109
		{
			yyVAL.dataType = &DataType{Type: "numeric"}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.dataType = &DataType{Type: "path"}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.dataType = &DataType{Type: "point"}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.dataType = &DataType{Type: "polygon"}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.dataType = &DataType{Type: "real"}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.dataType = &DataType{Type: "serial"}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.dataType = &DataType{Type: "smallint"}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.dataType = &DataType{Type: "smallserial"}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.dataType = &DataType{Type: "text"}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.dataType = &DataType{Type: "time"}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.dataType = &DataType{Type: "timestamp"}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.dataType = &DataType{Type: "uuid"}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.dataType = &DataType{Type: "xml"}
		}
	case 225:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNotNullStr}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNullStr}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1146
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrPrimaryKeyStr}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.colConstrs = ColConstrs{yyDollar[1].colConstr}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.colConstrs = append(yyVAL.colConstrs, yyDollar[2].colConstr)
		}
	case 230:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1167
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType, Constraints: yyDollar[3].colConstrs}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1171
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType}
		}
	case 232:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1177
		{
			yyVAL.colDefs = ColDefs{yyDollar[1].colDef}
		}
	case 233:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.colDefs = append(yyVAL.colDefs, yyDollar[3].colDef)
		}
	case 234:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1186
		{
			yyVAL.colDefs = nil
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1190
		{
			yyVAL.colDefs = yyDollar[2].colDefs
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1195
		{
			yyVAL.columns = nil
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1199
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 238:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1205
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1209
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 240:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1216
		{
			yyVAL.updateExprs = nil
		}
	case 241:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1220
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 242:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1226
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 243:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1230
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 244:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1236
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 245:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 246:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1246
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 247:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 248:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1256
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 249:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 250:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1266
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 253:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.byt = 0
		}
	case 254:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1277
		{
			yyVAL.byt = 1
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.empty = struct{}{}
		}
	case 256:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1282
		{
			yyVAL.empty = struct{}{}
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1285
		{
			yyVAL.str = ""
		}
	case 258:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1287
		{
			yyVAL.str = IgnoreStr
		}
	case 259:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1291
		{
			yyVAL.empty = struct{}{}
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1293
		{
			yyVAL.empty = struct{}{}
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1295
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1297
		{
			yyVAL.empty = struct{}{}
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1299
		{
			yyVAL.empty = struct{}{}
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1301
		{
			yyVAL.empty = struct{}{}
		}
	case 265:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1304
		{
			yyVAL.empty = struct{}{}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1306
		{
			yyVAL.empty = struct{}{}
		}
	case 267:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.empty = struct{}{}
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1311
		{
			yyVAL.empty = struct{}{}
		}
	case 269:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1314
		{
			yyVAL.empty = struct{}{}
		}
	case 270:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1316
		{
			yyVAL.empty = struct{}{}
		}
	case 271:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1320
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 272:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1326
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1332
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1341
		{
			decNesting(yylex)
		}
	case 275:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1346
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
