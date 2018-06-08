package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fileByte := readFile2Byte("fatetest.fate")
	fmt.Println(string(fileByte))
}

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
