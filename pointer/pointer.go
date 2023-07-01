package main

import "fmt"

func main() {
	// ポインタ変数
	var p *int

	n := 10

	// nのアドレスをpに代入
	p = &n

	fmt.Println(p)
	fmt.Println(*p)
}