package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go source("data1.dat", dataCh, wg)
	wg.Add(1)
	go source("data2.dat", dataCh, wg)

	oddCh, evenCh := splitter(dataCh)
	oddSumCh := sum(oddCh)
	evenSumCh := sum(evenCh)
	doneCh := merger("result.txt", oddSumCh, evenSumCh)

	wg.Wait()
	close(dataCh)

	<-doneCh
	fmt.Println("Done")
}

func source(fileName string, dataCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		if no, err := strconv.Atoi(txt); err == nil {
			dataCh <- no
		}
	}
}

func splitter(dataCh chan int) (<-chan int, <-chan int) {
	oddCh := make(chan int)
	evenCh := make(chan int)

	go func() {
		defer close(oddCh)
		defer close(evenCh)
		for data := range dataCh {
			if data%2 == 0 {
				evenCh <- data
			} else {
				oddCh <- data
			}
		}
	}()

	return oddCh, evenCh
}

func sum(ch <-chan int) <-chan int {
	resultCh := make(chan int)
	go func() {
		total := 0
		for no := range ch {
			total += no
		}
		resultCh <- total
	}()
	return resultCh
}

func merger(fileName string, oddSumCh, evenSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		for i := 0; i < 2; i++ {
			select {
			case oddSum := <-oddSumCh:
				fmt.Fprintf(file, "Odd Total : %d\n", oddSum)
			case evenSum := <-evenSumCh:
				fmt.Fprintf(file, "Even Total : %d\n", evenSum)
			}
		}
		close(doneCh)
	}()
	return doneCh
}
