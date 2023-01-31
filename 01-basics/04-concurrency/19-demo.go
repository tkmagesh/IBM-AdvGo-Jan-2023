package main

import (
	"fmt"
	"time"
)

func main() {
	fibCh := genFib()
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("[main function ends]")
}

func genFib() <-chan int {
	ch := make(chan int)
	stopCh := make(chan bool)

	go func() {
		fmt.Scanln()
		stopCh <- true
		fmt.Println("[Goroutine to accept stop signal - ends!]")
	}()

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
