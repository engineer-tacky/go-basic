package main

import "fmt"

func main() {
	m := map[string]int{"a": 100, "b": 200}

	m["b"] = 10

	fmt.Println(m["a"], m["b"])
}