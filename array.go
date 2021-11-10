package main

import "fmt"

func Fruit() {
	// 固定長の配列(2つまで)
	var fruit [2]string
	fruit[0] = "banana"
	fruit[1] = "apple"

	// 3つ目を追加すると、エラーとなる
	// first argument to append must be slice; have [2]string
	// fruit = append(fruit, "orange")

	fmt.Println(fruit)
}
