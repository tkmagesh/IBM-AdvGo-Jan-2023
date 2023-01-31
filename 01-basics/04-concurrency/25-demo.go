package main

import (
	"fmt"
)

func main() {

	resultCh, _ := divide(100, 0)
	select {
	case result := <-resultCh:
		fmt.Println("Result :", result)
	}

}

func divide(x, y int) (resultCh chan int, errorCh chan error) {
	resultCh = make(chan int)
	errorCh = make(chan error, 1) // created as a buffered channel to make the receive operation "optional" on the channel
	go func() {
		defer func() {
			if e := recover(); e != nil {
				errorCh <- e.(error)
			}
		}()
		result := x / y
		resultCh <- result
	}()
	return
}
