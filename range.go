package main

import "fmt"

func rangeStatement() {
	fruits := []string{"apple", "banana", "orange"}

	// indexとvalueのどちらも出力する場合
	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// indexだけ出力する場合
	for index := range fruits {
		fmt.Println(index)
	}

	// valueだけ出力する場合
	for _, value := range fruits {
		fmt.Println(value)
	}

	drinks := map[string]int{"cola": 100, "soda": 200, "tea": 300}

	// indexとvalueのどちらも出力する場合
	for key, value := range drinks {
		fmt.Println(key, value)
	}

	// keyだけ出力する場合
	for key := range drinks {
		fmt.Println(key)
	}

	// valueだけ出力する場合
	for _, value := range drinks {
		fmt.Println(value)
	}
}
