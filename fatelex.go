package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var fileLength int = 0
var currentChar int = 0
var tokenList []Token
var fileByte []byte
var currerntToken int = -1

func (tk *Token) Lex(lval *yySymType) int {
	currerntToken += 1
	if currerntToken >= len(tokenList) {
		return 0
	}
	if tokenList[currerntToken].tokenType == IDENTIFIER {
		lval.sval = tokenList[currerntToken].tokenName
	}
	//fmt.Println(currerntToken, "************************chenyao", tokenList[currerntToken].tokenType, tokenList[currerntToken].tokenName)
	//fmt.Println(currerntToken+1, "************************chenyao", tokenList[currerntToken+1].tokenType, tokenList[currerntToken+1].tokenName)
	return tokenList[currerntToken].tokenType
}

func (tk *Token) Error(s string) {
	log.Fatal(fmt.Errorf("Lex: %s", s))
}

func getLex(fileByte []byte) {
	var token Token
	for {
		if currentChar >= fileLength {
			break
		}
		c := fileByte[currentChar]
		currentChar += 1
		switch {
		case isSpace(int(c)):
		case isComment(int(c)):
			skipComment()
			//fmt.Println("This is a space or comment chenyao, and the value is ", string(c))
		case isNewline(int(c)):
			skipNewline()
			//fmt.Println("This is a new line chenyao")
		case isAlpha(int(c)):
			getWotd(c)
			//fmt.Println("This is a Letter or _ chenyao", string(c))
		case isAmbiguous(int(c)):
			ambigousSymbol(c)
		case isDigit(int(c)):
			getNumber(c)
		case isDoubleQuotation(int(c)):
			getLiteral(c)
		case isSingleQuotation(int(c)):
			getLiteral(c)
		default:
			token.tokenName = string(c)
			fmt.Println("<DEFAULT>", token.tokenName, "\t", token.tokenName)
			token.tokenType = int(c)
			tokenList = append(tokenList, token)
		}
		//fmt.Println(string(c))
	}
	// reg := regexp.MustCompile(`\d+\.\d+|\w+|\d+`)
	// file := reg.FindAll(fileByte, -1)
	// for _, v := range file {
	// 	println(string(v))
	// }
}

//读取文件
func readFile2Byte(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read file fail", err)
		return nil
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return nil
	}

	return fd
}

// 跳过注释
func skipComment() {
	for {
		c := fileByte[currentChar]
		currentChar += 1
		if isNewline(int(c)) || isEof(int(c)) {
			currentChar -= 1
			return
		}
	}
}

// 跳过空行
func skipNewline() {
	for {
		currentChar += 1
		if currentChar >= fileLength {
			return
		}
		c := fileByte[currentChar]
		if !isNewline(int(c)) {
			currentChar -= 1
			return
		}
	}
}

//获取关键字和标识符
func getWotd(c_char byte) {
	tokenName := []byte{c_char}
	for {
		c := fileByte[currentChar]
		currentChar += 1
		if !isAlnum(int(c)) && !isAlpha(int(c)) {
			currentChar -= 1
			var token Token
			token.tokenName = string(tokenName[:])
			if isKeyWord(&token) {
				fmt.Println("'", token.tokenName, token.tokenType, "'\t", token.tokenName)
				tokenList = append(tokenList, token)
				break
			} else {
				fmt.Println("<IDENTIFIER>\t", token.tokenName)
				token.tokenType = IDENTIFIER
				tokenList = append(tokenList, token)
				break
			}
		} else {
			tokenName = append(tokenName, c)
		}
	}
}

