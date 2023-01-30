package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) add(n int) {
	c.Lock()
	{
		c.count += n
	}
	c.Unlock()
}

var counter Counter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

func increment(wg *sync.WaitGroup) {
	counter.add(1)
	wg.Done()
}
