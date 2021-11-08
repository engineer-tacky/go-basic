package main

import "fmt"

func init() {
	fmt.Println("Init")
}

func Buzz() {
	fmt.Println("Bazz")
}

func main() {
	Buzz()
	fmt.Println("Hello world!")
}