//获取操作符
func ambigousSymbol(c_char byte) {
	var token Token
	token.tokenName = string(c_char)
	n_char := fileByte[currentChar]
	currentChar += 1
	switch c_char {
	case '<':
		switch n_char {
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<LE_OP>\t", token.tokenName)
			token.tokenType = LE_OP
			tokenList = append(tokenList, token)
			return
		case '<':
			token.tokenName += string(n_char)
			currentChar += 1
			n2_char := fileByte[currentChar]
			if n2_char == '=' {
				token.tokenName += string(n2_char)
				fmt.Println("<LEFT_ASSIGN>\t", token.tokenName)
				token.tokenType = LEFT_ASSIGN
				tokenList = append(tokenList, token)
				return
			} else {
				currentChar -= 1
				fmt.Println("<LEFT_OP>\t", token.tokenName)
				token.tokenType = LEFT_OP
				tokenList = append(tokenList, token)
				return
			}
		}
	case '>':
		switch n_char {
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<GE_OP>\t", token.tokenName)
			token.tokenType = GE_OP
			tokenList = append(tokenList, token)
			return
		case '>':
			token.tokenName += string(n_char)
			currentChar += 1
			n2_char := fileByte[currentChar]
			if n2_char == '=' {
				token.tokenName += string(n2_char)
				fmt.Println("<RIGHT_ASSIGN>\t", token.tokenName)
				token.tokenType = RIGHT_ASSIGN
				tokenList = append(tokenList, token)
				return
			} else {
				currentChar -= 1
				fmt.Println("<RIGHT_OP>\t", token.tokenName)
				token.tokenType = RIGHT_OP
				tokenList = append(tokenList, token)
				return
			}
		}
	case '=':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<EQ_OP>\t", token.tokenName)
			token.tokenType = EQ_OP
			tokenList = append(tokenList, token)
			return
		}
	case '+':
		switch n_char {
		case '+':
			token.tokenName += string(n_char)
			fmt.Println("<INC_OP>\t", token.tokenName)
			token.tokenType = INC_OP
			tokenList = append(tokenList, token)
			return
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<ADD_ASSIGN>\t", token.tokenName)
			token.tokenType = ADD_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '-':
		switch n_char {
		case '-':
			token.tokenName += string(n_char)
			fmt.Println("<DEC_OP>\t", token.tokenName)
			token.tokenType = DEC_OP
			tokenList = append(tokenList, token)
			return
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<SUB_ASSIGN>\t", token.tokenName)
			token.tokenType = SUB_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '*':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<MUL_ASSIGN>\t", token.tokenName)
			token.tokenType = MUL_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '/':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<DIV_ASSIGN>\t", token.tokenName)
			token.tokenType = DIV_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '%':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<MOD_ASSIGN>\t", token.tokenName)
			token.tokenType = MOD_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '^':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<XOR_ASSIGN>\t", token.tokenName)
			token.tokenType = XOR_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '!':
		if n_char == '=' {
			token.tokenName += string(n_char)
			fmt.Println("<NE_OP>\t", token.tokenName)
			token.tokenType = NE_OP
			tokenList = append(tokenList, token)
			return
		}
	case '&':
		switch n_char {
		case '&':
			token.tokenName += string(n_char)
			fmt.Println("<AND_OP>\t", token.tokenName)
			token.tokenType = AND_OP
			tokenList = append(tokenList, token)
			return
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<AND_ASSIGN>\t", token.tokenName)
			token.tokenType = AND_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	case '|':
		switch n_char {
		case '|':
			token.tokenName += string(n_char)
			fmt.Println("<OR_OP>\t", token.tokenName)
			token.tokenType = OR_OP
			tokenList = append(tokenList, token)
			return
		case '=':
			token.tokenName += string(n_char)
			fmt.Println("<OR_ASSIGN>\t", token.tokenName)
			token.tokenType = OR_ASSIGN
			tokenList = append(tokenList, token)
			return
		}
	}
	currentChar -= 1
	fmt.Println("<SIGLE_OP>", token.tokenName, "\t", token.tokenName)
	token.tokenType = int(c_char)
	tokenList = append(tokenList, token)
}

//获取数字
func getNumber(c_char byte) {
	var token Token
	token.tokenType = INT_CONSTANT
	token.tokenName += string(c_char)
	for {
		n_char := fileByte[currentChar]
		currentChar += 1
		if n_char == '.' && token.tokenType == INT_CONSTANT {
			token.tokenName += string(n_char)
			token.tokenType = FLOAT_CONSTANT
		} else if isDigit(int(n_char)) {
			token.tokenName += string(n_char)
		} else {
			currentChar -= 1
			fmt.Println("<NUMBER>", token.tokenName, "\t", token.tokenName)
			tokenList = append(tokenList, token)
			break
		}
	}
}

//获取字符和字符串
func getLiteral(c_char byte) {
	s := []byte{'`', byte(c_char)}
	for {
		n_char := fileByte[currentChar]
		currentChar += 1
		switch {
		case isEof(int(n_char)), isNewline(int(n_char)):
			fmt.Println("Fate: illegal literal constant format")
			return
		case c_char == n_char:
			s = append(s, []byte{byte(n_char), '`'}...)
			var token Token
			if v, err := strconv.Unquote(string(s)); err != nil {
				fmt.Println("Fate: errorerroreorror")
				return
			} else {
				token.tokenName = string(v[1 : len(v)-1])
			}
			//fmt.Println("****************", string(c_char))
			if c_char == '\'' {
				fmt.Println("<CHAR>", token.tokenName, "\t", token.tokenName)
				token.tokenType = CHAR_CONSTANT
				tokenList = append(tokenList, token)
				return
			} else if c_char == '"' {
				fmt.Println("<STRING>", token.tokenName, "\t", token.tokenName)
				token.tokenType = STRING_CONSTANT
				tokenList = append(tokenList, token)
				return
			}
		default:
			s = append(s, n_char)
		}
	}
}

//判断是否关键字
func isKeyWord(token *Token) bool {
	for _, v := range KeyWords {
		if token.tokenName == v.KeyWordsName {
			token.tokenType = v.KeyWordsType
			return true
		}
	}
	return false
}
