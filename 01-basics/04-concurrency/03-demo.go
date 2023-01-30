package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		var wg sync.WaitGroup
		wg.Add(1) // increment the counter by 1
		go f1(&wg)
	*/
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(wg)
	}
	f2()
	wg.Wait() // blocked until the counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrement the counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
