package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 10
	fmt.Println("len(ch) :", len(ch))
	ch <- 20
	fmt.Println("len(ch) :", len(ch))
	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))
	fmt.Println(<-ch)
	fmt.Println("len(ch) :", len(ch))
}
