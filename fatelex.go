package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var fileLength int = 0
var currentChar int = 0
var tokenList []Token
var fileByte []byte

func main() {
	fileByte = readFile2Byte("test.fate")
	fileLength = len(fileByte)
	Lex(fileByte)
	//fmt.Println(string(fileByte))
}

func Lex(fileByte []byte) {
	for {
		if currentChar == fileLength {
			break
		}
		c := fileByte[currentChar]
		currentChar += 1
		switch {
		case isSpace(int(c)):
		case isComment(int(c)):
			skipComment()
			fmt.Println("This is a space or comment chenyao, and the value is ", string(c))
		case isNewline(int(c)):
			skipNewline()
			fmt.Println("This is a new line chenyao")
		case isAlpha(int(c)):
			getToken(c)
			fmt.Println("This is a Letter or _ chenyao", string(c))
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
		currentChar += 1
		c := fileByte[currentChar]
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
		if currentChar == fileLength {
			return
		}
		c := fileByte[currentChar]
		if !isNewline(int(c)) {
			currentChar -= 1
			return
		}
	}
}

//获取token
func getToken(c_char byte) {
	tokenName := []byte{c_char}
	for {
		c := fileByte[currentChar]
		currentChar += 1
		if !isAlnum(int(c)) && !isAlpha(int(c)) {
			currentChar -= 1
			var token Token
			token.tokenName = string(tokenName[:])
			if isKeyWord(token.tokenName) {
				fmt.Println("'", token.tokenName, "'\t", token.tokenName)
				break
			} else {
				fmt.Println("<IDENTIFIER>\t", token.tokenName)
				break
			}
		} else {
			tokenName = append(tokenName, c)
		}
	}
}

//获取数字

//获取字符

//获取字符串

//判断是否关键字
func isKeyWord(tokenName string) bool {
	for _, v := range KeyWords {
		if tokenName == v {
			return true
		}
	}
	return false
}
