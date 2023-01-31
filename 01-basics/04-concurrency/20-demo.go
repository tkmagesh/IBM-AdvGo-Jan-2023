package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := getUserInterrupt()
	fibCh := genFib(stopCh)
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("[main function ends]")
}

func getUserInterrupt() <-chan struct{} {
	stopCh := make(chan struct{})
	go func() {
		fmt.Scanln()
		// stopCh <- true
		// stopCh <- struct{}{}
		close(stopCh)
		fmt.Println("[Goroutine to accept stop signal - ends!]")
	}()
	return stopCh
}

func genFib(stopCh <-chan struct{}) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case ch <- x:
				time.Sleep(500 * time.Millisecond)
				x, y = y, x+y
			case <-stopCh:
				break LOOP
			}
		}
		close(ch)
		fmt.Println("[Goroutine to generate fibnocci series - ends!]")
	}()
	return ch
}
