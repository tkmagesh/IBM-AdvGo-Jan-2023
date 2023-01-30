package main

import (
	"fmt"
)

func main() {
	var ch chan int
	ch = make(chan int)
	go divide(100, 0, ch)
	result := <-ch // blocked
	fmt.Println("Result :", result)
}

func divide(x, y int, ch chan int) {
	defer func() {
		if e := recover(); e != nil {
			ch <- 0
		}
	}()
	result := x / y
	ch <- result
	// time.Sleep(3 * time.Second)
	fmt.Println("Add operation completed")
}
