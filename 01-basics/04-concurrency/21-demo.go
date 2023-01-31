package main

import (
	"fmt"
	"time"
)

func main() {
	stopCh := time.After(10 * time.Second)
	fibCh := genFib(stopCh)
	for no := range fibCh {
		fmt.Println(no)
	}
	fmt.Println("[main function ends]")
}

func genFib(stopCh <-chan time.Time) <-chan int {
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
