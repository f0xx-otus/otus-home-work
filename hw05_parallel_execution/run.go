package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

var errCount int32
var taskChannel chan Task

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, n int, m int) error {
	maxErrCount := uint(m)
	goroutineCount := uint(n)
	taskChannel = make(chan Task)
	errCount = 0
	wg := &sync.WaitGroup{}
	wg.Add(n + 1)
	for i := uint(0); i < goroutineCount; i++ {
		go func() {
			defer wg.Done()
			worker(taskChannel)
		}()
	}

	go func() {
		defer wg.Done()
		publisher(tasks, maxErrCount)
	}()
	wg.Wait()
	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func publisher(tasks []Task, setErrCount uint) {
	for _, t := range tasks {
		if errCount == int32(setErrCount) {
			close(taskChannel)
			return
		}
		if errCount < int32(setErrCount) {
			taskChannel <- t
		}
	}
	close(taskChannel)
}

func worker(taskChannel chan Task) {
	for {
		task, ok := <-taskChannel
		if !ok {
			return
		}
		if task() != nil {
			atomic.AddInt32(&errCount, 1)
		}
	}
}
