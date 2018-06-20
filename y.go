//line fateyacc.y:1
package main

import __yyfmt__ "fmt"

//line fateyacc.y:3
//line fateyacc.y:5
type yySymType struct {
	yys                 int
	exprNode            ExprNode
	exprNode_list       []ExprNode
	unary_operator      byte
	assignment_operator string
	stmtNode            StmtNode
	blockItem           BlockItem
	variable_definition Variable
	type_specifier      int
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

//line fateyacc.y:570

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 488

var yyAct = [...]int{

	137, 148, 140, 126, 127, 62, 72, 23, 5, 135,
	70, 48, 67, 65, 47, 139, 69, 55, 68, 50,
	16, 18, 66, 181, 128, 60, 41, 114, 115, 116,
	74, 112, 113, 51, 103, 73, 120, 121, 108, 109,
	211, 102, 101, 181, 99, 71, 210, 218, 200, 198,
	130, 237, 129, 181, 97, 229, 59, 181, 223, 228,
	63, 181, 224, 181, 207, 212, 95, 213, 206, 93,
	201, 181, 194, 180, 94, 181, 91, 96, 54, 58,
	75, 57, 53, 46, 54, 52, 49, 235, 118, 204,
	203, 202, 125, 45, 98, 22, 100, 106, 107, 205,
	157, 131, 122, 123, 75, 136, 17, 117, 179, 119,
	124, 132, 156, 158, 160, 163, 164, 169, 170, 176,
	162, 161, 178, 159, 75, 75, 75, 75, 75, 75,
	75, 75, 75, 75, 75, 75, 75, 75, 75, 75,
	75, 75, 195, 136, 197, 173, 174, 175, 171, 172,
	63, 199, 165, 166, 167, 168, 20, 209, 184, 185,
	186, 187, 188, 189, 190, 191, 192, 193, 43, 12,
	75, 8, 9, 7, 13, 42, 104, 105, 110, 111,
	44, 21, 15, 92, 134, 214, 215, 2, 11, 14,
	10, 6, 4, 3, 1, 143, 142, 141, 138, 217,
	182, 219, 177, 216, 220, 221, 79, 222, 76, 19,
	64, 80, 0, 0, 0, 225, 0, 226, 0, 227,
	0, 56, 183, 61, 0, 230, 0, 0, 0, 231,
	232, 0, 234, 0, 236, 75, 238, 0, 239, 85,
	88, 89, 87, 86, 144, 77, 78, 26, 27, 28,
	29, 30, 31, 32, 33, 34, 35, 36, 37, 38,
	39, 40, 0, 0, 85, 88, 89, 87, 86, 144,
	77, 78, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 8, 9, 0, 0, 145, 146, 149,
	25, 150, 151, 152, 153, 154, 155, 90, 24, 147,
	0, 96, 196, 0, 0, 0, 0, 0, 8, 9,
	81, 82, 145, 146, 149, 83, 150, 151, 152, 153,
	154, 155, 90, 0, 147, 0, 96, 133, 0, 0,
	0, 0, 0, 0, 0, 81, 82, 0, 0, 0,
	83, 85, 88, 89, 87, 86, 144, 77, 78, 85,
	88, 89, 87, 86, 84, 77, 78, 0, 85, 88,
	89, 87, 86, 84, 77, 78, 85, 88, 89, 87,
	86, 84, 77, 78, 85, 88, 89, 87, 86, 84,
	77, 78, 0, 0, 0, 0, 0, 0, 0, 145,
	146, 149, 0, 150, 151, 152, 153, 154, 155, 90,
	0, 147, 0, 96, 0, 0, 0, 90, 233, 0,
	0, 0, 81, 82, 0, 0, 90, 83, 147, 0,
	81, 82, 0, 0, 90, 83, 208, 0, 0, 81,
	82, 0, 90, 0, 83, 0, 0, 81, 82, 0,
	0, 0, 83, 0, 0, 81, 82, 0, 0, 0,
	83, 26, 27, 28, 29, 30, 31, 32, 33, 34,
	35, 36, 37, 38, 39, 40, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
	40, 0, 0, 0, 0, 0, 0, 96,
}
var yyPact = [...]int{

	123, 123, -1000, -1000, -1000, -1000, -1000, 173, 97, 97,
	-1000, -1000, 147, 172, -1000, 33, 217, -1000, 217, 171,
	31, 38, 97, 22, -63, 436, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 21, 19, -1000, -67, 171, 15, 16, 217, -1000,
	-1000, -49, -1000, -1000, 171, 370, 13, 97, 421, -11,
	-1000, -1000, -1000, -1000, 25, 78, -28, -30, -38, 160,
	24, 166, -44, -50, -1000, -1000, 26, 370, 370, 370,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	370, -1000, -15, 217, -1000, 11, 260, 97, 370, 370,
	370, 370, 370, 370, 370, 370, 370, 370, 370, 370,
	370, 370, 370, 370, 370, 370, 370, 370, 370, 99,
	-1000, -1000, -1000, -1000, -1000, 10, -1000, -1000, 138, 8,
	97, -1000, -1000, -1000, 235, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -19, 370, -20, -1000, 6, 29,
	28, 27, 90, 4, 0, 362, -1000, -22, 78, -28,
	-30, -38, 160, 24, 24, 166, 166, 166, 166, -44,
	-44, -50, -50, -1000, -1000, -1000, -42, 2, -1000, -1000,
	-1000, 370, 370, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 217, -1000, -1000, 337, -21,
	337, -1000, 370, 370, 354, -6, -1000, -1000, -1000, -2,
	370, -1000, -1000, 370, -1000, -1000, -1000, -1000, 337, -1000,
	-4, -8, 354, -1000, -1000, -1000, -1000, -1000, 337, 337,
	345, 32, -1000, 337, -12, 337, -1000, 337, -1000, -1000,
}
var yyPgo = [...]int{

	0, 211, 1, 3, 5, 4, 210, 13, 22, 12,
	18, 16, 10, 45, 6, 30, 35, 24, 208, 11,
	206, 202, 200, 0, 198, 15, 2, 197, 196, 195,
	9, 7, 8, 194, 187, 193, 192, 191, 190, 188,
	188, 14, 184, 183, 175, 168,
}
var yyR1 = [...]int{

	0, 31, 31, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 33, 33, 34, 34,
	34, 35, 37, 37, 40, 40, 36, 36, 32, 32,
	41, 41, 41, 23, 23, 23, 23, 23, 23, 25,
	25, 42, 42, 30, 30, 24, 24, 24, 26, 26,
	27, 27, 27, 28, 28, 29, 29, 29, 29, 29,
	4, 5, 5, 6, 6, 7, 7, 8, 8, 9,
	9, 10, 10, 11, 11, 11, 12, 12, 12, 12,
	12, 13, 13, 13, 14, 14, 14, 16, 16, 16,
	16, 15, 20, 20, 20, 17, 17, 17, 17, 18,
	18, 18, 18, 18, 18, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 3, 3, 22, 22, 22, 22,
	22, 22, 22, 22, 22, 22, 22, 39, 43, 43,
	38, 38, 44, 44, 45, 45, 19, 19, 19, 21,
	21,
}
var yyR2 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 2, 1, 1,
	1, 1, 1, 1, 7, 8, 6, 7, 4, 4,
	0, 2, 4, 1, 1, 1, 1, 1, 1, 2,
	3, 1, 2, 1, 1, 3, 4, 3, 1, 2,
	5, 7, 5, 6, 7, 3, 2, 2, 2, 3,
	1, 1, 5, 1, 3, 1, 3, 1, 3, 1,
	3, 1, 3, 1, 3, 3, 1, 3, 3, 3,
	3, 1, 3, 3, 1, 3, 3, 1, 3, 3,
	3, 1, 1, 1, 1, 1, 2, 2, 2, 1,
	4, 4, 3, 2, 2, 1, 1, 1, 1, 1,
	1, 3, 1, 3, 1, 3, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 7, 2, 4,
	4, 5, 1, 3, 1, 3, 1, 3, 4, 1,
	3,
}
var yyChk = [...]int{

	-1000, -33, -34, -35, -36, -32, -37, 50, 48, 49,
	-38, -39, 46, 51, -34, 9, -19, 9, -19, 62,
	9, 9, 62, -31, 81, 73, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, -31, -44, -45, 9, 62, 45, -41, -19, 64,
	82, -31, 64, 63, 65, 84, -44, 66, 63, -31,
	74, -45, -4, -5, -6, -7, -8, -9, -10, -11,
	-12, -13, -14, -16, -15, -17, -18, 10, 11, -20,
	-1, 75, 76, 80, 9, 4, 8, 7, 5, 6,
	62, 63, -43, -19, -25, -31, 66, 65, 69, 19,
	18, 70, 71, 72, 16, 17, 73, 74, 14, 15,
	12, 13, 75, 76, 77, 78, 79, 81, 62, 83,
	10, 11, -17, -17, -15, -2, -3, -5, -17, 67,
	65, -31, -25, 67, -42, -30, -32, -23, -24, -25,
	-26, -27, -28, -29, 9, 52, 53, 64, -2, 54,
	56, 57, 58, 59, 60, 61, -41, -2, -7, -8,
	-9, -10, -11, -12, -12, -13, -13, -13, -13, -14,
	-14, -16, -16, -15, -15, -15, -2, -21, -3, 9,
	63, 65, -22, 84, 20, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 64, -19, 67, -30, 68, -4,
	68, 64, 62, 62, 62, 9, 64, 64, 64, -2,
	68, 82, 63, 65, -3, -3, -31, -23, 68, -23,
	-2, -2, -26, 64, 64, -5, -3, -23, 63, 63,
	-26, -23, -23, 63, -2, 55, -23, 63, -23, -23,
}
var yyDef = [...]int{

	0, -2, 16, 18, 19, 20, 21, 0, 0, 0,
	22, 23, 0, 0, 17, 0, 0, 136, 0, 0,
	0, 0, 30, 0, 0, 0, 1, 2, 3, 4,
	5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	15, 0, 0, 132, 134, 0, 0, 0, 0, 28,
	137, 0, 29, 130, 0, 0, 0, 0, 0, 31,
	138, 133, 135, 60, 61, 63, 65, 67, 69, 71,
	73, 76, 81, 84, 87, 91, 95, 0, 0, 0,
	99, 92, 93, 94, 105, 106, 107, 108, 109, 110,
	0, 131, 0, 0, 26, 0, 0, 30, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	103, 104, 96, 97, 98, 0, 112, 114, 91, 0,
	0, 128, 27, 39, 0, 41, 43, 44, 33, 34,
	35, 36, 37, 38, 105, 0, 0, 48, 0, 0,
	0, 0, 0, 0, 0, 0, 32, 0, 64, 66,
	68, 70, 72, 74, 75, 77, 78, 79, 80, 82,
	83, 85, 86, 88, 89, 90, 0, 0, 139, 102,
	111, 0, 0, 116, 117, 118, 119, 120, 121, 122,
	123, 124, 125, 126, 127, 0, 40, 42, 0, 0,
	0, 49, 0, 0, 0, 0, 56, 57, 58, 0,
	0, 100, 101, 0, 113, 115, 129, 45, 0, 47,
	0, 0, 0, 55, 59, 62, 140, 46, 0, 0,
	0, 50, 52, 0, 0, 0, 53, 0, 51, 54,
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

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:53
		{
			yyVAL.type_specifier = INT8
		}
	case 2:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:57
		{
			yyVAL.type_specifier = INT16
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:61
		{
			yyVAL.type_specifier = INT32
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:65
		{
			yyVAL.type_specifier = INT64
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:69
		{
			yyVAL.type_specifier = UINT8
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:73
		{
			yyVAL.type_specifier = UINT16
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:77
		{
			yyVAL.type_specifier = UINT32
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:81
		{
			yyVAL.type_specifier = UINT64
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:85
		{
			yyVAL.type_specifier = FLOAT32
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:89
		{
			yyVAL.type_specifier = FLOAT64
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:93
		{
			yyVAL.type_specifier = BOOL
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:97
		{
			yyVAL.type_specifier = CHAR
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:101
		{
			yyVAL.type_specifier = STRING
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:105
		{
			yyVAL.type_specifier = TYPE_CONST
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:109
		{
			yyVAL.type_specifier = TYPE_STRUCT
		}
	case 28:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line fateyacc.y:146
		{
			yyVAL.variable_definition = createDefinedVariable(LET, yyDollar[2].exprNode, yyDollar[3].type_specifier)
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line fateyacc.y:150
		{
			yyVAL.variable_definition = createDefinedVariable(SET, yyDollar[2].exprNode, yyDollar[3].type_specifier)
		}
	case 39:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:172
		{
			yyVAL.stmtNode = aa()
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:176
		{
			yyVAL.stmtNode = aa()
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:188
		{
			yyVAL.blockItem = createBlockItemFormVar(yyDollar[1].variable_definition)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:192
		{
			yyVAL.blockItem = createBlockItemFormStmt(yyDollar[1].stmtNode)
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:200
		{
			yyVAL.stmtNode = aa()
		}
	case 46:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line fateyacc.y:204
		{
			yyVAL.stmtNode = aa()
		}
	case 47:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:208
		{
			yyVAL.stmtNode = aa()
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:215
		{
			yyVAL.stmtNode = StmtNode{}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:219
		{
			yyVAL.stmtNode = createExprStmtNode(yyDollar[1].exprNode)
		}
	case 50:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line fateyacc.y:226
		{
			yyVAL.stmtNode = createIfNode(yyDollar[3].exprNode, yyDollar[5].stmtNode, (StmtNode{}))
		}
	case 51:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line fateyacc.y:230
		{
			yyVAL.stmtNode = createIfNode(yyDollar[3].exprNode, yyDollar[5].stmtNode, yyDollar[7].stmtNode)
		}
	case 52:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line fateyacc.y:234
		{
			yyVAL.stmtNode = aa()
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line fateyacc.y:241
		{
			yyVAL.stmtNode = createForNode(yyDollar[3].stmtNode, yyDollar[4].stmtNode, (ExprNode{}), yyDollar[6].stmtNode)
		}
	case 54:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line fateyacc.y:245
		{
			yyVAL.stmtNode = createForNode(yyDollar[3].stmtNode, yyDollar[4].stmtNode, yyDollar[5].exprNode, yyDollar[7].stmtNode)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:252
		{
			yyVAL.stmtNode = aa()
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:256
		{
			yyVAL.stmtNode = aa()
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:260
		{
			yyVAL.stmtNode = aa()
		}
	case 58:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:264
		{
			yyVAL.stmtNode = aa()
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:268
		{
			yyVAL.stmtNode = aa()
		}
	case 62:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line fateyacc.y:280
		{
			yyVAL.exprNode = createCondExprNode(yyDollar[1].exprNode, yyDollar[3].exprNode, yyDollar[5].exprNode)
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:288
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "||", yyDollar[3].exprNode)
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:296
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "&&", yyDollar[3].exprNode)
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:304
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "|", yyDollar[3].exprNode)
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:312
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "^", yyDollar[3].exprNode)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:320
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "&", yyDollar[3].exprNode)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:328
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "==", yyDollar[3].exprNode)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:332
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "!=", yyDollar[3].exprNode)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:340
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "<", yyDollar[3].exprNode)
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:344
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, ">", yyDollar[3].exprNode)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:348
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "<=", yyDollar[3].exprNode)
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:352
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, ">=", yyDollar[3].exprNode)
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:360
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "<<", yyDollar[3].exprNode)
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:364
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, ">>", yyDollar[3].exprNode)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:372
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "+", yyDollar[3].exprNode)
		}
	case 86:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:376
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "-", yyDollar[3].exprNode)
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:384
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "*", yyDollar[3].exprNode)
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:388
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "/", yyDollar[3].exprNode)
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:392
		{
			yyVAL.exprNode = ceateBinaryOpNode(yyDollar[1].exprNode, "%", yyDollar[3].exprNode)
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:403
		{
			yyVAL.unary_operator = '+'
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:407
		{
			yyVAL.unary_operator = '-'
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:411
		{
			yyVAL.unary_operator = '!'
		}
	case 96:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:419
		{
			yyVAL.exprNode = createPrefixOpNode(INC_OP, yyDollar[2].exprNode)
		}
	case 97:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:423
		{
			yyVAL.exprNode = createPrefixOpNode(DEC_OP, yyDollar[2].exprNode)
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:427
		{
			yyVAL.exprNode = createUnaryOpNode(yyDollar[1].unary_operator, yyDollar[2].exprNode)
		}
	case 100:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line fateyacc.y:435
		{
			yyVAL.exprNode = createArefNode(yyDollar[1].exprNode, yyDollar[3].exprNode)
		}
	case 101:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line fateyacc.y:439
		{
			yyVAL.exprNode = createFuncallNode(yyDollar[1].exprNode, yyDollar[3].exprNode_list)
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:443
		{
			yyVAL.exprNode = createMemberNode(yyDollar[1].exprNode, yyDollar[3].exprNode)
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:447
		{
			yyVAL.exprNode = createSuffixOpNode(INC_OP, yyDollar[1].exprNode)
		}
	case 104:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line fateyacc.y:451
		{
			yyVAL.exprNode = createSuffixOpNode(INC_OP, yyDollar[1].exprNode)
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:464
		{
			yyVAL.exprNode = yyDollar[2].exprNode
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:477
		{
			yyVAL.exprNode = createAssignNode(yyDollar[1].exprNode, yyDollar[2].assignment_operator, yyDollar[3].exprNode)
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:484
		{
			yyVAL.assignment_operator = "="
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:488
		{
			yyVAL.assignment_operator = "*="
		}
	case 118:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:492
		{
			yyVAL.assignment_operator = "/="
		}
	case 119:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:496
		{
			yyVAL.assignment_operator = "%="
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:500
		{
			yyVAL.assignment_operator = "+="
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:504
		{
			yyVAL.assignment_operator = "-="
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:508
		{
			yyVAL.assignment_operator = "<<="
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:512
		{
			yyVAL.assignment_operator = ">>="
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:516
		{
			yyVAL.assignment_operator = "&="
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:520
		{
			yyVAL.assignment_operator = "^="
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:524
		{
			yyVAL.assignment_operator = "|="
		}
	case 139:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line fateyacc.y:561
		{
			yyVAL.exprNode_list = bb()
		}
	case 140:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line fateyacc.y:565
		{
			yyVAL.exprNode_list = bb()
		}
	}
	goto yystack /* stack new state and value */
}
