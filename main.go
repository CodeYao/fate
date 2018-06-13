package main

func main() {
	var token *Token
	fileByte = readFile2Byte("test.fate")
	fileLength = len(fileByte)
	getLex(fileByte)

	yyParse(token)
	//fmt.Println(string(fileByte))

}
