//line yacc.y:2
package main

import __yyfmt__ "fmt"

//line yacc.y:2
//line yacc.y:5
type yySymType struct {
	yys  int
	bval bool
	cval byte
	ival uint64
	sval string
	fval float64
}

const IDENTIFIER = 57346
const INT_CONSTANT = 57347
const CHAR_CONSTANT = 57348
const BOOL_CONSTANT = 57349
const FLOAT_CONSTANT = 57350
const STRING_CONSTANT = 57351
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
	"IDENTIFIER",
	"INT_CONSTANT",
	"CHAR_CONSTANT",
	"BOOL_CONSTANT",
	"FLOAT_CONSTANT",
	"STRING_CONSTANT",
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

//line yacc.y:324

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
	113, 74, 102, 51, 108, 109, 73, 84, 85, 87,
	86, 88, 89, 77, 78, 212, 71, 114, 115, 116,
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
	12, 42, 8, 9, 7, 13, 215, 216, 46, 104,
	105, 206, 100, 110, 111, 17, 180, 44, 21, 15,
	218, 2, 220, 14, 217, 221, 222, 92, 223, 43,
	183, 178, 80, 76, 79, 64, 226, 134, 227, 144,
	228, 143, 142, 139, 184, 11, 231, 56, 10, 4,
	232, 233, 19, 235, 3, 237, 75, 239, 1, 240,
	145, 85, 87, 86, 88, 89, 77, 78, 26, 27,
	28, 29, 30, 31, 32, 33, 34, 35, 36, 37,
	38, 39, 40, 0, 61, 0, 0, 0, 0, 0,
	0, 0, 0, 145, 85, 87, 86, 88, 89, 77,
	78, 0, 12, 0, 8, 9, 0, 13, 146, 147,
	150, 25, 151, 152, 153, 154, 155, 156, 90, 24,
	148, 0, 96, 197, 0, 0, 0, 0, 0, 0,
	0, 81, 82, 0, 0, 12, 83, 8, 9, 0,
	13, 146, 147, 150, 0, 151, 152, 153, 154, 155,
	156, 90, 0, 148, 0, 96, 133, 0, 0, 0,
	0, 0, 0, 0, 81, 82, 0, 0, 0, 83,
	145, 85, 87, 86, 88, 89, 77, 78, 84, 85,
	87, 86, 88, 89, 77, 78, 0, 84, 85, 87,
	86, 88, 89, 77, 78, 84, 85, 87, 86, 88,
	89, 77, 78, 26, 27, 28, 29, 30, 31, 32,
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

	134, 134, -1000, -1000, -1000, -1000, -1000, 195, 191, 191,
	-1000, -1000, 170, 194, -1000, 117, 218, -1000, 218, 193,
	116, 143, 191, 109, -59, 422, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 108, 25, -1000, -67, 193, 5, 10, 218, -1000,
	-1000, -49, -1000, -1000, 193, 371, 22, 191, 353, 18,
	-1000, -1000, -1000, -1000, 38, 174, -20, -39, -45, 173,
	20, 181, -46, -30, -1000, -1000, 8, 371, 371, 371,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	371, -1000, -3, 218, -1000, -7, 269, 191, 371, 371,
	371, 371, 371, 371, 371, 371, 371, 371, 371, 371,
	371, 371, 371, 371, 371, 371, 371, 371, 371, 192,
	-1000, -1000, -1000, -1000, -1000, 21, -1000, -1000, 140, 106,
	191, -1000, -1000, -1000, 236, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -10, 371, -15, -1000, 17,
	115, 114, 113, 187, 46, 35, 363, -1000, -13, 174,
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

	0, 16, 238, 7, 201, 234, 229, 9, 8, 228,
	225, 61, 11, 0, 223, 2, 222, 221, 219, 217,
	10, 5, 1, 4, 215, 14, 22, 13, 54, 12,
	15, 46, 6, 36, 31, 24, 214, 213, 212, 211,
	3, 210, 207, 181, 209,
}
var yyR1 = [...]int{

	0, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 2, 4, 4,
	4, 5, 8, 8, 6, 6, 7, 7, 11, 11,
	11, 13, 13, 13, 13, 13, 13, 1, 1, 19,
	19, 20, 20, 20, 14, 14, 14, 15, 15, 16,
	16, 16, 17, 17, 18, 18, 18, 18, 18, 21,
	23, 23, 24, 24, 25, 25, 26, 26, 27, 27,
	28, 28, 29, 29, 29, 30, 30, 30, 30, 30,
	31, 31, 31, 32, 32, 32, 33, 33, 33, 33,
	34, 36, 36, 36, 35, 35, 35, 35, 37, 37,
	37, 37, 37, 37, 38, 38, 38, 38, 38, 38,
	38, 22, 22, 40, 40, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 10, 42, 42, 9,
	9, 43, 43, 44, 44, 12, 12, 12, 39, 39,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 6, 7, 4, 4, 0, 2,
	4, 1, 1, 1, 1, 1, 1, 2, 3, 1,
	2, 1, 1, 1, 3, 4, 3, 1, 2, 5,
	7, 5, 6, 7, 3, 2, 2, 2, 3, 1,
	1, 5, 1, 3, 1, 3, 1, 3, 1, 3,
	1, 3, 1, 3, 3, 1, 3, 3, 3, 3,
	1, 3, 3, 1, 3, 3, 1, 3, 3, 3,
	1, 1, 1, 1, 1, 2, 2, 2, 1, 4,
	4, 3, 2, 2, 1, 1, 1, 1, 1, 1,
	3, 1, 3, 1, 3, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 7, 2, 4, 4,
	5, 1, 3, 1, 3, 1, 3, 4, 1, 3,
}
var yyChk = [...]int{

	-1000, -2, -4, -5, -6, -7, -8, 50, 48, 49,
	-9, -10, 46, 51, -4, 4, -12, 4, -12, 62,
	4, 4, 62, -3, 81, 73, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, -3, -43, -44, 4, 62, 45, -11, -12, 64,
	82, -3, 64, 63, 65, 84, -43, 66, 63, -3,
	74, -44, -21, -23, -24, -25, -26, -27, -28, -29,
	-30, -31, -32, -33, -34, -35, -37, 10, 11, -36,
	-38, 75, 76, 80, 4, 5, 7, 6, 8, 9,
	62, 63, -42, -12, -1, -3, 66, 65, 69, 19,
	18, 70, 71, 72, 16, 17, 73, 74, 14, 15,
	12, 13, 75, 76, 77, 78, 79, 81, 62, 83,
	10, 11, -35, -35, -34, -22, -40, -23, -35, 67,
	65, -3, -1, 67, -19, -20, -8, -7, -13, -14,
	-1, -15, -16, -17, -18, 4, 52, 53, 64, -22,
	54, 56, 57, 58, 59, 60, 61, -11, -22, -25,
	-26, -27, -28, -29, -30, -30, -31, -31, -31, -31,
	-32, -32, -33, -33, -34, -34, -34, -22, -39, -40,
	4, 63, 65, -41, 84, 20, 21, 22, 23, 24,
	25, 26, 27, 28, 29, 64, -12, 67, -20, 68,
	-21, 68, 64, 62, 62, 62, 4, 64, 64, 64,
	-22, 68, 82, 63, 65, -40, -40, -3, -13, 68,
	-13, -22, -22, -15, 64, 64, -23, -40, -13, 63,
	63, -15, -13, -13, 63, -22, 55, -13, 63, -13,
	-13,
}
var yyDef = [...]int{

	0, -2, 16, 18, 19, 20, 21, 0, 0, 0,
	22, 23, 0, 0, 17, 0, 0, 135, 0, 0,
	0, 0, 28, 0, 0, 0, 1, 2, 3, 4,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 0, 0, 131, 133, 0, 0, 0, 0, 26,
	136, 0, 27, 129, 0, 0, 0, 0, 0, 29,
	137, 132, 134, 59, 60, 62, 64, 66, 68, 70,
	72, 75, 80, 83, 86, 90, 94, 0, 0, 0,
	98, 91, 92, 93, 104, 105, 106, 107, 108, 109,
	0, 130, 0, 0, 24, 0, 0, 28, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	102, 103, 95, 96, 97, 0, 111, 113, 90, 0,
	0, 127, 25, 37, 0, 39, 41, 42, 43, 31,
	32, 33, 34, 35, 36, 104, 0, 0, 47, 0,
	0, 0, 0, 0, 0, 0, 0, 30, 0, 63,
	65, 67, 69, 71, 73, 74, 76, 77, 78, 79,
	81, 82, 84, 85, 87, 88, 89, 0, 0, 138,
	101, 110, 0, 0, 115, 116, 117, 118, 119, 120,
	121, 122, 123, 124, 125, 126, 0, 38, 40, 0,
	0, 0, 48, 0, 0, 0, 0, 55, 56, 57,
	0, 0, 99, 100, 0, 112, 114, 128, 44, 0,
	46, 0, 0, 0, 54, 58, 61, 139, 45, 0,
	0, 0, 49, 51, 0, 0, 0, 52, 0, 50,
	53,
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

	case 24:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line yacc.y:84
		{
			createFuntion(yyDollar[2].sval, yyDollar[6].sval)
		}
	case 37:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line yacc.y:112
		{
			yyVAL.sval = "hahahahahaha"
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line yacc.y:116
		{
			yyVAL.sval = "hahahahahaha"
		}
	}
	goto yystack /* stack new state and value */
}
