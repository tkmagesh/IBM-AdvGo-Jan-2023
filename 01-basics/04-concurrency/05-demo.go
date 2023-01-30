package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()
	// count++
	wg.Done()
}
