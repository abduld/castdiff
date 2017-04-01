//line cc.y:34
package cc

import __yyfmt__ "fmt"

//line cc.y:34
import (
	//"runtime/debug"
	_ "fmt"

	. "github.com/abduld/castdiff/cc/ast"
)

type typeClass struct {
	c Storage
	q TypeQual
	t *Type
}

type idecor struct {
	d func(*Type) (*Type, *SymbolLiteral)
	i Expr
}

var id int = 0

func nextId() int {
	id += 1
	return id
}

//line cc.y:63
type yySymType struct {
	yys      int
	abdecor  func(*Type) *Type
	fundecl  *FuncStmt
	decl     Stmt
	decls    []DeclStmt
	decor    func(*Type) (*Type, *SymbolLiteral)
	decors   []func(*Type) (*Type, *SymbolLiteral)
	expr     Expr
	exprs    []Expr
	idec     idecor
	idecs    []idecor
	init     Expr
	inits    []Expr
	label    *Label
	labels   []*Label
	span     Span
	prefix   Expr
	prefixes []Expr
	stmt     Stmt
	stmts    []Stmt
	str      string
	strs     []string
	tc       typeClass
	tk       TypeType
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
const tokFor = 57364
const tokGoto = 57365
const tokIf = 57366
const tokInline = 57367
const tokName = 57368
const tokInteger = 57369
const tokReal = 57370
const tokInt = 57371
const tokFloat = 57372
const tokLitChar = 57373
const tokLong = 57374
const tokString = 57375
const tokOffsetof = 57376
const tokRegister = 57377
const tokReturn = 57378
const tokShort = 57379
const tokSigned = 57380
const tokStatic = 57381
const tokStruct = 57382
const tokSwitch = 57383
const tokTypeName = 57384
const tokTypedef = 57385
const tokUnion = 57386
const tokUnsigned = 57387
const tokVaArg = 57388
const tokVoid = 57389
const tokVolatile = 57390
const tokWhile = 57391
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

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
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
	"tokFor",
	"tokGoto",
	"tokIf",
	"tokInline",
	"tokName",
	"tokInteger",
	"tokReal",
	"tokInt",
	"tokFloat",
	"tokLitChar",
	"tokLong",
	"tokString",
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
	"'}'",
	"';'",
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

const yyPrivate = 57344

const yyLast = 1655

var yyAct = [...]int{

	321, 7, 119, 127, 354, 314, 269, 32, 234, 222,
	276, 290, 204, 118, 249, 105, 106, 107, 108, 109,
	110, 111, 112, 113, 228, 201, 6, 355, 246, 124,
	50, 5, 178, 226, 116, 130, 232, 320, 236, 4,
	140, 136, 143, 145, 389, 125, 387, 380, 379, 375,
	138, 117, 368, 366, 349, 348, 346, 308, 300, 299,
	195, 317, 303, 97, 33, 254, 36, 65, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 165, 374, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 177, 137,
	135, 392, 141, 357, 35, 307, 182, 183, 2, 3,
	103, 99, 356, 166, 101, 100, 102, 98, 198, 189,
	197, 353, 244, 192, 179, 179, 196, 181, 188, 197,
	351, 180, 345, 344, 191, 196, 252, 221, 134, 122,
	142, 121, 117, 115, 386, 28, 131, 378, 53, 184,
	185, 40, 58, 377, 219, 197, 259, 47, 39, 203,
	31, 196, 132, 376, 57, 373, 372, 306, 42, 46,
	295, 43, 294, 262, 56, 205, 41, 44, 54, 51,
	207, 37, 55, 52, 45, 206, 48, 59, 216, 224,
	311, 60, 61, 62, 63, 64, 214, 212, 137, 233,
	235, 141, 239, 131, 135, 187, 141, 186, 293, 128,
	230, 268, 251, 241, 242, 220, 216, 253, 213, 132,
	203, 233, 247, 219, 129, 217, 291, 292, 32, 360,
	257, 359, 302, 243, 225, 230, 240, 238, 210, 142,
	264, 284, 301, 265, 142, 281, 263, 261, 215, 66,
	200, 279, 241, 217, 270, 258, 256, 235, 260, 247,
	277, 209, 208, 194, 288, 388, 266, 211, 364, 271,
	123, 131, 104, 250, 273, 34, 278, 230, 193, 179,
	97, 267, 309, 280, 237, 285, 305, 132, 296, 298,
	1, 38, 297, 12, 313, 312, 304, 203, 202, 139,
	49, 310, 324, 289, 316, 279, 58, 257, 323, 133,
	242, 235, 315, 325, 277, 144, 146, 239, 318, 69,
	70, 71, 255, 287, 126, 282, 283, 103, 99, 274,
	329, 101, 100, 102, 98, 350, 275, 347, 248, 245,
	352, 59, 29, 358, 27, 60, 61, 62, 63, 64,
	239, 330, 227, 199, 30, 190, 365, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 361, 362, 0, 0, 0, 383, 384, 385,
	382, 367, 0, 0, 369, 370, 0, 53, 0, 391,
	40, 58, 390, 393, 0, 0, 47, 39, 0, 120,
	97, 0, 381, 57, 8, 9, 10, 42, 46, 11,
	43, 25, 22, 56, 0, 41, 44, 54, 51, 0,
	37, 55, 52, 45, 24, 48, 59, 0, 0, 0,
	60, 61, 62, 63, 64, 72, 73, 67, 68, 69,
	70, 71, 0, 0, 0, 0, 0, 103, 99, 0,
	0, 101, 100, 102, 98, 0, 14, 0, 0, 0,
	0, 0, 0, 0, 0, 15, 16, 13, 0, 0,
	0, 17, 18, 21, 97, 0, 0, 0, 0, 20,
	19, 0, 23, 0, 0, 86, 87, 88, 89, 90,
	91, 92, 93, 94, 95, 96, 85, 371, 84, 83,
	82, 81, 80, 78, 79, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	0, 103, 99, 0, 0, 101, 100, 102, 98, 331,
	0, 0, 328, 327, 0, 332, 341, 0, 0, 333,
	342, 334, 0, 0, 0, 0, 0, 335, 336, 337,
	0, 343, 9, 10, 0, 0, 11, 97, 25, 22,
	0, 338, 0, 0, 0, 0, 339, 0, 0, 0,
	0, 24, 0, 0, 340, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 270, 80, 78, 79, 74, 75,
	76, 77, 72, 73, 67, 68, 69, 70, 71, 0,
	0, 0, 0, 14, 103, 99, 0, 0, 101, 100,
	102, 98, 15, 16, 13, 0, 0, 0, 17, 18,
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
	0, 0, 0, 53, 103, 99, 40, 58, 101, 100,
	102, 98, 47, 39, 0, 120, 0, 0, 0, 57,
	0, 0, 0, 42, 46, 0, 43, 0, 0, 56,
	0, 41, 44, 54, 51, 0, 37, 55, 52, 45,
	0, 48, 59, 0, 0, 0, 60, 61, 62, 63,
	64, 53, 0, 0, 40, 58, 0, 0, 0, 0,
	47, 39, 0, 120, 0, 0, 0, 57, 0, 0,
	0, 42, 46, 0, 43, 0, 0, 56, 0, 41,
	44, 54, 51, 0, 37, 55, 52, 45, 0, 48,
	59, 0, 0, 0, 60, 61, 62, 63, 64, 0,
	0, 0, 322, 0, 0, 28, 0, 0, 53, 0,
	0, 40, 58, 0, 0, 0, 0, 47, 39, 0,
	31, 0, 0, 0, 57, 0, 0, 0, 42, 46,
	0, 43, 0, 0, 56, 97, 41, 44, 54, 51,
	0, 37, 55, 52, 45, 0, 48, 59, 0, 0,
	272, 60, 61, 62, 63, 64, 0, 85, 0, 84,
	83, 82, 81, 80, 78, 79, 74, 75, 76, 77,
	72, 73, 67, 68, 69, 70, 71, 0, 0, 0,
	0, 0, 103, 99, 97, 0, 101, 100, 102, 98,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 26, 0, 0, 83,
	82, 81, 80, 78, 79, 74, 75, 76, 77, 72,
	73, 67, 68, 69, 70, 71, 97, 0, 0, 0,
	0, 103, 99, 0, 0, 101, 100, 102, 98, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 82, 81, 80, 78, 79, 74, 75, 76,
	77, 72, 73, 67, 68, 69, 70, 71, 97, 0,
	0, 0, 0, 103, 99, 0, 0, 101, 100, 102,
	98, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 81, 80, 78, 79, 74,
	75, 76, 77, 72, 73, 67, 68, 69, 70, 71,
	0, 0, 0, 0, 0, 103, 99, 0, 0, 101,
	100, 102, 98, 8, 9, 10, 0, 0, 11, 97,
	25, 22, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 24, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 218, 0, 78, 79,
	74, 75, 76, 77, 72, 73, 67, 68, 69, 70,
	71, 0, 97, 0, 0, 14, 103, 99, 0, 0,
	101, 100, 102, 98, 15, 16, 13, 0, 0, 0,
	17, 18, 21, 0, 291, 292, 0, 0, 20, 19,
	0, 23, 79, 74, 75, 76, 77, 72, 73, 67,
	68, 69, 70, 71, 0, 0, 0, 97, 0, 103,
	99, 0, 0, 101, 100, 102, 98, 8, 9, 10,
	0, 0, 11, 0, 25, 22, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 24, 74, 75,
	76, 77, 72, 73, 67, 68, 69, 70, 71, 0,
	218, 0, 0, 0, 103, 99, 0, 0, 101, 100,
	102, 98, 0, 8, 9, 10, 0, 0, 11, 14,
	25, 22, 0, 0, 0, 0, 0, 0, 15, 16,
	13, 0, 0, 24, 17, 18, 21, 0, 0, 0,
	0, 0, 20, 19, 0, 23, 8, 9, 10, 0,
	0, 11, 0, 25, 22, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 24, 0, 0, 0,
	0, 0, 0, 0, 15, 16, 13, 0, 0, 0,
	17, 18, 21, 0, 0, 8, 9, 10, 20, 19,
	11, 23, 25, 22, 97, 0, 0, 0, 14, 0,
	0, 0, 0, 0, 0, 24, 0, 15, 16, 13,
	0, 0, 0, 17, 18, 21, 0, 0, 218, 0,
	0, 20, 19, 0, 114, 0, 0, 0, 0, 0,
	0, 67, 68, 69, 70, 71, 0, 0, 0, 0,
	0, 103, 99, 0, 0, 101, 100, 102, 98, 0,
	0, 0, 17, 18, 21, 0, 0, 0, 0, 0,
	20, 19, 53, 23, 0, 40, 58, 0, 0, 0,
	231, 47, 39, 0, 120, 0, 0, 0, 57, 229,
	0, 0, 42, 46, 0, 43, 0, 0, 56, 0,
	41, 44, 54, 51, 0, 37, 55, 52, 45, 0,
	48, 59, 0, 0, 0, 60, 61, 62, 63, 64,
	363, 0, 0, 0, 53, 0, 0, 40, 58, 0,
	0, 0, 0, 47, 39, 0, 120, 0, 0, 0,
	57, 0, 0, 0, 42, 46, 0, 43, 0, 0,
	56, 0, 41, 44, 54, 51, 0, 37, 55, 52,
	45, 0, 48, 59, 0, 0, 0, 60, 61, 62,
	63, 64, 53, 0, 0, 40, 58, 0, 0, 0,
	0, 47, 39, 0, 120, 0, 0, 0, 57, 0,
	0, 0, 42, 46, 0, 43, 0, 0, 56, 0,
	41, 44, 54, 51, 0, 37, 55, 52, 45, 0,
	48, 59, 0, 0, 0, 60, 61, 62, 63, 64,
	53, 0, 0, 40, 58, 0, 0, 0, 0, 47,
	0, 0, 120, 0, 0, 0, 57, 0, 0, 0,
	42, 46, 0, 43, 0, 0, 56, 0, 41, 44,
	54, 0, 0, 0, 55, 0, 45, 53, 48, 59,
	0, 58, 0, 60, 61, 62, 63, 64, 0, 120,
	0, 0, 0, 57, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 56, 0, 0, 0, 54, 0, 0,
	0, 55, 0, 0, 0, 0, 59, 0, 0, 0,
	60, 61, 62, 63, 64,
}
var yyPact = [...]int{

	3, -1000, -1000, 1247, 899, -40, 189, 707, -1000, -1000,
	-1000, -1000, 239, 1247, 1247, 1247, 1247, 1247, 1247, 1247,
	1247, 1280, 39, 378, 37, -1000, -1000, -1000, 35, -1000,
	-1000, 237, 120, 1513, 1598, 1561, -1000, -1000, 245, 245,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 1247, 1247, 1247, 1247,
	1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247,
	1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247,
	1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247, 1247,
	-1000, -1000, 245, 245, -1000, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 378, 1513, 107, 105, 30, -1000,
	-1000, 1247, 252, 204, -49, 57, 190, -1000, 293, 120,
	-1000, -1000, -1000, 1598, 1561, -1000, -1000, 1598, -1000, 1561,
	-1000, -1000, -1000, -1000, 203, -1000, 202, 707, 230, 230,
	13, 13, 13, 1304, 1304, 350, 350, 350, 350, 1122,
	1167, 1079, 507, 1018, 976, 934, 165, 707, 707, 707,
	707, 707, 707, 707, 707, 707, 707, 707, 216, 189,
	97, 119, -1000, -1000, 96, 188, 1201, -1000, 125, 293,
	33, 30, 663, 89, -1000, -1000, 1413, 1247, 1201, 1513,
	120, 120, 293, -1000, 22, -1000, -1000, -1000, 1513, 247,
	1247, 32, -1000, -1000, 1319, 1247, 13, -1000, -43, 1247,
	30, 1413, 56, 1513, -1000, 139, 73, 186, -1000, -1000,
	177, -1000, 112, 707, -1000, 707, -1000, 195, -1000, 120,
	-1000, 57, 31, -1000, -1000, 842, -1000, 120, 185, -1000,
	180, 885, 1247, 618, -1000, 1097, 109, 125, 72, -1000,
	70, -1000, -1000, 1413, 125, 31, 293, 177, -1000, -1000,
	-1000, -50, -1000, -1000, -51, 182, -1000, 31, 159, -1000,
	-46, 247, -1000, -1000, 1247, 67, -1000, -3, -1000, 129,
	-1000, 245, 1247, -1000, -1000, -1000, -1000, 177, -1000, -1000,
	-1000, 120, 1247, -1000, -1000, 707, -1000, -1000, -47, 1201,
	-1000, -1000, -1000, 574, 794, -1000, 707, -1000, -1000, -1000,
	-1000, -1000, -1000, 525, -1000, -1000, -1000, 29, 28, -1000,
	-53, -1000, -54, -55, -1000, 26, 245, 17, 1247, 8,
	-1, 1247, 158, 156, 1247, 1247, -1000, 1465, -1000, -1000,
	219, 1247, -56, 1247, -57, -1000, 1247, 1247, 424, -1000,
	-1000, 66, 65, -1000, -17, -60, -1000, 63, -1000, 53,
	47, -1000, -61, -62, 1247, 1247, -1000, -1000, -1000, -1000,
	-1000, 44, -63, 207, -1000, -1000, -65, 1247, -1000, -1000,
	1, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 9, 355, 354, 14, 24, 37, 353, 352, 33,
	39, 344, 342, 28, 339, 338, 12, 10, 336, 329,
	1, 36, 27, 4, 326, 325, 26, 32, 35, 324,
	29, 8, 323, 38, 322, 313, 308, 11, 303, 302,
	6, 0, 5, 3, 300, 30, 104, 66, 40, 276,
	64, 41, 299, 50, 298, 25, 293, 2, 291, 34,
	13, 275, 290, 289, 284, 283, 282,
}
var yyR1 = [...]int{

	0, 62, 62, 10, 10, 10, 22, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 42, 42, 42, 63, 40, 35, 35, 35,
	41, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 39, 1, 1, 1,
	2, 2, 2, 16, 16, 16, 16, 16, 5, 5,
	5, 5, 28, 28, 44, 44, 44, 44, 44, 44,
	45, 45, 45, 45, 45, 45, 45, 46, 46, 46,
	46, 46, 46, 46, 46, 46, 47, 47, 48, 48,
	61, 57, 57, 57, 57, 57, 60, 59, 6, 12,
	11, 11, 11, 64, 3, 43, 43, 58, 58, 17,
	17, 13, 61, 61, 37, 20, 20, 61, 61, 4,
	24, 31, 31, 33, 33, 33, 34, 34, 32, 32,
	37, 66, 66, 65, 65, 38, 38, 49, 49, 23,
	23, 21, 21, 26, 26, 27, 27, 7, 7, 36,
	36, 8, 8, 9, 9, 29, 29, 30, 30, 54,
	54, 55, 55, 50, 50, 51, 51, 52, 52, 53,
	53, 18, 18, 19, 19, 14, 14, 25, 25, 15,
	15, 56, 56,
}
var yyR2 = [...]int{

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
var yyChk = [...]int{

	-1000, -62, 105, 106, -10, -22, -26, -20, 26, 27,
	28, 31, -56, 89, 78, 87, 88, 93, 94, 102,
	101, 95, 34, 104, 46, 33, 107, -11, 6, -12,
	-3, 21, -57, -50, -61, -46, -47, 42, -58, 19,
	12, 37, 29, 32, 38, 45, 30, 18, 47, -44,
	-45, 40, 44, 9, 39, 43, 35, 25, 13, 48,
	52, 53, 54, 55, 56, 107, 60, 87, 88, 89,
	90, 91, 85, 86, 81, 82, 83, 84, 79, 80,
	78, 77, 76, 75, 74, 72, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 71, 50, 104, 98,
	102, 101, 103, 97, 33, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, 104, 104, -59, -22, -60, -57,
	21, 104, 104, 33, -30, -16, -29, -43, 89, 104,
	-28, 26, 42, -61, -46, -47, -51, -50, -53, -52,
	-48, -47, -46, -43, -49, -43, -49, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -22, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -27, -26,
	-27, -22, -43, -43, -59, -59, 100, 100, -1, 89,
	-2, 104, -20, 26, 59, 109, 104, 98, 61, -7,
	60, -55, -54, -45, -16, -51, -53, -48, 59, 59,
	73, 51, 100, 99, 100, 60, -20, -33, 59, 98,
	-55, 104, -1, 60, 100, -10, -9, -8, -5, 26,
	-60, 17, -21, -20, -31, -20, -33, -64, -6, -57,
	-28, -16, -16, -45, 100, -14, -13, -60, -15, -4,
	26, -20, 104, -20, 108, -34, -21, -1, -9, 100,
	-59, 108, 100, 60, -1, -16, 89, 104, 99, -40,
	59, -30, 108, -13, -19, -18, -17, -16, -49, -43,
	-65, 60, -25, -24, 61, -27, 100, -32, -31, -38,
	-37, 97, 98, 99, 100, 100, -5, -55, -63, 109,
	109, 60, 73, 108, -4, -20, 100, 108, 60, -66,
	-37, 61, -43, -20, -42, -17, -20, 108, -31, 99,
	-6, -41, 108, -36, -39, -35, 109, 8, 7, -40,
	-22, 4, 10, 14, 16, 22, 23, 24, 36, 41,
	49, 11, 15, 26, 104, 104, 109, -42, 109, 109,
	-41, 104, -43, 104, -23, -22, 104, 104, -20, 73,
	73, -22, -22, 5, 49, -23, 109, -22, 109, -22,
	-22, 73, 100, 100, 104, 109, 100, 100, 100, 109,
	109, -22, -23, -41, -41, -41, 100, 109, 58, 109,
	-23, -41, 100, -41,
}
var yyDef = [...]int{

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
var yyTok1 = [...]int{

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
var yyTok2 = [...]int{

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
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:224
		{
			yylex.(*lexer).prog = &Prog{Decls: yyDollar[2].stmts, Id: nextId()}
			return 0
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:229
		{
			return 0
		}
	case 3:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:234
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:239
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmts...)
		}
	case 5:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:244
		{
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:249
		{
			yyVAL.span = yyDollar[1].span
			if len(yyDollar[1].exprs) == 1 {
				yyVAL.expr = yyDollar[1].exprs[0]
				break
			}
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Callee: &SymbolLiteral{Value: ","}, Args: yyDollar[1].exprs}
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:261
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &SymbolLiteral{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:270
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &IntegerLiteral{
				Value:      yyDollar[1].intlit.Value,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:279
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &RealLiteral{
				Value:      yyDollar[1].reallit.Value,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:288
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &CharLiteral{
				Value:      yyDollar[1].charlit.Value,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:297
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = &TupleExpr{
				Args:       yyDollar[1].exprs,
				Brackets:   None,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:307
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Add,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 13:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:318
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Sub,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:329
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Mul,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 15:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:340
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Div,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:351
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Mod,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 17:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:362
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Lsh,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:373
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Rsh,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:384
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Lt,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 20:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:395
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Gt,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:406
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         LtEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 22:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:417
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         GtEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:428
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         EqEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:439
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         NotEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:450
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         And,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 26:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:461
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Xor,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 27:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:472
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Or,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:483
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         AndAnd,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 29:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:494
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &BinaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         OrOr,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 30:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:505
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)

			yyVAL.expr = &CondExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Cond:       yyDollar[1].expr,
				Then:       yyDollar[3].expr,
				Else:       yyDollar[5].expr,
			}
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:517
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Eq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:528
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         AddEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:539
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         SubEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 34:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:550
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         MulEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:561
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         DivEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:572
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         ModEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:583
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         LshEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:594
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         RshEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 39:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:605
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         AndEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:616
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         XorEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:627
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &AssignExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         OrEq,
				Left:       yyDollar[1].expr,
				Right:      yyDollar[3].expr,
			}
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:638
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Indir,
				Arg:        yyDollar[2].expr,
			}
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:648
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Addr,
				Arg:        yyDollar[2].expr,
			}
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:658
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Plus,
				Arg:        yyDollar[2].expr,
			}
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:668
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Minus,
				Arg:        yyDollar[2].expr,
			}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:678
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Not,
				Arg:        yyDollar[2].expr,
			}
		}
	case 47:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:688
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         Twid,
				Arg:        yyDollar[2].expr,
			}
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:698
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         PreInc,
				Arg:        yyDollar[2].expr,
			}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:708
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Op:         PreDec,
				Arg:        yyDollar[2].expr,
			}
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:718
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Callee: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "sizeof",
				},
				Args: []Expr{yyDollar[2].expr},
			}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:732
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Callee: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "sizeof", // sizeof type
				},
				Args: []Expr{yyDollar[3].typ},
			}
		}
	case 52:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:746
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Callee: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "offsetof",
				},
				Args: []Expr{yyDollar[3].typ, yyDollar[5].expr},
			}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:760
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &CastExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Type:       yyDollar[2].typ,
				Expr:       yyDollar[4].expr,
			}
		}
	case 54:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:770
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &CastExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Type:       yyDollar[2].typ,
				Expr: &TupleExpr{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Brackets:   Brace,
					Args:       yyDollar[4].inits,
					Id:         nextId(),
				},
			}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:785
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &ParenExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Expr:       yyDollar[2].expr,
				Id:         nextId(),
			}
		}
	case 56:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:794
		{

			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.expr = &CUDALaunchExpr{
				SyntaxInfo:   SyntaxInfo{Span: yyVAL.span},
				Callee:       yyDollar[1].expr,
				LaunchParams: yyDollar[3].exprs,
				Args:         yyDollar[6].exprs,
				Id:           nextId(),
			}
		}
	case 57:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:806
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Callee:     yyDollar[1].expr,
				Args:       yyDollar[3].exprs,
				Id:         nextId(),
			}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:816
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.expr = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Callee: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "array_index",
				},
				Args: []Expr{yyDollar[1].expr, yyDollar[3].expr},
				Id:   nextId(),
			}
		}
	case 59:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:830
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Op:         PostInc,
				Arg:        yyDollar[1].expr,
				Id:         nextId(),
			}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:840
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = &UnaryExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Op:         PostDec,
				Arg:        yyDollar[1].expr,
				Id:         nextId(),
			}
		}
	case 61:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:850
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.expr = &VaArgExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Type:       yyDollar[5].typ,
				Arg:        yyDollar[3].expr,
				Id:         nextId(),
			}
		}
	case 62:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:861
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 63:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:866
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = yyDollar[1].stmts
			for _, d := range yyDollar[2].stmts {
				yyVAL.stmts = append(yyVAL.stmts, d)
			}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:874
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmt)
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:881
		{
			yylex.(*lexer).pushScope()
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:885
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &BlockStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Stmts: yyDollar[3].stmts}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:893
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Name: yyDollar[2].expr, IsCase: true}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:898
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Name: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "default",
				},
			}
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:911
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.label = &Label{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Name: &SymbolLiteral{
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Value:      yyDollar[1].str,
				},
			}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:926
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = yyDollar[2].stmt
			yyVAL.stmt = &LabeledStmt{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Labels:     yyDollar[1].labels,
				Expr:       yyDollar[2].stmt,
			}
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:939
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:944
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()}
		}
	case 73:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:949
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &EmptyStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()}
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:954
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:959
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &ExprStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Expr: yyDollar[1].expr}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:964
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &BlockStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Stmts: yyDollar[2].stmts}
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:969
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &BreakStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:974
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &ContinueStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()}
		}
	case 79:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:979
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &DoStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Body: yyDollar[2].stmt, Cond: yyDollar[5].expr}
		}
	case 80:
		yyDollar = yyS[yypt-9 : yypt+1]
		//line cc.y:984
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[9].span)
			yyVAL.stmt = &ForStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:     nextId(),
				Init:   yyDollar[3].expr,
				Cond:   yyDollar[5].expr,
				Update: yyDollar[7].expr,
				Body:   yyDollar[9].stmt,
			}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:995
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &GotoStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Target: yyDollar[2].symlit}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1000
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &IfStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Cond: yyDollar[3].expr, Then: yyDollar[5].stmt}
		}
	case 83:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line cc.y:1005
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[7].span)
			yyVAL.stmt = &IfStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Cond: yyDollar[3].expr, Then: yyDollar[5].stmt, Else: yyDollar[7].stmt}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1010
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmt = &ReturnStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Value: yyDollar[2].expr}
		}
	case 85:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1015
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &SwitchStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Cond: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1020
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = &WhileStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId(), Cond: yyDollar[3].expr, Body: yyDollar[5].stmt}
		}
	case 87:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1027
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1032
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].syntaxs)
			abdecor := yyDollar[3].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: Ptr, Base: t, Qual: q, Id: nextId()})
			}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1041
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.abdecor = yyDollar[1].abdecor
		}
	case 90:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1048
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			abdecor := yyDollar[1].abdecor
			decls := yyDollar[3].stmts
			span := yyVAL.span
			for _, decl := range decls {
				switch decl := decl.(type) {
				default:
					break
				case *DeclStmt:
					t := decl.Type
					if t != nil {
						if t.Type == TypedefType && t.Base != nil {
							t = t.Base
						}
						if t.Type == Array {
							if t.Size == nil {
								t = t.Base
							}
							decl.Type = &Type{Type: Ptr, Base: t, Id: nextId()}
						}
					}
				}
			}
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Func, Base: t, Stmts: decls, Id: nextId()})
			}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1077
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			abdecor := yyDollar[1].abdecor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Array, Base: t, Size: expr, Id: nextId()})
			}

		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1088
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.abdecor = yyDollar[2].abdecor
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1096
		{
			yyVAL.span = yyDollar[1].span
			name := yyDollar[1].symlit
			yyVAL.decor = func(t *Type) (*Type, *SymbolLiteral) { return t, name }
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1102
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			_, q, _ := splitTypeWords(yyDollar[2].syntaxs)
			decor := yyDollar[3].decor
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, *SymbolLiteral) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Ptr, Base: t, Qual: q, Id: nextId()})
			}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1112
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decor = yyDollar[2].decor
		}
	case 96:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1117
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			decls := yyDollar[3].stmts
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, *SymbolLiteral) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Func, Base: t, Stmts: decls, Id: nextId()})
			}
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1127
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			decor := yyDollar[1].decor
			span := yyVAL.span
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, *SymbolLiteral) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Type: Array, Base: t, Size: expr, Id: nextId()})
			}
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1140
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &SymbolLiteral{
					Value:      yyDollar[1].str,
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				},
				Id: nextId(),
			}
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1153
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmt = &DeclStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyDollar[2].abdecor(yyDollar[1].typ), Id: nextId()}
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1158
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			typ, name := yyDollar[2].decor(yyDollar[1].typ)
			yyVAL.stmt = &DeclStmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Id: nextId()}
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1164
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmt = &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &SymbolLiteral{
					Value:      "...",
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
				},
				Id: nextId(),
			}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1180
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idec = idecor{yyDollar[1].decor, nil}
		}
	case 103:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1185
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idec = idecor{yyDollar[1].decor, yyDollar[3].init}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1193
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1202
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1211
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1220
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1229
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1238
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1250
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 111:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1259
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 112:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1268
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1277
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1286
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1295
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &SymbolLiteral{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1304
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				Value:      yyDollar[1].str,
				Id:         nextId(),
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
			}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1316
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyDollar[1].str,
				Id:         nextId(),
			}
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1325
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyDollar[1].str,
				Id:         nextId(),
			}
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1334
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "int",
				Id:         nextId(),
			}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1343
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "long",
				Id:         nextId(),
			}
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1352
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "signed",
				Id:         nextId(),
			}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1361
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "unsigned",
				Id:         nextId(),
			}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1370
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "float",
				Id:         nextId(),
			}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1379
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "real",
				Id:         nextId(),
			}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1388
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = &LanguageKeyword{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      "void",
				Id:         nextId(),
			}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1399
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = yyDollar[1].syntax
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1404
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = yyDollar[1].syntax
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1411
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = yyDollar[1].syntax
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1416
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = yyDollar[1].syntax
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1424
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.typ = yyDollar[1].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{
					Type: TypedefType,
					Name: &SymbolLiteral{
						Value:      yyDollar[1].str,
						Id:         nextId(),
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					},
					Id: nextId(),
				}
			}
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1448
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(
				append(yyDollar[1].syntaxs, &LanguageKeyword{
					Value:      "int",
					Id:         nextId(),
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				}))
		}
	case 132:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1458
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyDollar[1].syntaxs, yyDollar[3].syntaxs...))
			yyVAL.tc.t = yyDollar[2].typ
		}
	case 133:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1464
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyDollar[1].syntaxs = append(yyDollar[1].syntaxs, yyDollar[2].syntax)
			yyDollar[1].syntaxs = append(yyDollar[1].syntaxs, yyDollar[3].syntaxs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyDollar[1].syntaxs)
		}
	case 134:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1471
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyDollar[2].syntaxs)
			yyVAL.tc.t = yyDollar[1].typ
		}
	case 135:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1477
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var ts []Syntax
			ts = append(ts, yyDollar[1].syntax)
			ts = append(ts, yyDollar[2].syntaxs...)
			//PrintStack()
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 136:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1489
		{
			yyVAL.span = yyDollar[1].span
			if yyDollar[1].tc.c != 0 {
				yylex.(*lexer).Errorf("%v not allowed here", yyDollar[1].tc.c)
			}
			if yyDollar[1].tc.q != 0 && yyDollar[1].tc.q != Const && yyDollar[1].tc.q != Volatile {
				yylex.(*lexer).Errorf("%v ignored here (TODO)?", yyDollar[1].tc.q)
			}
			yyVAL.typ = yyDollar[1].tc.t
		}
	case 137:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1502
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yyDollar[2].abdecor(yyDollar[1].typ)
		}
	case 138:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1510
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			// TODO: use $1.q
			yyVAL.stmts = nil
			for _, idec := range yyDollar[2].idecs {
				typ, name := idec.d(yyDollar[1].tc.t)
				d := &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       name,
					Type:       typ,
					Storage:    yyDollar[1].tc.c,
					Init:       idec.i,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.stmts = append(yyVAL.stmts, d)
			}
			if yyDollar[2].idecs == nil {
				d := &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       &SymbolLiteral{},
					Type:       yyDollar[1].tc.t,
					Storage:    yyDollar[1].tc.c,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.stmts = append(yyVAL.stmts, d)
			}
		}
	case 139:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1543
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			// TODO: use $1.q
			yyVAL.stmts = nil
			for _, idec := range yyDollar[2].idecs {
				typ, name := idec.d(yyDollar[1].tc.t)
				d := lx.lookupDecl(name)
				if d == nil {
					d = &DeclStmt{
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
						Name:       name,
						Type:       typ,
						Storage:    yyDollar[1].tc.c,
						Init:       idec.i,
						Id:         nextId(),
					}
					lx.pushDecl(d)
				} else {
					//d.Span = $<span>$
					if idec.i != nil {
						switch d := d.(type) {
						case *DeclStmt:
							d.Init = idec.i
						}
					}
				}
				yyVAL.stmts = append(yyVAL.stmts, d)
			}
			if yyDollar[2].idecs == nil {
				d := &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       &SymbolLiteral{},
					Type:       yyDollar[1].tc.t,
					Storage:    yyDollar[1].tc.c,
					Id:         nextId(),
				}
				lx.pushDecl(d)
				yyVAL.stmts = append(yyVAL.stmts, d)
			}
		}
	case 140:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1587
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 141:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1592
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = []Stmt{yyDollar[1].stmt}
		}
	case 142:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1597
		{
			yyVAL.stmts = yyDollar[4].stmts
		}
	case 143:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1603
		{
			lx := yylex.(*lexer)
			typ, name := yyDollar[2].decor(yyDollar[1].tc.t)
			if typ.Type != Func {
				yylex.(*lexer).Errorf("invalid function definition")
				return 0
			}
			d := lx.lookupDecl(name)
			if d == nil {
				d = &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       name,
					Type:       typ,
					Storage:    yyDollar[1].tc.c,
					Id:         nextId(),
				}
				lx.pushDecl(d)
			} else {
				switch d := d.(type) {
				case *DeclStmt:
					d.Type = typ
				}
			}
			yyVAL.decl = d
			lx.pushScope()
			for _, decl := range typ.Stmts {
				lx.pushDecl(decl)
			}
		}
	case 144:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1633
		{
			yylex.(*lexer).popScope()
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.stmt = yyDollar[4].decl
			if yyDollar[3].stmts != nil {
				yylex.(*lexer).Errorf("cannot use pre-prototype definitions")
			}
			switch decl := yyVAL.stmt.(type) {
			case *DeclStmt:
				yyVAL.stmt = &FuncStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       decl.Name,
					ReturnType: decl.Type,
					IsDecl:     false,
					Storage:    decl.Storage,
					Body:       yyDollar[5].stmt,
				}
			case *FuncStmt:
				yyVAL.stmt = &FuncStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       decl.Name,
					ReturnType: decl.ReturnType,
					Args:       decl.Args,
					IsDecl:     false,
					Storage:    decl.Storage,
					Body:       yyDollar[5].stmt,
				}
			}
		}
	case 145:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1665
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.symlit = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyDollar[1].str,
				Id:         nextId(),
			}
		}
	case 146:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1674
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.symlit = &SymbolLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Value:      yyDollar[1].str,
				Id:         nextId(),
			}
		}
	case 147:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1686
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Struct
		}
	case 148:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1691
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.tk = Union
		}
	case 149:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1698
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decor = yyDollar[1].decor
		}
	case 150:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1703
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			name := yyDollar[1].syntax
			expr := yyDollar[3].expr
			yyVAL.decor = func(t *Type) (*Type, *SymbolLiteral) {
				t.Size = expr
				switch name := name.(type) {
				case *SymbolLiteral:
					return t, name
				default:
					return t, &SymbolLiteral{
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
						Value:      name.String(),
						Id:         nextId(),
					}
				}
			}
		}
	case 151:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1724
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmts = nil
			for _, decor := range yyDollar[2].decors {
				typ, name := decor(yyDollar[1].typ)
				yyVAL.stmts = append(yyVAL.stmts, &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Name:       name,
					Type:       typ,
					Id:         nextId(),
				})
			}
			if yyDollar[2].decors == nil {
				yyVAL.stmts = append(yyVAL.stmts, &DeclStmt{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Type:       yyDollar[1].typ,
					Id:         nextId(),
				})
			}
		}
	case 152:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1747
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Type:       yyDollar[1].tk,
				Tag:        yyDollar[2].symlit,
				Id:         nextId(),
			})
		}
	case 153:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1757
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Type:       yyDollar[1].tk,
				Tag:        yyDollar[2].syntax,
				Stmts:      yyDollar[4].stmts,
				Id:         nextId(),
			})
		}
	case 154:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1770
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefix = &DotExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Struct: &EmptyLiteral{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
				},
				Member: yyDollar[2].symlit,
			}
		}
	case 155:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1785
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &ArrowExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Ptr:        yyDollar[1].expr,
				Member:     yyDollar[3].symlit,
			}
		}
	case 156:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1795
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.expr = &DotExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Id:         nextId(),
				Struct:     yyDollar[1].expr,
				Member:     yyDollar[3].symlit,
			}
		}
	case 157:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1808
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: Enum, Tag: yyDollar[2].symlit, Id: nextId()})
		}
	case 158:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line cc.y:1813
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[6].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: Enum, Tag: yyDollar[2].syntax, Stmts: yyDollar[4].stmts, Id: nextId()})
		}
	case 159:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1820
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			var x *Init
			if yyDollar[2].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[2].expr, Id: nextId()}
			}
			yyVAL.stmt = &DeclStmt{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Name: &SymbolLiteral{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Value:      yyDollar[1].str,
					Id:         nextId(),
				},
				Init: x,
				Id:   nextId(),
			}
			yylex.(*lexer).pushDecl(yyVAL.stmt)
		}
	case 160:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1841
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.expr = yyDollar[2].expr
		}
	case 161:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1849
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyDollar[1].expr, Id: nextId()}
		}
	case 162:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1854
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyDollar[1].inits, Id: nextId()}
		}
	case 163:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1861
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.inits = []Expr{
				&BracedExpr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Id: nextId()},
			}
		}
	case 164:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line cc.y:1868
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[4].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 165:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line cc.y:1873
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[5].span)
			yyVAL.inits = append(yyDollar[2].inits, yyDollar[3].init)
		}
	case 166:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1879
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 167:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1884
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.inits = append(yyDollar[1].inits, yyDollar[2].init)
		}
	case 168:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1891
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.init = yyDollar[1].init
		}
	case 169:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1896
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.init = &TupleExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Brackets:   None,
				Id:         nextId(),
				Args:       append(yyDollar[1].prefixes, yyDollar[3].init),
			}
		}
	case 170:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1908
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.prefix = &CallExpr{
				SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Callee: &LanguageKeyword{
					SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					Id:         nextId(),
					Value:      "array_index",
				},
				Args: []Expr{yyDollar[2].expr},
				Id:   nextId(),
			}
		}
	case 171:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1923
		{
			yyVAL.span = Span{}
		}
	case 172:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1927
		{
			yyVAL.span = yyDollar[1].span
		}
	case 173:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1932
		{
			yyVAL.span = Span{}
		}
	case 174:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1936
		{
			yyVAL.span = yyDollar[1].span
		}
	case 175:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1945
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.prefixes = []Expr{yyDollar[1].prefix}
		}
	case 176:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:1950
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.prefixes = append(yyDollar[1].prefixes, yyDollar[2].prefix)
		}
	case 177:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1956
		{
			yyVAL.span = Span{}
			yyVAL.syntax = &EmptyLiteral{}
		}
	case 178:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1961
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntax = yyDollar[1].symlit
		}
	case 179:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1967
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 180:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1972
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 181:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:1978
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 182:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1983
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 183:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:1990
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = []Expr{yyDollar[1].expr}
		}
	case 184:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:1995
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.exprs = append(yyDollar[1].exprs, yyDollar[3].expr)
		}
	case 185:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2001
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 186:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2006
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.exprs = yyDollar[1].exprs
		}
	case 187:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2012
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 188:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2017
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmts...)
		}
	case 189:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2023
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 190:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2028
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.labels = append(yyDollar[1].labels, yyDollar[2].label)
		}
	case 191:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2035
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = []Stmt{yyDollar[1].stmt}
		}
	case 192:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:2040
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 193:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2046
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 194:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2051
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 195:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2058
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = []idecor{yyDollar[1].idec}
		}
	case 196:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:2063
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.idecs = append(yyDollar[1].idecs, yyDollar[3].idec)
		}
	case 197:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2069
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 198:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2074
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.idecs = yyDollar[1].idecs
		}
	case 199:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2081
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = []Syntax{yyDollar[1].syntax}
		}
	case 200:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2086
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.syntaxs = append(yyDollar[1].syntaxs, yyDollar[2].syntax)
		}
	case 201:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2092
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 202:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2097
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = yyDollar[1].syntaxs
		}
	case 203:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2104
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = []Syntax{yyDollar[1].syntax}
		}
	case 204:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2109
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.syntaxs = append(yyDollar[1].syntaxs, yyDollar[2].syntax)
		}
	case 205:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2115
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 206:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2120
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = yyDollar[1].syntaxs
		}
	case 207:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2127
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = []Syntax{yyDollar[1].syntax}
		}
	case 208:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2132
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.syntaxs = append(yyDollar[1].syntaxs, yyDollar[2].syntax)
		}
	case 209:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2138
		{
			yyVAL.span = Span{}
			yyVAL.syntaxs = nil
		}
	case 210:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2143
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.syntaxs = yyDollar[1].syntaxs
		}
	case 211:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2150
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyDollar[1].decor)
		}
	case 212:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:2156
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.decors = append(yyDollar[1].decors, yyDollar[3].decor)
		}
	case 213:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2162
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 214:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2167
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.decors = yyDollar[1].decors
		}
	case 215:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2174
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = yyDollar[1].stmts
		}
	case 216:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2179
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[2].stmts...)
		}
	case 217:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line cc.y:2185
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 218:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2190
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.expr = yyDollar[1].expr
		}
	case 219:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2197
		{
			yyVAL.span = yyDollar[1].span
			yyVAL.stmts = []Stmt{yyDollar[1].stmt}
		}
	case 220:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line cc.y:2202
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[3].span)
			yyVAL.stmts = append(yyDollar[1].stmts, yyDollar[3].stmt)
		}
	case 221:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line cc.y:2209
		{
			yyVAL.span = yyDollar[1].span
			if yyDollar[1].strlit == nil {
				yyVAL.exprs = []Expr{
					&StringLiteral{
						Value:      "",
						Id:         nextId(),
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					},
				}
			} else {
				yyVAL.exprs = []Expr{
					&StringLiteral{
						Value:      yyDollar[1].strlit.Value,
						Id:         nextId(),
						SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
					},
				}
			}
		}
	case 222:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line cc.y:2230
		{
			yyVAL.span = span(yyDollar[1].span, yyDollar[2].span)
			yyVAL.exprs = append(yyDollar[1].exprs, &StringLiteral{
				SyntaxInfo: SyntaxInfo{Span: yyDollar[2].span},
				Value:      yyDollar[2].strlit.Value,
				Id:         nextId(),
			})
		}
	}
	goto yystack /* stack new state and value */
}
