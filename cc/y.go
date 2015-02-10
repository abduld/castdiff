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
const tokShift = 57393
const tokElse = 57394
const tokAddEq = 57395
const tokSubEq = 57396
const tokMulEq = 57397
const tokDivEq = 57398
const tokModEq = 57399
const tokLshEq = 57400
const tokRshEq = 57401
const tokAndEq = 57402
const tokXorEq = 57403
const tokOrEq = 57404
const tokOrOr = 57405
const tokAndAnd = 57406
const tokEqEq = 57407
const tokNotEq = 57408
const tokLtEq = 57409
const tokGtEq = 57410
const tokLsh = 57411
const tokRsh = 57412
const tokCast = 57413
const tokSizeof = 57414
const tokUnary = 57415
const tokDec = 57416
const tokInc = 57417
const tokArrow = 57418
const startExpr = 57419
const startProg = 57420
const tokEOF = 57421

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
	-1, 119,
	54, 101,
	103, 101,
	-2, 181,
	-1, 137,
	53, 172,
	-2, 146,
	-1, 139,
	53, 172,
	-2, 151,
	-1, 241,
	103, 207,
	-2, 171,
	-1, 273,
	67, 172,
	-2, 92,
}

const yyNprod = 217
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 1481

var yyAct = []int{

	315, 7, 113, 121, 348, 308, 263, 31, 228, 216,
	270, 284, 198, 112, 99, 100, 101, 102, 103, 104,
	105, 106, 107, 222, 243, 195, 6, 349, 240, 49,
	172, 5, 118, 220, 110, 124, 226, 314, 230, 4,
	134, 137, 139, 132, 119, 130, 383, 32, 381, 374,
	111, 373, 369, 362, 360, 343, 342, 340, 302, 294,
	293, 34, 141, 142, 143, 144, 145, 146, 147, 148,
	149, 150, 151, 152, 153, 154, 155, 156, 157, 158,
	159, 131, 161, 162, 163, 164, 165, 166, 167, 168,
	169, 170, 171, 189, 128, 311, 136, 297, 248, 59,
	176, 177, 3, 2, 35, 386, 301, 160, 91, 192,
	10, 380, 8, 9, 21, 372, 371, 186, 173, 173,
	370, 175, 182, 174, 125, 367, 23, 366, 125, 191,
	24, 190, 238, 300, 126, 212, 111, 129, 126, 135,
	289, 288, 213, 178, 179, 253, 191, 191, 190, 190,
	97, 93, 197, 92, 13, 95, 94, 96, 256, 218,
	208, 206, 181, 14, 15, 12, 180, 368, 351, 16,
	17, 20, 200, 199, 201, 131, 22, 122, 19, 18,
	183, 260, 210, 350, 347, 345, 339, 338, 123, 246,
	136, 185, 261, 227, 229, 136, 233, 215, 116, 115,
	109, 287, 262, 305, 224, 207, 245, 235, 236, 214,
	210, 247, 213, 197, 354, 227, 241, 353, 296, 211,
	204, 278, 31, 295, 251, 275, 237, 257, 219, 224,
	234, 232, 209, 135, 258, 60, 129, 259, 135, 285,
	286, 194, 264, 203, 202, 273, 235, 211, 188, 252,
	250, 229, 254, 241, 271, 382, 205, 117, 282, 10,
	98, 8, 9, 21, 358, 244, 265, 187, 267, 57,
	33, 224, 125, 173, 303, 23, 272, 279, 274, 24,
	299, 290, 126, 231, 212, 292, 291, 1, 307, 306,
	197, 37, 11, 196, 133, 304, 48, 318, 310, 273,
	298, 251, 58, 127, 236, 229, 309, 283, 271, 317,
	319, 233, 312, 91, 138, 140, 249, 281, 16, 17,
	20, 120, 276, 277, 323, 22, 268, 19, 18, 344,
	269, 341, 242, 239, 346, 28, 26, 352, 221, 193,
	29, 184, 0, 0, 233, 324, 0, 63, 64, 65,
	359, 0, 0, 52, 0, 97, 93, 57, 92, 0,
	95, 94, 96, 0, 0, 114, 355, 356, 0, 0,
	56, 377, 378, 379, 376, 361, 0, 55, 363, 364,
	0, 53, 0, 385, 0, 54, 384, 387, 0, 0,
	58, 91, 0, 0, 0, 0, 375, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 79, 365,
	78, 77, 76, 75, 74, 72, 73, 68, 69, 70,
	71, 66, 67, 61, 62, 63, 64, 65, 0, 0,
	0, 0, 0, 97, 93, 91, 92, 0, 95, 94,
	96, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 79, 0, 78, 77, 76, 75, 74, 72,
	73, 68, 69, 70, 71, 66, 67, 61, 62, 63,
	64, 65, 0, 0, 0, 0, 0, 97, 93, 313,
	92, 91, 95, 94, 96, 0, 0, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 79, 0,
	78, 77, 76, 75, 74, 72, 73, 68, 69, 70,
	71, 66, 67, 61, 62, 63, 64, 65, 0, 0,
	0, 0, 0, 97, 93, 0, 92, 280, 95, 94,
	96, 91, 0, 0, 0, 0, 217, 80, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 79, 0,
	78, 77, 76, 75, 74, 72, 73, 68, 69, 70,
	71, 66, 67, 61, 62, 63, 64, 65, 0, 0,
	0, 0, 0, 97, 93, 0, 92, 0, 95, 94,
	96, 325, 0, 0, 322, 321, 0, 326, 335, 0,
	0, 327, 336, 328, 0, 0, 0, 0, 0, 0,
	329, 330, 331, 0, 0, 10, 0, 337, 9, 21,
	0, 332, 0, 0, 0, 0, 333, 0, 0, 52,
	0, 23, 39, 57, 334, 24, 0, 0, 46, 38,
	264, 114, 45, 0, 0, 0, 56, 41, 0, 42,
	0, 0, 0, 55, 0, 40, 43, 53, 50, 13,
	36, 54, 51, 44, 0, 47, 58, 0, 14, 15,
	12, 0, 0, 0, 16, 17, 20, 0, 0, 0,
	0, 22, 0, 19, 18, 91, 0, 0, 0, 0,
	320, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 79, 0, 78, 77, 76, 75, 74, 72,
	73, 68, 69, 70, 71, 66, 67, 61, 62, 63,
	64, 65, 316, 0, 0, 0, 0, 97, 93, 0,
	92, 0, 95, 94, 96, 52, 0, 0, 39, 57,
	0, 0, 0, 0, 46, 38, 91, 114, 45, 0,
	0, 0, 56, 41, 10, 42, 8, 9, 21, 55,
	0, 40, 43, 53, 50, 0, 36, 54, 51, 44,
	23, 47, 58, 0, 24, 0, 66, 67, 61, 62,
	63, 64, 65, 0, 0, 0, 0, 0, 97, 93,
	0, 92, 0, 95, 94, 96, 0, 0, 13, 0,
	0, 0, 0, 0, 0, 0, 0, 14, 15, 12,
	0, 0, 0, 16, 17, 20, 0, 0, 27, 0,
	22, 52, 19, 18, 39, 57, 0, 0, 0, 0,
	46, 38, 0, 30, 45, 0, 0, 0, 56, 41,
	0, 42, 91, 0, 0, 55, 0, 40, 43, 53,
	50, 0, 36, 54, 51, 44, 0, 47, 58, 79,
	0, 78, 77, 76, 75, 74, 72, 73, 68, 69,
	70, 71, 66, 67, 61, 62, 63, 64, 65, 0,
	0, 0, 0, 0, 97, 93, 0, 92, 0, 95,
	94, 96, 52, 0, 0, 39, 57, 0, 0, 0,
	0, 46, 38, 0, 114, 45, 0, 0, 0, 56,
	41, 0, 42, 0, 255, 0, 55, 0, 40, 43,
	53, 50, 0, 36, 54, 51, 44, 27, 47, 58,
	52, 0, 0, 39, 57, 0, 0, 0, 0, 46,
	38, 0, 30, 45, 0, 0, 0, 56, 41, 0,
	42, 0, 0, 0, 55, 0, 40, 43, 53, 50,
	0, 36, 54, 51, 44, 0, 47, 58, 91, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 266, 0, 0, 77, 76,
	75, 74, 72, 73, 68, 69, 70, 71, 66, 67,
	61, 62, 63, 64, 65, 91, 0, 0, 0, 0,
	97, 93, 0, 92, 0, 95, 94, 96, 0, 0,
	0, 0, 25, 0, 0, 0, 76, 75, 74, 72,
	73, 68, 69, 70, 71, 66, 67, 61, 62, 63,
	64, 65, 91, 0, 0, 0, 0, 97, 93, 0,
	92, 0, 95, 94, 96, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 75, 74, 72, 73, 68, 69,
	70, 71, 66, 67, 61, 62, 63, 64, 65, 91,
	0, 0, 0, 0, 97, 93, 0, 92, 0, 95,
	94, 96, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 74, 72, 73, 68, 69, 70, 71, 66,
	67, 61, 62, 63, 64, 65, 91, 0, 0, 0,
	0, 97, 93, 0, 92, 0, 95, 94, 96, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	72, 73, 68, 69, 70, 71, 66, 67, 61, 62,
	63, 64, 65, 91, 0, 0, 0, 0, 97, 93,
	0, 92, 0, 95, 94, 96, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 73, 68,
	69, 70, 71, 66, 67, 61, 62, 63, 64, 65,
	10, 0, 8, 9, 21, 97, 93, 0, 92, 0,
	95, 94, 96, 0, 52, 0, 23, 39, 57, 0,
	24, 0, 0, 46, 0, 212, 114, 45, 0, 0,
	0, 56, 41, 0, 42, 0, 0, 0, 55, 0,
	40, 43, 53, 0, 13, 91, 54, 0, 44, 0,
	47, 58, 0, 14, 15, 12, 0, 0, 0, 16,
	17, 20, 0, 285, 286, 0, 22, 0, 19, 18,
	0, 68, 69, 70, 71, 66, 67, 61, 62, 63,
	64, 65, 10, 0, 8, 9, 21, 97, 93, 0,
	92, 0, 95, 94, 96, 0, 0, 0, 23, 0,
	0, 10, 24, 8, 9, 21, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 23, 0, 0,
	0, 24, 0, 0, 0, 0, 13, 0, 0, 0,
	0, 0, 0, 0, 0, 14, 15, 12, 0, 91,
	0, 16, 17, 20, 0, 13, 0, 0, 22, 0,
	19, 18, 0, 0, 14, 15, 12, 0, 0, 0,
	16, 17, 20, 0, 0, 0, 0, 108, 0, 19,
	18, 61, 62, 63, 64, 65, 0, 0, 0, 0,
	0, 97, 93, 0, 92, 0, 95, 94, 96, 52,
	0, 0, 39, 57, 0, 0, 0, 225, 46, 38,
	0, 114, 45, 0, 0, 0, 56, 41, 0, 42,
	223, 0, 0, 55, 0, 40, 43, 53, 50, 0,
	36, 54, 51, 44, 357, 47, 58, 0, 52, 0,
	0, 39, 57, 0, 0, 0, 0, 46, 38, 0,
	114, 45, 0, 0, 0, 56, 41, 0, 42, 0,
	0, 0, 55, 0, 40, 43, 53, 50, 0, 36,
	54, 51, 44, 52, 47, 58, 39, 57, 0, 0,
	0, 0, 46, 38, 0, 114, 45, 0, 0, 0,
	56, 41, 0, 42, 0, 0, 0, 55, 0, 40,
	43, 53, 50, 0, 36, 54, 51, 44, 0, 47,
	58,
}
var yyPact = []int{

	3, -1000, -1000, 1234, 911, -2, 181, 626, -1000, -1000,
	-1000, 212, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234,
	1253, 106, 716, 105, -1000, -1000, -1000, 104, -1000, -1000,
	209, 94, 1434, 344, 1185, -1000, -1000, 242, 242, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234,
	1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234,
	1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234, 1234,
	1234, 1234, 1234, 1234, -1000, -1000, 242, 242, -1000, 59,
	59, 59, 59, 59, 59, 59, 59, 59, 716, 1434,
	71, 67, 97, -1000, -1000, 1234, 237, 195, -10, 54,
	187, -1000, 256, 94, -1000, -1000, -1000, 344, 1185, -1000,
	-1000, 344, -1000, 1185, -1000, -1000, -1000, -1000, 191, -1000,
	190, 626, 264, 264, 59, 59, 59, 1270, 1270, 687,
	687, 687, 687, 1094, 1176, 1057, 1020, 983, 946, 909,
	153, 626, 626, 626, 626, 626, 626, 626, 626, 626,
	626, 626, 206, 181, 66, 112, -1000, -1000, 65, 178,
	82, -1000, 120, 256, 103, 97, 482, 64, -1000, -1000,
	1360, 1234, 82, 1434, 94, 94, 256, -1000, 37, -1000,
	-1000, -1000, 1434, 235, 1234, 95, -1000, -1000, 231, 1234,
	59, -1000, -4, 1234, 97, 1360, 50, 1434, -1000, 802,
	63, 173, -1000, -1000, 98, -1000, 109, 626, -1000, 626,
	-1000, 189, -1000, 94, -1000, 54, 55, -1000, -1000, 873,
	-1000, 94, 171, -1000, 166, 783, 1234, 432, -1000, 1152,
	108, 120, 46, -1000, 45, -1000, -1000, 1360, 120, 55,
	256, 98, -1000, -1000, -1000, -43, -1000, -1000, -44, 169,
	-1000, 55, 151, -1000, -5, 235, -1000, -1000, 1234, 38,
	-1000, 4, -1000, 148, -1000, 242, 1234, -1000, -1000, -1000,
	-1000, 98, -1000, -1000, -1000, 94, 1234, -1000, -1000, 626,
	-1000, -1000, -7, 82, -1000, -1000, -1000, 386, 610, -1000,
	626, -1000, -1000, -1000, -1000, -1000, -1000, 577, -1000, -1000,
	-1000, 93, 92, -1000, -46, -1000, -47, -48, -1000, 91,
	242, 90, 1234, 89, 74, 1234, 150, 147, 1234, 1234,
	-1000, 1399, -1000, -1000, 217, 1234, -49, 1234, -50, -1000,
	1234, 1234, 342, -1000, -1000, 32, 30, -1000, 73, -51,
	-1000, 25, -1000, 21, 20, -1000, -52, -54, 1234, 1234,
	-1000, -1000, -1000, -1000, -1000, 16, -55, 203, -1000, -1000,
	-57, 1234, -1000, -1000, 10, -1000, -1000, -1000,
}
var yyPgo = []int{

	0, 9, 341, 23, 340, 24, 37, 339, 338, 33,
	39, 336, 335, 28, 333, 332, 12, 10, 330, 326,
	1, 36, 27, 4, 323, 322, 26, 30, 35, 321,
	32, 8, 317, 38, 316, 310, 309, 11, 307, 297,
	6, 0, 5, 296, 29, 61, 104, 40, 3, 276,
	47, 45, 294, 43, 293, 25, 292, 2, 291, 34,
	13, 270, 287, 285, 283, 278, 274,
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
	44, 45, 45, 45, 45, 45, 45, 45, 45, 45,
	46, 46, 47, 47, 61, 57, 57, 57, 57, 57,
	60, 59, 6, 12, 11, 11, 11, 64, 4, 48,
	48, 58, 58, 17, 17, 13, 61, 61, 37, 20,
	20, 61, 61, 5, 24, 31, 31, 33, 33, 33,
	34, 34, 32, 32, 37, 66, 66, 65, 65, 38,
	38, 49, 49, 23, 23, 21, 21, 26, 26, 27,
	27, 7, 7, 36, 36, 8, 8, 9, 9, 29,
	29, 30, 30, 54, 54, 55, 55, 50, 50, 51,
	51, 52, 52, 53, 53, 18, 18, 19, 19, 14,
	14, 25, 25, 15, 15, 56, 56,
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
	1, 1, 1, 1, 1, 1, 3, 3, 2, 2,
	1, 2, 3, 3, 1, 1, 5, 0, 5, 1,
	1, 1, 1, 1, 3, 3, 2, 5, 2, 3,
	3, 2, 6, 2, 2, 1, 1, 2, 4, 5,
	0, 3, 1, 3, 3, 0, 1, 0, 1, 1,
	2, 0, 1, 0, 1, 0, 1, 1, 3, 0,
	1, 0, 2, 0, 2, 1, 3, 0, 1, 1,
	3, 0, 1, 1, 2, 0, 1, 1, 2, 0,
	1, 1, 2, 0, 1, 1, 3, 0, 1, 1,
	2, 0, 1, 1, 3, 1, 2,
}
var yyChk = []int{

	-1000, -62, 100, 99, -10, -22, -26, -20, 30, 31,
	28, -56, 83, 72, 81, 82, 87, 88, 97, 96,
	89, 32, 94, 44, 48, 101, -11, 6, -12, -4,
	21, -57, -50, -61, -45, -46, 40, -58, 19, 12,
	35, 27, 29, 36, 43, 22, 18, 45, -43, -44,
	38, 42, 9, 37, 41, 33, 26, 13, 46, 101,
	54, 81, 82, 83, 84, 85, 79, 80, 75, 76,
	77, 78, 73, 74, 72, 71, 70, 69, 68, 66,
	55, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 49, 94, 92, 97, 96, 98, 91, 48, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, 94, 94,
	-59, -22, -60, -57, 21, 94, 94, 48, -30, -16,
	-29, -48, 83, 94, -28, 30, 40, -61, -45, -46,
	-51, -50, -53, -52, -47, -46, -45, -48, -49, -48,
	-49, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-22, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -27, -26, -27, -22, -48, -48, -59, -59,
	95, 95, -1, 83, -2, 94, -20, 30, 53, 103,
	94, 92, 55, -7, 54, -55, -54, -44, -16, -51,
	-53, -47, 53, 53, 67, 50, 95, 93, 95, 54,
	-20, -33, 53, 92, -55, 94, -1, 54, 95, -10,
	-9, -8, -3, 30, -60, 17, -21, -20, -31, -20,
	-33, -64, -6, -57, -28, -16, -16, -44, 95, -14,
	-13, -60, -15, -5, 30, -20, 94, -20, 102, -34,
	-21, -1, -9, 95, -59, 102, 95, 54, -1, -16,
	83, 94, 93, -40, 53, -30, 102, -13, -19, -18,
	-17, -16, -49, -48, -65, 54, -25, -24, 55, -27,
	95, -32, -31, -38, -37, 91, 92, 93, 95, 95,
	-3, -55, -63, 103, 103, 54, 67, 102, -5, -20,
	95, 102, 54, -66, -37, 55, -48, -20, -42, -17,
	-20, 102, -31, 93, -6, -41, 102, -36, -39, -35,
	103, 8, 7, -40, -22, 4, 10, 14, 16, 23,
	24, 25, 34, 39, 47, 11, 15, 30, 94, 94,
	103, -42, 103, 103, -41, 94, -48, 94, -23, -22,
	94, 94, -20, 67, 67, -22, -22, 5, 47, -23,
	103, -22, 103, -22, -22, 67, 95, 95, 94, 103,
	95, 95, 95, 103, 103, -22, -23, -41, -41, -41,
	95, 103, 52, 103, -23, -41, 95, -41,
}
var yyDef = []int{

	0, -2, 3, 0, 0, 0, 6, 177, 7, 8,
	9, 10, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 215, 1, 4, 0, 134, 135,
	105, 191, 125, 199, 203, 197, 124, 171, 171, 111,
	112, 113, 114, 115, 116, 117, 118, 119, 120, 121,
	141, 142, 103, 104, 106, 107, 108, 109, 110, 2,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 179, 179, 0, 58, 59, 0, 0, 216, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 0, 0,
	0, 0, 86, 130, 105, 0, 0, 0, 0, -2,
	192, 92, 195, 0, 189, 139, 140, 199, 203, 198,
	128, 200, 129, 204, 201, 122, 123, -2, 0, -2,
	0, 178, 11, 12, 13, 14, 15, 16, 17, 18,
	19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
	0, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 0, 180, 0, 0, 149, 150, 0, 0,
	0, 54, 131, 195, 88, 86, 0, 0, 3, 133,
	187, 175, 0, 137, 0, 0, 196, 193, 0, 126,
	127, 202, 0, 0, 0, 0, 56, 57, 50, 0,
	52, 53, 160, 175, 86, 187, 0, 0, 5, 0,
	0, 188, 185, 97, 86, 100, 0, 176, 102, 155,
	156, 0, 182, 191, 190, 101, 93, 194, 94, 0,
	209, -2, 167, 213, 211, 29, 179, 0, 157, 0,
	0, 87, 0, 91, 0, 136, 95, 0, 98, 99,
	195, 86, 96, 138, 64, 0, 147, 210, 0, 208,
	205, 143, 0, -2, 0, 168, 153, 212, 0, 0,
	51, 0, 162, 165, 169, 0, 0, 90, 89, 60,
	186, 86, 61, 132, 145, 171, 0, 152, 214, 154,
	55, 158, 161, 0, 170, 166, 148, 0, 183, 206,
	144, 159, 163, 164, 62, 63, 65, 0, 69, 184,
	70, 0, 0, 73, 0, 61, 0, 0, 183, 0,
	0, 0, 173, 0, 0, 0, 0, 7, 0, 0,
	74, 183, 76, 77, 0, 173, 0, 0, 0, 174,
	0, 0, 0, 67, 68, 0, 0, 75, 0, 0,
	80, 0, 83, 0, 0, 66, 0, 0, 0, 173,
	183, 183, 183, 71, 72, 0, 0, 81, 84, 85,
	0, 173, 183, 78, 0, 82, 183, 79,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 87, 3, 3, 3, 85, 72, 3,
	94, 95, 83, 81, 54, 82, 91, 84, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 67, 103,
	75, 55, 76, 66, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 92, 3, 93, 71, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 53, 70, 102, 88,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 56, 57, 58, 59, 60, 61, 62, 63, 64,
	65, 68, 69, 73, 74, 77, 78, 79, 80, 86,
	89, 90, 96, 97, 98, 99, 100, 101,
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
		//line cc.y:189
		{
			yylex.(*lexer).prog = &Prog{Decls: yyS[yypt-1].decls}
			return 0
		}
	case 2:
		//line cc.y:194
		{
			yylex.(*lexer).expr = yyS[yypt-1].expr
			return 0
		}
	case 3:
		//line cc.y:200
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 4:
		//line cc.y:205
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 5:
		//line cc.y:210
		{
		}
	case 6:
		//line cc.y:215
		{
			yyVAL.span = yyS[yypt-0].span
			if len(yyS[yypt-0].exprs) == 1 {
				yyVAL.expr = yyS[yypt-0].exprs[0]
				break
			}
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Comma, List: yyS[yypt-0].exprs}
		}
	case 7:
		//line cc.y:226
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Name, Text: yyS[yypt-0].str, XDecl: yyS[yypt-0].decl}
		}
	case 8:
		//line cc.y:231
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyS[yypt-0].str}
		}
	case 9:
		//line cc.y:236
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Number, Text: yyS[yypt-0].str}
		}
	case 10:
		//line cc.y:241
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: String, Texts: yyS[yypt-0].strs}
		}
	case 11:
		//line cc.y:246
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Add, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 12:
		//line cc.y:251
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Sub, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 13:
		//line cc.y:256
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mul, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 14:
		//line cc.y:261
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Div, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 15:
		//line cc.y:266
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Mod, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 16:
		//line cc.y:271
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 17:
		//line cc.y:276
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Rsh, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 18:
		//line cc.y:281
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Lt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 19:
		//line cc.y:286
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Gt, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 20:
		//line cc.y:291
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 21:
		//line cc.y:296
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: GtEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 22:
		//line cc.y:301
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: EqEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 23:
		//line cc.y:306
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: NotEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 24:
		//line cc.y:311
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: And, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 25:
		//line cc.y:316
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Xor, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 26:
		//line cc.y:321
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Or, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 27:
		//line cc.y:326
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndAnd, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 28:
		//line cc.y:331
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrOr, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 29:
		//line cc.y:336
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cond, List: []*Expr{yyS[yypt-4].expr, yyS[yypt-2].expr, yyS[yypt-0].expr}}
		}
	case 30:
		//line cc.y:341
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Eq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 31:
		//line cc.y:346
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AddEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 32:
		//line cc.y:351
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SubEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 33:
		//line cc.y:356
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: MulEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 34:
		//line cc.y:361
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: DivEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 35:
		//line cc.y:366
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ModEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 36:
		//line cc.y:371
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 37:
		//line cc.y:376
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: RshEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 38:
		//line cc.y:381
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: AndEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 39:
		//line cc.y:386
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: XorEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 40:
		//line cc.y:391
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: OrEq, Left: yyS[yypt-2].expr, Right: yyS[yypt-0].expr}
		}
	case 41:
		//line cc.y:396
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Indir, Left: yyS[yypt-0].expr}
		}
	case 42:
		//line cc.y:401
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Addr, Left: yyS[yypt-0].expr}
		}
	case 43:
		//line cc.y:406
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Plus, Left: yyS[yypt-0].expr}
		}
	case 44:
		//line cc.y:411
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Minus, Left: yyS[yypt-0].expr}
		}
	case 45:
		//line cc.y:416
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Not, Left: yyS[yypt-0].expr}
		}
	case 46:
		//line cc.y:421
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Twid, Left: yyS[yypt-0].expr}
		}
	case 47:
		//line cc.y:426
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreInc, Left: yyS[yypt-0].expr}
		}
	case 48:
		//line cc.y:431
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PreDec, Left: yyS[yypt-0].expr}
		}
	case 49:
		//line cc.y:436
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofExpr, Left: yyS[yypt-0].expr}
		}
	case 50:
		//line cc.y:441
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: SizeofType, Type: yyS[yypt-1].typ}
		}
	case 51:
		//line cc.y:446
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Offsetof, Type: yyS[yypt-3].typ, Left: yyS[yypt-1].expr}
		}
	case 52:
		//line cc.y:451
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Cast, Type: yyS[yypt-2].typ, Left: yyS[yypt-0].expr}
		}
	case 53:
		//line cc.y:456
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CastInit, Type: yyS[yypt-2].typ, Init: &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits}}
		}
	case 54:
		//line cc.y:461
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Paren, Left: yyS[yypt-1].expr}
		}
	case 55:
		//line cc.y:466
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: CUDACall, Left: yyS[yypt-6].expr, LaunchParams: yyS[yypt-4].exprs, List: yyS[yypt-1].exprs}
		}
	case 56:
		//line cc.y:471
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Call, Left: yyS[yypt-3].expr, List: yyS[yypt-1].exprs}
		}
	case 57:
		//line cc.y:476
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Index, Left: yyS[yypt-3].expr, Right: yyS[yypt-1].expr}
		}
	case 58:
		//line cc.y:481
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostInc, Left: yyS[yypt-1].expr}
		}
	case 59:
		//line cc.y:486
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: PostDec, Left: yyS[yypt-1].expr}
		}
	case 60:
		//line cc.y:491
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: VaArg, Left: yyS[yypt-3].expr, Type: yyS[yypt-1].typ}
		}
	case 61:
		//line cc.y:497
		{
			yyVAL.span = Span{}
			yyVAL.stmts = nil
		}
	case 62:
		//line cc.y:502
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = yyS[yypt-1].stmts
			for _, d := range yyS[yypt-0].decls {
				yyVAL.stmts = append(yyVAL.stmts, &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtDecl, Decl: d})
			}
		}
	case 63:
		//line cc.y:510
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmts = append(yyS[yypt-1].stmts, yyS[yypt-0].stmt)
		}
	case 64:
		//line cc.y:517
		{
			yylex.(*lexer).pushScope()
		}
	case 65:
		//line cc.y:521
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yylex.(*lexer).popScope()
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Block, Block: yyS[yypt-1].stmts}
		}
	case 66:
		//line cc.y:529
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Case, Expr: yyS[yypt-1].expr}
		}
	case 67:
		//line cc.y:534
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Default}
		}
	case 68:
		//line cc.y:539
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.label = &Label{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: LabelName, Name: yyS[yypt-1].str}
		}
	case 69:
		//line cc.y:546
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = yyS[yypt-0].stmt
			yyVAL.stmt.Labels = yyS[yypt-1].labels
		}
	case 70:
		//line cc.y:554
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 71:
		//line cc.y:559
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 72:
		//line cc.y:564
		{
			yyVAL.span = yyS[yypt-4].span
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Empty}
		}
	case 73:
		//line cc.y:569
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.stmt = yyS[yypt-0].stmt
		}
	case 74:
		//line cc.y:574
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: StmtExpr, Expr: yyS[yypt-1].expr}
		}
	case 75:
		//line cc.y:579
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: ARGBEGIN, Block: yyS[yypt-1].stmts}
		}
	case 76:
		//line cc.y:584
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Break}
		}
	case 77:
		//line cc.y:589
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Continue}
		}
	case 78:
		//line cc.y:594
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Do, Body: yyS[yypt-5].stmt, Expr: yyS[yypt-2].expr}
		}
	case 79:
		//line cc.y:599
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
		//line cc.y:610
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Goto, Text: yyS[yypt-1].str}
		}
	case 81:
		//line cc.y:615
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 82:
		//line cc.y:620
		{
			yyVAL.span = span(yyS[yypt-6].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: If, Expr: yyS[yypt-4].expr, Body: yyS[yypt-2].stmt, Else: yyS[yypt-0].stmt}
		}
	case 83:
		//line cc.y:625
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Return, Expr: yyS[yypt-1].expr}
		}
	case 84:
		//line cc.y:630
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Switch, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 85:
		//line cc.y:635
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.stmt = &Stmt{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: While, Expr: yyS[yypt-2].expr, Body: yyS[yypt-0].stmt}
		}
	case 86:
		//line cc.y:642
		{
			yyVAL.span = Span{}
			yyVAL.abdecor = func(t *Type) *Type { return t }
		}
	case 87:
		//line cc.y:647
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			_, q, _ := splitTypeWords(yyS[yypt-1].strs)
			abdecor := yyS[yypt-0].abdecor
			yyVAL.abdecor = func(t *Type) *Type {
				return abdecor(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Ptr, Base: t, Qual: q})
			}
		}
	case 88:
		//line cc.y:656
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.abdecor = yyS[yypt-0].abdecor
		}
	case 89:
		//line cc.y:663
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
		//line cc.y:687
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
		//line cc.y:698
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.abdecor = yyS[yypt-1].abdecor
		}
	case 92:
		//line cc.y:706
		{
			yyVAL.span = yyS[yypt-0].span
			name := yyS[yypt-0].str
			yyVAL.decor = func(t *Type) (*Type, string) { return t, name }
		}
	case 93:
		//line cc.y:712
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
		//line cc.y:722
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decor = yyS[yypt-1].decor
		}
	case 95:
		//line cc.y:727
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
		//line cc.y:737
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
		//line cc.y:750
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyS[yypt-0].str}
		}
	case 98:
		//line cc.y:755
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Type: yyS[yypt-0].abdecor(yyS[yypt-1].typ)}
		}
	case 99:
		//line cc.y:760
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			typ, name := yyS[yypt-0].decor(yyS[yypt-1].typ)
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: name, Type: typ}
		}
	case 100:
		//line cc.y:766
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: "..."}
		}
	case 101:
		//line cc.y:774
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idec = idecor{yyS[yypt-0].decor, nil}
		}
	case 102:
		//line cc.y:779
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idec = idecor{yyS[yypt-2].decor, yyS[yypt-0].init}
		}
	case 103:
		//line cc.y:787
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 104:
		//line cc.y:792
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 105:
		//line cc.y:797
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 106:
		//line cc.y:802
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 107:
		//line cc.y:807
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 108:
		//line cc.y:812
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 109:
		//line cc.y:820
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 110:
		//line cc.y:825
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 111:
		//line cc.y:833
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 112:
		//line cc.y:838
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 113:
		//line cc.y:843
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 114:
		//line cc.y:848
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 115:
		//line cc.y:853
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 116:
		//line cc.y:858
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 117:
		//line cc.y:863
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 118:
		//line cc.y:868
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 119:
		//line cc.y:873
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 120:
		//line cc.y:880
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 121:
		//line cc.y:885
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 122:
		//line cc.y:892
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 123:
		//line cc.y:897
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 124:
		//line cc.y:905
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.typ = yyS[yypt-0].typ
			if yyVAL.typ == nil {
				yyVAL.typ = &Type{Kind: TypedefType, Name: yyS[yypt-0].str}
			}
		}
	case 125:
		//line cc.y:921
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(append(yyS[yypt-0].strs, "int"))
		}
	case 126:
		//line cc.y:926
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(append(yyS[yypt-2].strs, yyS[yypt-0].strs...))
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 127:
		//line cc.y:932
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyS[yypt-2].strs = append(yyS[yypt-2].strs, yyS[yypt-1].str)
			yyS[yypt-2].strs = append(yyS[yypt-2].strs, yyS[yypt-0].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(yyS[yypt-2].strs)
		}
	case 128:
		//line cc.y:939
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.tc.c, yyVAL.tc.q, _ = splitTypeWords(yyS[yypt-0].strs)
			yyVAL.tc.t = yyS[yypt-1].typ
		}
	case 129:
		//line cc.y:945
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var ts []string
			ts = append(ts, yyS[yypt-1].str)
			ts = append(ts, yyS[yypt-0].strs...)
			yyVAL.tc.c, yyVAL.tc.q, yyVAL.tc.t = splitTypeWords(ts)
		}
	case 130:
		//line cc.y:956
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
	case 131:
		//line cc.y:969
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yyS[yypt-0].abdecor(yyS[yypt-1].typ)
		}
	case 132:
		//line cc.y:977
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
	case 133:
		//line cc.y:997
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
	case 134:
		//line cc.y:1025
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 135:
		//line cc.y:1030
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 136:
		//line cc.y:1035
		{
			yyVAL.decls = yyS[yypt-1].decls
		}
	case 137:
		//line cc.y:1041
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
	case 138:
		//line cc.y:1062
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
	case 139:
		//line cc.y:1075
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 140:
		//line cc.y:1080
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 141:
		//line cc.y:1088
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Struct
		}
	case 142:
		//line cc.y:1093
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.tk = Union
		}
	case 143:
		//line cc.y:1100
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decor = yyS[yypt-0].decor
		}
	case 144:
		//line cc.y:1105
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			name := yyS[yypt-2].str
			expr := yyS[yypt-0].expr
			yyVAL.decor = func(t *Type) (*Type, string) {
				t.Width = expr
				return t, name
			}
		}
	case 145:
		//line cc.y:1117
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
	case 146:
		//line cc.y:1131
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyS[yypt-1].tk, Tag: yyS[yypt-0].str})
		}
	case 147:
		//line cc.y:1136
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: yyS[yypt-4].tk, Tag: yyS[yypt-3].str, Decls: yyS[yypt-1].decls})
		}
	case 148:
		//line cc.y:1143
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Dot: yyS[yypt-0].str}
		}
	case 149:
		//line cc.y:1150
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Arrow, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].str}
		}
	case 150:
		//line cc.y:1155
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.expr = &Expr{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Op: Dot, Left: yyS[yypt-2].expr, Text: yyS[yypt-0].str}
		}
	case 151:
		//line cc.y:1163
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-0].str})
		}
	case 152:
		//line cc.y:1168
		{
			yyVAL.span = span(yyS[yypt-5].span, yyS[yypt-0].span)
			yyVAL.typ = yylex.(*lexer).pushType(&Type{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Kind: Enum, Tag: yyS[yypt-4].str, Decls: yyS[yypt-2].decls})
		}
	case 153:
		//line cc.y:1175
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			var x *Init
			if yyS[yypt-0].expr != nil {
				x = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr}
			}
			yyVAL.decl = &Decl{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Name: yyS[yypt-1].str, Init: x}
			yylex.(*lexer).pushDecl(yyVAL.decl)
		}
	case 154:
		//line cc.y:1187
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 155:
		//line cc.y:1195
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Expr: yyS[yypt-0].expr}
		}
	case 156:
		//line cc.y:1200
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = &Init{SyntaxInfo: SyntaxInfo{Span: yyVAL.span}, Braced: yyS[yypt-0].inits}
		}
	case 157:
		//line cc.y:1207
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.inits = []*Init{}
		}
	case 158:
		//line cc.y:1212
		{
			yyVAL.span = span(yyS[yypt-3].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 159:
		//line cc.y:1217
		{
			yyVAL.span = span(yyS[yypt-4].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-3].inits, yyS[yypt-2].init)
		}
	case 160:
		//line cc.y:1223
		{
			yyVAL.span = Span{}
			yyVAL.inits = nil
		}
	case 161:
		//line cc.y:1228
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.inits = append(yyS[yypt-2].inits, yyS[yypt-1].init)
		}
	case 162:
		//line cc.y:1235
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.init = yyS[yypt-0].init
		}
	case 163:
		//line cc.y:1240
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.init = yyS[yypt-0].init
			yyVAL.init.Prefix = yyS[yypt-2].prefixes
		}
	case 164:
		//line cc.y:1248
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.prefix = &Prefix{Span: yyVAL.span, Index: yyS[yypt-1].expr}
		}
	case 165:
		//line cc.y:1254
		{
			yyVAL.span = Span{}
		}
	case 166:
		//line cc.y:1258
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 167:
		//line cc.y:1263
		{
			yyVAL.span = Span{}
		}
	case 168:
		//line cc.y:1267
		{
			yyVAL.span = yyS[yypt-0].span
		}
	case 169:
		//line cc.y:1276
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.prefixes = []*Prefix{yyS[yypt-0].prefix}
		}
	case 170:
		//line cc.y:1281
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.prefixes = append(yyS[yypt-1].prefixes, yyS[yypt-0].prefix)
		}
	case 171:
		//line cc.y:1287
		{
			yyVAL.span = Span{}
			yyVAL.str = ""
		}
	case 172:
		//line cc.y:1292
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.str = yyS[yypt-0].str
		}
	case 173:
		//line cc.y:1298
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 174:
		//line cc.y:1303
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 175:
		//line cc.y:1309
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 176:
		//line cc.y:1314
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 177:
		//line cc.y:1321
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = []*Expr{yyS[yypt-0].expr}
		}
	case 178:
		//line cc.y:1326
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.exprs = append(yyS[yypt-2].exprs, yyS[yypt-0].expr)
		}
	case 179:
		//line cc.y:1332
		{
			yyVAL.span = Span{}
			yyVAL.exprs = nil
		}
	case 180:
		//line cc.y:1337
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.exprs = yyS[yypt-0].exprs
		}
	case 181:
		//line cc.y:1343
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 182:
		//line cc.y:1348
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 183:
		//line cc.y:1354
		{
			yyVAL.span = Span{}
			yyVAL.labels = nil
		}
	case 184:
		//line cc.y:1359
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.labels = append(yyS[yypt-1].labels, yyS[yypt-0].label)
		}
	case 185:
		//line cc.y:1366
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 186:
		//line cc.y:1371
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 187:
		//line cc.y:1377
		{
			yyVAL.span = Span{}
			yyVAL.decls = nil
		}
	case 188:
		//line cc.y:1382
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 189:
		//line cc.y:1389
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = []idecor{yyS[yypt-0].idec}
		}
	case 190:
		//line cc.y:1394
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.idecs = append(yyS[yypt-2].idecs, yyS[yypt-0].idec)
		}
	case 191:
		//line cc.y:1400
		{
			yyVAL.span = Span{}
			yyVAL.idecs = nil
		}
	case 192:
		//line cc.y:1405
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.idecs = yyS[yypt-0].idecs
		}
	case 193:
		//line cc.y:1412
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 194:
		//line cc.y:1417
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 195:
		//line cc.y:1423
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 196:
		//line cc.y:1428
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 197:
		//line cc.y:1435
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 198:
		//line cc.y:1440
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 199:
		//line cc.y:1446
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 200:
		//line cc.y:1451
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 201:
		//line cc.y:1458
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 202:
		//line cc.y:1463
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	case 203:
		//line cc.y:1469
		{
			yyVAL.span = Span{}
			yyVAL.strs = nil
		}
	case 204:
		//line cc.y:1474
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = yyS[yypt-0].strs
		}
	case 205:
		//line cc.y:1481
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = nil
			yyVAL.decors = append(yyVAL.decors, yyS[yypt-0].decor)
		}
	case 206:
		//line cc.y:1487
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decors = append(yyS[yypt-2].decors, yyS[yypt-0].decor)
		}
	case 207:
		//line cc.y:1493
		{
			yyVAL.span = Span{}
			yyVAL.decors = nil
		}
	case 208:
		//line cc.y:1498
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decors = yyS[yypt-0].decors
		}
	case 209:
		//line cc.y:1505
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = yyS[yypt-0].decls
		}
	case 210:
		//line cc.y:1510
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-1].decls, yyS[yypt-0].decls...)
		}
	case 211:
		//line cc.y:1516
		{
			yyVAL.span = Span{}
			yyVAL.expr = nil
		}
	case 212:
		//line cc.y:1521
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.expr = yyS[yypt-0].expr
		}
	case 213:
		//line cc.y:1528
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.decls = []*Decl{yyS[yypt-0].decl}
		}
	case 214:
		//line cc.y:1533
		{
			yyVAL.span = span(yyS[yypt-2].span, yyS[yypt-0].span)
			yyVAL.decls = append(yyS[yypt-2].decls, yyS[yypt-0].decl)
		}
	case 215:
		//line cc.y:1540
		{
			yyVAL.span = yyS[yypt-0].span
			yyVAL.strs = []string{yyS[yypt-0].str}
		}
	case 216:
		//line cc.y:1545
		{
			yyVAL.span = span(yyS[yypt-1].span, yyS[yypt-0].span)
			yyVAL.strs = append(yyS[yypt-1].strs, yyS[yypt-0].str)
		}
	}
	goto yystack /* stack new state and value */
}
