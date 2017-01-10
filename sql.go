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
const CAST = 57353
const WHERE = 57354
const GROUP = 57355
const HAVING = 57356
const ORDER = 57357
const BY = 57358
const LIMIT = 57359
const FOR = 57360
const OFFSET = 57361
const ALL = 57362
const DISTINCT = 57363
const AS = 57364
const EXISTS = 57365
const ASC = 57366
const DESC = 57367
const INTO = 57368
const DUPLICATE = 57369
const KEY = 57370
const DEFAULT = 57371
const SET = 57372
const LOCK = 57373
const VALUES = 57374
const LAST_INSERT_ID = 57375
const NEXT = 57376
const VALUE = 57377
const JOIN = 57378
const STRAIGHT_JOIN = 57379
const LEFT = 57380
const RIGHT = 57381
const INNER = 57382
const OUTER = 57383
const CROSS = 57384
const NATURAL = 57385
const USE = 57386
const FORCE = 57387
const ON = 57388
const ID = 57389
const STRING = 57390
const NUMBER = 57391
const VALUE_ARG = 57392
const LIST_ARG = 57393
const COMMENT = 57394
const NULL = 57395
const TRUE = 57396
const FALSE = 57397
const OR = 57398
const AND = 57399
const NOT = 57400
const BETWEEN = 57401
const CASE = 57402
const WHEN = 57403
const THEN = 57404
const ELSE = 57405
const LE = 57406
const GE = 57407
const NE = 57408
const NULL_SAFE_EQUAL = 57409
const IS = 57410
const LIKE = 57411
const REGEXP = 57412
const IN = 57413
const SHIFT_LEFT = 57414
const SHIFT_RIGHT = 57415
const UNARY = 57416
const INTERVAL = 57417
const RIGHT_ARROW = 57418
const END = 57419
const CREATE = 57420
const ALTER = 57421
const DROP = 57422
const RENAME = 57423
const ANALYZE = 57424
const TABLE = 57425
const INDEX = 57426
const VIEW = 57427
const TO = 57428
const IGNORE = 57429
const IF = 57430
const UNIQUE = 57431
const USING = 57432
const SHOW = 57433
const DESCRIBE = 57434
const EXPLAIN = 57435
const PRIMARY = 57436
const UNUSED = 57437
const SMALLINT = 57438
const INTEGER = 57439
const BIGINT = 57440
const FLOAT = 57441
const DECIMAL = 57442
const NUMERIC = 57443
const REAL = 57444
const DOUBLE = 57445
const SMALLSERIAL = 57446
const SERIAL = 57447
const BIGSERIAL = 57448
const MONEY = 57449
const CHAR_VARYING = 57450
const CHAR = 57451
const TEXT = 57452
const BYTEA = 57453
const TIMESTAMP = 57454
const DATE = 57455
const TIME = 57456
const BOOLEAN = 57457
const ENUM = 57458
const POINT = 57459
const LINE = 57460
const LSEG = 57461
const BOX = 57462
const PATH = 57463
const POLYGON = 57464
const CIRCLE = 57465
const CIDR = 57466
const INET = 57467
const MACADDR = 57468
const BIT = 57469
const BIT_VARYING = 57470
const UUID = 57471
const XML = 57472
const JSON = 57473
const JSONB = 57474

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
	"CAST",
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
	"RIGHT_ARROW",
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
	47, 274,
	91, 274,
	92, 274,
	-2, 273,
	-1, 120,
	61, 169,
	62, 169,
	67, 169,
	68, 169,
	69, 169,
	70, 169,
	71, 169,
	72, 169,
	73, 169,
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
	87, 169,
	-2, 99,
	-1, 121,
	61, 170,
	62, 170,
	67, 170,
	68, 170,
	69, 170,
	70, 170,
	71, 170,
	72, 170,
	73, 170,
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
	87, 170,
	-2, 100,
}

const yyNprod = 278
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 924

