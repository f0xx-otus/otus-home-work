package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var taskChannel chan Task
var errCount int32
var taskMu sync.Mutex
var cond *sync.Cond
var wg sync.WaitGroup

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error {
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < N; i++ {
			fmt.Println(i)
		}
	}()
	go func() {
		defer wg.Done()
		fmt.Println("pub")
	}()
	wg.Wait()
	fmt.Println(errCount)
	if errCount > int32(M) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func publisher(tasks []Task, setErrCount int) {
	if errCount == int32(setErrCount) {
		close(taskChannel)
	}
	for _, t := range tasks {
		if errCount < int32(setErrCount) {
			taskMu.Lock()
			taskChannel <- t
			taskMu.Unlock()
			cond.Broadcast()
		}
	}
}

func worker() {
	taskMu.Lock()
	task, ok := <-taskChannel
	for !ok {
		cond.Wait()
	}
	result := task()
	if result != nil {
		return
	}
	taskMu.Unlock()
	atomic.AddInt32(&errCount, 1)
}
