package main

import "fmt"

func printArray() {
	// 固定長の配列(2つまで)
	var fruits [2]string
	fruits[0] = "banana"
	fruits[1] = "apple"

	// 3つ目を追加すると、エラーとなる
	// firsts argument to append must be slice; have [2]string
	// fruits = append(fruits, "orange")

	fmt.Println(fruits)

	// 可変長の配列
	var money []int = []int{1000, 2000}
	money = append(money, 3000)

	fmt.Println(money)
}
