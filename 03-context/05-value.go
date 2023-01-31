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
	valCtx := context.WithValue(rootCtx, "root-key", "root-val")

	cancelCtx, cancel := context.WithCancel(valCtx)
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
	fmt.Println("[generateNos] root-key =", ctx.Value("root-key"))
	wg.Add(1)
	// cancelCtx2, _ := context.WithCancel(ctx)
	// go printNos(cancelCtx2, wg)

	genValCtx := context.WithValue(ctx, "gen-key", "gen-val")

	timeoutCtx, cancel2 := context.WithTimeout(genValCtx, 10*time.Second)
	defer func() {
		fmt.Println("cancelling printNos")
		cancel2()
	}()
	go printNos(timeoutCtx, wg)

	wg.Add(1)
	cancelCtx3, cancel3 := context.WithCancel(genValCtx)
	defer func() {
		fmt.Println("cancelling printFib")
		cancel3()
	}()
	go printFib(cancelCtx3, wg)

	select {
	case <-ctx.Done():
		fmt.Println("generateNos completed")
	}
}

func printNos(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("[printNos] root-key =", ctx.Value("root-key"))
	fmt.Println("[printNos] gen-key =", ctx.Value("gen-key"))
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
	fmt.Println("[printFib] root-key =", ctx.Value("root-key"))
	fmt.Println("[printFib] gen-key =", ctx.Value("gen-key"))
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
