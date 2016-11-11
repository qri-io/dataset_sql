//line sql.y:6
package dataset_sql

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
	limitOffset *LimitOffset
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
const OFFSET = 57360
const ALL = 57361
const DISTINCT = 57362
const AS = 57363
const EXISTS = 57364
const ASC = 57365
const DESC = 57366
const INTO = 57367
const DUPLICATE = 57368
const KEY = 57369
const DEFAULT = 57370
const SET = 57371
const LOCK = 57372
const VALUES = 57373
const LAST_INSERT_ID = 57374
const NEXT = 57375
const VALUE = 57376
const JOIN = 57377
const STRAIGHT_JOIN = 57378
const LEFT = 57379
const RIGHT = 57380
const INNER = 57381
const OUTER = 57382
const CROSS = 57383
const NATURAL = 57384
const USE = 57385
const FORCE = 57386
const ON = 57387
const ID = 57388
const STRING = 57389
const NUMBER = 57390
const VALUE_ARG = 57391
const LIST_ARG = 57392
const COMMENT = 57393
const NULL = 57394
const TRUE = 57395
const FALSE = 57396
const OR = 57397
const AND = 57398
const NOT = 57399
const BETWEEN = 57400
const CASE = 57401
const WHEN = 57402
const THEN = 57403
const ELSE = 57404
const LE = 57405
const GE = 57406
const NE = 57407
const NULL_SAFE_EQUAL = 57408
const IS = 57409
const LIKE = 57410
const REGEXP = 57411
const IN = 57412
const SHIFT_LEFT = 57413
const SHIFT_RIGHT = 57414
const UNARY = 57415
const INTERVAL = 57416
const END = 57417
const CREATE = 57418
const ALTER = 57419
const DROP = 57420
const RENAME = 57421
const ANALYZE = 57422
const TABLE = 57423
const INDEX = 57424
const VIEW = 57425
const TO = 57426
const IGNORE = 57427
const IF = 57428
const UNIQUE = 57429
const USING = 57430
const SHOW = 57431
const DESCRIBE = 57432
const EXPLAIN = 57433
const PRIMARY = 57434
const UNUSED = 57435
const SMALLINT = 57436
const INTEGER = 57437
const BIGINT = 57438
const FLOAT = 57439
const DECIMAL = 57440
const NUMERIC = 57441
const REAL = 57442
const DOUBLE = 57443
const SMALLSERIAL = 57444
const SERIAL = 57445
const BIGSERIAL = 57446
const MONEY = 57447
const CHAR_VARYING = 57448
const CHAR = 57449
const TEXT = 57450
const BYTEA = 57451
const TIMESTAMP = 57452
const DATE = 57453
const TIME = 57454
const BOOLEAN = 57455
const ENUM = 57456
const POINT = 57457
const LINE = 57458
const LSEG = 57459
const BOX = 57460
const PATH = 57461
const POLYGON = 57462
const CIRCLE = 57463
const CIDR = 57464
const INET = 57465
const MACADDR = 57466
const BIT = 57467
const BIT_VARYING = 57468
const UUID = 57469
const XML = 57470
const JSON = 57471
const JSONB = 57472

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
	"OFFSET",
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
	46, 274,
	90, 274,
	-2, 273,
	-1, 119,
	60, 169,
	61, 169,
	66, 169,
	67, 169,
	68, 169,
	69, 169,
	70, 169,
	71, 169,
	72, 169,
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
	86, 169,
	-2, 99,
	-1, 120,
	60, 170,
	61, 170,
	66, 170,
	67, 170,
	68, 170,
	69, 170,
	70, 170,
	71, 170,
	72, 170,
	74, 170,
	75, 170,
	76, 170,
	77, 170,
	78, 170,
	79, 170,
	80, 170,
	81, 170,
	82, 170,
	83, 170,
	84, 170,
	85, 170,
	86, 170,
	-2, 100,
}

const yyNprod = 278
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 838

