package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var count int
	wg := &sync.WaitGroup{}
	rand.Seed(20)
	flag.IntVar(&count, "count", 0, "Number of goroutines")
	flag.Parse()
	fmt.Printf("Starting %d goroutines.. Hit ENTER to start...", count)
	fmt.Scanln()
	for i := 1; i <= count; i++ {
		wg.Add(1)
		go fn(i, wg)
	}
	wg.Wait() // blocked until the counter becomes 0
	fmt.Println("Hit ENTER to shutdown")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
	wg.Done() // decrement the counter by 1
}
