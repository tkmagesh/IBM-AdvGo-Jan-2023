package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		data, isOpen := <-ch
		if !isOpen {
			break
		}
		fmt.Println(data)
	}
}

func fn(ch chan int) {
	for i := 1; i <= 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
