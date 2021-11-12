package main

import "fmt"

func ifStatement() {
	// 基本構文
	result := "OK"
	if result == "OK" {
		fmt.Println("OKです！")
	} else if result == "NG" {
		fmt.Println("NGです...")
	} else {
		fmt.Println("OKでもNGでもないです")
	}

	// && と ||
	max, min := 10, 1
	if max == 10 && min == 1 {
		fmt.Println("両方がtrueの時")
	} else if max == 10 || min == 1 {
		fmt.Println("どちらかがtrueの時")
	} else {
		fmt.Println("どちらもfalseの時")
	}
}
