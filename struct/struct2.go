package main

import "fmt"

type MyStruct struct {
	a string
	b, c int
}

type MyStruct2 struct {
	MyStruct
	d int
}

func main() {
	var st2 MyStruct2
	st2.a = "Hello"
	st2.d = 150
	fmt.Println(st2.a, st2.b, st2.c, st2.d)
}