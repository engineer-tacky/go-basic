package main

import (
	"log"
)

func logStatement() {
	// 基本的なログ出力
	log.Println("ログ出力")
	log.Printf("%T %v", "ログ出力", "ログ出力")

	// FatallnやFatalfで出力すると、それ以降は処理されない
	log.Fatalln("エラーログ出力")
	log.Fatalf("%T %v", "エラーログ出力", "エラーログ出力")
}
