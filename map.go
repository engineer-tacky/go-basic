package main

import "fmt"

func printMap() {
	// Mapの初期化
	fruits := map[string]int{"banana": 100, "apple": 200}

	// Mapの中身を全て表示
	fmt.Println(fruits)

	// 特定のキーの値を表示
	fmt.Println(fruits["banana"])

	// 特定のキーの値を書き換える
	fruits["banana"] = 300
	fmt.Println(fruits)

	// 新しい要素の追加
	fruits["orange"] = 500
	fmt.Println(fruits)
}
