package main

import "fmt"

func main() {
	ch := make(chan int)

	go send(ch)

	c := <-ch
	fmt.Printf("[値を受信: %d]\n", c)
}

func send(ch chan int) {
	fmt.Println("[値を送信: 1]")
	ch <- 1
}