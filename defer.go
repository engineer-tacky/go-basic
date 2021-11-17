package main

import (
	"fmt"
)

func deferStatement() {
	defer fmt.Println("こんにちは")
	defer fmt.Println("こんばんは")
	defer fmt.Println("おやすみ")

	fmt.Println("おはよう")
}
