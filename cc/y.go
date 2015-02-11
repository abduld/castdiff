//line cc.y:34
package cc

import __yyfmt__ "fmt"

//line cc.y:34
type typeClass struct {
	c Storage
	q TypeQual
	t *Type
}

type idecor struct {
	d func(*Type) (*Type, string)
	i *Init
}

//line cc.y:49
type yySymType struct {
	yys      int
	abdecor  func(*Type) *Type
	decl     *Decl
	decls    []*Decl
	decor    func(*Type) (*Type, string)
	decors   []func(*Type) (*Type, string)
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
const tokNumber = 57373
const tokOffsetof = 57374
const tokRegister = 57375
const tokReturn = 57376
const tokShort = 57377
const tokSigned = 57378
const tokStatic = 57379
const tokStruct = 57380
const tokSwitch = 57381
const tokTypeName = 57382
const tokTypedef = 57383
const tokUnion = 57384
const tokUnsigned = 57385
const tokVaArg = 57386
const tokVoid = 57387
const tokVolatile = 57388
const tokWhile = 57389
const tokString = 57390
const tokLCuBrk = 57391
const tokRCuBrk = 57392
const tokDevice = 57393
const tokHost = 57394
const tokGlobal = 57395
const tokShared = 57396
const tokRestrict = 57397
const tokShift = 57398
const tokElse = 57399
const tokAddEq = 57400
const tokSubEq = 57401
const tokMulEq = 57402
const tokDivEq = 57403
const tokModEq = 57404
const tokLshEq = 57405
const tokRshEq = 57406
const tokAndEq = 57407
const tokXorEq = 57408
const tokOrEq = 57409
const tokOrOr = 57410
const tokAndAnd = 57411
const tokEqEq = 57412
const tokNotEq = 57413
const tokLtEq = 57414
const tokGtEq = 57415
const tokLsh = 57416
const tokRsh = 57417
const tokCast = 57418
const tokSizeof = 57419
const tokUnary = 57420
const tokDec = 57421
const tokInc = 57422
const tokArrow = 57423
const startExpr = 57424
const startProg = 57425
const tokEOF = 57426

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
	"tokNumber",
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
	"'('",
	"')'",
	"tokDec",
	"tokInc",
	"tokArrow",
	"startExpr",
	"startProg",
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
	-1, 124,
	59, 101,
	108, 101,
	-2, 186,
	-1, 142,
	58, 177,
	-2, 151,
	-1, 144,
	58, 177,
	-2, 156,
	-1, 246,
	108, 212,
	-2, 176,
	-1, 278,
	72, 177,
	-2, 92,
}

const yyNprod = 222
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1659

