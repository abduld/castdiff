//line cc.y:34
package cc

import __yyfmt__ "fmt"

//line cc.y:34
import (
// "runtime/debug"
)

type typeClass struct {
	c Storage
	q TypeQual
	t *Type
}

type idecor struct {
	d func(*Type) (*Type, Syntax)
	i *Init
}

var id int = 0

func nextId() int {
	id += 1
	return id
}

//line cc.y:60
type yySymType struct {
	yys      int
	abdecor  func(*Type) *Type
	decl     *Decl
	decls    []*Decl
	decor    func(*Type) (*Type, Syntax)
	decors   []func(*Type) (*Type, Syntax)
	expr     *Expr
	exprs    []*Expr
	idec     idecor
	idecs    []idecor
	init     *Init
	inits    []*Init
	label    *Label
	labels   []*Label
	span     Span
	prefix   *Prefix
	prefixes []*Prefix
	stmt     *Stmt
	stmts    []*Stmt
	str      string
	strs     []string
	tc       typeClass
	tk       TypeKind
	typ      *Type
	symlit   *SymbolLiteral
	boollit  *BooleanLiteral
	reallit  *RealLiteral
	intlit   *IntegerLiteral
	charlit  *CharLiteral
	strlit   *StringLiteral
	syntax   Syntax
	syntaxs  []Syntax
	langkey  *LanguageKeyword
}

const tokARGBEGIN = 57346
const tokARGEND = 57347
const tokAUTOLIB = 57348
const tokSET = 57349
const tokUSED = 57350
const tokAuto = 57351
const tokBreak = 57352
const tokCase = 57353
const tokChar = 57354
const tokConst = 57355
const tokContinue = 57356
const tokDefault = 57357
const tokDo = 57358
const tokDotDotDot = 57359
const tokDouble = 57360
const tokEnum = 57361
const tokError = 57362
const tokExtern = 57363
const tokFloat = 57364
const tokFor = 57365
const tokGoto = 57366
const tokIf = 57367
const tokInline = 57368
const tokInt = 57369
const tokLitChar = 57370
const tokLong = 57371
const tokName = 57372
const tokInteger = 57373
const tokReal = 57374
const tokOffsetof = 57375
const tokRegister = 57376
const tokReturn = 57377
const tokShort = 57378
const tokSigned = 57379
const tokStatic = 57380
const tokStruct = 57381
const tokSwitch = 57382
const tokTypeName = 57383
const tokTypedef = 57384
const tokUnion = 57385
const tokUnsigned = 57386
const tokVaArg = 57387
const tokVoid = 57388
const tokVolatile = 57389
const tokWhile = 57390
const tokString = 57391
const tokLCuBrk = 57392
const tokRCuBrk = 57393
const tokDevice = 57394
const tokHost = 57395
const tokGlobal = 57396
const tokShared = 57397
const tokRestrict = 57398
const tokShift = 57399
const tokElse = 57400
const tokAddEq = 57401
const tokSubEq = 57402
const tokMulEq = 57403
const tokDivEq = 57404
const tokModEq = 57405
const tokLshEq = 57406
const tokRshEq = 57407
const tokAndEq = 57408
const tokXorEq = 57409
const tokOrEq = 57410
const tokOrOr = 57411
const tokAndAnd = 57412
const tokEqEq = 57413
const tokNotEq = 57414
const tokLtEq = 57415
const tokGtEq = 57416
const tokLsh = 57417
const tokRsh = 57418
const tokCast = 57419
const tokSizeof = 57420
const tokUnary = 57421
const tokDec = 57422
const tokInc = 57423
const tokArrow = 57424
const startProg = 57425
const startExpr = 57426
const tokEOF = 57427

var yyToknames = []string{
	"tokARGBEGIN",
	"tokARGEND",
	"tokAUTOLIB",
	"tokSET",
	"tokUSED",
	"tokAuto",
	"tokBreak",
	"tokCase",
	"tokChar",
	"tokConst",
	"tokContinue",
	"tokDefault",
	"tokDo",
	"tokDotDotDot",
	"tokDouble",
	"tokEnum",
	"tokError",
	"tokExtern",
	"tokFloat",
	"tokFor",
	"tokGoto",
	"tokIf",
	"tokInline",
	"tokInt",
	"tokLitChar",
	"tokLong",
	"tokName",
	"tokInteger",
	"tokReal",
	"tokOffsetof",
	"tokRegister",
	"tokReturn",
	"tokShort",
	"tokSigned",
	"tokStatic",
	"tokStruct",
	"tokSwitch",
	"tokTypeName",
	"tokTypedef",
	"tokUnion",
	"tokUnsigned",
	"tokVaArg",
	"tokVoid",
	"tokVolatile",
	"tokWhile",
	"tokString",
	"tokLCuBrk",
	"tokRCuBrk",
	"tokDevice",
	"tokHost",
	"tokGlobal",
	"tokShared",
	"tokRestrict",
	"tokShift",
	"tokElse",
	"'{'",
	"','",
	"'='",
	"tokAddEq",
	"tokSubEq",
	"tokMulEq",
	"tokDivEq",
	"tokModEq",
	"tokLshEq",
	"tokRshEq",
	"tokAndEq",
	"tokXorEq",
	"tokOrEq",
	"'?'",
	"':'",
	"tokOrOr",
	"tokAndAnd",
	"'|'",
	"'^'",
	"'&'",
	"tokEqEq",
	"tokNotEq",
	"'<'",
	"'>'",
	"tokLtEq",
	"tokGtEq",
	"tokLsh",
	"tokRsh",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"tokCast",
	"'!'",
	"'~'",
	"tokSizeof",
	"tokUnary",
	"'.'",
	"'['",
	"']'",
	"')'",
	"tokDec",
	"tokInc",
	"tokArrow",
	"'('",
	"startProg",
	"startExpr",
	"tokEOF",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 125,
	60, 102,
	109, 102,
	-2, 187,
	-1, 143,
	59, 178,
	-2, 152,
	-1, 145,
	59, 178,
	-2, 157,
	-1, 247,
	109, 213,
	-2, 177,
	-1, 279,
	73, 178,
	-2, 93,
}

const yyNprod = 223
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1658

