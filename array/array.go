package main

import "fmt"

func main() {
	var a [3]int
	b := [2]string{"banana", "apple"}

	// 初期値を指定すると、...で要素数を省略できる
	c := [...]int{1, 2}

	// インデックスを指定して初期化
	d := [3]int{1: 10, 2: 20}

	fmt.Println(a[0])
	fmt.Println(b[1])
	fmt.Println(d[0], d[1], d[2])

	// 配列の要素数を取得
	fmt.Println(len(a))
	fmt.Println(len(c))
}