var yyAct = []int{

	320, 7, 118, 126, 353, 313, 268, 31, 233, 221,
	275, 289, 203, 117, 104, 105, 106, 107, 108, 109,
	110, 111, 112, 227, 248, 200, 6, 354, 245, 49,
	177, 5, 123, 225, 115, 129, 231, 319, 235, 4,
	139, 142, 144, 137, 124, 135, 388, 386, 379, 378,
	116, 374, 32, 367, 365, 348, 347, 345, 299, 307,
	298, 194, 316, 302, 253, 64, 35, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 160, 161, 162, 163, 164, 136, 166, 167, 168,
	169, 170, 171, 172, 173, 174, 175, 176, 34, 134,
	197, 140, 3, 2, 391, 181, 182, 306, 196, 385,
	195, 243, 165, 218, 377, 96, 258, 376, 375, 372,
	371, 305, 191, 178, 178, 294, 180, 187, 179, 293,
	196, 133, 195, 141, 261, 188, 223, 196, 213, 195,
	373, 116, 211, 96, 186, 185, 190, 356, 183, 184,
	71, 72, 66, 67, 68, 69, 70, 202, 355, 352,
	350, 130, 102, 98, 344, 97, 343, 100, 99, 101,
	310, 131, 251, 220, 121, 120, 114, 205, 204, 206,
	66, 67, 68, 69, 70, 136, 292, 215, 267, 212,
	102, 98, 218, 97, 359, 100, 99, 101, 232, 234,
	140, 238, 358, 134, 301, 140, 290, 291, 209, 229,
	283, 250, 240, 241, 219, 215, 252, 300, 202, 127,
	232, 246, 280, 262, 216, 214, 65, 31, 57, 256,
	128, 242, 141, 224, 229, 239, 237, 141, 130, 263,
	199, 269, 264, 208, 207, 193, 387, 210, 131, 122,
	278, 240, 216, 103, 257, 255, 234, 259, 246, 276,
	363, 58, 249, 287, 130, 33, 59, 60, 61, 62,
	63, 270, 96, 272, 131, 192, 229, 308, 178, 279,
	236, 297, 284, 1, 37, 304, 295, 11, 201, 277,
	138, 296, 48, 312, 311, 202, 265, 323, 132, 288,
	309, 322, 324, 315, 278, 303, 256, 266, 254, 241,
	234, 314, 96, 276, 286, 125, 238, 317, 281, 102,
	98, 282, 97, 273, 100, 99, 101, 143, 145, 328,
	274, 247, 244, 28, 349, 26, 346, 226, 198, 351,
	29, 189, 357, 0, 0, 0, 0, 0, 0, 238,
	329, 68, 69, 70, 0, 364, 0, 0, 0, 102,
	98, 0, 97, 0, 100, 99, 101, 0, 0, 0,
	0, 360, 361, 0, 0, 0, 382, 383, 384, 381,
	366, 0, 0, 368, 369, 0, 52, 0, 390, 39,
	57, 389, 392, 0, 0, 46, 38, 0, 119, 45,
	0, 380, 0, 56, 41, 10, 42, 8, 9, 21,
	55, 0, 40, 43, 53, 50, 0, 36, 54, 51,
	44, 23, 47, 58, 0, 24, 0, 0, 59, 60,
	61, 62, 63, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 13, 0, 0, 0, 0, 0,
	0, 0, 0, 14, 15, 12, 0, 0, 0, 16,
	17, 20, 96, 0, 0, 0, 22, 0, 19, 18,
	0, 0, 0, 85, 86, 87, 88, 89, 90, 91,
	92, 93, 94, 95, 84, 370, 83, 82, 81, 80,
	79, 77, 78, 73, 74, 75, 76, 71, 72, 66,
	67, 68, 69, 70, 0, 0, 96, 0, 0, 102,
	98, 0, 97, 0, 100, 99, 101, 85, 86, 87,
	88, 89, 90, 91, 92, 93, 94, 95, 84, 0,
	83, 82, 81, 80, 79, 77, 78, 73, 74, 75,
	76, 71, 72, 66, 67, 68, 69, 70, 0, 0,
	96, 0, 0, 102, 98, 318, 97, 0, 100, 99,
	101, 85, 86, 87, 88, 89, 90, 91, 92, 93,
	94, 95, 84, 0, 83, 82, 81, 80, 79, 77,
	78, 73, 74, 75, 76, 71, 72, 66, 67, 68,
	69, 70, 0, 0, 0, 96, 0, 102, 98, 0,
	97, 285, 100, 99, 101, 222, 85, 86, 87, 88,
	89, 90, 91, 92, 93, 94, 95, 84, 0, 83,
	82, 81, 80, 79, 77, 78, 73, 74, 75, 76,
	71, 72, 66, 67, 68, 69, 70, 0, 0, 0,
	0, 0, 102, 98, 0, 97, 0, 100, 99, 101,
	330, 0, 0, 327, 326, 0, 331, 340, 0, 0,
	332, 341, 333, 0, 0, 0, 0, 0, 0, 334,
	335, 336, 0, 0, 10, 0, 342, 9, 21, 0,
	337, 0, 0, 52, 0, 338, 39, 57, 0, 0,
	23, 0, 46, 339, 24, 119, 45, 0, 0, 0,
	56, 41, 0, 42, 269, 0, 0, 55, 0, 40,
	43, 53, 0, 0, 0, 54, 0, 44, 0, 47,
	58, 0, 0, 13, 0, 59, 60, 61, 62, 63,
	0, 0, 14, 15, 12, 0, 0, 0, 16, 17,
	20, 0, 0, 0, 96, 22, 0, 19, 18, 0,
	0, 0, 0, 0, 325, 85, 86, 87, 88, 89,
	90, 91, 92, 93, 94, 95, 84, 0, 83, 82,
	81, 80, 79, 77, 78, 73, 74, 75, 76, 71,
	72, 66, 67, 68, 69, 70, 0, 0, 0, 0,
	0, 102, 98, 0, 97, 0, 100, 99, 101, 27,
	0, 0, 52, 0, 0, 39, 57, 0, 0, 0,
	0, 46, 38, 0, 30, 45, 0, 0, 0, 56,
	41, 0, 42, 0, 0, 0, 55, 0, 40, 43,
	53, 50, 0, 36, 54, 51, 44, 0, 47, 58,
	0, 0, 0, 0, 59, 60, 61, 62, 63, 52,
	0, 0, 39, 57, 0, 0, 0, 0, 46, 38,
	0, 119, 45, 0, 0, 0, 56, 41, 0, 42,
	0, 0, 0, 55, 0, 40, 43, 53, 50, 0,
	36, 54, 51, 44, 0, 47, 58, 0, 0, 0,
	0, 59, 60, 61, 62, 63, 0, 0, 52, 0,
	260, 39, 57, 0, 0, 0, 0, 46, 38, 0,
	119, 45, 0, 0, 0, 56, 41, 0, 42, 0,
	0, 0, 55, 0, 40, 43, 53, 50, 0, 36,
	54, 51, 44, 0, 47, 58, 0, 0, 0, 0,
	59, 60, 61, 62, 63, 27, 0, 321, 52, 0,
	0, 39, 57, 0, 0, 0, 0, 46, 38, 0,
	30, 45, 0, 0, 0, 56, 41, 0, 42, 0,
	0, 0, 55, 96, 40, 43, 53, 50, 0, 36,
	54, 51, 44, 0, 47, 58, 0, 0, 0, 0,
	59, 60, 61, 62, 63, 84, 271, 83, 82, 81,
	80, 79, 77, 78, 73, 74, 75, 76, 71, 72,
	66, 67, 68, 69, 70, 0, 0, 0, 0, 0,
	102, 98, 96, 97, 0, 100, 99, 101, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 25, 0, 82, 81, 80,
	79, 77, 78, 73, 74, 75, 76, 71, 72, 66,
	67, 68, 69, 70, 96, 0, 0, 0, 0, 102,
	98, 0, 97, 0, 100, 99, 101, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	81, 80, 79, 77, 78, 73, 74, 75, 76, 71,
	72, 66, 67, 68, 69, 70, 96, 0, 0, 0,
	0, 102, 98, 0, 97, 0, 100, 99, 101, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 80, 79, 77, 78, 73, 74, 75,
	76, 71, 72, 66, 67, 68, 69, 70, 96, 0,
	0, 0, 0, 102, 98, 0, 97, 0, 100, 99,
	101, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 96, 0, 79, 77, 78, 73,
	74, 75, 76, 71, 72, 66, 67, 68, 69, 70,
	0, 0, 0, 0, 0, 102, 98, 0, 97, 96,
	100, 99, 101, 77, 78, 73, 74, 75, 76, 71,
	72, 66, 67, 68, 69, 70, 0, 0, 0, 0,
	0, 102, 98, 0, 97, 0, 100, 99, 101, 78,
	73, 74, 75, 76, 71, 72, 66, 67, 68, 69,
	70, 10, 0, 8, 9, 21, 102, 98, 0, 97,
	52, 100, 99, 101, 57, 0, 0, 23, 0, 0,
	0, 24, 119, 0, 0, 0, 0, 56, 0, 0,
	0, 217, 0, 0, 55, 0, 0, 0, 53, 0,
	0, 0, 54, 0, 0, 96, 0, 58, 0, 0,
	13, 0, 59, 60, 61, 62, 63, 0, 0, 14,
	15, 12, 0, 0, 0, 16, 17, 20, 0, 290,
	291, 0, 22, 0, 19, 18, 73, 74, 75, 76,
	71, 72, 66, 67, 68, 69, 70, 10, 0, 8,
	9, 21, 102, 98, 0, 97, 0, 100, 99, 101,
	0, 0, 0, 23, 0, 0, 10, 24, 8, 9,
	21, 0, 0, 0, 0, 0, 0, 217, 0, 0,
	0, 0, 23, 0, 0, 0, 24, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 15, 12, 0, 0,
	0, 16, 17, 20, 0, 13, 0, 0, 22, 0,
	19, 18, 0, 0, 14, 15, 12, 0, 0, 0,
	16, 17, 20, 0, 0, 0, 0, 22, 0, 19,
	18, 10, 0, 8, 9, 21, 0, 0, 0, 0,
	0, 0, 10, 0, 8, 9, 21, 23, 0, 0,
	0, 24, 0, 0, 0, 0, 0, 0, 23, 0,
	0, 0, 24, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 217, 0, 0, 0, 0, 0, 0, 0,
	13, 0, 0, 0, 0, 0, 0, 0, 0, 14,
	15, 12, 0, 0, 0, 16, 17, 20, 0, 0,
	0, 0, 113, 0, 19, 18, 16, 17, 20, 0,
	0, 0, 0, 22, 52, 19, 18, 39, 57, 0,
	0, 0, 230, 46, 38, 0, 119, 45, 0, 0,
	0, 56, 41, 0, 42, 228, 0, 0, 55, 0,
	40, 43, 53, 50, 0, 36, 54, 51, 44, 0,
	47, 58, 0, 0, 0, 0, 59, 60, 61, 62,
	63, 362, 0, 0, 0, 52, 0, 0, 39, 57,
	0, 0, 0, 0, 46, 38, 0, 119, 45, 0,
	0, 0, 56, 41, 0, 42, 0, 0, 0, 55,
	0, 40, 43, 53, 50, 0, 36, 54, 51, 44,
	0, 47, 58, 0, 0, 0, 0, 59, 60, 61,
	62, 63, 52, 0, 0, 39, 57, 0, 0, 0,
	0, 46, 38, 0, 119, 45, 0, 0, 0, 56,
	41, 0, 42, 0, 0, 0, 55, 0, 40, 43,
	53, 50, 0, 36, 54, 51, 44, 0, 47, 58,
	0, 0, 0, 0, 59, 60, 61, 62, 63,
}
var yyPact = []int{

	-2, -1000, -1000, 1328, 949, -41, 167, 705, -1000, -1000,
	-1000, 205, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1403, 77, 377, 76, -1000, -1000, -1000, 75, -1000, -1000,
	201, 131, 1603, 1251, 684, -1000, -1000, 234, 234, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 1328, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328,
	1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, 1328, -1000,
	-1000, 234, 234, -1000, 223, 223, 223, 223, 223, 223,
	223, 223, 223, 377, 1603, 45, 44, 47, -1000, -1000,
	1328, 245, 187, -47, 40, 181, -1000, 215, 131, -1000,
	-1000, -1000, 1251, 684, -1000, -1000, 1251, -1000, 684, -1000,
	-1000, -1000, -1000, 186, -1000, 185, 705, 263, 263, 223,
	223, 223, 94, 94, 66, 66, 66, 66, 1160, 1246,
	1135, 1109, 1067, 1025, 983, 136, 705, 705, 705, 705,
	705, 705, 705, 705, 705, 705, 705, 197, 167, 42,
	91, -1000, -1000, 38, 166, 1309, -1000, 95, 215, 74,
	47, 556, 36, -1000, -1000, 1505, 1328, 1309, 1603, 131,
	131, 215, -1000, 11, -1000, -1000, -1000, 1603, 232, 1328,
	73, -1000, -1000, 1414, 1328, 223, -1000, -43, 1328, 47,
	1505, 16, 1603, -1000, 803, 34, 164, -1000, -1000, 208,
	-1000, 90, 705, -1000, 705, -1000, 183, -1000, 131, -1000,
	40, 33, -1000, -1000, 899, -1000, 131, 163, -1000, 150,
	934, 1328, 511, -1000, 1223, 88, 95, 29, -1000, 25,
	-1000, -1000, 1505, 95, 33, 215, 208, -1000, -1000, -1000,
	-48, -1000, -1000, -50, 158, -1000, 33, 132, -1000, -44,
	232, -1000, -1000, 1328, 21, -1000, 0, -1000, 110, -1000,
	234, 1328, -1000, -1000, -1000, -1000, 208, -1000, -1000, -1000,
	131, 1328, -1000, -1000, 705, -1000, -1000, -45, 1309, -1000,
	-1000, -1000, 467, 850, -1000, 705, -1000, -1000, -1000, -1000,
	-1000, -1000, 656, -1000, -1000, -1000, 67, 65, -1000, -51,
	-1000, -52, -53, -1000, 61, 234, 60, 1328, 59, 48,
	1328, 130, 122, 1328, 1328, -1000, 1556, -1000, -1000, 213,
	1328, -54, 1328, -55, -1000, 1328, 1328, 423, -1000, -1000,
	20, 19, -1000, 41, -57, -1000, 18, -1000, 17, 14,
	-1000, -59, -60, 1328, 1328, -1000, -1000, -1000, -1000, -1000,
	9, -61, 189, -1000, -1000, -62, 1328, -1000, -1000, 4,
	-1000, -1000, -1000,
}
var yyPgo = []int{

	0, 9, 341, 23, 340, 24, 37, 338, 337, 33,
	39, 335, 333, 28, 332, 331, 12, 10, 330, 323,
	1, 36, 27, 4, 321, 318, 26, 30, 35, 315,
	32, 8, 314, 38, 308, 302, 301, 11, 299, 297,
	6, 0, 5, 292, 29, 98, 66, 40, 3, 289,
	52, 45, 290, 43, 288, 25, 287, 2, 284, 34,
	13, 265, 283, 281, 280, 279, 277,
}
var yyR1 = []int{

	0, 62, 62, 10, 10, 10, 22, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 20, 20, 20, 20, 20, 20, 20, 20, 20,
	20, 42, 42, 42, 63, 40, 35, 35, 35, 41,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 39,
	39, 39, 39, 39, 39, 39, 1, 1, 1, 2,
	2, 2, 16, 16, 16, 16, 16, 3, 3, 3,
	3, 28, 28, 43, 43, 43, 43, 43, 43, 44,
	44, 44, 44, 44, 44, 44, 45, 45, 45, 45,
	45, 45, 45, 45, 45, 46, 46, 47, 47, 61,
	57, 57, 57, 57, 57, 60, 59, 6, 12, 11,
	11, 11, 64, 4, 48, 48, 58, 58, 17, 17,
	13, 61, 61, 37, 20, 20, 61, 61, 5, 24,
	31, 31, 33, 33, 33, 34, 34, 32, 32, 37,
	66, 66, 65, 65, 38, 38, 49, 49, 23, 23,
	21, 21, 26, 26, 27, 27, 7, 7, 36, 36,
	8, 8, 9, 9, 29, 29, 30, 30, 54, 54,
	55, 55, 50, 50, 51, 51, 52, 52, 53, 53,
	18, 18, 19, 19, 14, 14, 25, 25, 15, 15,
	56, 56,
}
var yyR2 = []int{

	0, 3, 3, 0, 2, 5, 1, 1, 1, 1,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	4, 6, 4, 4, 3, 7, 4, 4, 2, 2,
	6, 0, 2, 2, 0, 4, 3, 2, 2, 2,
	1, 5, 5, 1, 2, 3, 2, 2, 7, 9,
	3, 5, 7, 3, 5, 5, 0, 3, 1, 4,
	4, 3, 1, 3, 3, 4, 4, 1, 2, 2,
	1, 1, 3, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 3, 3, 2, 2, 1, 2, 3, 3, 1,
	1, 5, 0, 5, 1, 1, 1, 1, 1, 3,
	3, 2, 5, 2, 3, 3, 2, 6, 2, 2,
	1, 1, 2, 4, 5, 0, 3, 1, 3, 3,
	0, 1, 0, 1, 1, 2, 0, 1, 0, 1,
	0, 1, 1, 3, 0, 1, 0, 2, 0, 2,
	1, 3, 0, 1, 1, 3, 0, 1, 1, 2,
	0, 1, 1, 2, 0, 1, 1, 2, 0, 1,
	1, 3, 0, 1, 1, 2, 0, 1, 1, 3,
	1, 2,
}
var yyChk = []int{

	-1000, -62, 105, 104, -10, -22, -26, -20, 30, 31,
	28, -56, 88, 77, 86, 87, 92, 93, 102, 101,
	94, 32, 99, 44, 48, 106, -11, 6, -12, -4,
	21, -57, -50, -61, -45, -46, 40, -58, 19, 12,
	35, 27, 29, 36, 43, 22, 18, 45, -43, -44,
	38, 42, 9, 37, 41, 33, 26, 13, 46, 51,
	52, 53, 54, 55, 106, 59, 86, 87, 88, 89,
	90, 84, 85, 80, 81, 82, 83, 78, 79, 77,
	76, 75, 74, 73, 71, 60, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 49, 99, 97, 102,
	101, 103, 96, 48, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, 99, 99, -59, -22, -60, -57, 21,
	99, 99, 48, -30, -16, -29, -48, 88, 99, -28,
	30, 40, -61, -45, -46, -51, -50, -53, -52, -47,
	-46, -45, -48, -49, -48, -49, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -22, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -27, -26, -27,
	-22, -48, -48, -59, -59, 100, 100, -1, 88, -2,
	99, -20, 30, 58, 108, 99, 97, 60, -7, 59,
	-55, -54, -44, -16, -51, -53, -47, 58, 58, 72,
	50, 100, 98, 100, 59, -20, -33, 58, 97, -55,
	99, -1, 59, 100, -10, -9, -8, -3, 30, -60,
	17, -21, -20, -31, -20, -33, -64, -6, -57, -28,
	-16, -16, -44, 100, -14, -13, -60, -15, -5, 30,
	-20, 99, -20, 107, -34, -21, -1, -9, 100, -59,
	107, 100, 59, -1, -16, 88, 99, 98, -40, 58,
	-30, 107, -13, -19, -18, -17, -16, -49, -48, -65,
	59, -25, -24, 60, -27, 100, -32, -31, -38, -37,
	96, 97, 98, 100, 100, -3, -55, -63, 108, 108,
	59, 72, 107, -5, -20, 100, 107, 59, -66, -37,
	60, -48, -20, -42, -17, -20, 107, -31, 98, -6,
	-41, 107, -36, -39, -35, 108, 8, 7, -40, -22,
	4, 10, 14, 16, 23, 24, 25, 34, 39, 47,
	11, 15, 30, 99, 99, 108, -42, 108, 108, -41,
	99, -48, 99, -23, -22, 99, 99, -20, 72, 72,
	-22, -22, 5, 47, -23, 108, -22, 108, -22, -22,
	72, 100, 100, 99, 108, 100, 100, 100, 108, 108,
	-22, -23, -41, -41, -41, 100, 108, 57, 108, -23,
	-41, 100, -41,
}
var yyDef = []int{

	0, -2, 3, 0, 0, 0, 6, 182, 7, 8,
	9, 10, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 220, 1, 4, 0, 139, 140,
	105, 196, 130, 204, 208, 202, 129, 176, 176, 116,
	117, 118, 119, 120, 121, 122, 123, 124, 125, 126,
	146, 147, 103, 104, 106, 107, 108, 109, 110, 111,
	112, 113, 114, 115, 2, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 184, 184, 0, 58,
	59, 0, 0, 221, 41, 42, 43, 44, 45, 46,
	47, 48, 49, 0, 0, 0, 0, 86, 135, 105,
	0, 0, 0, 0, -2, 197, 92, 200, 0, 194,
	144, 145, 204, 208, 203, 133, 205, 134, 209, 206,
	127, 128, -2, 0, -2, 0, 183, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 0, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 0, 185, 0,
	0, 154, 155, 0, 0, 0, 54, 136, 200, 88,
	86, 0, 0, 3, 138, 192, 180, 0, 142, 0,
	0, 201, 198, 0, 131, 132, 207, 0, 0, 0,
	0, 56, 57, 50, 0, 52, 53, 165, 180, 86,
	192, 0, 0, 5, 0, 0, 193, 190, 97, 86,
	100, 0, 181, 102, 160, 161, 0, 187, 196, 195,
	101, 93, 199, 94, 0, 214, -2, 172, 218, 216,
	29, 184, 0, 162, 0, 0, 87, 0, 91, 0,
	141, 95, 0, 98, 99, 200, 86, 96, 143, 64,
	0, 152, 215, 0, 213, 210, 148, 0, -2, 0,
	173, 158, 217, 0, 0, 51, 0, 167, 170, 174,
	0, 0, 90, 89, 60, 191, 86, 61, 137, 150,
	176, 0, 157, 219, 159, 55, 163, 166, 0, 175,
	171, 153, 0, 188, 211, 149, 164, 168, 169, 62,
	63, 65, 0, 69, 189, 70, 0, 0, 73, 0,
	61, 0, 0, 188, 0, 0, 0, 178, 0, 0,
	0, 0, 7, 0, 0, 74, 188, 76, 77, 0,
	178, 0, 0, 0, 179, 0, 0, 0, 67, 68,
	0, 0, 75, 0, 0, 80, 0, 83, 0, 0,
	66, 0, 0, 0, 178, 188, 188, 188, 71, 72,
	0, 0, 81, 84, 85, 0, 178, 188, 78, 0,
	82, 188, 79,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 92, 3, 3, 3, 90, 77, 3,
	99, 100, 88, 86, 59, 87, 96, 89, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 72, 108,
	80, 60, 81, 71, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 97, 3, 98, 76, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 58, 75, 107, 93,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 61, 62, 63, 64,
	65, 66, 67, 68, 69, 70, 73, 74, 78, 79,
	82, 83, 84, 85, 91, 94, 95, 101, 102, 103,
	104, 105, 106,
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
		//line cc.y:195
		{
			yylex.(*lexer).prog = &Prog{Decls: yyS[yypt-1].decls}
			return 0
		}
	case 2:
		//line cc.y:200
		{
			yylex.(*lexer).expr = yyS[yypt-1].expr
			return 0
		}
	case 3:
		//line cc.y:206
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 4:
		//line cc.y:211
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 5:
		//line cc.y:216
		{
		}
	case 6:
		//line cc.y:221
		{
			yyVAL.span = yyS[yypt-0].span
			if len(yyS[yypt-0].exprs) == 1 {
				yyVAL.expr = yyS[yypt-0].exprs[0]
				break
			}
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Comma, List: yyS[yypt-0].exprs}
		}
	case 7:
		//line cc.y:232
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Name, Text: yyS[yypt-0].str, XDecl: yyS[yypt-0].decl}
		}
	case 8:
		//line cc.y:237
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyS[yypt-0].str}
		}
	case 9:
		//line cc.y:242
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyS[yypt-0].str}
		}
	case 10:
		//line cc.y:247
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: String, Texts: yyS[yypt-0].strs}
		}
	case 11:
		//line cc.y:252
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Add, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 12:
		//line cc.y:257
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Sub, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 13:
		//line cc.y:262
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mul, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 14:
		//line cc.y:267
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Div, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 15:
		//line cc.y:272
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mod, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 16:
		//line cc.y:277
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 17:
		//line cc.y:282
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Rsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 18:
		//line cc.y:287
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 19:
		//line cc.y:292
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Gt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 20:
		//line cc.y:297
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 21:
		//line cc.y:302
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: GtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 22:
		//line cc.y:307
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: EqEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 23:
		//line cc.y:312
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: NotEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 24:
		//line cc.y:317
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: And, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 25:
		//line cc.y:322
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Xor, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 26:
		//line cc.y:327
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Or, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 27:
		//line cc.y:332
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndAnd, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 28:
		//line cc.y:337
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrOr, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 29:
		//line cc.y:342
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cond, List: []*Expr{yyS[yypt-4].expr, yyS[yypt-2].expr, yyS[yypt-0].expr}}
		}
	case 30:
		//line cc.y:347
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Eq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 31:
		//line cc.y:352
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AddEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 32:
		//line cc.y:357
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SubEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 33:
		//line cc.y:362
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: MulEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 34:
		//line cc.y:367
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: DivEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 35:
		//line cc.y:372
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ModEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 36:
		//line cc.y:377
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 37:
		//line cc.y:382
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: RshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 38:
		//line cc.y:387
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 39:
		//line cc.y:392
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: XorEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 40:
		//line cc.y:397
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 41:
		//line cc.y:402
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Indir, Left: yyS[yypt-0].expr}
		}
	case 42:
		//line cc.y:407
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Addr, Left: yyS[yypt-0].expr}
		}
	case 43:
		//line cc.y:412
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Plus, Left: yyS[yypt-0].expr}
		}
	case 44:
		//line cc.y:417
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Minus, Left: yyS[yypt-0].expr}
		}
	case 45:
		//line cc.y:422
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Not, Left: yyS[yypt-0].expr}
		}
	case 46:
		//line cc.y:427
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Twid, Left: yyS[yypt-0].expr}
		}
	case 47:
		//line cc.y:432
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreInc, Left: yyS[yypt-0].expr}
		}
	case 48:
		//line cc.y:437
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreDec, Left: yyS[yypt-0].expr}
		}
	case 49:
		//line cc.y:442
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofExpr, Left: yyS[yypt-0].expr}
		}
	case 50:
		//line cc.y:447
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofType, Type: yyS[yypt-1].typ}
		}
	case 51:
		//line cc.y:452
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Offsetof, Type: yyS[yypt-3].typ, Left: yyS[yypt-1].expr}
		}
	case 52:
		//line cc.y:457
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cast, Type: yyS[yypt-2].typ, Left: yyS[yypt-0].expr}
		}
	case 53:
		//line cc.y:462
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CastInit, Type: yyS[yypt-2].typ, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits}}
		}
	case 54:
		//line cc.y:467
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Paren, Left: yyS[yypt-1].expr}
		}
	case 55:
		//line cc.y:472
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CUDACall, Left: yyS[yypt-6].expr, LaunchParams: yyS[yypt-4].exprs, List: yyS[yypt-1].exprs}
		}
	case 56:
		//line cc.y:477
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Call, Left: yyS[yypt-3].expr, List: yyS[yypt-1].exprs}
		}
	case 57:
		//line cc.y:482
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Index, Left: yyS[yypt-3].expr, Right: yyS[yypt-1].expr}
		}
	case 58:
		//line cc.y:487
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostInc, Left: yyS[yypt-1].expr}
		}
	case 59:
		//line cc.y:492
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostDec, Left: yyS[yypt-1].expr}
		}
	case 60:
		//line cc.y:497
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: VaArg, Left: yyS[yypt-3].expr, Type: yyS[yypt-1].typ}
		}
	case 61:
		//line cc.y:503
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 62:
		//line cc.y:508
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = yyS[yypt-1].stmts
			for _, d := range yyS[yypt-0].decls {
				yyVAL.stmts = append(yyVAL.stmts, &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtDecl, Decl: d})
			}
		}
	case 63:
		//line cc.y:516
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = append(yyS[yypt-1].stmts, yyS[yypt-0].stmt)
		}
	case 64:
		//line cc.y:523
		{
			yylex.(*lexer).pushScope()
		}
	case 65:
		//line cc.y:527
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Block, Block: yyS[yypt-1].stmts}
		}
	case 66:
		//line cc.y:535
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Case, Expr: yyS[yypt-1].expr}
		}
	case 67:
		//line cc.y:540
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Default}
		}
	case 68:
		//line cc.y:545
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LabelName, Name: yyS[yypt-1].str}
		}
	case 69:
		//line cc.y:552
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = yyS[yypt-0].stmt
			yyVAL.stmt.Labels = yyS[yypt-1].labels
		}
	case 70:
		//line cc.y:560
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 71:
		//line cc.y:565
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 72:
		//line cc.y:570
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 73:
		//line cc.y:575
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = yyS[yypt-0].stmt
		}
	case 74:
		//line cc.y:580
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtExpr, Expr: yyS[yypt-1].expr}
		}
	case 75:
		//line cc.y:585
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ARGBEGIN, Block: yyS[yypt-1].stmts}
		}
	case 76:
		//line cc.y:590
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Break}
		}
	case 77:
		//line cc.y:595
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Continue}
		}
	case 78:
		//line cc.y:600
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Do, Body: yyS[yypt-5].stmt, Expr: yyS[yypt-2].expr}
		}
	case 79:
		//line cc.y:605
		{
			yyVAL.span = span(yyS[yypt-8].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span},
				Op:   For,
				Pre:  yyS[yypt-6].expr,
				Expr: yyS[yypt-4].expr,
				Post: yyS[yypt-2].expr,
				Body: yyS[yypt-0].stmt,
			}
		}
	case 80:
		//line cc.y:616
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Goto, Text: yyS[yypt-1].str}
		}
	case 81:
		//line cc.y:621
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 82:
		//line cc.y:626
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyS[yypt-4].expr, Body: yyS[yypt-2].stmt, Else: yyS[yypt-0].stmt}
		}
	case 83:
		//line cc.y:631
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Return, Expr: yyS[yypt-1].expr}
		}
	case 84:
		//line cc.y:636
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Switch, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 85:
		//line cc.y:641
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: While, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 86:
		//line cc.y:648
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 87:
		//line cc.y:653
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			_, q, _ := splitTypeWords(yyS[yypt-1].strs)
			abdecor := yyS[yypt-0].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 88:
		//line cc.y:662
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.abdecor = yyS[yypt-0].abdecor
		}
	case 89:
		//line cc.y:669
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
						decl.Type = &Type{Kind: Ptr, Base: t}
					}
				}
			}
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls})
			}
		}
	case 90:
		//line cc.y:693
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			abdecor := yyS[yypt-3].abdecor
			span := yyVAL.span
			expr := yyS[yypt-1].expr
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}

		}
	case 91:
		//line cc.y:704
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.abdecor = yyS[yypt-1].abdecor
		}
	case 92:
		//line cc.y:712
		{
			yyVAL.span = yyS[yypt-0].span
			name := yyS[yypt-0].str
			yyVAL.decor = func(t *Type) (*Type, string) { return t, name }
		}
	case 93:
		//line cc.y:718
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			_, q, _ := splitTypeWords(yyS[yypt-1].strs)
			decor := yyS[yypt-0].decor
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 94:
		//line cc.y:728
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decor = yyS[yypt-1].decor
		}
	case 95:
		//line cc.y:733
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			decor := yyS[yypt-3].decor
			decls := yyS[yypt-1].decls
			span := yyVAL.span
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Func, Base: t, Decls: decls})
			}
		}
	case 96:
		//line cc.y:743
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			decor := yyS[yypt-3].decor
			span := yyVAL.span
			expr := yyS[yypt-1].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				return decor(&Type{SyntaxInfo: SyntaxInfo{Span: span}, Kind: Array, Base: t, Width: expr})
			}
		}
	case 97:
		//line cc.y:756
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyS[yypt-0].str}
		}
	case 98:
		//line cc.y:761
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyS[yypt-0].abdecor(yyS[yypt-1].typ)}
		}
	case 99:
		//line cc.y:766
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			typ, name := yyS[yypt-0].decor(yyS[yypt-1].typ)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ}
		}
	case 100:
		//line cc.y:772
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "..."}
		}
	case 101:
		//line cc.y:780
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idec = idecor{yyS[yypt-0].decor, nil}
		}
	case 102:
		//line cc.y:785
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idec = idecor{yyS[yypt-2].decor, yyS[yypt-0].init}
		}
	case 103:
		//line cc.y:793
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 104:
		//line cc.y:798
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 105:
		//line cc.y:803
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 106:
		//line cc.y:808
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 107:
		//line cc.y:813
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 108:
		//line cc.y:818
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 109:
		//line cc.y:826
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 110:
		//line cc.y:831
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 111:
		//line cc.y:836
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 112:
		//line cc.y:841
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 113:
		//line cc.y:846
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 114:
		//line cc.y:851
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 115:
		//line cc.y:856
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 116:
		//line cc.y:864
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 117:
		//line cc.y:869
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 118:
		//line cc.y:874
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 119:
		//line cc.y:879
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 120:
		//line cc.y:884
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 121:
		//line cc.y:889
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 122:
		//line cc.y:894
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 123:
		//line cc.y:899
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 124:
		//line cc.y:904
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 125:
		//line cc.y:911
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 126:
		//line cc.y:916
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 127:
		//line cc.y:923
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 128:
		//line cc.y:928
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 129:
		//line cc.y:936
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.typ = yyS[yypt-0].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{Kind: TypedefType, Name: yyS[yypt-0].str}
			}
		}
	case 130:
		//line cc.y:952
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(append(yyS[yypt-0].strs, "int"))
		}
	case 131:
		//line cc.y:957
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyS[yypt-2].strs, yyS[yypt-0].strs...))
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 132:
		//line cc.y:963
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyS[yypt-2].strs = append(yyS[yypt-2].strs, yyS[yypt-1].str)
			yyS[yypt-2].strs = append(yyS[yypt-2].strs, yyS[yypt-0].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyS[yypt-2].strs)
		}
	case 133:
		//line cc.y:970
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyS[yypt-0].strs)
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 134:
		//line cc.y:976
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var ts []string
			ts = append(ts, yyS[yypt-1].str)
			ts = append(ts, yyS[yypt-0].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 135:
		//line cc.y:987
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
	case 136:
		//line cc.y:1000
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yyS[yypt-0].abdecor(yyS[yypt-1].typ)
		}
	case 137:
		//line cc.y:1008
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyS[yypt-1].idecs {
				typ, name := idec.d(yyS[yypt-2].tc.t)
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Storage: yyS[yypt-2].tc.c, Init: idec.i}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
			if yyS[yypt-1].idecs == nil {
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "", Type: yyS[yypt-2].tc.t, Storage: yyS[yypt-2].tc.c}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 138:
		//line cc.y:1028
		{
			lx := yylex.(*lexer)
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			// TODO: use $1.q
			yyVAL.decls = nil
			for _, idec := range yyS[yypt-1].idecs {
				typ, name := idec.d(yyS[yypt-2].tc.t)
				d := lx.lookupDecl(name)
				if d == nil {
					d = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ, Storage: yyS[yypt-2].tc.c, Init: idec.i}
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
				d := &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "", Type: yyS[yypt-2].tc.t, Storage: yyS[yypt-2].tc.c}
				lx.pushDecl(d)
				yyVAL.decls = append(yyVAL.decls, d)
			}
		}
	case 139:
		//line cc.y:1056
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 140:
		//line cc.y:1061
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 141:
		//line cc.y:1066
		{
			yyVAL.decls = yyS[yypt-1].decls
		}
	case 142:
		//line cc.y:1072
		{
			lx := yylex.(*lexer)
			typ, name := yyS[yypt-1].decor(yyS[yypt-2].tc.t)
			if typ.Kind != Func {
				yylex.(*lexer).Errorf("invalid function definition")
				return 0
			}
			d := lx.lookupDecl(name)
			if d == nil {
				d = &Decl{Name: name, Type: typ, Storage: yyS[yypt-2].tc.c}
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
	case 143:
		//line cc.y:1093
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
	case 144:
		//line cc.y:1106
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 145:
		//line cc.y:1111
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 146:
		//line cc.y:1119
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Struct
		}
	case 147:
		//line cc.y:1124
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Union
		}
	case 148:
		//line cc.y:1131
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decor = yyS[yypt-0].decor
		}
	case 149:
		//line cc.y:1136
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			name := yyS[yypt-2].str
			expr := yyS[yypt-0].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				t.Width = expr
				return t, name
			}
		}
	case 150:
		//line cc.y:1148
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = nil
			for _, decor := range yyS[yypt-1].decors {
				typ, name := decor(yyS[yypt-2].typ)
				yyVAL.decls = append(yyVAL.decls, &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ})
			}
			if yyS[yypt-1].decors == nil {
				yyVAL.decls = append(yyVAL.decls, &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyS[yypt-2].typ})
			}
		}
	case 151:
		//line cc.y:1162
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyS[yypt-1].tk, Tag: yyS[yypt-0].str})
		}
	case 152:
		//line cc.y:1167
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyS[yypt-4].tk, Tag: yyS[yypt-3].str, Decls: yyS[yypt-1].decls})
		}
	case 153:
		//line cc.y:1174
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Dot: yyS[yypt-0].str}
		}
	case 154:
		//line cc.y:1181
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Arrow, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].str}
		}
	case 155:
		//line cc.y:1186
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Dot, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].str}
		}
	case 156:
		//line cc.y:1194
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-0].str})
		}
	case 157:
		//line cc.y:1199
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-4].str, Decls: yyS[yypt-2].decls})
		}
	case 158:
		//line cc.y:1206
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var x *Init
			if yyS[yypt-0].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr}
			}
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyS[yypt-1].str, Init: x}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 159:
		//line cc.y:1218
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 160:
		//line cc.y:1226
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr}
		}
	case 161:
		//line cc.y:1231
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits}
		}
	case 162:
		//line cc.y:1238
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.inits = []*Init{}
		}
	case 163:
		//line cc.y:1243
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 164:
		//line cc.y:1248
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-3].inits, yyS[yypt-2].init)
		}
	case 165:
		//line cc.y:1254
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 166:
		//line cc.y:1259
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 167:
		//line cc.y:1266
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = yyS[yypt-0].init
		}
	case 168:
		//line cc.y:1271
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.init = yyS[yypt-0].init
			yyVAL.init.Prefix = yyS[yypt-2].prefixes
		}
	case 169:
		//line cc.y:1279
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Index: yyS[yypt-1].expr}
		}
	case 170:
		//line cc.y:1285
		{
			yyVAL.span = Span{}
		}
	case 171:
		//line cc.y:1289
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 172:
		//line cc.y:1294
		{
			yyVAL.span = Span{}
		}
	case 173:
		//line cc.y:1298
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 174:
		//line cc.y:1307
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.prefixes = []*Prefix{yyS[yypt-0].prefix}
		}
	case 175:
		//line cc.y:1312
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefixes = append(yyS[yypt-1].prefixes, yyS[yypt-0].prefix)
		}
	case 176:
		//line cc.y:1318
		{
			yyVAL.span = Span{}
			yyVAL.str = ""
		}
	case 177:
		//line cc.y:1323
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 178:
		//line cc.y:1329
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 179:
		//line cc.y:1334
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 180:
		//line cc.y:1340
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 181:
		//line cc.y:1345
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 182:
		//line cc.y:1352
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = []*Expr{yyS[yypt-0].expr}
		}
	case 183:
		//line cc.y:1357
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 184:
		//line cc.y:1363
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 185:
		//line cc.y:1368
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = yyS[yypt-0].exprs
		}
	case 186:
		//line cc.y:1374
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 187:
		//line cc.y:1379
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 188:
		//line cc.y:1385
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 189:
		//line cc.y:1390
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.labels = append(yyS[yypt-1].labels, yyS[yypt-0].label)
		}
	case 190:
		//line cc.y:1397
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 191:
		//line cc.y:1402
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 192:
		//line cc.y:1408
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 193:
		//line cc.y:1413
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 194:
		//line cc.y:1420
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = []idecor{yyS[yypt-0].idec}
		}
	case 195:
		//line cc.y:1425
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idecs = append(yyS[yypt-2].idecs, yyS[yypt-0].idec)
		}
	case 196:
		//line cc.y:1431
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 197:
		//line cc.y:1436
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = yyS[yypt-0].idecs
		}
	case 198:
		//line cc.y:1443
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 199:
		//line cc.y:1448
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 200:
		//line cc.y:1454
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 201:
		//line cc.y:1459
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 202:
		//line cc.y:1466
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 203:
		//line cc.y:1471
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 204:
		//line cc.y:1477
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 205:
		//line cc.y:1482
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 206:
		//line cc.y:1489
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 207:
		//line cc.y:1494
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 208:
		//line cc.y:1500
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 209:
		//line cc.y:1505
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 210:
		//line cc.y:1512
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyS[yypt-0].decor)
		}
	case 211:
		//line cc.y:1518
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decors = append(yyS[yypt-2].decors, yyS[yypt-0].decor)
		}
	case 212:
		//line cc.y:1524
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 213:
		//line cc.y:1529
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = yyS[yypt-0].decors
		}
	case 214:
		//line cc.y:1536
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 215:
		//line cc.y:1541
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 216:
		//line cc.y:1547
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 217:
		//line cc.y:1552
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 218:
		//line cc.y:1559
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 219:
		//line cc.y:1564
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 220:
		//line cc.y:1571
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 221:
		//line cc.y:1576
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	}
	goto yystack /* stack new state and value */
}
