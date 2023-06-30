package main

import "fmt"

type MyStruct struct {
	a string
	b, c int
}

func main() {
	var st MyStruct
	st.a = "Hello"
	st.b = 50
	st.c = 100
	fmt.Println(st)
}