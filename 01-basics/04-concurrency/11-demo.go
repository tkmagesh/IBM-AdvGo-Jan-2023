package main

import (
	"fmt"
)

func main() {
	var ch chan int
	ch = make(chan int)
	go add(100, 200, ch)
	result := <-ch // blocked
	fmt.Println("Result :", result)
}

func add(x, y int, ch chan int) {
	result := x + y
	ch <- result
	// time.Sleep(3 * time.Second)
	fmt.Println("Add operation completed")
}
