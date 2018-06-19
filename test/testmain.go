package main

import "fmt"

type test1 struct {
	a int
	d test2
}
type test2 struct {
	c int
	e *test1
}

func main() {
	var test test2
	test.c = 5
	test.e = &test1{}
	test.e.a = 5
	fmt.Println("Hello ChenYao!", test, test.e.d)
}
