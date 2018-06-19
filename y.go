//line fateyacc.y:1
package main

import __yyfmt__ "fmt"

//line fateyacc.y:3
//line fateyacc.y:5
type yySymType struct {
	yys      int
	exprNode ExprNode
}

const INT_LITERAL = 57346
const FLOAT_LITERAL = 57347
const STRING_LITERAL = 57348
const CHAR_LITERAL = 57349
const BOOL_LITERAL = 57350
const IDENTIFIER = 57351
const INC_OP = 57352
const DEC_OP = 57353
const LEFT_OP = 57354
const RIGHT_OP = 57355
const LE_OP = 57356
const GE_OP = 57357
const EQ_OP = 57358
const NE_OP = 57359
const AND_OP = 57360
const OR_OP = 57361
const MUL_ASSIGN = 57362
const DIV_ASSIGN = 57363
const MOD_ASSIGN = 57364
const ADD_ASSIGN = 57365
const SUB_ASSIGN = 57366
const LEFT_ASSIGN = 57367
const RIGHT_ASSIGN = 57368
const AND_ASSIGN = 57369
const XOR_ASSIGN = 57370
const OR_ASSIGN = 57371
const INT8 = 57372
const INT16 = 57373
const INT32 = 57374
const INT64 = 57375
const UINT8 = 57376
const UINT16 = 57377
const UINT32 = 57378
const UINT64 = 57379
const FLOAT32 = 57380
const FLOAT64 = 57381
const BOOL = 57382
const CHAR = 57383
const STRING = 57384
const TYPE_CONST = 57385
const TYPE_STRUCT = 57386
const STRUCT = 57387
const CONST = 57388
const EXTERN = 57389
const LET = 57390
const SET = 57391
const FUNC = 57392
const TYPE = 57393
const CASE = 57394
const DEFAULT = 57395
const IF = 57396
const ELSE = 57397
const SWITCH = 57398
const FOR = 57399
const GOTO = 57400
const CONTINUE = 57401
const BREAK = 57402
const RETURN = 57403

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"INT_LITERAL",
	"FLOAT_LITERAL",
	"STRING_LITERAL",
	"CHAR_LITERAL",
	"BOOL_LITERAL",
	"IDENTIFIER",
	"INC_OP",
	"DEC_OP",
	"LEFT_OP",
	"RIGHT_OP",
	"LE_OP",
	"GE_OP",
	"EQ_OP",
	"NE_OP",
	"AND_OP",
	"OR_OP",
	"MUL_ASSIGN",
	"DIV_ASSIGN",
	"MOD_ASSIGN",
	"ADD_ASSIGN",
	"SUB_ASSIGN",
	"LEFT_ASSIGN",
	"RIGHT_ASSIGN",
	"AND_ASSIGN",
	"XOR_ASSIGN",
	"OR_ASSIGN",
	"INT8",
	"INT16",
	"INT32",
	"INT64",
	"UINT8",
	"UINT16",
	"UINT32",
	"UINT64",
	"FLOAT32",
	"FLOAT64",
	"BOOL",
	"CHAR",
	"STRING",
	"TYPE_CONST",
	"TYPE_STRUCT",
	"STRUCT",
	"CONST",
	"EXTERN",
	"LET",
	"SET",
	"FUNC",
	"TYPE",
	"CASE",
	"DEFAULT",
	"IF",
	"ELSE",
	"SWITCH",
	"FOR",
	"GOTO",
	"CONTINUE",
	"BREAK",
	"RETURN",
	"'('",
	"')'",
	"';'",
	"','",
	"'{'",
	"'}'",
	"':'",
	"'?'",
	"'|'",
	"'^'",
	"'&'",
	"'<'",
	"'>'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"'!'",
	"'['",
	"']'",
	"'.'",
	"'='",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line fateyacc.y:331

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 467