var yyAct = []int{

	321, 7, 119, 127, 354, 314, 269, 32, 234, 222,
	276, 290, 204, 118, 249, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 228, 201, 6, 355, 246, 124,
	50, 5, 178, 226, 116, 130, 232, 320, 236, 4,
	140, 138, 143, 145, 136, 125, 389, 387, 380, 379,
	375, 117, 33, 368, 366, 349, 348, 346, 300, 299,
	308, 195, 317, 303, 254, 65, 36, 374, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 165, 137, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 197,
	135, 131, 141, 131, 35, 196, 182, 183, 307, 2,
	3, 189, 132, 166, 132, 357, 11, 97, 8, 9,
	10, 22, 356, 192, 179, 179, 191, 181, 188, 198,
	197, 180, 244, 24, 353, 351, 196, 25, 134, 345,
	142, 344, 117, 252, 221, 122, 121, 218, 115, 184,
	185, 392, 72, 73, 67, 68, 69, 70, 71, 203,
	128, 219, 266, 259, 103, 99, 197, 219, 101, 100,
	102, 98, 196, 360, 386, 129, 206, 267, 205, 378,
	207, 17, 18, 21, 377, 376, 137, 373, 216, 20,
	19, 372, 23, 306, 97, 311, 295, 294, 262, 233,
	235, 141, 239, 224, 135, 214, 141, 212, 187, 186,
	230, 293, 251, 241, 242, 220, 216, 253, 268, 213,
	203, 233, 247, 359, 302, 217, 210, 284, 32, 301,
	257, 291, 292, 243, 225, 230, 240, 238, 281, 142,
	264, 103, 99, 265, 142, 101, 100, 102, 98, 263,
	215, 279, 241, 217, 66, 258, 256, 235, 260, 247,
	277, 200, 270, 209, 288, 208, 194, 388, 211, 271,
	123, 104, 364, 131, 273, 34, 278, 230, 250, 179,
	97, 193, 309, 280, 132, 285, 305, 237, 296, 298,
	1, 38, 297, 12, 313, 312, 304, 203, 202, 139,
	49, 310, 324, 289, 316, 279, 323, 257, 325, 133,
	242, 235, 315, 255, 277, 144, 146, 239, 318, 69,
	70, 71, 287, 58, 126, 282, 283, 103, 99, 274,
	329, 101, 100, 102, 98, 350, 275, 347, 248, 245,
	352, 29, 27, 358, 227, 199, 30, 190, 0, 0,
	239, 330, 0, 0, 0, 0, 365, 59, 0, 0,
	0, 0, 60, 61, 62, 63, 64, 0, 0, 0,
	0, 0, 361, 362, 0, 0, 0, 383, 384, 385,
	382, 367, 0, 0, 369, 370, 0, 53, 0, 391,
	40, 58, 390, 393, 0, 0, 47, 39, 97, 120,
	46, 0, 381, 0, 57, 42, 11, 43, 8, 9,
	10, 22, 56, 0, 41, 44, 54, 51, 0, 37,
	55, 52, 45, 24, 48, 59, 0, 25, 0, 0,
	60, 61, 62, 63, 64, 67, 68, 69, 70, 71,
	0, 0, 0, 0, 0, 103, 99, 0, 0, 101,
	100, 102, 98, 0, 0, 0, 14, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 16, 13, 0, 0,
	0, 17, 18, 21, 97, 0, 0, 0, 0, 20,
	19, 0, 23, 0, 0, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 85, 371, 84, 83,
	82, 81, 80, 78, 79, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	0, 103, 99, 0, 0, 101, 100, 102, 98, 331,
	0, 0, 328, 327, 0, 332, 341, 0, 0, 333,
	342, 334, 0, 0, 0, 0, 0, 0, 335, 336,
	337, 0, 0, 11, 97, 343, 9, 10, 22, 0,
	338, 0, 0, 0, 0, 339, 0, 0, 0, 0,
	24, 0, 0, 340, 25, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 270, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	0, 103, 99, 14, 0, 101, 100, 102, 98, 0,
	0, 0, 15, 16, 13, 0, 0, 0, 17, 18,
	21, 0, 0, 0, 97, 0, 20, 19, 0, 23,
	0, 0, 0, 0, 326, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 85, 0, 84, 83,
	82, 81, 80, 78, 79, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 0, 0, 97, 0,
	0, 103, 99, 319, 0, 101, 100, 102, 98, 86,
	87, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	85, 0, 84, 83, 82, 81, 80, 78, 79, 74,
	75, 76, 77, 72, 73, 67, 68, 69, 70, 71,
	0, 0, 0, 97, 0, 103, 99, 0, 286, 101,
	100, 102, 98, 223, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 96, 85, 0, 84, 83, 82,
	81, 80, 78, 79, 74, 75, 76, 77, 72, 73,
	67, 68, 69, 70, 71, 0, 0, 97, 0, 0,
	103, 99, 0, 0, 101, 100, 102, 98, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 96, 85,
	0, 84, 83, 82, 81, 80, 78, 79, 74, 75,
	76, 77, 72, 73, 67, 68, 69, 70, 71, 0,
	28, 0, 0, 53, 103, 99, 40, 58, 101, 100,
	102, 98, 47, 39, 0, 31, 46, 0, 0, 0,
	57, 42, 0, 43, 0, 0, 0, 0, 56, 0,
	41, 44, 54, 51, 0, 37, 55, 52, 45, 0,
	48, 59, 0, 0, 0, 0, 60, 61, 62, 63,
	64, 53, 0, 0, 40, 58, 0, 0, 0, 0,
	47, 39, 0, 120, 46, 0, 0, 0, 57, 42,
	0, 43, 0, 0, 0, 0, 56, 0, 41, 44,
	54, 51, 0, 37, 55, 52, 45, 0, 48, 59,
	0, 0, 0, 0, 60, 61, 62, 63, 64, 0,
	53, 0, 261, 40, 58, 0, 0, 0, 0, 47,
	39, 0, 120, 46, 0, 0, 0, 57, 42, 0,
	43, 0, 0, 0, 0, 56, 0, 41, 44, 54,
	51, 0, 37, 55, 52, 45, 0, 48, 59, 0,
	0, 0, 0, 60, 61, 62, 63, 64, 28, 0,
	322, 53, 0, 0, 40, 58, 0, 0, 0, 0,
	47, 39, 0, 31, 46, 0, 0, 0, 57, 42,
	0, 43, 0, 0, 0, 0, 56, 0, 41, 44,
	54, 51, 97, 37, 55, 52, 45, 0, 48, 59,
	0, 0, 0, 0, 60, 61, 62, 63, 64, 272,
	0, 0, 0, 0, 85, 0, 84, 83, 82, 81,
	80, 78, 79, 74, 75, 76, 77, 72, 73, 67,
	68, 69, 70, 71, 0, 97, 0, 0, 0, 103,
	99, 0, 0, 101, 100, 102, 98, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 26,
	83, 82, 81, 80, 78, 79, 74, 75, 76, 77,
	72, 73, 67, 68, 69, 70, 71, 97, 0, 0,
	0, 0, 103, 99, 0, 0, 101, 100, 102, 98,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 82, 81, 80, 78, 79, 74, 75,
	76, 77, 72, 73, 67, 68, 69, 70, 71, 97,
	0, 0, 0, 0, 103, 99, 0, 0, 101, 100,
	102, 98, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 81, 80, 78, 79,
	74, 75, 76, 77, 72, 73, 67, 68, 69, 70,
	71, 97, 0, 0, 0, 0, 103, 99, 0, 0,
	101, 100, 102, 98, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 80,
	78, 79, 74, 75, 76, 77, 72, 73, 67, 68,
	69, 70, 71, 0, 0, 0, 0, 0, 103, 99,
	0, 0, 101, 100, 102, 98, 11, 0, 8, 9,
	10, 22, 0, 0, 53, 0, 0, 0, 58, 0,
	0, 0, 0, 24, 0, 0, 120, 25, 0, 0,
	0, 57, 0, 0, 0, 0, 0, 218, 0, 56,
	0, 0, 0, 54, 0, 0, 0, 55, 0, 0,
	0, 0, 59, 0, 97, 0, 14, 60, 61, 62,
	63, 64, 0, 0, 0, 15, 16, 13, 0, 0,
	0, 17, 18, 21, 0, 291, 292, 0, 0, 20,
	19, 97, 23, 78, 79, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	0, 103, 99, 0, 0, 101, 100, 102, 98, 0,
	0, 79, 74, 75, 76, 77, 72, 73, 67, 68,
	69, 70, 71, 0, 0, 0, 0, 0, 103, 99,
	0, 0, 101, 100, 102, 98, 11, 0, 8, 9,
	10, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 24, 0, 0, 11, 25, 8, 9,
	10, 22, 0, 0, 0, 0, 0, 218, 0, 0,
	0, 0, 0, 24, 0, 0, 0, 25, 0, 0,
	0, 0, 0, 0, 0, 0, 14, 11, 0, 8,
	9, 10, 22, 0, 0, 15, 16, 13, 0, 0,
	0, 17, 18, 21, 24, 0, 14, 0, 25, 20,
	19, 0, 23, 0, 0, 15, 16, 13, 0, 0,
	0, 17, 18, 21, 0, 0, 0, 0, 0, 20,
	19, 0, 23, 0, 0, 0, 0, 14, 0, 0,
	0, 0, 0, 0, 0, 0, 15, 16, 13, 0,
	0, 0, 17, 18, 21, 0, 0, 0, 0, 0,
	20, 19, 53, 114, 0, 40, 58, 0, 0, 0,
	231, 47, 39, 0, 120, 46, 0, 0, 0, 57,
	42, 0, 43, 229, 0, 0, 0, 56, 0, 41,
	44, 54, 51, 0, 37, 55, 52, 45, 0, 48,
	59, 0, 0, 0, 0, 60, 61, 62, 63, 64,
	363, 0, 0, 0, 53, 0, 0, 40, 58, 0,
	0, 0, 0, 47, 39, 0, 120, 46, 0, 0,
	0, 57, 42, 0, 43, 0, 0, 0, 0, 56,
	0, 41, 44, 54, 51, 0, 37, 55, 52, 45,
	0, 48, 59, 0, 0, 0, 0, 60, 61, 62,
	63, 64, 53, 0, 0, 40, 58, 0, 0, 0,
	0, 47, 39, 0, 120, 46, 0, 0, 0, 57,
	42, 0, 43, 0, 0, 0, 0, 56, 0, 41,
	44, 54, 51, 0, 37, 55, 52, 45, 0, 48,
	59, 0, 0, 0, 0, 60, 61, 62, 63, 64,
	53, 0, 0, 40, 58, 0, 0, 0, 0, 47,
	0, 0, 120, 46, 0, 0, 0, 57, 42, 0,
	43, 0, 0, 0, 0, 56, 0, 41, 44, 54,
	0, 0, 0, 55, 0, 45, 0, 48, 59, 0,
	0, 0, 0, 60, 61, 62, 63, 64,
}
var yyPact = []int{

	4, -1000, -1000, 1328, 942, -42, 194, 707, -1000, -1000,
	-1000, -1000, 222, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1328, 1359, 44, 378, 42, -1000, -1000, -1000, 41, -1000,
	-1000, 221, 71, 1553, 1205, 1601, -1000, -1000, 243, 243,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	-1000, -1000, 243, 243, -1000, 144, 144, 144, 144, 144,
	144, 144, 144, 144, 378, 1553, 109, 108, 22, -1000,
	-1000, 1328, 251, 207, -48, 68, 201, -1000, 310, 71,
	-1000, -1000, -1000, 1205, 1601, -1000, -1000, 1205, -1000, 1601,
	-1000, -1000, -1000, -1000, 206, -1000, 204, 707, 230, 230,
	144, 144, 144, 348, 348, 67, 67, 67, 67, 1231,
	504, 1204, 1101, 1059, 1017, 975, 153, 707, 707, 707,
	707, 707, 707, 707, 707, 707, 707, 707, 217, 194,
	107, 120, -1000, -1000, 105, 190, 1308, -1000, 69, 310,
	40, 22, 663, 103, -1000, -1000, 1453, 1328, 1308, 1553,
	71, 71, 310, -1000, 32, -1000, -1000, -1000, 1553, 248,
	1328, 39, -1000, -1000, 88, 1328, 144, -1000, -44, 1328,
	22, 1453, 63, 1553, -1000, 794, 98, 189, -1000, -1000,
	73, -1000, 119, 707, -1000, 707, -1000, 203, -1000, 71,
	-1000, 68, 1, -1000, -1000, 891, -1000, 71, 178, -1000,
	166, 932, 1328, 618, -1000, 1178, 112, 69, 97, -1000,
	96, -1000, -1000, 1453, 69, 1, 310, 73, -1000, -1000,
	-1000, -50, -1000, -1000, -51, 169, -1000, 1, 151, -1000,
	-45, 248, -1000, -1000, 1328, 93, -1000, 0, -1000, 134,
	-1000, 243, 1328, -1000, -1000, -1000, -1000, 73, -1000, -1000,
	-1000, 71, 1328, -1000, -1000, 707, -1000, -1000, -46, 1308,
	-1000, -1000, -1000, 574, 842, -1000, 707, -1000, -1000, -1000,
	-1000, -1000, -1000, 525, -1000, -1000, -1000, 37, 35, -1000,
	-52, -1000, -53, -54, -1000, 31, 243, 30, 1328, 18,
	11, 1328, 150, 100, 1328, 1328, -1000, 1505, -1000, -1000,
	224, 1328, -55, 1328, -56, -1000, 1328, 1328, 424, -1000,
	-1000, 91, 87, -1000, -37, -59, -1000, 85, -1000, 84,
	79, -1000, -60, -61, 1328, 1328, -1000, -1000, -1000, -1000,
	-1000, 74, -62, 209, -1000, -1000, -63, 1328, -1000, -1000,
	51, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 9, 347, 24, 346, 14, 37, 345, 344, 33,
	39, 342, 341, 28, 339, 338, 12, 10, 336, 329,
	1, 36, 27, 4, 326, 325, 26, 32, 35, 324,
	29, 8, 322, 38, 313, 308, 306, 11, 303, 302,
	6, 0, 5, 3, 300, 30, 104, 66, 40, 276,
	52, 44, 299, 41, 298, 25, 293, 2, 291, 34,
	13, 275, 290, 289, 287, 283, 282,
}
var yyR1 = []int{

	0, 62, 62, 10, 10, 10, 22, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 42, 42, 42, 63, 40, 35, 35, 35,
	41, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 1, 1, 1,
	2, 2, 2, 16, 16, 16, 16, 16, 3, 3,
	3, 3, 28, 28, 44, 44, 44, 44, 44, 44,
	45, 45, 45, 45, 45, 45, 45, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 47, 47, 48, 48,
	61, 57, 57, 57, 57, 57, 60, 59, 6, 12,
	11, 11, 11, 64, 4, 43, 43, 58, 58, 17,
	17, 13, 61, 61, 37, 20, 20, 61, 61, 5,
	24, 31, 31, 33, 33, 33, 34, 34, 32, 32,
	37, 66, 66, 65, 65, 38, 38, 49, 49, 23,
	23, 21, 21, 26, 26, 27, 27, 7, 7, 36,
	36, 8, 8, 9, 9, 29, 29, 30, 30, 54,
	54, 55, 55, 50, 50, 51, 51, 52, 52, 53,
	53, 18, 18, 19, 19, 14, 14, 25, 25, 15,
	15, 56, 56,
}
var yyR2 = []int{

	0, 3, 3, 0, 2, 5, 1, 1, 1, 1,
	1, 1, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	5, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 4, 6, 4, 4, 3, 7, 4, 4, 2,
	2, 6, 0, 2, 2, 0, 4, 3, 2, 2,
	2, 1, 5, 5, 1, 2, 3, 2, 2, 7,
	9, 3, 5, 7, 3, 5, 5, 0, 3, 1,
	4, 4, 3, 1, 3, 3, 4, 4, 1, 2,
	2, 1, 1, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 3, 2, 2, 1, 2, 3, 3,
	1, 1, 5, 0, 5, 1, 1, 1, 1, 1,
	3, 3, 2, 5, 2, 3, 3, 2, 6, 2,
	2, 1, 1, 2, 4, 5, 0, 3, 1, 3,
	3, 0, 1, 0, 1, 1, 2, 0, 1, 0,
	1, 0, 1, 1, 3, 0, 1, 0, 2, 0,
	2, 1, 3, 0, 1, 1, 3, 0, 1, 1,
	2, 0, 1, 1, 2, 0, 1, 1, 2, 0,
	1, 1, 3, 0, 1, 1, 2, 0, 1, 1,
	3, 1, 2,
}
var yyChk = []int{

	-1000, -62, 105, 106, -10, -22, -26, -20, 30, 31,
	32, 28, -56, 89, 78, 87, 88, 93, 94, 102,
	101, 95, 33, 104, 45, 49, 107, -11, 6, -12,
	-4, 21, -57, -50, -61, -46, -47, 41, -58, 19,
	12, 36, 27, 29, 37, 44, 22, 18, 46, -44,
	-45, 39, 43, 9, 38, 42, 34, 26, 13, 47,
	52, 53, 54, 55, 56, 107, 60, 87, 88, 89,
	90, 91, 85, 86, 81, 82, 83, 84, 79, 80,
	78, 77, 76, 75, 74, 72, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 50, 104, 98,
	102, 101, 103, 97, 49, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, 104, 104, -59, -22, -60, -57,
	21, 104, 104, 49, -30, -16, -29, -43, 89, 104,
	-28, 30, 41, -61, -46, -47, -51, -50, -53, -52,
	-48, -47, -46, -43, -49, -43, -49, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -22, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -27, -26,
	-27, -22, -43, -43, -59, -59, 100, 100, -1, 89,
	-2, 104, -20, 30, 59, 109, 104, 98, 61, -7,
	60, -55, -54, -45, -16, -51, -53, -48, 59, 59,
	73, 51, 100, 99, 100, 60, -20, -33, 59, 98,
	-55, 104, -1, 60, 100, -10, -9, -8, -3, 30,
	-60, 17, -21, -20, -31, -20, -33, -64, -6, -57,
	-28, -16, -16, -45, 100, -14, -13, -60, -15, -5,
	30, -20, 104, -20, 108, -34, -21, -1, -9, 100,
	-59, 108, 100, 60, -1, -16, 89, 104, 99, -40,
	59, -30, 108, -13, -19, -18, -17, -16, -49, -43,
	-65, 60, -25, -24, 61, -27, 100, -32, -31, -38,
	-37, 97, 98, 99, 100, 100, -3, -55, -63, 109,
	109, 60, 73, 108, -5, -20, 100, 108, 60, -66,
	-37, 61, -43, -20, -42, -17, -20, 108, -31, 99,
	-6, -41, 108, -36, -39, -35, 109, 8, 7, -40,
	-22, 4, 10, 14, 16, 23, 24, 25, 35, 40,
	48, 11, 15, 30, 104, 104, 109, -42, 109, 109,
	-41, 104, -43, 104, -23, -22, 104, 104, -20, 73,
	73, -22, -22, 5, 48, -23, 109, -22, 109, -22,
	-22, 73, 100, 100, 104, 109, 100, 100, 100, 109,
	109, -22, -23, -41, -41, -41, 100, 109, 58, 109,
	-23, -41, 100, -41,
}
var yyDef = []int{

	0, -2, 3, 0, 0, 0, 6, 183, 7, 8,
	9, 10, 11, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 221, 1, 4, 0, 140,
	141, 106, 197, 131, 205, 209, 203, 130, 177, 177,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	127, 147, 148, 104, 105, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 116, 2, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 185, 185, 0,
	59, 60, 0, 0, 222, 42, 43, 44, 45, 46,
	47, 48, 49, 50, 0, 0, 0, 0, 87, 136,
	106, 0, 0, 0, 0, -2, 198, 93, 201, 0,
	195, 145, 146, 205, 209, 204, 134, 206, 135, 210,
	207, 128, 129, -2, 0, -2, 0, 184, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 0, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 0, 186,
	0, 0, 155, 156, 0, 0, 0, 55, 137, 201,
	89, 87, 0, 0, 3, 139, 193, 181, 0, 143,
	0, 0, 202, 199, 0, 132, 133, 208, 0, 0,
	0, 0, 57, 58, 51, 0, 53, 54, 166, 181,
	87, 193, 0, 0, 5, 0, 0, 194, 191, 98,
	87, 101, 0, 182, 103, 161, 162, 0, 188, 197,
	196, 102, 94, 200, 95, 0, 215, -2, 173, 219,
	217, 30, 185, 0, 163, 0, 0, 88, 0, 92,
	0, 142, 96, 0, 99, 100, 201, 87, 97, 144,
	65, 0, 153, 216, 0, 214, 211, 149, 0, -2,
	0, 174, 159, 218, 0, 0, 52, 0, 168, 171,
	175, 0, 0, 91, 90, 61, 192, 87, 62, 138,
	151, 177, 0, 158, 220, 160, 56, 164, 167, 0,
	176, 172, 154, 0, 189, 212, 150, 165, 169, 170,
	63, 64, 66, 0, 70, 190, 71, 0, 0, 74,
	0, 62, 0, 0, 189, 0, 0, 0, 179, 0,
	0, 0, 0, 7, 0, 0, 75, 189, 77, 78,
	0, 179, 0, 0, 0, 180, 0, 0, 0, 68,
	69, 0, 0, 76, 0, 0, 81, 0, 84, 0,
	0, 67, 0, 0, 0, 179, 189, 189, 189, 72,
	73, 0, 0, 82, 85, 86, 0, 179, 189, 79,
	0, 83, 189, 80,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 93, 3, 3, 3, 91, 78, 3,
	104, 100, 89, 87, 60, 88, 97, 90, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 73, 109,
	81, 61, 82, 72, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 98, 3, 99, 77, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 59, 76, 108, 94,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 74, 75, 79,
	80, 83, 84, 85, 86, 92, 95, 96, 101, 102,
	103, 105, 106, 107,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
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

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
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
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
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
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
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
			if yyn < 0 || yyn == yychar {
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
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
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
		//line cc.y:219
		{
			yylex.(*lexer).prog = &Prog{Decls: yyS[yypt-1].decls, Id: nextId()}
			return 0
		}
	case 2:
		//line cc.y:224
		{
			yylex.(*lexer).expr = yyS[yypt-1].expr
			return 0
		}
	case 3:
		//line cc.y:230
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 4:
		//line cc.y:235
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 5:
		//line cc.y:240
		{
		}
	case 6:
		//line cc.y:245
		{
			yyVAL.span = yyS[yypt-0].span
			if len(yyS[yypt-0].exprs) == 1 {
				yyVAL.expr = yyS[yypt-0].exprs[0]
				break
			}
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Comma, List: yyS[yypt-0].exprs}
		}
	case 7:
		//line cc.y:256
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Name,
				Text: &SymbolLiteral{
					Value:      yyS[yypt-0].str,
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				},
				XDecl: yyS[yypt-0].decl,
			}
		}
	case 8:
		//line cc.y:271
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Literal,
				Text:       yyS[yypt-0].intlit,
			}
		}
	case 9:
		//line cc.y:281
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Literal,
				Text:       yyS[yypt-0].reallit,
			}
		}
	case 10:
		//line cc.y:291
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Literal,
				Text: &CharLiteral{
					Value:      yyS[yypt-0].str[0],
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				}}
		}
	case 11:
		//line cc.y:304
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: String, Texts: yyS[yypt-0].syntaxs}
		}
	case 12:
		//line cc.y:309
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Add, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 13:
		//line cc.y:314
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Sub, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 14:
		//line cc.y:319
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Mul, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 15:
		//line cc.y:324
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Div, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 16:
		//line cc.y:329
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Mod, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 17:
		//line cc.y:334
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Lsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 18:
		//line cc.y:339
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Rsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 19:
		//line cc.y:344
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Lt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 20:
		//line cc.y:349
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Gt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 21:
		//line cc.y:354
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: LtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 22:
		//line cc.y:359
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: GtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 23:
		//line cc.y:364
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: EqEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 24:
		//line cc.y:369
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: NotEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 25:
		//line cc.y:374
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: And, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 26:
		//line cc.y:379
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Xor, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 27:
		//line cc.y:384
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Or, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 28:
		//line cc.y:389
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: AndAnd, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 29:
		//line cc.y:394
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: OrOr, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 30:
		//line cc.y:399
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Cond, List: []*Expr{yyS[yypt-4].expr, yyS[yypt-2].expr, yyS[yypt-0].expr}}
		}
	case 31:
		//line cc.y:404
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Eq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 32:
		//line cc.y:409
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: AddEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 33:
		//line cc.y:414
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: SubEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 34:
		//line cc.y:419
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: MulEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 35:
		//line cc.y:424
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: DivEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 36:
		//line cc.y:429
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: ModEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 37:
		//line cc.y:434
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: LshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 38:
		//line cc.y:439
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: RshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 39:
		//line cc.y:444
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: AndEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 40:
		//line cc.y:449
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: XorEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 41:
		//line cc.y:454
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: OrEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 42:
		//line cc.y:459
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Indir, Left: yyS[yypt-0].expr}
		}
	case 43:
		//line cc.y:464
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Addr, Left: yyS[yypt-0].expr}
		}
	case 44:
		//line cc.y:469
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Plus, Left: yyS[yypt-0].expr}
		}
	case 45:
		//line cc.y:474
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Minus, Left: yyS[yypt-0].expr}
		}
	case 46:
		//line cc.y:479
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Not, Left: yyS[yypt-0].expr}
		}
	case 47:
		//line cc.y:484
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Twid, Left: yyS[yypt-0].expr}
		}
	case 48:
		//line cc.y:489
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: PreInc, Left: yyS[yypt-0].expr}
		}
	case 49:
		//line cc.y:494
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: PreDec, Left: yyS[yypt-0].expr}
		}
	case 50:
		//line cc.y:499
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: SizeofExpr, Left: yyS[yypt-0].expr}
		}
	case 51:
		//line cc.y:504
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: SizeofType, Type: yyS[yypt-1].typ}
		}
	case 52:
		//line cc.y:509
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Offsetof, Type: yyS[yypt-3].typ, Left: yyS[yypt-1].expr}
		}
	case 53:
		//line cc.y:514
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Cast, Type: yyS[yypt-2].typ, Left: yyS[yypt-0].expr}
		}
	case 54:
		//line cc.y:519
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: CastInit, Type: yyS[yypt-2].typ, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits, Id: nextId()}}
		}
	case 55:
		//line cc.y:524
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Paren, Left: yyS[yypt-1].expr}
		}
	case 56:
		//line cc.y:529
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: CUDACall, Left: yyS[yypt-6].expr, LaunchParams: yyS[yypt-4].exprs, List: yyS[yypt-1].exprs}
		}
	case 57:
		//line cc.y:534
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Call, Left: yyS[yypt-3].expr, List: yyS[yypt-1].exprs}
		}
	case 58:
		//line cc.y:539
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Index, Left: yyS[yypt-3].expr, Right: yyS[yypt-1].expr}
		}
	case 59:
		//line cc.y:544
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: PostInc, Left: yyS[yypt-1].expr}
		}
	case 60:
		//line cc.y:549
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: PostDec, Left: yyS[yypt-1].expr}
		}
	case 61:
		//line cc.y:554
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: VaArg, Left: yyS[yypt-3].expr, Type: yyS[yypt-1].typ}
		}
	case 62:
		//line cc.y:560
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 63:
		//line cc.y:565
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = yyS[yypt-1].stmts
			for _, d := range yyS[yypt-0].decls {
				yyVAL.stmts = append(yyVAL.stmts, &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: StmtDecl, Decl: d})
			}
		}
	case 64:
		//line cc.y:573
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = append(yyS[yypt-1].stmts, yyS[yypt-0].stmt)
		}
	case 65:
		//line cc.y:580
		{
			yylex.(*lexer).pushScope()
		}
	case 66:
		//line cc.y:584
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Block, Block: yyS[yypt-1].stmts}
		}
	case 67:
		//line cc.y:592
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Case, Expr: yyS[yypt-1].expr}
		}
	case 68:
		//line cc.y:597
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Default}
		}
	case 69:
		//line cc.y:602
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         LabelName,
				Name: &SymbolLiteral{
					Value:      yyS[yypt-1].str,
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				},
			}
		}
	case 70:
		//line cc.y:618
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = yyS[yypt-0].stmt
			yyVAL.stmt.Labels = yyS[yypt-1].labels
		}
	case 71:
		//line cc.y:626
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Empty}
		}
	case 72:
		//line cc.y:631
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Empty}
		}
	case 73:
		//line cc.y:636
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Empty}
		}
	case 74:
		//line cc.y:641
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = yyS[yypt-0].stmt
		}
	case 75:
		//line cc.y:646
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: StmtExpr, Expr: yyS[yypt-1].expr}
		}
	case 76:
		//line cc.y:651
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: ARGBEGIN, Block: yyS[yypt-1].stmts}
		}
	case 77:
		//line cc.y:656
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Break}
		}
	case 78:
		//line cc.y:661
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Continue}
		}
	case 79:
		//line cc.y:666
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Do, Body: yyS[yypt-5].stmt, Expr: yyS[yypt-2].expr}
		}
	case 80:
		//line cc.y:671
		{
			yyVAL.span = span(yyS[yypt-8].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id: nextId(), Op: For,
				Pre:  yyS[yypt-6].expr,
				Expr: yyS[yypt-4].expr,
				Post: yyS[yypt-2].expr,
				Body: yyS[yypt-0].stmt,
			}
		}
	case 81:
		//line cc.y:682
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Goto, Text: yyS[yypt-1].symlit}
		}
	case 82:
		//line cc.y:687
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: If, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 83:
		//line cc.y:692
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: If, Expr: yyS[yypt-4].expr, Body: yyS[yypt-2].stmt, Else: yyS[yypt-0].stmt}
		}
	case 84:
		//line cc.y:697
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Return, Expr: yyS[yypt-1].expr}
		}
	case 85:
		//line cc.y:702
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Switch, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 86:
		//line cc.y:707
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: While, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 87:
		//line cc.y:714
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 88:
		//line cc.y:719
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			_, q, _ := splitTypeWords(yyS[yypt-1].syntaxs)
			abdecor := yyS[yypt-0].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Ptr, Base: t, Qual: q, Id: nextId()})
			}
		}
	case 89:
		//line cc.y:728
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.abdecor = yyS[yypt-0].abdecor
		}
	case 90:
		//line cc.y:735
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			abdecor := yyS[yypt-3].abdecor
			decls := yyS[yypt-1].decls
			span := yyVAL.span
			for _, decl := range decls {
				t := decl.Type
				if t != nil {
					if t.Kind == TypedefType && t.Base != nil {
						t = t.Base
					}
					if t.Kind == Array {
						if t.Width == nil {
							t = t.Base
						}
						decl.Type = &Type{Kind: Ptr, Base: t, Id: nextId()}
					}
				}
			}
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls, Id: nextId()})
			}
		}
	case 91:
		//line cc.y:759
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			abdecor := yyS[yypt-3].abdecor
			span := yyVAL.span
			expr := yyS[yypt-1].expr
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr, Id: nextId()})
			}

		}
	case 92:
		//line cc.y:770
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.abdecor = yyS[yypt-1].abdecor
		}
	case 93:
		//line cc.y:778
		{
			yyVAL.span = yyS[yypt-0].span
			name := yyS[yypt-0].symlit
			yyVAL.decor = func(t *Type) (*Type, Syntax) { return t, name }
		}
	case 94:
		//line cc.y:784
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			_, q, _ := splitTypeWords(yyS[yypt-1].syntaxs)
			decor := yyS[yypt-0].decor
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, Syntax) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Ptr, Base: t, Qual: q, Id: nextId()})
			}
		}
	case 95:
		//line cc.y:794
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decor = yyS[yypt-1].decor
		}
	case 96:
		//line cc.y:799
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			decor := yyS[yypt-3].decor
			decls := yyS[yypt-1].decls
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, Syntax) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls, Id: nextId()})
			}
		}
	case 97:
		//line cc.y:809
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			decor := yyS[yypt-3].decor
			span := yyVAL.span
			expr := yyS[yypt-1].expr
			yyVAL.decor = func(t *Type) (*Type, Syntax) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr, Id: nextId()})
			}
		}
	case 98:
		//line cc.y:822
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &SymbolLiteral{
					Value:      yyS[yypt-0].str,
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				},
				Id: nextId(),
			}
		}
	case 99:
		//line cc.y:835
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyS[yypt-0].abdecor(yyS[yypt-1].typ), Id: nextId()}
		}
	case 100:
		//line cc.y:840
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			typ, name := yyS[yypt-0].decor(yyS[yypt-1].typ)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Id: nextId()}
		}
	case 101:
		//line cc.y:846
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &LanguageKeyword{
					Value:      "...",
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
				},
				Id: nextId(),
			}
		}
	case 102:
		//line cc.y:862
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idec = idecor{yyS[yypt-0].decor, nil}
		}
	case 103:
		//line cc.y:867
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idec = idecor{yyS[yypt-2].decor, yyS[yypt-0].init}
		}
	case 104:
		//line cc.y:875
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 105:
		//line cc.y:884
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 106:
		//line cc.y:893
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 107:
		//line cc.y:902
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 108:
		//line cc.y:911
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 109:
		//line cc.y:920
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 110:
		//line cc.y:932
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 111:
		//line cc.y:941
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 112:
		//line cc.y:950
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 113:
		//line cc.y:959
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 114:
		//line cc.y:968
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 115:
		//line cc.y:977
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 116:
		//line cc.y:986
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 117:
		//line cc.y:998
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 118:
		//line cc.y:1007
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 119:
		//line cc.y:1016
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 120:
		//line cc.y:1025
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 121:
		//line cc.y:1034
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 122:
		//line cc.y:1043
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 123:
		//line cc.y:1052
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 124:
		//line cc.y:1061
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 125:
		//line cc.y:1070
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 126:
		//line cc.y:1081
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = yyS[yypt-0].syntax
		}
	case 127:
		//line cc.y:1086
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = yyS[yypt-0].syntax
		}
	case 128:
		//line cc.y:1093
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = yyS[yypt-0].syntax
		}
	case 129:
		//line cc.y:1098
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = yyS[yypt-0].syntax
		}
	case 130:
		//line cc.y:1106
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.typ = yyS[yypt-0].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{
					Kind: TypedefType,
					Name: &SymbolLiteral{
						Value:      yyS[yypt-0].str,
						Id:         nextId(),
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					},
					Id: nextId(),
				}
			}
		}
	case 131:
		//line cc.y:1130
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(
				append(yyS[yypt-0].syntaxs, &SymbolLiteral{
					Value:      "int",
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				}))
		}
	case 132:
		//line cc.y:1140
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyS[yypt-2].syntaxs, yyS[yypt-0].syntaxs...))
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 133:
		//line cc.y:1146
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyS[yypt-2].syntaxs = append(yyS[yypt-2].syntaxs, yyS[yypt-1].syntax)
			yyS[yypt-2].syntaxs = append(yyS[yypt-2].syntaxs, yyS[yypt-0].syntaxs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyS[yypt-2].syntaxs)
		}
	case 134:
		//line cc.y:1153
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyS[yypt-0].syntaxs)
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 135:
		//line cc.y:1159
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var ts []Syntax
			ts = append(ts, yyS[yypt-1].syntax)
			ts = append(ts, yyS[yypt-0].syntaxs...)
			//PrintStack()
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 136:
		//line cc.y:1171
		{
			yyVAL.span = yyS[yypt-0].span
			if yyS[yypt-0].tc.c != 0 {
				yylex.(*lexer).Errorf("%v not allowed here", yyS[yypt-0].tc.c)
			}
			if yyS[yypt-0].tc.q != 0 && yyS[yypt-0].tc.q != Const && yyS[yypt-0].tc.q != Volatile {
				yylex.(*lexer).Errorf("%v ignored here (TODO)?", yyS[yypt-0].tc.q)
			}
			yyVAL.typ = yyS[yypt-0].tc.t
		}
	case 137:
		//line cc.y:1184
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yyS[yypt-0].abdecor(yyS[yypt-1].typ)
		}
	case 138:
		//line cc.y:1192
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyS[yypt-1].idecs {
				typ, name := idec.d(yyS[yypt-2].tc.t)
				d := &Decl{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       name,
					Type:       typ,
					Storage:    yyS[yypt-2].tc.c,
					Init:       idec.i,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
			if yyS[yypt-1].idecs == nil {
				d := &Decl{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       &SymbolLiteral{},
					Type:       yyS[yypt-2].tc.t,
					Storage:    yyS[yypt-2].tc.c,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 139:
		//line cc.y:1225
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyS[yypt-1].idecs {
				typ, name := idec.d(yyS[yypt-2].tc.t)
				d := lx.lookupDecl(name)
				if d == nil {
					d = &Decl{
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
						Name:       name,
						Type:       typ,
						Storage:    yyS[yypt-2].tc.c,
						Init:       idec.i,
						Id:         nextId(),
					}
					lx.pushDecl(d)
				} else {
					d.Span = yyVAL.span
					if idec.i != nil {
						d.Init = idec.i
					}
				}
				yyVAL.decls = append(yyVAL.decls, d)
			}
			if yyS[yypt-1].idecs == nil {
				d := &Decl{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       &SymbolLiteral{},
					Type:       yyS[yypt-2].tc.t,
					Storage:    yyS[yypt-2].tc.c,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 140:
		//line cc.y:1266
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 141:
		//line cc.y:1271
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 142:
		//line cc.y:1276
		{
			yyVAL.decls = yyS[yypt-1].decls
		}
	case 143:
		//line cc.y:1282
		{
			lx := yylex.(*lexer)
			typ, name := yyS[yypt-1].decor(yyS[yypt-2].tc.t)
			if typ.Kind != Func {
				yylex.(*lexer).Errorf("invalid function definition")
				return 0
			}
			d := lx.lookupDecl(name)
			if d == nil {
				d = &Decl{Name: name, Type: typ, Storage: yyS[yypt-2].tc.c, Id: nextId()}
				lx.pushDecl(d)
			} else {
				d.Type = typ
			}
			yyVAL.decl = d
			lx.pushScope()
			for _, decl := range typ.Decls {
				lx.pushDecl(decl)
			}
		}
	case 144:
		//line cc.y:1303
		{
			yylex.(*lexer).popScope()
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.decl = yyS[yypt-1].decl
			yyVAL.decl.Span = yyVAL.span
			if yyS[yypt-2].decls != nil {
				yylex.(*lexer).Errorf("cannot use pre-prototype definitions")
			}
			yyVAL.decl.Body = yyS[yypt-0].stmt
		}
	case 145:
		//line cc.y:1316
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.symlit = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 146:
		//line cc.y:1325
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.symlit = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			}
		}
	case 147:
		//line cc.y:1337
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Struct
		}
	case 148:
		//line cc.y:1342
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Union
		}
	case 149:
		//line cc.y:1349
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decor = yyS[yypt-0].decor
		}
	case 150:
		//line cc.y:1354
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			name := yyS[yypt-2].syntax
			expr := yyS[yypt-0].expr
			yyVAL.decor = func(t *Type) (*Type, Syntax) {
				t.Width = expr
				return t, name
			}
		}
	case 151:
		//line cc.y:1366
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = nil
			for _, decor := range yyS[yypt-1].decors {
				typ, name := decor(yyS[yypt-2].typ)
				yyVAL.decls = append(yyVAL.decls, &Decl{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       name,
					Type:       typ,
					Id:         nextId(),
				})
			}
			if yyS[yypt-1].decors == nil {
				yyVAL.decls = append(yyVAL.decls, &Decl{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Type:       yyS[yypt-2].typ,
					Id:         nextId(),
				})
			}
		}
	case 152:
		//line cc.y:1389
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Kind:       yyS[yypt-1].tk,
				Tag:        yyS[yypt-0].symlit,
				Id:         nextId(),
			})
		}
	case 153:
		//line cc.y:1399
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Kind:       yyS[yypt-4].tk,
				Tag:        yyS[yypt-3].syntax,
				Decls:      yyS[yypt-1].decls,
				Id:         nextId(),
			})
		}
	case 154:
		//line cc.y:1412
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Dot: yyS[yypt-0].symlit}
		}
	case 155:
		//line cc.y:1419
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Arrow, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].symlit}
		}
	case 156:
		//line cc.y:1424
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Op: Dot, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].symlit}
		}
	case 157:
		//line cc.y:1432
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-0].symlit, Id: nextId()})
		}
	case 158:
		//line cc.y:1437
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-4].syntax, Decls: yyS[yypt-2].decls, Id: nextId()})
		}
	case 159:
		//line cc.y:1444
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var x *Init
			if yyS[yypt-0].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr, Id: nextId()}
			}
			yyVAL.decl = &Decl{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &SymbolLiteral{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Value:      yyS[yypt-1].str,
					Id:         nextId(),
				},
				Init: x,
				Id:   nextId(),
			}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 160:
		//line cc.y:1465
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 161:
		//line cc.y:1473
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr, Id: nextId()}
		}
	case 162:
		//line cc.y:1478
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits, Id: nextId()}
		}
	case 163:
		//line cc.y:1485
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.inits = []*Init{}
		}
	case 164:
		//line cc.y:1490
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 165:
		//line cc.y:1495
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-3].inits, yyS[yypt-2].init)
		}
	case 166:
		//line cc.y:1501
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 167:
		//line cc.y:1506
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 168:
		//line cc.y:1513
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = yyS[yypt-0].init
		}
	case 169:
		//line cc.y:1518
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.init = yyS[yypt-0].init
			yyVAL.init.Prefix = yyS[yypt-2].prefixes
		}
	case 170:
		//line cc.y:1526
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Index: yyS[yypt-1].expr}
		}
	case 171:
		//line cc.y:1532
		{
			yyVAL.span = Span{}
		}
	case 172:
		//line cc.y:1536
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 173:
		//line cc.y:1541
		{
			yyVAL.span = Span{}
		}
	case 174:
		//line cc.y:1545
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 175:
		//line cc.y:1554
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.prefixes = []*Prefix{yyS[yypt-0].prefix}
		}
	case 176:
		//line cc.y:1559
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefixes = append(yyS[yypt-1].prefixes, yyS[yypt-0].prefix)
		}
	case 177:
		//line cc.y:1565
		{
			yyVAL.span = Span{}
			yyVAL.syntax = &EmptyLiteral{}
		}
	case 178:
		//line cc.y:1570
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntax = yyS[yypt-0].symlit
		}
	case 179:
		//line cc.y:1576
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 180:
		//line cc.y:1581
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 181:
		//line cc.y:1587
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 182:
		//line cc.y:1592
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 183:
		//line cc.y:1599
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = []*Expr{yyS[yypt-0].expr}
		}
	case 184:
		//line cc.y:1604
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 185:
		//line cc.y:1610
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 186:
		//line cc.y:1615
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = yyS[yypt-0].exprs
		}
	case 187:
		//line cc.y:1621
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 188:
		//line cc.y:1626
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 189:
		//line cc.y:1632
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 190:
		//line cc.y:1637
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.labels = append(yyS[yypt-1].labels, yyS[yypt-0].label)
		}
	case 191:
		//line cc.y:1644
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 192:
		//line cc.y:1649
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 193:
		//line cc.y:1655
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 194:
		//line cc.y:1660
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 195:
		//line cc.y:1667
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = []idecor{yyS[yypt-0].idec}
		}
	case 196:
		//line cc.y:1672
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idecs = append(yyS[yypt-2].idecs, yyS[yypt-0].idec)
		}
	case 197:
		//line cc.y:1678
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 198:
		//line cc.y:1683
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = yyS[yypt-0].idecs
		}
	case 199:
		//line cc.y:1690
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = []Syntax{yyS[yypt-0].syntax}
		}
	case 200:
		//line cc.y:1695
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.syntaxs = append(yyS[yypt-1].syntaxs, yyS[yypt-0].syntax)
		}
	case 201:
		//line cc.y:1701
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 202:
		//line cc.y:1706
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = yyS[yypt-0].syntaxs
		}
	case 203:
		//line cc.y:1713
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = []Syntax{yyS[yypt-0].syntax}
		}
	case 204:
		//line cc.y:1718
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.syntaxs = append(yyS[yypt-1].syntaxs, yyS[yypt-0].syntax)
		}
	case 205:
		//line cc.y:1724
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 206:
		//line cc.y:1729
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = yyS[yypt-0].syntaxs
		}
	case 207:
		//line cc.y:1736
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = []Syntax{yyS[yypt-0].syntax}
		}
	case 208:
		//line cc.y:1741
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.syntaxs = append(yyS[yypt-1].syntaxs, yyS[yypt-0].syntax)
		}
	case 209:
		//line cc.y:1747
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 210:
		//line cc.y:1752
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = yyS[yypt-0].syntaxs
		}
	case 211:
		//line cc.y:1759
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyS[yypt-0].decor)
		}
	case 212:
		//line cc.y:1765
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decors = append(yyS[yypt-2].decors, yyS[yypt-0].decor)
		}
	case 213:
		//line cc.y:1771
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 214:
		//line cc.y:1776
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = yyS[yypt-0].decors
		}
	case 215:
		//line cc.y:1783
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 216:
		//line cc.y:1788
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 217:
		//line cc.y:1794
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 218:
		//line cc.y:1799
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 219:
		//line cc.y:1806
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 220:
		//line cc.y:1811
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 221:
		//line cc.y:1818
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.syntaxs = []Syntax{
				&StringLiteral{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Value:      yyS[yypt-0].str,
					Id:         nextId(),
				},
			}
		}
	case 222:
		//line cc.y:1829
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.syntaxs = append(yyS[yypt-1].syntaxs, &StringLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyS[yypt-0].span},
				Value:      yyS[yypt-0].str,
				Id:         nextId(),
			})
		}
	}
	goto yystack /* stack new state and value */
}
