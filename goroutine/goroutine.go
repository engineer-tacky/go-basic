package main

import (
	"fmt"
	"time"
)

var n = 0

func main() {
	fmt.Println("main start")
	fmt.Println(&n)

	go sub()
	time.Sleep(time.Millisecond * 3)
}

func sub() {
	fmt.Println("sub start")
	fmt.Println(&n)

	time.Sleep(time.Millisecond * 10)

	// mainが先に終了するのでこれは実行されない
	fmt.Println("sub end")
}