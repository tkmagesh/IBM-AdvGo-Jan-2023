package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rootCtx := context.Background()
	cancelCtx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	go func() {
		fmt.Println("Hit ENTER to stop...")
		fmt.Scanln()
		cancel()
	}()

	wg.Add(1)
	go generateNos(cancelCtx, wg)
	wg.Wait()
}

func generateNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	go printNos(ctx, wg)

	wg.Add(1)
	go printFib(ctx, wg)

	select {
	case <-ctx.Done():
		fmt.Println("generateNos completed")
	}
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
LOOP:
	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("printNos completed")
			break LOOP
		default:
			fmt.Println("no :", i*10)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func printFib(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	x, y := 0, 1
LOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("printFib completed")
			break LOOP
		default:
			fmt.Println("fib No :", x)
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
	}
}
