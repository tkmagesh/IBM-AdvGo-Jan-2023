package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()

	go func() {
		time.Sleep(5 * time.Second)
		ch2 <- 200
	}()

	resultCh := func() chan int {
		resultCh := make(chan int)
		wg := sync.WaitGroup{}
		go func() {
			wg.Add(1)
			go func() {
				resultCh <- <-ch1
				wg.Done()
			}()

			wg.Add(1)
			go func() {
				resultCh <- <-ch2
				wg.Done()
			}()
			wg.Wait()
			close(resultCh)
		}()
		return resultCh
	}()

	for data := range resultCh {
		fmt.Println(data)
	}
}
