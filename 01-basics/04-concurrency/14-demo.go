package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go fn(ch, wg)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	wg.Wait()
}

func fn(ch chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
	wg.Done()
}