var yyAct = [...]int{

	128, 243, 62, 295, 200, 218, 197, 386, 390, 464,
	343, 300, 112, 233, 225, 376, 255, 232, 138, 47,
	286, 234, 199, 3, 392, 249, 150, 74, 99, 391,
	67, 84, 100, 156, 35, 64, 37, 40, 69, 41,
	38, 71, 94, 445, 41, 48, 49, 154, 231, 444,
	50, 443, 410, 412, 68, 80, 70, 130, 46, 42,
	113, 84, 421, 104, 43, 44, 45, 354, 158, 83,
	245, 139, 122, 215, 123, 84, 210, 98, 393, 181,
	123, 477, 87, 123, 64, 105, 136, 64, 171, 91,
	123, 93, 287, 89, 141, 60, 123, 428, 246, 106,
	124, 125, 126, 165, 434, 127, 120, 121, 81, 84,
	108, 411, 131, 153, 155, 152, 201, 65, 161, 221,
	65, 203, 204, 205, 206, 167, 427, 84, 63, 157,
	84, 114, 115, 101, 84, 169, 84, 116, 212, 117,
	187, 188, 182, 183, 184, 185, 186, 181, 226, 171,
	84, 228, 222, 118, 241, 217, 180, 179, 187, 188,
	182, 183, 184, 185, 186, 181, 223, 240, 242, 145,
	105, 229, 196, 198, 287, 105, 374, 60, 123, 254,
	208, 246, 263, 264, 265, 258, 267, 268, 269, 270,
	271, 272, 273, 274, 275, 276, 239, 244, 170, 169,
	277, 278, 280, 220, 266, 123, 123, 282, 60, 281,
	90, 284, 75, 171, 105, 140, 64, 293, 291, 166,
	237, 86, 242, 294, 473, 246, 299, 184, 185, 186,
	181, 257, 279, 247, 248, 429, 283, 290, 182, 183,
	184, 185, 186, 181, 170, 169, 358, 65, 109, 28,
	423, 340, 105, 425, 129, 246, 357, 262, 147, 171,
	377, 363, 86, 222, 439, 170, 169, 365, 366, 367,
	260, 261, 259, 168, 359, 139, 364, 377, 109, 109,
	171, 227, 97, 202, 162, 371, 369, 58, 207, 360,
	361, 362, 209, 246, 381, 139, 237, 72, 383, 226,
	14, 77, 379, 384, 387, 216, 373, 380, 73, 109,
	79, 341, 168, 388, 279, 246, 85, 257, 250, 252,
	253, 88, 163, 251, 442, 92, 163, 246, 95, 341,
	246, 86, 238, 109, 431, 432, 441, 135, 109, 109,
	109, 123, 402, 256, 60, 382, 246, 298, 246, 159,
	401, 375, 160, 449, 76, 397, 435, 399, 394, 415,
	14, 416, 408, 55, 398, 417, 400, 345, 348, 349,
	350, 346, 424, 347, 351, 419, 54, 109, 82, 405,
	403, 422, 420, 426, 406, 404, 289, 213, 180, 179,
	187, 188, 182, 183, 184, 185, 186, 181, 433, 237,
	237, 237, 237, 345, 348, 349, 350, 346, 238, 347,
	351, 470, 39, 440, 148, 109, 407, 236, 349, 350,
	96, 356, 447, 134, 471, 446, 51, 52, 296, 256,
	448, 133, 438, 297, 451, 452, 387, 453, 450, 219,
	437, 396, 454, 456, 57, 222, 61, 476, 142, 14,
	15, 16, 17, 462, 14, 28, 30, 109, 463, 1,
	465, 465, 465, 64, 389, 468, 472, 224, 474, 475,
	466, 467, 146, 18, 478, 355, 352, 164, 479, 149,
	480, 59, 338, 36, 230, 339, 151, 66, 132, 29,
	292, 59, 214, 236, 469, 59, 430, 385, 436, 395,
	372, 455, 211, 457, 458, 31, 32, 33, 34, 285,
	59, 238, 238, 238, 238, 59, 119, 111, 14, 59,
	378, 110, 59, 113, 288, 172, 370, 103, 107, 409,
	344, 59, 342, 137, 235, 122, 102, 19, 20, 22,
	21, 23, 78, 59, 53, 27, 59, 56, 13, 12,
	24, 25, 26, 11, 10, 9, 8, 7, 6, 123,
	5, 4, 106, 124, 125, 126, 2, 0, 127, 120,
	121, 0, 0, 108, 0, 131, 0, 0, 0, 0,
	0, 59, 180, 179, 187, 188, 182, 183, 184, 185,
	186, 181, 0, 0, 114, 115, 236, 236, 236, 236,
	116, 0, 117, 0, 0, 0, 0, 109, 0, 109,
	109, 59, 103, 459, 460, 461, 118, 103, 331, 318,
	301, 316, 313, 325, 329, 314, 332, 330, 302, 324,
	309, 308, 333, 307, 335, 312, 334, 305, 315, 327,
	321, 322, 306, 326, 328, 311, 310, 317, 323, 303,
	304, 336, 337, 319, 320, 0, 103, 113, 180, 179,
	187, 188, 182, 183, 184, 185, 186, 181, 0, 122,
	0, 0, 113, 0, 0, 0, 59, 0, 0, 59,
	0, 0, 0, 0, 122, 353, 0, 59, 0, 0,
	0, 0, 0, 123, 103, 0, 106, 124, 125, 126,
	0, 0, 127, 120, 121, 0, 0, 108, 123, 131,
	0, 106, 124, 125, 126, 0, 0, 127, 120, 121,
	0, 0, 108, 0, 131, 0, 0, 0, 114, 115,
	101, 0, 14, 0, 116, 0, 117, 113, 0, 0,
	0, 0, 0, 114, 115, 0, 0, 0, 0, 116,
	118, 117, 0, 0, 0, 0, 0, 0, 0, 0,
	113, 0, 0, 0, 0, 118, 0, 0, 0, 0,
	0, 0, 0, 123, 0, 0, 106, 124, 125, 126,
	0, 0, 127, 143, 144, 0, 0, 0, 0, 131,
	59, 59, 59, 59, 0, 0, 123, 0, 0, 106,
	124, 125, 126, 413, 414, 127, 143, 144, 114, 115,
	0, 0, 131, 0, 116, 0, 117, 179, 187, 188,
	182, 183, 184, 185, 186, 181, 0, 0, 0, 0,
	118, 114, 115, 0, 0, 0, 0, 116, 0, 117,
	0, 0, 0, 0, 0, 0, 0, 0, 174, 177,
	0, 0, 0, 118, 189, 190, 191, 192, 193, 194,
	195, 178, 175, 176, 173, 180, 179, 187, 188, 182,
	183, 184, 185, 186, 181, 418, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 368, 65, 0, 0, 0,
	0, 0, 0, 180, 179, 187, 188, 182, 183, 184,
	185, 186, 181, 180, 179, 187, 188, 182, 183, 184,
	185, 186, 181, 0, 180, 179, 187, 188, 182, 183,
	184, 185, 186, 181,
}
var yyPact = [...]int{

	443, -1000, -1000, 450, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -65,
	-64, -40, -35, -41, -1000, -1000, -1000, 448, 406, 342,
	-1000, -59, 127, 436, 67, -74, -46, 67, -1000, -43,
	67, -1000, 127, -77, 162, -77, 127, -1000, -1000, -1000,
	-1000, -1000, -1000, 273, 67, -1000, 53, 352, 39, -1000,
	-1000, 127, 173, -1000, 15, -1000, 127, 32, 160, -1000,
	127, -1000, -60, 127, 397, 236, 67, -30, 646, -1000,
	413, -1000, 127, 67, 127, 59, 67, 749, 36, 391,
	-80, -1000, 18, -1000, 127, -30, -1000, 127, -1000, 274,
	-1000, -1000, 197, 33, 139, 787, -1000, -1000, 661, 512,
	-1000, -1000, -1000, 159, 749, 749, 749, 749, 159, -1000,
	-1000, -1000, 159, -1000, -1000, -1000, -1000, -1000, -1000, -16,
	-1000, 749, 127, -1000, -1000, 43, 283, -1000, 424, 661,
	-1000, 580, 27, -1000, -1000, 726, -1000, 67, -1000, 235,
	67, -1000, -54, -1000, -1000, -1000, -1000, -1000, -1000, -30,
	-30, -1000, 158, 646, -1000, -1000, 67, 70, 49, 661,
	661, 262, 749, 131, 195, 749, 749, 749, 262, 749,
	749, 749, 749, 749, 749, 749, 749, 749, 749, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 14, 787, 206, 244,
	266, 787, 749, -1000, -1000, -1000, 836, 646, -1000, 448,
	67, 28, 580, -30, 354, 67, 67, 424, 411, 417,
	139, 67, 580, -1000, 299, -1000, 506, 127, -1000, -1000,
	127, -1000, 263, 331, -1000, -1000, 45, 399, 294, -1000,
	-1000, -1000, -1000, -1000, 278, 646, -1000, 14, 75, -1000,
	-1000, 233, -1000, -1000, 580, -1000, 726, -1000, -1000, 131,
	749, 749, 749, 580, 580, 825, -1000, 60, 738, -1000,
	143, 143, -8, -8, -8, 156, 156, -1000, -1000, 749,
	-1000, 504, -1000, 278, -1000, 110, -1000, 661, 231, 159,
	450, 214, 297, -1000, 411, -1000, 749, 749, 67, -1000,
	-32, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -30, -30,
	428, 158, 158, 158, 158, -1000, 314, 306, -1000, 344,
	343, 380, 8, -1000, 127, 127, -1000, 281, -1000, 278,
	-1000, -1000, -1000, 266, -1000, 580, 580, 815, 749, 580,
	506, -1000, -31, -1000, 749, 185, -1000, 345, 205, -1000,
	-1000, -1000, 67, -1000, 78, 187, -1000, 310, -1000, -32,
	-1000, 48, -1000, 328, -1000, 426, 416, 331, 218, 367,
	-1000, -1000, -1000, -1000, 300, -1000, 288, -1000, -1000, -1000,
	-49, -51, -57, -1000, -1000, -1000, -1000, -1000, 749, 580,
	132, -1000, 580, 749, 325, 159, -1000, 749, 749, 749,
	-1000, -1000, -1000, -1000, -1000, -1000, 424, 661, 749, 661,
	661, -1000, -1000, 159, 159, 159, 580, -1000, 580, 445,
	-1000, 580, 580, -1000, 411, 139, 184, 139, 139, 67,
	67, 67, 67, 393, 176, -1000, 176, 176, 173, -1000,
	439, 4, -1000, 67, -1000, -1000, -1000, 67, -1000, 67,
	-1000,
}
var yyPgo = [...]int{

	0, 566, 22, 561, 560, 558, 557, 556, 555, 554,
	553, 549, 548, 489, 547, 545, 544, 542, 28, 32,
	536, 17, 13, 21, 534, 532, 10, 530, 254, 529,
	9, 18, 63, 528, 525, 524, 521, 6, 25, 16,
	4, 520, 12, 57, 517, 516, 509, 20, 502, 500,
	499, 498, 5, 497, 7, 496, 3, 494, 492, 490,
	15, 2, 128, 488, 412, 308, 487, 486, 484, 483,
	479, 0, 477, 448, 476, 475, 19, 14, 472, 467,
	8, 464, 11, 459, 456, 169, 1,
}
var yyR1 = [...]int{

	0, 83, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 2, 3, 3, 4,
	5, 6, 7, 7, 7, 8, 8, 8, 9, 10,
	10, 10, 11, 12, 12, 12, 84, 13, 14, 14,
	15, 15, 15, 16, 16, 17, 17, 18, 18, 19,
	19, 19, 20, 20, 72, 72, 72, 21, 21, 22,
	22, 23, 23, 23, 24, 24, 24, 24, 75, 75,
	74, 74, 74, 25, 25, 25, 25, 26, 26, 26,
	26, 27, 27, 28, 28, 29, 29, 29, 29, 30,
	30, 31, 31, 32, 32, 32, 32, 32, 32, 33,
	33, 33, 33, 33, 33, 33, 33, 33, 33, 33,
	33, 33, 38, 38, 38, 38, 38, 38, 34, 34,
	34, 34, 34, 34, 34, 39, 39, 39, 43, 40,
	40, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 37, 37, 37, 37, 37, 37,
	37, 37, 37, 37, 45, 48, 48, 46, 46, 47,
	49, 49, 44, 44, 44, 36, 36, 36, 36, 36,
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
	1, 1, 1, 5, 13, 6, 3, 8, 8, 8,
	7, 3, 6, 8, 4, 6, 7, 4, 5, 4,
	5, 5, 3, 2, 2, 2, 0, 2, 0, 2,
	1, 2, 2, 0, 1, 0, 1, 1, 3, 1,
	2, 3, 1, 1, 0, 1, 2, 1, 3, 1,
	1, 3, 3, 3, 3, 5, 5, 3, 0, 1,
	0, 1, 2, 1, 2, 2, 1, 2, 3, 2,
	3, 2, 2, 1, 3, 0, 5, 5, 5, 1,
	3, 0, 2, 1, 3, 3, 2, 3, 3, 1,
	1, 3, 3, 4, 3, 4, 3, 4, 5, 6,
	3, 2, 1, 2, 1, 2, 1, 2, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 1, 3, 1,
	3, 1, 1, 1, 6, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 2, 2, 2, 3, 3,
	4, 5, 4, 1, 5, 0, 1, 1, 2, 4,
	0, 2, 1, 3, 3, 1, 1, 1, 1, 1,
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
	-9, -10, -11, -12, 6, 7, 8, 9, 30, 94,
	95, 97, 96, 98, 107, 108, 109, -15, 5, -13,
	-84, -13, -13, -13, -13, 99, -69, 101, 105, -64,
	101, 103, 99, 99, 100, 101, 99, -76, -76, -76,
	-2, 20, 21, -16, 34, 21, -14, -64, -28, -73,
	50, 10, -61, -62, -71, 50, -66, 104, 100, -71,
	99, -71, -28, -65, 104, 50, -65, -28, -17, 37,
	-71, 55, 26, 30, 91, -28, 48, 67, -28, 61,
	50, -76, -28, -76, 102, -28, 23, 46, -71, -18,
	-19, 84, -20, -73, -32, -37, 50, -33, 61, -85,
	-36, -44, -42, 11, 82, 83, 88, 90, 104, -45,
	57, 58, 23, 47, 51, 52, 53, 56, -71, -28,
	-43, 63, -63, 18, 10, -28, -61, -73, -31, 12,
	-62, -37, -73, 57, 58, -85, -78, -85, 23, -70,
	106, -67, 97, 95, 29, 96, 15, 111, 50, -28,
	-28, -76, 10, 48, -72, -71, 22, 92, -85, 60,
	59, 74, -34, 77, 61, 75, 76, 62, 74, 79,
	78, 87, 82, 83, 84, 85, 86, 80, 81, 67,
	68, 69, 70, 71, 72, 73, -32, -37, -32, -2,
	-40, -37, -85, -37, -37, -37, -37, -85, -43, -85,
	92, -48, -37, -28, -58, 30, -85, -31, -52, 15,
	-32, 92, -37, -76, -79, -77, -71, 46, -71, -76,
	-68, 102, -21, -22, -23, -24, -28, -43, -85, -19,
	-71, 84, -71, -86, -18, 21, 49, -32, -32, -38,
	56, 61, 57, 58, -37, -39, -85, -43, 54, 77,
	75, 76, 62, -37, -37, -37, -38, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -86, -86, 48,
	-86, -37, -71, -18, -71, -46, -47, 64, -35, 32,
	-2, -61, -59, -71, -52, -56, 17, 16, 48, -86,
	-82, 114, 122, 143, 144, 131, 136, 127, 125, 124,
	140, 139, 129, 116, 119, 132, 115, 141, 113, 147,
	148, 134, 135, 142, 123, 117, 137, 133, 138, 118,
	121, 112, 120, 126, 130, 128, 145, 146, -28, -28,
	-31, 48, -25, -26, -27, 36, 40, 42, 37, 38,
	39, 43, -74, -73, 22, -75, 22, -21, -86, -18,
	56, 57, 58, -40, -39, -37, -37, -37, 60, -37,
	22, -86, -49, -47, 66, -32, -60, 46, -41, -42,
	-60, -86, 48, -56, -37, -53, -54, -37, -77, -81,
	-80, 61, 56, 110, -76, -50, 13, -22, -23, -22,
	-23, 36, 36, 36, 41, 36, 41, 36, -26, -29,
	44, 103, 45, -73, -73, -86, -86, -86, 60, -37,
	-82, 93, -37, 65, 27, 48, -71, 48, 19, 48,
	-55, 24, 25, -80, 56, 28, -51, 14, 16, 46,
	46, 36, 36, 100, 100, 100, -37, -86, -37, 28,
	-42, -37, -37, -54, -52, -32, -40, -32, -32, -85,
	-85, -85, 8, -56, -30, -71, -30, -30, -61, -57,
	18, 31, -86, 48, -86, -86, 8, 77, -71, -71,
	-71,
}
var yyDef = [...]int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 8,
	9, 10, 11, 12, 36, 36, 36, 36, 36, 269,
	259, 0, 0, 0, 277, 277, 277, 0, 40, 43,
	38, 259, 0, 0, 0, 257, 0, 0, 270, 0,
	0, 260, 0, 255, 0, 255, 0, 33, 34, 35,
	16, 41, 42, 45, 0, 44, 37, 0, 0, 83,
	274, 0, 21, 250, 0, 273, 0, 0, 0, 277,
	0, 277, 0, 0, 0, 0, 0, 32, 0, 46,
	0, 39, 0, 0, 0, 91, 0, 0, 236, 0,
	271, 24, 0, 27, 0, 29, 256, 0, 277, 13,
	47, 49, 54, 83, 52, 53, -2, 93, 0, 0,
	131, 132, 133, 0, 0, 0, 0, 0, 0, 153,
	-2, -2, 0, 275, 165, 166, 167, 168, 162, 0,
	249, 155, 0, 253, 254, 238, 91, 84, 175, 0,
	251, 252, 83, 169, 170, 0, 277, 0, 258, 0,
	0, 277, 267, 261, 262, 263, 264, 265, 266, 28,
	30, 31, 0, 0, 50, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 118,
	119, 120, 121, 122, 123, 124, 96, 0, 0, 0,
	0, 129, 0, 145, 146, 147, 0, 0, 111, 0,
	0, 0, 156, 15, 0, 0, 0, 175, 183, 0,
	92, 0, 129, 22, 0, 234, 0, 0, 272, 25,
	0, 268, 91, 57, 59, 60, 70, 68, 0, 48,
	56, 51, 163, 149, 0, 0, 276, 94, 95, 98,
	112, 0, 114, 116, 101, 102, 0, 126, 127, 0,
	0, 0, 0, 104, 106, 0, 110, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 97, 128, 0,
	248, 0, 148, 0, 164, 160, 157, 0, 242, 0,
	245, 242, 0, 240, 183, 20, 0, 0, 0, 237,
	233, 190, 191, 192, 193, 194, 195, 196, 197, 198,
	199, 200, 201, 202, 203, 204, 205, 206, 207, 208,
	209, 210, 211, 212, 213, 214, 215, 216, 217, 218,
	219, 220, 221, 222, 223, 224, 225, 226, 277, 26,
	171, 0, 0, 0, 0, 73, 0, 0, 76, 0,
	0, 0, 85, 71, 0, 0, 69, 0, 150, 0,
	113, 115, 117, 0, 103, 105, 107, 0, 0, 130,
	0, 152, 0, 158, 0, 0, 17, 0, 244, 246,
	18, 239, 0, 19, 184, 176, 177, 180, 235, 232,
	230, 0, 228, 0, 23, 173, 0, 58, 64, 0,
	67, 74, 75, 77, 0, 79, 0, 81, 82, 61,
	0, 0, 0, 72, 62, 63, 151, 125, 0, 108,
	0, 154, 161, 0, 0, 0, 241, 0, 0, 0,
	179, 181, 182, 231, 227, 229, 175, 0, 0, 0,
	0, 78, 80, 0, 0, 0, 109, 134, 159, 0,
	247, 185, 186, 178, 183, 174, 172, 65, 66, 0,
	0, 0, 0, 187, 0, 89, 0, 0, 243, 14,
	0, 0, 86, 0, 87, 88, 188, 0, 90, 0,
	189,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 86, 79, 3,
	47, 49, 84, 82, 48, 83, 91, 85, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	68, 67, 69, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 87, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 78, 3, 88,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 50, 51, 52, 53, 54,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 66, 70, 71, 72, 73, 74, 75, 76, 77,
	80, 81, 89, 90, 92, 93, 94, 95, 96, 97,
	98, 99, 100, 101, 102, 103, 104, 105, 106, 107,
	108, 109, 110, 111, 112, 113, 114, 115, 116, 117,
	118, 119, 120, 121, 122, 123, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 137,
	138, 139, 140, 141, 142, 143, 144, 145, 146, 147,
	148,
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
		//line sql.y:199
		{
			setParseTree(yylex, yyDollar[1].statement)
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:205
		{
			yyVAL.statement = yyDollar[1].selStmt
		}
	case 13:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:221
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs}
		}
	case 14:
		yyDollar = yyS[yypt-13 : yypt+1]
		//line sql.y:225
		{
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), Distinct: yyDollar[3].str, Hints: yyDollar[4].str, SelectExprs: yyDollar[5].selectExprs, From: yyDollar[7].tableExprs, Where: NewWhere(WhereStr, yyDollar[8].boolExpr), GroupBy: GroupBy(yyDollar[9].valExprs), Having: NewWhere(HavingStr, yyDollar[10].boolExpr), OrderBy: yyDollar[11].orderBy, LimitOffset: yyDollar[12].limitOffset, Lock: yyDollar[13].str}
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:229
		{
			if yyDollar[4].colIdent.Lowered() != "value" {
				yylex.Error("expecting value after next")
				return 1
			}
			yyVAL.selStmt = &Select{Comments: Comments(yyDollar[2].bytes2), SelectExprs: SelectExprs{Nextval{}}, From: TableExprs{&AliasedTableExpr{Expr: yyDollar[6].tableName}}}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:237
		{
			yyVAL.selStmt = &Union{Type: yyDollar[2].str, Left: yyDollar[1].selStmt, Right: yyDollar[3].selStmt}
		}
	case 17:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:243
		{
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: yyDollar[6].columns, Rows: yyDollar[7].insRows, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 18:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:247
		{
			cols := make(Columns, 0, len(yyDollar[7].updateExprs))
			vals := make(ValTuple, 0, len(yyDollar[7].updateExprs))
			for _, updateList := range yyDollar[7].updateExprs {
				cols = append(cols, updateList.Name)
				vals = append(vals, updateList.Expr)
			}
			yyVAL.statement = &Insert{Comments: Comments(yyDollar[2].bytes2), Ignore: yyDollar[3].str, Table: yyDollar[5].tableName, Columns: cols, Rows: Values{vals}, OnDup: OnDup(yyDollar[8].updateExprs)}
		}
	case 19:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:259
		{
			yyVAL.statement = &Update{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[3].tableName, Exprs: yyDollar[5].updateExprs, Where: NewWhere(WhereStr, yyDollar[6].boolExpr), OrderBy: yyDollar[7].orderBy, LimitOffset: yyDollar[8].limitOffset}
		}
	case 20:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:265
		{
			yyVAL.statement = &Delete{Comments: Comments(yyDollar[2].bytes2), Table: yyDollar[4].tableName, Where: NewWhere(WhereStr, yyDollar[5].boolExpr), OrderBy: yyDollar[6].orderBy, LimitOffset: yyDollar[7].limitOffset}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:271
		{
			yyVAL.statement = &Set{Comments: Comments(yyDollar[2].bytes2), Exprs: yyDollar[3].updateExprs}
		}
	case 22:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:277
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: yyDollar[4].tableName, ColDefs: yyDollar[5].colDefs}
		}
	case 23:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line sql.y:281
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[7].tableName, NewName: yyDollar[7].tableName}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:286
		{
			yyVAL.statement = &DDL{Action: CreateStr, NewName: TableName{TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 25:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:292
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[4].tableName, NewName: yyDollar[4].tableName}
		}
	case 26:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line sql.y:296
		{
			// Change this to a rename statement
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[4].tableName, NewName: yyDollar[7].tableName}
		}
	case 27:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:301
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: TableName{TableIdent(yyDollar[3].colIdent.Lowered())}, NewName: TableName{TableIdent(yyDollar[3].colIdent.Lowered())}}
		}
	case 28:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:307
		{
			yyVAL.statement = &DDL{Action: RenameStr, Table: yyDollar[3].tableName, NewName: yyDollar[5].tableName}
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:313
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: yyDollar[4].tableName, IfExists: exists}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:321
		{
			// Change this to an alter statement
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[5].tableName, NewName: yyDollar[5].tableName}
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:326
		{
			var exists bool
			if yyDollar[3].byt != 0 {
				exists = true
			}
			yyVAL.statement = &DDL{Action: DropStr, Table: TableName{TableIdent(yyDollar[4].colIdent.Lowered())}, IfExists: exists}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:336
		{
			yyVAL.statement = &DDL{Action: AlterStr, Table: yyDollar[3].tableName, NewName: yyDollar[3].tableName}
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:342
		{
			yyVAL.statement = &Other{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:346
		{
			yyVAL.statement = &Other{}
		}
	case 35:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:350
		{
			yyVAL.statement = &Other{}
		}
	case 36:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:355
		{
			setAllowComments(yylex, true)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:359
		{
			yyVAL.bytes2 = yyDollar[2].bytes2
			setAllowComments(yylex, false)
		}
	case 38:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:365
		{
			yyVAL.bytes2 = nil
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:369
		{
			yyVAL.bytes2 = append(yyDollar[1].bytes2, yyDollar[2].bytes)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:375
		{
			yyVAL.str = UnionStr
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:379
		{
			yyVAL.str = UnionAllStr
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:383
		{
			yyVAL.str = UnionDistinctStr
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:388
		{
			yyVAL.str = ""
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:392
		{
			yyVAL.str = DistinctStr
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:397
		{
			yyVAL.str = ""
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:401
		{
			yyVAL.str = StraightJoinHint
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:407
		{
			yyVAL.selectExprs = SelectExprs{yyDollar[1].selectExpr}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:411
		{
			yyVAL.selectExprs = append(yyVAL.selectExprs, yyDollar[3].selectExpr)
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:417
		{
			yyVAL.selectExpr = &StarExpr{}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:421
		{
			yyVAL.selectExpr = &NonStarExpr{Expr: yyDollar[1].expr, As: yyDollar[2].colIdent}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:425
		{
			yyVAL.selectExpr = &StarExpr{TableName: yyDollar[1].tableIdent}
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:431
		{
			yyVAL.expr = yyDollar[1].boolExpr
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:435
		{
			yyVAL.expr = yyDollar[1].valExpr
		}
	case 54:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:440
		{
			yyVAL.colIdent = ColIdent{}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:444
		{
			yyVAL.colIdent = yyDollar[1].colIdent
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:448
		{
			yyVAL.colIdent = yyDollar[2].colIdent
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:454
		{
			yyVAL.tableExprs = TableExprs{yyDollar[1].tableExpr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:458
		{
			yyVAL.tableExprs = append(yyVAL.tableExprs, yyDollar[3].tableExpr)
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:468
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].tableName, As: yyDollar[2].tableIdent, Hints: yyDollar[3].indexHints}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:472
		{
			yyVAL.tableExpr = &AliasedTableExpr{Expr: yyDollar[1].subquery, As: yyDollar[3].tableIdent}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:476
		{
			yyVAL.tableExpr = &ParenTableExpr{Exprs: yyDollar[2].tableExprs}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:489
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:493
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 66:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:497
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr, On: yyDollar[5].boolExpr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:501
		{
			yyVAL.tableExpr = &JoinTableExpr{LeftExpr: yyDollar[1].tableExpr, Join: yyDollar[2].str, RightExpr: yyDollar[3].tableExpr}
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:506
		{
			yyVAL.empty = struct{}{}
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:508
		{
			yyVAL.empty = struct{}{}
		}
	case 70:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:511
		{
			yyVAL.tableIdent = ""
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:515
		{
			yyVAL.tableIdent = yyDollar[1].tableIdent
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:519
		{
			yyVAL.tableIdent = yyDollar[2].tableIdent
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:525
		{
			yyVAL.str = JoinStr
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:529
		{
			yyVAL.str = JoinStr
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:533
		{
			yyVAL.str = JoinStr
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:537
		{
			yyVAL.str = StraightJoinStr
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:543
		{
			yyVAL.str = LeftJoinStr
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:547
		{
			yyVAL.str = LeftJoinStr
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:551
		{
			yyVAL.str = RightJoinStr
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:555
		{
			yyVAL.str = RightJoinStr
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:561
		{
			yyVAL.str = NaturalJoinStr
		}
	case 82:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:565
		{
			if yyDollar[2].str == LeftJoinStr {
				yyVAL.str = NaturalLeftJoinStr
			} else {
				yyVAL.str = NaturalRightJoinStr
			}
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:575
		{
			yyVAL.tableName = TableName{yyDollar[1].tableIdent}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:579
		{
			yyVAL.tableName = append(yyDollar[1].tableName, yyDollar[3].tableIdent)
		}
	case 85:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:584
		{
			yyVAL.indexHints = nil
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:588
		{
			yyVAL.indexHints = &IndexHints{Type: UseStr, Indexes: yyDollar[4].colIdents}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:592
		{
			yyVAL.indexHints = &IndexHints{Type: IgnoreStr, Indexes: yyDollar[4].colIdents}
		}
	case 88:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:596
		{
			yyVAL.indexHints = &IndexHints{Type: ForceStr, Indexes: yyDollar[4].colIdents}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:602
		{
			yyVAL.colIdents = []ColIdent{yyDollar[1].colIdent}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:606
		{
			yyVAL.colIdents = append(yyDollar[1].colIdents, yyDollar[3].colIdent)
		}
	case 91:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:611
		{
			yyVAL.boolExpr = nil
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:615
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:622
		{
			yyVAL.boolExpr = &AndExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:626
		{
			yyVAL.boolExpr = &OrExpr{Left: yyDollar[1].boolExpr, Right: yyDollar[3].boolExpr}
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:630
		{
			yyVAL.boolExpr = &NotExpr{Expr: yyDollar[2].boolExpr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:634
		{
			yyVAL.boolExpr = &ParenBoolExpr{Expr: yyDollar[2].boolExpr}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:638
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].boolExpr}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:644
		{
			yyVAL.boolExpr = BoolVal(true)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:648
		{
			yyVAL.boolExpr = BoolVal(false)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:652
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: yyDollar[2].str, Right: yyDollar[3].valExpr}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:656
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: InStr, Right: yyDollar[3].colTuple}
		}
	case 103:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:660
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotInStr, Right: yyDollar[4].colTuple}
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:664
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: LikeStr, Right: yyDollar[3].valExpr}
		}
	case 105:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:668
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotLikeStr, Right: yyDollar[4].valExpr}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:672
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: RegexpStr, Right: yyDollar[3].valExpr}
		}
	case 107:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:676
		{
			yyVAL.boolExpr = &ComparisonExpr{Left: yyDollar[1].valExpr, Operator: NotRegexpStr, Right: yyDollar[4].valExpr}
		}
	case 108:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:680
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: BetweenStr, From: yyDollar[3].valExpr, To: yyDollar[5].valExpr}
		}
	case 109:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:684
		{
			yyVAL.boolExpr = &RangeCond{Left: yyDollar[1].valExpr, Operator: NotBetweenStr, From: yyDollar[4].valExpr, To: yyDollar[6].valExpr}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:688
		{
			yyVAL.boolExpr = &IsExpr{Operator: yyDollar[3].str, Expr: yyDollar[1].valExpr}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:692
		{
			yyVAL.boolExpr = &ExistsExpr{Subquery: yyDollar[2].subquery}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:698
		{
			yyVAL.str = IsNullStr
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:702
		{
			yyVAL.str = IsNotNullStr
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:706
		{
			yyVAL.str = IsTrueStr
		}
	case 115:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:710
		{
			yyVAL.str = IsNotTrueStr
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:714
		{
			yyVAL.str = IsFalseStr
		}
	case 117:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:718
		{
			yyVAL.str = IsNotFalseStr
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:724
		{
			yyVAL.str = EqualStr
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:728
		{
			yyVAL.str = LessThanStr
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:732
		{
			yyVAL.str = GreaterThanStr
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:736
		{
			yyVAL.str = LessEqualStr
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:740
		{
			yyVAL.str = GreaterEqualStr
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:744
		{
			yyVAL.str = NotEqualStr
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:748
		{
			yyVAL.str = NullSafeEqualStr
		}
	case 125:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:754
		{
			yyVAL.colTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:758
		{
			yyVAL.colTuple = yyDollar[1].subquery
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:762
		{
			yyVAL.colTuple = ListArg(yyDollar[1].bytes)
		}
	case 128:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:768
		{
			yyVAL.subquery = &Subquery{yyDollar[2].selStmt}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:774
		{
			yyVAL.valExprs = ValExprs{yyDollar[1].valExpr}
		}
	case 130:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:778
		{
			yyVAL.valExprs = append(yyDollar[1].valExprs, yyDollar[3].valExpr)
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:784
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:788
		{
			yyVAL.valExpr = yyDollar[1].colName
		}
	case 133:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:796
		{
			yyVAL.valExpr = yyDollar[1].rowTuple
		}
	case 134:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line sql.y:800
		{
			yyVAL.valExpr = &CastValExpr{Val: yyDollar[3].valExpr, Type: yyDollar[5].dataType}
		}
	case 135:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:804
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitAndStr, Right: yyDollar[3].valExpr}
		}
	case 136:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:808
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitOrStr, Right: yyDollar[3].valExpr}
		}
	case 137:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:812
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: BitXorStr, Right: yyDollar[3].valExpr}
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:816
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: PlusStr, Right: yyDollar[3].valExpr}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:820
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MinusStr, Right: yyDollar[3].valExpr}
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:824
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: MultStr, Right: yyDollar[3].valExpr}
		}
	case 141:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:828
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: DivStr, Right: yyDollar[3].valExpr}
		}
	case 142:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:832
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ModStr, Right: yyDollar[3].valExpr}
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:836
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftLeftStr, Right: yyDollar[3].valExpr}
		}
	case 144:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:840
		{
			yyVAL.valExpr = &BinaryExpr{Left: yyDollar[1].valExpr, Operator: ShiftRightStr, Right: yyDollar[3].valExpr}
		}
	case 145:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:844
		{
			if num, ok := yyDollar[2].valExpr.(NumVal); ok {
				yyVAL.valExpr = num
			} else {
				yyVAL.valExpr = &UnaryExpr{Operator: UPlusStr, Expr: yyDollar[2].valExpr}
			}
		}
	case 146:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:852
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
	case 147:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:865
		{
			yyVAL.valExpr = &UnaryExpr{Operator: TildaStr, Expr: yyDollar[2].valExpr}
		}
	case 148:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:869
		{
			// This rule prevents the usage of INTERVAL
			// as a function. If support is needed for that,
			// we'll need to revisit this. The solution
			// will be non-trivial because of grammar conflicts.
			yyVAL.valExpr = &IntervalExpr{Expr: yyDollar[2].valExpr, Unit: yyDollar[3].colIdent}
		}
	case 149:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:877
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent)}
		}
	case 150:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:881
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Exprs: yyDollar[3].selectExprs}
		}
	case 151:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:885
		{
			yyVAL.valExpr = &FuncExpr{Name: string(yyDollar[1].tableIdent), Distinct: true, Exprs: yyDollar[4].selectExprs}
		}
	case 152:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:889
		{
			yyVAL.valExpr = &FuncExpr{Name: "if", Exprs: yyDollar[3].selectExprs}
		}
	case 153:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:893
		{
			yyVAL.valExpr = yyDollar[1].caseExpr
		}
	case 154:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:899
		{
			yyVAL.caseExpr = &CaseExpr{Expr: yyDollar[2].valExpr, Whens: yyDollar[3].whens, Else: yyDollar[4].valExpr}
		}
	case 155:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:904
		{
			yyVAL.valExpr = nil
		}
	case 156:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:908
		{
			yyVAL.valExpr = yyDollar[1].valExpr
		}
	case 157:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:914
		{
			yyVAL.whens = []*When{yyDollar[1].when}
		}
	case 158:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:918
		{
			yyVAL.whens = append(yyDollar[1].whens, yyDollar[2].when)
		}
	case 159:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:924
		{
			yyVAL.when = &When{Cond: yyDollar[2].boolExpr, Val: yyDollar[4].valExpr}
		}
	case 160:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:929
		{
			yyVAL.valExpr = nil
		}
	case 161:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:933
		{
			yyVAL.valExpr = yyDollar[2].valExpr
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:939
		{
			yyVAL.colName = &ColName{Name: yyDollar[1].colIdent}
		}
	case 163:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:943
		{
			yyVAL.colName = &ColName{Qualifier: TableName{yyDollar[1].tableIdent}, Name: yyDollar[3].colIdent}
		}
	case 164:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:947
		{
			yyVAL.colName = &ColName{Qualifier: yyDollar[1].tableName, Name: yyDollar[3].colIdent}
		}
	case 165:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:953
		{
			yyVAL.valExpr = StrVal(yyDollar[1].bytes)
		}
	case 166:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:957
		{
			yyVAL.valExpr = NumVal(yyDollar[1].bytes)
		}
	case 167:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:961
		{
			yyVAL.valExpr = ValArg(yyDollar[1].bytes)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:965
		{
			yyVAL.valExpr = &NullVal{}
		}
	case 169:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:969
		{
			yyVAL.valExpr = BoolVal(true)
		}
	case 170:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:973
		{
			yyVAL.valExpr = BoolVal(false)
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:978
		{
			yyVAL.valExprs = nil
		}
	case 172:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:982
		{
			yyVAL.valExprs = yyDollar[3].valExprs
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:987
		{
			yyVAL.boolExpr = nil
		}
	case 174:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:991
		{
			yyVAL.boolExpr = yyDollar[2].boolExpr
		}
	case 175:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:996
		{
			yyVAL.orderBy = nil
		}
	case 176:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1000
		{
			yyVAL.orderBy = yyDollar[3].orderBy
		}
	case 177:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1006
		{
			yyVAL.orderBy = OrderBy{yyDollar[1].order}
		}
	case 178:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1010
		{
			yyVAL.orderBy = append(yyDollar[1].orderBy, yyDollar[3].order)
		}
	case 179:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1016
		{
			yyVAL.order = &Order{Expr: yyDollar[1].valExpr, Direction: yyDollar[2].str}
		}
	case 180:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1021
		{
			yyVAL.str = AscScr
		}
	case 181:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1025
		{
			yyVAL.str = AscScr
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1029
		{
			yyVAL.str = DescScr
		}
	case 183:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1034
		{
			yyVAL.limitOffset = nil
		}
	case 184:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1038
		{
			yyVAL.limitOffset = &LimitOffset{Rowcount: yyDollar[2].valExpr}
		}
	case 185:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1042
		{
			yyVAL.limitOffset = &LimitOffset{Offset: yyDollar[2].valExpr, Rowcount: yyDollar[4].valExpr}
		}
	case 186:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1046
		{
			yyVAL.limitOffset = &LimitOffset{Rowcount: yyDollar[2].valExpr, Offset: yyDollar[4].valExpr}
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1051
		{
			yyVAL.str = ""
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1055
		{
			yyVAL.str = ForUpdateStr
		}
	case 189:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line sql.y:1059
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
		//line sql.y:1073
		{
			yyVAL.dataType = &DataType{Type: "bigint"}
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1075
		{
			yyVAL.dataType = &DataType{Type: "bigserial"}
		}
	case 192:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1077
		{
			yyVAL.dataType = &DataType{Type: "bit"}
		}
	case 193:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1079
		{
			yyVAL.dataType = &DataType{Type: "bit_varying"}
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1081
		{
			yyVAL.dataType = &DataType{Type: "boolean"}
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1083
		{
			yyVAL.dataType = &DataType{Type: "box"}
		}
	case 196:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1085
		{
			yyVAL.dataType = &DataType{Type: "bytea"}
		}
	case 197:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1087
		{
			yyVAL.dataType = &DataType{Type: "char"}
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1089
		{
			yyVAL.dataType = &DataType{Type: "char_varying"}
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1091
		{
			yyVAL.dataType = &DataType{Type: "cidr"}
		}
	case 200:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1093
		{
			yyVAL.dataType = &DataType{Type: "circle"}
		}
	case 201:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1095
		{
			yyVAL.dataType = &DataType{Type: "date"}
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1097
		{
			yyVAL.dataType = &DataType{Type: "decimal"}
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1099
		{
			yyVAL.dataType = &DataType{Type: "double"}
		}
	case 204:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1101
		{
			yyVAL.dataType = &DataType{Type: "enum"}
		}
	case 205:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1103
		{
			yyVAL.dataType = &DataType{Type: "float"}
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1105
		{
			yyVAL.dataType = &DataType{Type: "inet"}
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1107
		{
			yyVAL.dataType = &DataType{Type: "integer"}
		}
	case 208:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1111
		{
			yyVAL.dataType = &DataType{Type: "json"}
		}
	case 209:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1113
		{
			yyVAL.dataType = &DataType{Type: "jsonb"}
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1115
		{
			yyVAL.dataType = &DataType{Type: "line"}
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1117
		{
			yyVAL.dataType = &DataType{Type: "lseg"}
		}
	case 212:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1119
		{
			yyVAL.dataType = &DataType{Type: "macaddr"}
		}
	case 213:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1121
		{
			yyVAL.dataType = &DataType{Type: "money"}
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1123
		{
			yyVAL.dataType = &DataType{Type: "numeric"}
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1125
		{
			yyVAL.dataType = &DataType{Type: "path"}
		}
	case 216:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1127
		{
			yyVAL.dataType = &DataType{Type: "point"}
		}
	case 217:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1129
		{
			yyVAL.dataType = &DataType{Type: "polygon"}
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1131
		{
			yyVAL.dataType = &DataType{Type: "real"}
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1133
		{
			yyVAL.dataType = &DataType{Type: "serial"}
		}
	case 220:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1135
		{
			yyVAL.dataType = &DataType{Type: "smallint"}
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1137
		{
			yyVAL.dataType = &DataType{Type: "smallserial"}
		}
	case 222:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1139
		{
			yyVAL.dataType = &DataType{Type: "text"}
		}
	case 223:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1141
		{
			yyVAL.dataType = &DataType{Type: "time"}
		}
	case 224:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1143
		{
			yyVAL.dataType = &DataType{Type: "timestamp"}
		}
	case 225:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1145
		{
			yyVAL.dataType = &DataType{Type: "uuid"}
		}
	case 226:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1147
		{
			yyVAL.dataType = &DataType{Type: "xml"}
		}
	case 227:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1152
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNotNullStr}
		}
	case 228:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1156
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrNullStr}
		}
	case 229:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1160
		{
			yyVAL.colConstr = &ColConstr{Constraint: ColConstrPrimaryKeyStr}
		}
	case 230:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1170
		{
			yyVAL.colConstrs = ColConstrs{yyDollar[1].colConstr}
		}
	case 231:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1174
		{
			yyVAL.colConstrs = append(yyVAL.colConstrs, yyDollar[2].colConstr)
		}
	case 232:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1181
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType, Constraints: yyDollar[3].colConstrs}
		}
	case 233:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1185
		{
			yyVAL.colDef = &ColDef{ColName: yyDollar[1].colIdent, ColType: yyDollar[2].dataType}
		}
	case 234:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1191
		{
			yyVAL.colDefs = ColDefs{yyDollar[1].colDef}
		}
	case 235:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1195
		{
			yyVAL.colDefs = append(yyVAL.colDefs, yyDollar[3].colDef)
		}
	case 236:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1200
		{
			yyVAL.colDefs = nil
		}
	case 237:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1204
		{
			yyVAL.colDefs = yyDollar[2].colDefs
		}
	case 238:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1209
		{
			yyVAL.columns = nil
		}
	case 239:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1213
		{
			yyVAL.columns = yyDollar[2].columns
		}
	case 240:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1219
		{
			yyVAL.columns = Columns{yyDollar[1].colIdent}
		}
	case 241:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1223
		{
			yyVAL.columns = append(yyVAL.columns, yyDollar[3].colIdent)
		}
	case 242:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1230
		{
			yyVAL.updateExprs = nil
		}
	case 243:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line sql.y:1234
		{
			yyVAL.updateExprs = yyDollar[5].updateExprs
		}
	case 244:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1240
		{
			yyVAL.insRows = yyDollar[2].values
		}
	case 245:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1244
		{
			yyVAL.insRows = yyDollar[1].selStmt
		}
	case 246:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1250
		{
			yyVAL.values = Values{yyDollar[1].rowTuple}
		}
	case 247:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1254
		{
			yyVAL.values = append(yyDollar[1].values, yyDollar[3].rowTuple)
		}
	case 248:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1260
		{
			yyVAL.rowTuple = ValTuple(yyDollar[2].valExprs)
		}
	case 249:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1264
		{
			yyVAL.rowTuple = yyDollar[1].subquery
		}
	case 250:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1270
		{
			yyVAL.updateExprs = UpdateExprs{yyDollar[1].updateExpr}
		}
	case 251:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1274
		{
			yyVAL.updateExprs = append(yyDollar[1].updateExprs, yyDollar[3].updateExpr)
		}
	case 252:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1280
		{
			yyVAL.updateExpr = &UpdateExpr{Name: yyDollar[1].colIdent, Expr: yyDollar[3].valExpr}
		}
	case 255:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1289
		{
			yyVAL.byt = 0
		}
	case 256:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1291
		{
			yyVAL.byt = 1
		}
	case 257:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1294
		{
			yyVAL.empty = struct{}{}
		}
	case 258:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line sql.y:1296
		{
			yyVAL.empty = struct{}{}
		}
	case 259:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1299
		{
			yyVAL.str = ""
		}
	case 260:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1301
		{
			yyVAL.str = IgnoreStr
		}
	case 261:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1305
		{
			yyVAL.empty = struct{}{}
		}
	case 262:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1307
		{
			yyVAL.empty = struct{}{}
		}
	case 263:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1309
		{
			yyVAL.empty = struct{}{}
		}
	case 264:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1311
		{
			yyVAL.empty = struct{}{}
		}
	case 265:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1313
		{
			yyVAL.empty = struct{}{}
		}
	case 266:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1315
		{
			yyVAL.empty = struct{}{}
		}
	case 267:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1318
		{
			yyVAL.empty = struct{}{}
		}
	case 268:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1320
		{
			yyVAL.empty = struct{}{}
		}
	case 269:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1323
		{
			yyVAL.empty = struct{}{}
		}
	case 270:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1325
		{
			yyVAL.empty = struct{}{}
		}
	case 271:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1328
		{
			yyVAL.empty = struct{}{}
		}
	case 272:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line sql.y:1330
		{
			yyVAL.empty = struct{}{}
		}
	case 273:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1334
		{
			yyVAL.colIdent = NewColIdent(string(yyDollar[1].bytes))
		}
	case 274:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1340
		{
			yyVAL.tableIdent = TableIdent(yyDollar[1].bytes)
		}
	case 275:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1346
		{
			if incNesting(yylex) {
				yylex.Error("max nesting level reached")
				return 1
			}
		}
	case 276:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line sql.y:1355
		{
			decNesting(yylex)
		}
	case 277:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line sql.y:1360
		{
			forceEOF(yylex)
		}
	}
	goto yystack /* stack new state and value */
}
