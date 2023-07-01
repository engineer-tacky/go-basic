package main

import "fmt"

func main() {
	// ポインタ変数 ゼロ値で初期化
	var p *int = new(int)

	fmt.Println(p)
	fmt.Println(*p)
}