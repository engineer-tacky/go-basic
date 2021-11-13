package main

import "fmt"

func forStatement() {
	// 基本的なループ
	for i := 0; i < 10; i++ {
		if i == 3 {
			// 処理をスキップ
			continue
		}
		if i > 5 {
			// 処理を終了
			break
		}
		fmt.Println(i)
	}

	// 省略して書くこともできる
	number := 0
	for number < 10 {
		number++
		fmt.Println(number)
	}

}
