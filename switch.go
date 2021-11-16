package main

import (
	"fmt"
	"time"
)

func switchStatement() {
	// 基本的な記述
	fruits := "apple"
	switch fruits {
	case "apple":
		fmt.Println("apple!")
	case "banana":
		fmt.Println("banana!")
	default:
		fmt.Println("default!")
	}

	// caseに条件を記述
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("午前")
	case t.Hour() > 12:
		fmt.Println("午後")
	}

}