var yyAct = [...]int{

	127, 104, 241, 292, 463, 62, 195, 198, 383, 387,
	47, 215, 230, 340, 373, 222, 282, 112, 229, 253,
	247, 136, 148, 231, 74, 197, 3, 99, 389, 67,
	243, 41, 121, 388, 154, 64, 48, 49, 69, 228,
	100, 71, 40, 35, 41, 37, 407, 409, 152, 38,
	43, 44, 45, 50, 94, 80, 122, 443, 244, 106,
	123, 124, 125, 442, 441, 126, 119, 120, 68, 156,
	108, 70, 129, 46, 128, 42, 419, 98, 444, 355,
	91, 390, 93, 122, 64, 105, 122, 64, 214, 134,
	84, 113, 114, 101, 139, 179, 169, 115, 106, 116,
	233, 476, 87, 163, 408, 182, 183, 184, 179, 159,
	194, 196, 117, 151, 153, 150, 199, 283, 283, 371,
	200, 201, 202, 203, 89, 168, 167, 218, 63, 155,
	165, 421, 238, 58, 244, 432, 208, 81, 260, 217,
	169, 122, 65, 72, 168, 167, 223, 77, 256, 225,
	219, 258, 259, 257, 143, 220, 213, 168, 167, 169,
	226, 167, 85, 472, 244, 237, 239, 88, 105, 245,
	246, 92, 169, 105, 95, 169, 122, 252, 106, 60,
	261, 262, 263, 133, 265, 266, 267, 268, 269, 270,
	271, 272, 273, 274, 242, 157, 205, 264, 158, 275,
	276, 278, 236, 60, 279, 180, 181, 182, 183, 184,
	179, 105, 64, 289, 351, 138, 14, 287, 90, 239,
	164, 248, 250, 251, 296, 290, 249, 358, 359, 360,
	75, 209, 280, 109, 137, 234, 286, 277, 244, 161,
	244, 86, 60, 145, 277, 356, 255, 28, 65, 160,
	105, 337, 338, 244, 354, 427, 122, 137, 166, 60,
	423, 219, 361, 109, 109, 363, 364, 365, 379, 244,
	338, 357, 204, 295, 244, 122, 206, 362, 374, 437,
	86, 429, 430, 368, 367, 372, 161, 140, 212, 211,
	244, 378, 109, 86, 380, 166, 223, 374, 370, 73,
	381, 384, 377, 376, 224, 79, 122, 97, 402, 400,
	234, 385, 440, 403, 401, 235, 109, 439, 399, 398,
	59, 109, 109, 109, 83, 335, 254, 14, 336, 447,
	59, 433, 255, 422, 59, 178, 177, 185, 186, 180,
	181, 182, 183, 184, 179, 76, 391, 39, 82, 59,
	146, 394, 285, 396, 59, 55, 413, 412, 59, 109,
	415, 59, 405, 395, 416, 397, 103, 96, 54, 353,
	59, 469, 135, 418, 404, 29, 346, 347, 420, 57,
	424, 132, 59, 436, 470, 59, 51, 52, 131, 293,
	235, 31, 32, 33, 34, 294, 431, 216, 109, 14,
	15, 16, 17, 177, 185, 186, 180, 181, 182, 183,
	184, 179, 254, 234, 234, 234, 234, 435, 59, 393,
	137, 61, 18, 475, 445, 461, 14, 30, 446, 28,
	1, 297, 449, 450, 384, 386, 451, 453, 109, 455,
	456, 448, 221, 219, 454, 460, 452, 144, 59, 103,
	352, 349, 162, 240, 103, 147, 462, 36, 464, 464,
	464, 227, 64, 465, 466, 149, 471, 467, 473, 474,
	66, 130, 288, 477, 210, 468, 428, 478, 382, 479,
	434, 392, 369, 207, 281, 19, 20, 22, 21, 23,
	118, 111, 103, 235, 235, 235, 235, 375, 24, 25,
	26, 110, 291, 284, 170, 107, 240, 342, 345, 346,
	347, 343, 59, 344, 348, 59, 406, 438, 341, 339,
	232, 350, 102, 59, 78, 53, 27, 56, 13, 12,
	11, 103, 328, 315, 298, 313, 310, 322, 326, 311,
	329, 327, 299, 321, 306, 305, 330, 304, 332, 309,
	331, 302, 312, 324, 318, 319, 303, 323, 325, 308,
	307, 314, 320, 300, 301, 333, 334, 316, 317, 185,
	186, 180, 181, 182, 183, 184, 179, 121, 122, 10,
	9, 106, 123, 124, 125, 8, 7, 126, 141, 142,
	109, 6, 109, 109, 129, 5, 457, 458, 459, 4,
	2, 122, 0, 0, 106, 123, 124, 125, 0, 14,
	126, 119, 120, 113, 114, 108, 0, 129, 0, 115,
	0, 116, 0, 0, 0, 121, 59, 59, 59, 59,
	0, 0, 0, 0, 117, 426, 113, 114, 101, 410,
	411, 0, 115, 414, 116, 0, 0, 0, 121, 122,
	0, 0, 106, 123, 124, 125, 0, 117, 126, 119,
	120, 0, 0, 108, 425, 129, 0, 0, 0, 0,
	0, 0, 122, 0, 14, 106, 123, 124, 125, 0,
	0, 126, 119, 120, 113, 114, 108, 0, 129, 0,
	115, 0, 116, 0, 178, 177, 185, 186, 180, 181,
	182, 183, 184, 179, 0, 117, 0, 113, 114, 0,
	0, 0, 0, 115, 122, 116, 417, 106, 123, 124,
	125, 0, 0, 126, 141, 142, 0, 0, 117, 0,
	129, 0, 0, 0, 178, 177, 185, 186, 180, 181,
	182, 183, 184, 179, 0, 0, 0, 0, 0, 113,
	114, 0, 0, 0, 0, 115, 0, 116, 0, 0,
	0, 0, 0, 0, 0, 172, 175, 65, 0, 0,
	117, 187, 188, 189, 190, 191, 192, 193, 176, 173,
	174, 171, 178, 177, 185, 186, 180, 181, 182, 183,
	184, 179, 366, 0, 0, 178, 177, 185, 186, 180,
	181, 182, 183, 184, 179, 0, 0, 0, 0, 0,
	178, 177, 185, 186, 180, 181, 182, 183, 184, 179,
	178, 177, 185, 186, 180, 181, 182, 183, 184, 179,
	342, 345, 346, 347, 343, 0, 344, 348,
}
var yyPact = [...]int{

	393, -1000, -1000, 424, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -54,
	-57, -22, -47, -24, -1000, -1000, -1000, 420, 367, 335,
	-1000, -70, 154, 411, 93, -73, -30, 93, -1000, -26,
	93, -1000, 154, -78, 181, -78, 154, -1000, -1000, -1000,
	-1000, -1000, -1000, 269, 93, -1000, 83, 323, 295, 0,
	-1000, 154, 194, -1000, 36, -1000, 154, 64, 169, -1000,
	154, -1000, -46, 154, 345, 262, 93, -1000, 555, -1000,
	371, -1000, 154, 93, 154, 409, 93, 532, 229, 328,
	-82, -1000, 20, -1000, 154, -1000, -1000, 154, -1000, 239,
	-1000, -1000, 199, 40, 99, 705, -1000, -1000, 626, 603,
	-1000, -1000, -1000, 532, 532, 532, 532, 229, -1000, -1000,
	-1000, 229, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 532,
	154, -1000, -1000, 260, 246, -2, 383, 626, -1000, 743,
	37, -1000, -1000, 668, -1000, 93, -1000, 259, 93, -1000,
	-61, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	130, 555, -1000, -1000, 93, 49, 10, 626, 626, 166,
	532, 95, 77, 532, 532, 532, 166, 532, 532, 532,
	532, 532, 532, 532, 532, 532, 532, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 23, 705, 86, 242, 190, 705,
	-1000, -1000, -1000, 718, 555, -1000, 420, 55, 743, -1000,
	321, 93, 93, 383, 154, 373, 380, 99, 129, 743,
	-1000, 226, -1000, 422, 154, -1000, -1000, 154, -1000, 223,
	795, -1000, -1000, 193, 348, 210, -1000, -1000, -1000, -1000,
	-11, -1000, 192, 555, -1000, 23, 102, -1000, -1000, 172,
	-1000, -1000, 743, -1000, 668, -1000, -1000, 95, 532, 532,
	532, 743, 743, 733, -1000, 490, 325, -1000, 22, 22,
	9, 9, 9, 124, 124, -1000, -1000, 532, -1000, -1000,
	192, 54, -1000, 626, 252, 229, 424, 233, 221, -1000,
	373, -1000, -1000, 532, 532, 93, -1000, -27, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 407, 130, 130,
	130, 130, -1000, 284, 283, -1000, 274, 273, 339, 3,
	-1000, 154, 154, -1000, 205, 129, -1000, 192, -1000, -1000,
	-1000, 190, -1000, 743, 743, 657, 532, 743, -1000, -15,
	-1000, 532, 67, -1000, 307, 213, -1000, -1000, -1000, 93,
	-1000, 617, 208, -1000, 258, -1000, -27, -1000, 80, -1000,
	304, -1000, 404, 368, 795, 234, 472, -1000, -1000, -1000,
	-1000, 282, -1000, 277, -1000, -1000, -1000, -34, -35, -41,
	-1000, -1000, -1000, -1000, -12, -1000, -1000, 532, 743, -1000,
	743, 532, 302, 229, -1000, 532, 532, 532, -1000, -1000,
	-1000, -1000, -1000, -1000, 383, 626, 532, 626, 626, -1000,
	-1000, 229, 229, 229, 93, 743, 743, 417, -1000, 743,
	743, -1000, 373, 99, 197, 99, 99, 93, 93, 93,
	-1000, 93, 354, 116, -1000, 116, 116, 194, -1000, 415,
	25, -1000, 93, -1000, -1000, -1000, 93, -1000, 93, -1000,
}
var yyPgo = [...]int{

	0, 600, 25, 599, 595, 591, 586, 585, 580, 579,
	530, 529, 528, 375, 527, 526, 525, 524, 27, 40,
	522, 18, 12, 23, 520, 519, 13, 518, 100, 516,
	4, 21, 1, 505, 504, 503, 501, 6, 20, 19,
	7, 497, 17, 74, 491, 490, 484, 16, 483, 482,
	481, 480, 11, 478, 8, 476, 3, 475, 474, 472,
	14, 5, 128, 471, 347, 299, 470, 465, 461, 457,
	455, 0, 452, 287, 451, 450, 10, 15, 447, 442,
	9, 435, 431, 430, 427, 154, 2,
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
	49, 44, 44, 44, 44, 36, 36, 36, 36, 36,
	36, 50, 50, 51, 51, 52, 52, 53, 53, 54,
	55, 55, 55, 56, 56, 56, 56, 57, 57, 57,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 82, 82, 82,
	82, 82, 82, 82, 82, 82, 82, 80, 80, 80,
	81, 81, 77, 77, 79, 79, 78, 78, 58, 58,
	59, 59, 60, 60, 35, 35, 41, 41, 42, 42,
	61, 61, 62, 63, 63, 65, 65, 66, 66, 64,
	64, 67, 67, 67, 67, 67, 67, 68, 68, 69,
	69, 70, 70, 71, 73, 85, 86, 76,
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
	2, 1, 3, 5, 7, 1, 1, 1, 1, 1,
	1, 0, 3, 0, 2, 0, 3, 1, 3, 2,
	0, 1, 1, 0, 2, 4, 4, 0, 2, 4,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 2,
	1, 2, 3, 2, 1, 3, 0, 3, 0, 3,
	1, 3, 0, 5, 2, 1, 1, 3, 3, 1,
	1, 3, 3, 1, 1, 0, 2, 0, 3, 0,
	1, 1, 1, 1, 1, 1, 1, 0, 1, 0,
	1, 0, 2, 1, 1, 1, 1, 0,
}
var yyChk = [...]int{

	-1000, -83, -1, -2, -3, -4, -5, -6, -7, -8,
	-9, -10, -11, -12, 6, 7, 8, 9, 29, 92,
	93, 95, 94, 96, 105, 106, 107, -15, 5, -13,
	-84, -13, -13, -13, -13, 97, -69, 99, 103, -64,
	99, 101, 97, 97, 98, 99, 97, -76, -76, -76,
	-2, 19, 20, -16, 33, 20, -14, -64, -28, -73,
	49, 10, -61, -62, -71, 49, -66, 102, 98, -71,
	97, -71, -28, -65, 102, 49, -65, -28, -17, 36,
	-71, 54, 25, 29, 90, -28, 47, 66, -28, 60,
	49, -76, -28, -76, 100, -28, 22, 45, -71, -18,
	-19, 83, -20, -73, -32, -37, 49, -33, 60, -85,
	-36, -44, -42, 81, 82, 87, 89, 102, -45, 56,
	57, 22, 46, 50, 51, 52, 55, -71, -43, 62,
	-63, 17, 10, -28, -61, -73, -31, 11, -62, -37,
	-73, 56, 57, -85, -78, -85, 22, -70, 104, -67,
	95, 93, 28, 94, 14, 109, 49, -28, -28, -76,
	10, 47, -72, -71, 21, 90, -85, 59, 58, 73,
	-34, 76, 60, 74, 75, 61, 73, 78, 77, 86,
	81, 82, 83, 84, 85, 79, 80, 66, 67, 68,
	69, 70, 71, 72, -32, -37, -32, -2, -40, -37,
	-37, -37, -37, -37, -85, -43, -85, -48, -37, -28,
	-58, 29, -85, -31, 90, -52, 14, -32, 90, -37,
	-76, -79, -77, -71, 45, -71, -76, -68, 100, -21,
	-22, -23, -24, -28, -43, -85, -19, -71, 83, -71,
	-73, -86, -18, 20, 48, -32, -32, -38, 55, 60,
	56, 57, -37, -39, -85, -43, 53, 76, 74, 75,
	61, -37, -37, -37, -38, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -86, -86, 47, -86, -71,
	-18, -46, -47, 63, -35, 31, -2, -61, -59, -71,
	-52, -73, -56, 16, 15, 47, -86, -82, 112, 120,
	141, 142, 129, 134, 125, 123, 122, 138, 137, 127,
	114, 117, 130, 113, 139, 111, 145, 146, 132, 133,
	140, 121, 115, 135, 131, 136, 116, 119, 110, 118,
	124, 128, 126, 143, 144, -28, -28, -31, 47, -25,
	-26, -27, 35, 39, 41, 36, 37, 38, 42, -74,
	-73, 21, -75, 21, -21, 90, -86, -18, 55, 56,
	57, -40, -39, -37, -37, -37, 59, -37, -86, -49,
	-47, 65, -32, -60, 45, -41, -42, -60, -86, 47,
	-56, -37, -53, -54, -37, -77, -81, -80, 60, 55,
	108, -76, -50, 12, -22, -23, -22, -23, 35, 35,
	35, 40, 35, 40, 35, -26, -29, 43, 101, 44,
	-73, -73, -86, -71, -73, -86, -86, 59, -37, 91,
	-37, 64, 26, 47, -71, 47, 18, 47, -55, 23,
	24, -80, 55, 27, -51, 13, 15, 45, 45, 35,
	35, 98, 98, 98, 90, -37, -37, 27, -42, -37,
	-37, -54, -52, -32, -40, -32, -32, -85, -85, -85,
	-71, 8, -56, -30, -71, -30, -30, -61, -57, 17,
	30, -86, 47, -86, -86, 8, 76, -71, -71, -71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 35, 35, 35, 35, 35, 269,
	259, 0, 0, 0, 277, 277, 277, 0, 39, 42,
	37, 259, 0, 0, 0, 257, 0, 0, 270, 0,
	0, 260, 0, 255, 0, 255, 0, 32, 33, 34,
	15, 40, 41, 44, 0, 43, 36, 0, 0, 82,
	274, 0, 20, 250, 0, 273, 0, 0, 0, 277,
	0, 277, 0, 0, 0, 0, 0, 31, 0, 45,
	0, 38, 0, 0, 0, 91, 0, 0, 236, 0,
	271, 23, 0, 26, 0, 28, 256, 0, 277, 0,
	46, 48, 53, 0, 51, 52, -2, 93, 0, 0,
	131, 132, 133, 0, 0, 0, 0, 0, 152, -2,
	-2, 0, 275, 165, 166, 167, 168, 161, 249, 154,
	0, 253, 254, 238, 91, 83, 175, 0, 251, 252,
	0, 169, 170, 0, 277, 0, 258, 0, 0, 277,
	267, 261, 262, 263, 264, 265, 266, 27, 29, 30,
	0, 0, 49, 54, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 118, 119, 120,
	121, 122, 123, 124, 96, 0, 0, 0, 0, 129,
	144, 145, 146, 0, 0, 111, 0, 0, 155, 14,
	0, 0, 0, 175, 0, 183, 0, 92, 0, 129,
	21, 0, 234, 0, 0, 272, 24, 0, 268, 91,
	56, 58, 59, 69, 67, 0, 47, 55, 50, 162,
	0, 148, 0, 0, 276, 94, 95, 98, 112, 0,
	114, 116, 101, 102, 0, 126, 127, 0, 0, 0,
	0, 104, 106, 0, 110, 134, 135, 136, 137, 138,
	139, 140, 141, 142, 143, 97, 128, 0, 248, 147,
	0, 159, 156, 0, 242, 0, 245, 242, 0, 240,
	183, 84, 19, 0, 0, 0, 237, 233, 190, 191,
	192, 193, 194, 195, 196, 197, 198, 199, 200, 201,
	202, 203, 204, 205, 206, 207, 208, 209, 210, 211,
	212, 213, 214, 215, 216, 217, 218, 219, 220, 221,
	222, 223, 224, 225, 226, 277, 25, 171, 0, 0,
	0, 0, 72, 0, 0, 75, 0, 0, 0, 85,
	70, 0, 0, 68, 0, 0, 149, 0, 113, 115,
	117, 0, 103, 105, 107, 0, 0, 130, 151, 0,
	157, 0, 0, 16, 0, 244, 246, 17, 239, 0,
	18, 184, 176, 177, 180, 235, 232, 230, 0, 228,
	0, 22, 173, 0, 57, 63, 0, 66, 73, 74,
	76, 0, 78, 0, 80, 81, 60, 0, 0, 0,
	71, 61, 62, 163, 0, 150, 125, 0, 108, 153,
	160, 0, 0, 0, 241, 0, 0, 0, 179, 181,
	182, 231, 227, 229, 175, 0, 0, 0, 0, 77,
	79, 0, 0, 0, 0, 109, 158, 0, 247, 185,
	186, 178, 183, 174, 172, 64, 65, 0, 0, 0,
	164, 0, 187, 0, 89, 0, 0, 243, 13, 0,
	0, 86, 0, 87, 88, 188, 0, 90, 0, 189,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 85, 78, 3,
	46, 48, 83, 81, 47, 82, 90, 84, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	67, 66, 68, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 86, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 77, 3, 87,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 49, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 69, 70, 71, 72, 73, 74, 75, 76, 79,
	80, 88, 89, 91, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146,
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
		//line sql.y:198
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:204
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:220
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, LimitOffset: yyDollar[12].limitOffset, Lock: yyDollar[13].str}
		}
	case 14:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:224
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:232
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 16:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:238
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:242
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
		//line sql.y:254
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, LimitOffset: yyDollar[8].limitOffset}
		}
	case 19:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:260
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, LimitOffset: yyDollar[7].limitOffset}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:266
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 21:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:272
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableName, ColDefs: yyDollar[5].colDefs}
		}
	case 22:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:276
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 23:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:281
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:287
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:291
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:296
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}, NewName: &TableName{Name: TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:302
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:308
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:316
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:321
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: &TableName{Name: TableIdent(yyDollar[4].colIdent.Lowered())}, IfExists: exists}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:331
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 32:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:337
		{
			yyVAL.statement = &Other{}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:341
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:345
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:350
		{
			setAllowComments(yylex, true)
		}
	case 36:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:354
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 37:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:360
		{
			yyVAL.bytes2 = nil
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:364
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:370
		{
			yyVAL.str = UnionStr
		}
	case 40:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:374
		{
			yyVAL.str = UnionAllStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:378
		{
			yyVAL.str = UnionDistinctStr
		}
	case 42:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:383
		{
			yyVAL.str = ""
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:387
		{
			yyVAL.str = DistinctStr
		}
	case 44:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:392
		{
			yyVAL.str = ""
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:396
		{
			yyVAL.str = StraightJoinHint
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:402
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:406
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:412
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:416
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:420
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:426
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:430
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 53:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:435
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:439
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 55:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:443
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:449
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:453
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:463
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:467
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:471
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:484
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 64:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:488
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:492
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:496
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 67:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:501
		{
			yyVAL.empty = struct{}{}
		}
	case 68:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:503
		{
			yyVAL.empty = struct{}{}
		}
	case 69:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:506
		{
			yyVAL.tableIdent = ""
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:510
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 71:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:514
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:520
		{
			yyVAL.str = JoinStr
		}
	case 73:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:524
		{
			yyVAL.str = JoinStr
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:528
		{
			yyVAL.str = JoinStr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:532
		{
			yyVAL.str = StraightJoinStr
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:538
		{
			yyVAL.str = LeftJoinStr
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:542
		{
			yyVAL.str = LeftJoinStr
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:546
		{
			yyVAL.str = RightJoinStr
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:550
		{
			yyVAL.str = RightJoinStr
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:556
		{
			yyVAL.str = NaturalJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:560
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:570
		{
			yyVAL.tableName = &TableName{Name: yyDollar[1].tableIdent}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:574
		{
			yyVAL.tableName = &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:578
		{
			yyVAL.tableName = &TableName{User: yyDollar[1].tableIdent, Qualifier: yyDollar[3].tableIdent, Name: yyDollar[5].tableIdent}
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:583
		{
			yyVAL.indexHints = nil
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:587
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:591
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:595
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:601
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:605
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:610
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:614
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:621
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:625
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:629
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:633
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:637
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:643
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:647
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:651
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:655
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:659
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:663
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:667
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:671
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:675
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:679
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:683
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:687
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:691
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:697
		{
			yyVAL.str = IsNullStr
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:701
		{
			yyVAL.str = IsNotNullStr
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:705
		{
			yyVAL.str = IsTrueStr
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:709
		{
			yyVAL.str = IsNotTrueStr
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:713
		{
			yyVAL.str = IsFalseStr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:717
		{
			yyVAL.str = IsNotFalseStr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:723
		{
			yyVAL.str = EqualStr
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:727
		{
			yyVAL.str = LessThanStr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:731
		{
			yyVAL.str = GreaterThanStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:735
		{
			yyVAL.str = LessEqualStr
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:739
		{
			yyVAL.str = GreaterEqualStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:743
		{
			yyVAL.str = NotEqualStr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:747
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:753
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:757
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:761
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:767
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:773
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:777
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:783
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:787
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:791
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 134:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:795
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:799
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:803
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:807
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:811
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:815
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:819
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:823
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:827
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:831
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:835
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:843
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
		//line sql.y:856
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 147:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:860
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:868
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 149:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:872
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 150:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:876
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:880
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 152:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:884
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:890
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 154:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:895
		{
			yyVAL.valExpr = nil
		}
	case 155:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:899
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:905
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:909
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 158:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:915
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 159:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:920
		{
			yyVAL.valExpr = nil
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:924
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:930
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 162:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:934
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Name: yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 163:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:938
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{Qualifier: yyDollar[1].tableIdent, Name: yyDollar[3].tableIdent}, Name: yyDollar[5].colIdent}
		}
	case 164:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:942
		{
			yyVAL.colName = &ColName{Qualifier: &TableName{User: yyDollar[1].tableIdent, Qualifier: yyDollar[3].tableIdent, Name: yyDollar[5].tableIdent}, Name: yyDollar[7].colIdent}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:948
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:952
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:956
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:960
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:964
		{
			yyVAL.valExpr = BoolVal(true)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:968
		{
			yyVAL.valExpr = BoolVal(false)
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:973
		{
			yyVAL.valExprs = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:977
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:982
		{
			yyVAL.boolExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:986
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:991
		{
			yyVAL.orderBy = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:995
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1001
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1005
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1011
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.str = AscScr
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1020
		{
			yyVAL.str = AscScr
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1024
		{
			yyVAL.str = DescScr
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.limitOffset = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1033
		{
			yyVAL.limitOffset = &LimitOffset{Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1037
		{
			yyVAL.limitOffset = &LimitOffset{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1041
		{
			yyVAL.limitOffset = &LimitOffset{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.str = ""
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1050
		{
			yyVAL.str = ForUpdateStr
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1054
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
	case 190:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1068
		{
			yyVAL.dataType = &DataType{Type: "bigint"}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1070
		{
			yyVAL.dataType = &DataType{Type: "bigserial"}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1072
		{
			yyVAL.dataType = &DataType{Type: "bit"}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1074
		{
			yyVAL.dataType = &DataType{Type: "bit_varying"}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1076
		{
			yyVAL.dataType = &DataType{Type: "boolean"}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1078
		{
			yyVAL.dataType = &DataType{Type: "box"}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1080
		{
			yyVAL.dataType = &DataType{Type: "bytea"}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1082
		{
			yyVAL.dataType = &DataType{Type: "char"}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1084
		{
			yyVAL.dataType = &DataType{Type: "char_varying"}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1086
		{
			yyVAL.dataType = &DataType{Type: "cidr"}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1088
		{
			yyVAL.dataType = &DataType{Type: "circle"}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1090
		{
			yyVAL.dataType = &DataType{Type: "date"}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1092
		{
			yyVAL.dataType = &DataType{Type: "decimal"}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1094
		{
			yyVAL.dataType = &DataType{Type: "double"}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1096
		{
			yyVAL.dataType = &DataType{Type: "enum"}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1098
		{
			yyVAL.dataType = &DataType{Type: "float"}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1100
		{
			yyVAL.dataType = &DataType{Type: "inet"}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1102
		{
			yyVAL.dataType = &DataType{Type: "integer"}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1106
		{
			yyVAL.dataType = &DataType{Type: "json"}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1108
		{
			yyVAL.dataType = &DataType{Type: "jsonb"}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1110
		{
			yyVAL.dataType = &DataType{Type: "line"}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1112
		{
			yyVAL.dataType = &DataType{Type: "lseg"}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1114
		{
			yyVAL.dataType = &DataType{Type: "macaddr"}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1116
		{
			yyVAL.dataType = &DataType{Type: "money"}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1118
		{
			yyVAL.dataType = &DataType{Type: "numeric"}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1120
		{
			yyVAL.dataType = &DataType{Type: "path"}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1122
		{
			yyVAL.dataType = &DataType{Type: "point"}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1124
		{
			yyVAL.dataType = &DataType{Type: "polygon"}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1126
		{
			yyVAL.dataType = &DataType{Type: "real"}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1128
		{
			yyVAL.dataType = &DataType{Type: "serial"}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1130
		{
			yyVAL.dataType = &DataType{Type: "smallint"}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1132
		{
			yyVAL.dataType = &DataType{Type: "smallserial"}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1134
		{
			yyVAL.dataType = &DataType{Type: "text"}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1136
		{
			yyVAL.dataType = &DataType{Type: "time"}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1138
		{
			yyVAL.dataType = &DataType{Type: "timestamp"}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1140
		{
			yyVAL.dataType = &DataType{Type: "uuid"}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1142
		{
			yyVAL.dataType = &DataType{Type: "xml"}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNotNullStr}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1151
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNullStr}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1155
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrPrimaryKeyStr}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1165
		{
			yyVAL.colConstrs = ColConstrs{yyDollar[1].colConstr}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1169
		{
			yyVAL.colConstrs = append(yyVAL.colConstrs, yyDollar[2].colConstr)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1176
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType, Constraints: yyDollar[3].colConstrs}
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1180
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType}
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1186
		{
			yyVAL.colDefs = ColDefs{yyDollar[1].colDef}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1190
		{
			yyVAL.colDefs = append(yyVAL.colDefs, yyDollar[3].colDef)
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1195
		{
			yyVAL.colDefs = nil
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1199
		{
			yyVAL.colDefs = yyDollar[2].colDefs
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1204
		{
			yyVAL.columns = nil
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1208
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1214
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1218
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1225
		{
			yyVAL.updateExprs = nil
		}
	case 243:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1229
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1235
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1239
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1245
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1249
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 248:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1255
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1259
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1265
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 251:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1269
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 252:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1275
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1284
		{
			yyVAL.byt = 0
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1286
		{
			yyVAL.byt = 1
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1289
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1291
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.str = ""
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1296
		{
			yyVAL.str = IgnoreStr
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1300
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1302
		{
			yyVAL.empty = struct{}{}
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1304
		{
			yyVAL.empty = struct{}{}
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1306
		{
			yyVAL.empty = struct{}{}
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1308
		{
			yyVAL.empty = struct{}{}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1310
		{
			yyVAL.empty = struct{}{}
		}
	case 267:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1313
		{
			yyVAL.empty = struct{}{}
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1315
		{
			yyVAL.empty = struct{}{}
		}
	case 269:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1318
		{
			yyVAL.empty = struct{}{}
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1320
		{
			yyVAL.empty = struct{}{}
		}
	case 271:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 272:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1329
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1335
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1341
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1350
		{
			decNesting(yylex)
		}
	case 277:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1355
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
