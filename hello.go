package main

import "fmt"

const Name = "nekopunch"

const (
	Country = "Japan"
	Birth   = 2021
)

func main() {
	// 再代入は、エラーになります。
	// cannot assign to Name (declared const)
	// Name = "aaaaa"
	fmt.Println(Name, Country, Birth)
}
