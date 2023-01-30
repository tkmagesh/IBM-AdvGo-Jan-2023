package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	for data := range ch {
		fmt.Println(data)
		time.Sleep(500 * time.Millisecond)
	}
}

func fn(ch chan int) {
	for i := 1; i <= 10; i++ {
		// time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	close(ch)
}