var yyAct = [...]int{

	138, 149, 141, 126, 127, 62, 72, 23, 6, 5,
	135, 48, 69, 67, 65, 70, 140, 55, 120, 121,
	16, 18, 66, 50, 128, 60, 41, 103, 182, 112,
	113, 74, 102, 51, 108, 109, 73, 85, 88, 89,
	87, 86, 84, 77, 78, 212, 71, 114, 115, 116,
	101, 219, 182, 201, 68, 211, 59, 99, 199, 96,
	63, 47, 130, 238, 129, 182, 95, 225, 182, 93,
	118, 57, 230, 58, 182, 94, 229, 213, 182, 214,
	75, 202, 182, 97, 181, 91, 182, 54, 53, 117,
	54, 119, 125, 106, 107, 90, 234, 236, 224, 208,
	158, 131, 122, 123, 75, 136, 137, 98, 81, 82,
	207, 124, 132, 83, 159, 161, 163, 170, 171, 177,
	164, 165, 179, 160, 75, 75, 75, 75, 75, 75,
	75, 75, 75, 75, 75, 75, 75, 75, 75, 75,
	75, 75, 196, 136, 137, 198, 174, 175, 176, 172,
	173, 63, 200, 166, 167, 168, 169, 162, 210, 157,
	185, 186, 187, 188, 189, 190, 191, 192, 193, 194,
	195, 75, 52, 49, 20, 205, 204, 203, 45, 22,
	12, 46, 8, 9, 7, 13, 215, 216, 42, 104,
	105, 206, 100, 110, 111, 17, 180, 44, 21, 15,
	218, 2, 220, 14, 217, 221, 222, 92, 223, 43,
	183, 178, 79, 134, 144, 143, 226, 142, 227, 139,
	228, 11, 10, 4, 184, 3, 231, 19, 1, 76,
	232, 233, 64, 235, 56, 237, 75, 239, 80, 240,
	85, 88, 89, 87, 86, 145, 77, 78, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 39, 40, 0, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 85, 88, 89, 87, 86, 145, 77,
	78, 0, 12, 0, 8, 9, 0, 13, 146, 147,
	150, 25, 151, 152, 153, 154, 155, 156, 90, 24,
	148, 0, 96, 197, 0, 0, 0, 0, 0, 0,
	0, 81, 82, 0, 0, 12, 83, 8, 9, 0,
	13, 146, 147, 150, 0, 151, 152, 153, 154, 155,
	156, 90, 0, 148, 0, 96, 133, 0, 0, 0,
	0, 0, 0, 0, 81, 82, 0, 0, 0, 83,
	85, 88, 89, 87, 86, 145, 77, 78, 85, 88,
	89, 87, 86, 84, 77, 78, 0, 85, 88, 89,
	87, 86, 84, 77, 78, 85, 88, 89, 87, 86,
	84, 77, 78, 26, 27, 28, 29, 30, 31, 32,
	33, 34, 35, 36, 37, 38, 39, 40, 146, 147,
	150, 0, 151, 152, 153, 154, 155, 156, 90, 0,
	148, 0, 96, 0, 0, 0, 90, 0, 148, 96,
	0, 81, 82, 0, 0, 90, 83, 209, 0, 81,
	82, 0, 0, 90, 83, 0, 0, 0, 81, 82,
	0, 0, 0, 83, 0, 0, 81, 82, 0, 0,
	0, 83, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40,
}
var yyPact = [...]int{

	134, 134, -1000, -1000, -1000, -1000, -1000, 190, 186, 186,
	-1000, -1000, 165, 189, -1000, 117, 218, -1000, 218, 188,
	116, 136, 186, 109, -59, 422, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 108, 25, -1000, -67, 188, 5, 10, 218, -1000,
	-1000, -49, -1000, -1000, 188, 371, 22, 186, 353, 18,
	-1000, -1000, -1000, -1000, 38, 174, -20, -39, -45, 173,
	20, 181, -46, -30, -1000, -1000, 8, 371, 371, 371,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	371, -1000, -3, 218, -1000, -7, 269, 186, 371, 371,
	371, 371, 371, 371, 371, 371, 371, 371, 371, 371,
	371, 371, 371, 371, 371, 371, 371, 371, 371, 187,
	-1000, -1000, -1000, -1000, -1000, 21, -1000, -1000, 140, 106,
	186, -1000, -1000, -1000, 236, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -10, 371, -15, -1000, 17,
	115, 114, 113, 182, 46, 35, 363, -1000, -13, 174,
	-20, -39, -45, 173, 20, 20, 181, 181, 181, 181,
	-46, -46, -30, -30, -1000, -1000, -1000, -37, 14, -1000,
	-1000, -1000, 371, 371, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 218, -1000, -1000, 346,
	-17, 346, -1000, 371, 371, 354, 34, -1000, -1000, -1000,
	3, 371, -1000, -1000, 371, -1000, -1000, -1000, -1000, 346,
	-1000, 13, 9, 354, -1000, -1000, -1000, -1000, -1000, 346,
	346, 33, 42, -1000, 346, 0, 346, -1000, 346, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 238, 1, 3, 2, 5, 4, 232, 14, 22,
	13, 54, 12, 15, 46, 6, 31, 36, 24, 229,
	228, 7, 201, 225, 223, 9, 8, 222, 221, 221,
	61, 16, 11, 0, 219, 217, 215, 214, 213, 10,
	212, 211, 210, 207, 188, 209,
}
var yyR1 = [...]int{

	0, 21, 21, 21, 21, 21, 21, 21, 21, 21,
	21, 21, 21, 21, 21, 21, 20, 20, 22, 22,
	22, 23, 26, 26, 29, 29, 24, 24, 25, 25,
	30, 30, 30, 33, 33, 33, 33, 33, 33, 31,
	31, 38, 38, 39, 39, 39, 34, 34, 34, 4,
	4, 35, 35, 35, 36, 36, 37, 37, 37, 37,
	37, 5, 6, 6, 7, 7, 8, 8, 9, 9,
	10, 10, 11, 11, 12, 12, 12, 13, 13, 13,
	13, 13, 14, 14, 14, 15, 15, 15, 17, 17,
	17, 17, 16, 40, 40, 40, 18, 18, 18, 18,
	19, 19, 19, 19, 19, 19, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 3, 3, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 28, 43,
	43, 27, 27, 44, 44, 45, 45, 32, 32, 32,
	41, 41,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 7, 8, 6, 7, 4, 4,
	0, 2, 4, 1, 1, 1, 1, 1, 1, 2,
	3, 1, 2, 1, 1, 1, 3, 4, 3, 1,
	2, 5, 7, 5, 6, 7, 3, 2, 2, 2,
	3, 1, 1, 5, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 3, 1, 3, 3, 1, 3, 3,
	3, 3, 1, 3, 3, 1, 3, 3, 1, 3,
	3, 3, 1, 1, 1, 1, 1, 2, 2, 2,
	1, 4, 4, 3, 2, 2, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 7, 2,
	4, 4, 5, 1, 3, 1, 3, 1, 3, 4,
	1, 3,
}
var yyChk = [...]int{

	-1000, -20, -22, -23, -24, -25, -26, 50, 48, 49,
	-27, -28, 46, 51, -22, 9, -32, 9, -32, 62,
	9, 9, 62, -21, 81, 73, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, -21, -44, -45, 9, 62, 45, -30, -32, 64,
	82, -21, 64, 63, 65, 84, -44, 66, 63, -21,
	74, -45, -5, -6, -7, -8, -9, -10, -11, -12,
	-13, -14, -15, -17, -16, -18, -19, 10, 11, -40,
	-1, 75, 76, 80, 9, 4, 8, 7, 5, 6,
	62, 63, -43, -32, -31, -21, 66, 65, 69, 19,
	18, 70, 71, 72, 16, 17, 73, 74, 14, 15,
	12, 13, 75, 76, 77, 78, 79, 81, 62, 83,
	10, 11, -18, -18, -16, -2, -3, -6, -18, 67,
	65, -21, -31, 67, -38, -39, -26, -25, -33, -34,
	-31, -4, -35, -36, -37, 9, 52, 53, 64, -2,
	54, 56, 57, 58, 59, 60, 61, -30, -2, -8,
	-9, -10, -11, -12, -13, -13, -14, -14, -14, -14,
	-15, -15, -17, -17, -16, -16, -16, -2, -41, -3,
	9, 63, 65, -42, 84, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 64, -32, 67, -39, 68,
	-5, 68, 64, 62, 62, 62, 9, 64, 64, 64,
	-2, 68, 82, 63, 65, -3, -3, -21, -33, 68,
	-33, -2, -2, -4, 64, 64, -6, -3, -33, 63,
	63, -4, -33, -33, 63, -2, 55, -33, 63, -33,
	-33,
}
var yyDef = [...]int{

	0, -2, 16, 18, 19, 20, 21, 0, 0, 0,
	22, 23, 0, 0, 17, 0, 0, 137, 0, 0,
	0, 0, 30, 0, 0, 0, 1, 2, 3, 4,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 0, 0, 133, 135, 0, 0, 0, 0, 28,
	138, 0, 29, 131, 0, 0, 0, 0, 0, 31,
	139, 134, 136, 61, 62, 64, 66, 68, 70, 72,
	74, 77, 82, 85, 88, 92, 96, 0, 0, 0,
	100, 93, 94, 95, 106, 107, 108, 109, 110, 111,
	0, 132, 0, 0, 26, 0, 0, 30, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	104, 105, 97, 98, 99, 0, 113, 115, 92, 0,
	0, 129, 27, 39, 0, 41, 43, 44, 45, 33,
	34, 35, 36, 37, 38, 106, 0, 0, 49, 0,
	0, 0, 0, 0, 0, 0, 0, 32, 0, 65,
	67, 69, 71, 73, 75, 76, 78, 79, 80, 81,
	83, 84, 86, 87, 89, 90, 91, 0, 0, 140,
	103, 112, 0, 0, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 128, 0, 40, 42, 0,
	0, 0, 50, 0, 0, 0, 0, 57, 58, 59,
	0, 0, 101, 102, 0, 114, 116, 130, 46, 0,
	48, 0, 0, 0, 56, 60, 63, 141, 47, 0,
	0, 0, 51, 53, 0, 0, 0, 54, 0, 52,
	55,
}
var yyTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 80, 3, 3, 3, 79, 72, 3,
	62, 63, 77, 75, 65, 76, 83, 78, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 68, 64,
	73, 84, 74, 69, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 81, 3, 82, 71, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 66, 70, 67,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
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

	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:129
		{
			yyVAL.exprNode = aa()
		}
	case 50:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:133
		{
			yyVAL.exprNode = yyDollar[1].exprNode
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:237
		{
			yyVAL.exprNode = aa()
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:241
		{
			yyVAL.exprNode = aa()
		}
	case 99:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:245
		{
			yyVAL.exprNode = aa()
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:267
		{
			yyVAL.exprNode = yyDollar[2].exprNode
		}
	}
	goto yystack /* stack new state and value */
}
