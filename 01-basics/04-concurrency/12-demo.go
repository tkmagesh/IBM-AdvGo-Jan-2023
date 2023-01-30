package main

import (
	"fmt"
)

//consumer
func main() {
	ch := add(100, 200)
	result := <-ch // blocked
	fmt.Println("Result :", result)
}

//producer
func add(x, y int) <-chan int {
	var ch chan int
	ch = make(chan int)
	go func() {
		result := x + y
		ch <- result
	}()
	return ch
}
