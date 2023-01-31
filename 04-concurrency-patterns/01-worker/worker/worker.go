/* fill in the blanks */
package worker

import (
	"fmt"
	"sync"
)

type Work interface {
	Task()
}

type Worker struct {
	workQueue chan Work
	wg        sync.WaitGroup
}

func (w *Worker) Add(work Work) {
	w.workQueue <- work
}

func (w *Worker) Shutdown() {
	close(w.workQueue)
	w.wg.Wait()
}

func New(count int) *Worker {
	worker := &Worker{
		workQueue: make(chan Work),
	}
	worker.wg.Add(count)
	for i := 1; i <= count; i++ {
		go func(id int) {
			defer worker.wg.Done()
			fmt.Printf("Worker %d started...\n", id)
			for wk := range worker.workQueue {
				wk.Task()
			}
			fmt.Printf("Worker %d shutdown...\n", id)
		}(i)
	}
	return worker
}
