package main

import (
	"fmt"
)

func main() {

	resultCh, errorCh := divide(100, 0)
	select {
	case result := <-resultCh:
		fmt.Println("Result :", result)
	case e := <-errorCh:
		fmt.Println("something went wrong..", e)
	}
}

func divide(x, y int) (resultCh chan int, errorCh chan error) {
	resultCh = make(chan int)
	errorCh = make(chan error)
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
