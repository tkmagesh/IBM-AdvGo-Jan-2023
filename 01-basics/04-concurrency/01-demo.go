package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //schedule the func to be executed through the scheduler
	f2()

	// DO NOT DO THE FOLLOWING
	// time.Sleep(time.Second) // block the execution of main fn for 1 sec and there by give the opportunity for the scheduler to go and look for the goroutines scheduled and execute them
	fmt.Scanln() // blocked for some I/O
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
