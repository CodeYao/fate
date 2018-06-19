package main

//type TokenType int
type BasicType int

// const (
// 	NULL                TokenType = iota //空字符
// 	EXCEPTION                            //异常字符
// 	NUMBER_TOKEN                         //数字标识符
// 	CHAR_TOKEN                           //字符标识符
// 	CHAR_SIGN_TOKEN                      //单引号'标识符
// 	STRING_TOKEN                         //字符串标识符
// 	STRING_SIGN_TOKEN                    //双引号"标识符
// 	BOOL_TOKEN                           //布尔标识符
// 	IF_TOKEN                             //if标识符
// 	ELSE_TOKEN                           //else标识符
// 	FOR_TOKEN                            //for标识符
// 	FUNC_TOKEN                           //func标识符
// 	PARAM_TOKEN                          //变量标识符
// 	STATE_TOKEN                          //声明变量标识符
// 	STATE_TYPE_TOKEN                     //声明类型标识符let,set
// 	TOKEN_TYPE_TOKEN                     //变量类型标识符int,float...
// 	IDENTIFIER_TOKEN                     //变量
// 	ADD_OPERATOR_TOKEN                   //加法+
// 	SUB_OPERATOR_TOKEN                   //减法-
// 	MUL_OPERATOR_TOKEN                   //乘法*
// 	DIV_OPERATOR_TOKEN                   //除法/
// 	MOD_OPERATOR_TOKEN                   //取余%
// 	ASS_OPERATOR_TOKEN                   //赋值=
// 	LEFT_PAREN_TOKEN                     //左括号(
// 	RIGHT_PAREN_TOKEN                    //右括号)
// 	LEFT_BRACKET_TOKEN                   //左中括号[
// 	RIGHT_BRACKET_TOKEN                  //右中括号]
// 	LEFT_BRACES_TOKEN                    //左大括号{
// 	RIGHT_BRACES_TOKEN                   //右大括号}
// 	END_OF_LINE_TOKEN                    //行结束符;
// 	EQ_TOKEN                             // ==
// 	NE_TOKEN                             // !=
// 	GT_TOKEN                             // >
// 	GE_TOKEN                             // >=
// 	LT_TOKEN                             // <
// 	LE_TOKEN                             // <=
// 	LOGICAL_AND_TOKEN                    // &&
// 	LOGICAL_OR_TOKEN                     // ||

// )

const (
	INT8_T BasicType = iota
	INT16_T
	INT32_T
	INT64_T
	UINT8_T
	UINT16_T
	UINT32_T
	UINT64_T
	FLOAT32_T
	FLOAT64_T
	BOOL_T
	STRING_T
	CHAR_T
	NULL_T
	ERRORTYPE_T
)

type KeyWordsToken struct {
	KeyWordsName string
	KeyWordsType int
}

var KeyWords = []KeyWordsToken{{"int8", INT8}, {"int16", INT16}, {"int32", INT32}, {"int64", INT64}, {"uint8", UINT8}, {"uint16", UINT16}, {"uint32", UINT32}, {"uint64", UINT64}, {"bool", BOOL}, {"float32", FLOAT32}, {"float64", FLOAT64}, {"string", STRING}, {"char", CHAR}, {"if", IF}, {"else", ELSE}, {"for", FOR}, {"switch", SWITCH}, {"case", CASE}, {"break", BREAK}, {"return", RETURN}, {"let", LET}, {"set", SET}, {"func", FUNC}, {"false", BOOL_LITERAL}, {"true", BOOL_LITERAL}}

type Token struct {
	tokenName string
	tokenType int
}
