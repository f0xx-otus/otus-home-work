package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
var taskChannel chan Task
var mutex sync.Mutex

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, n int, m int) error {
	var errCount int32
	maxErrCount := uint(m)
	goroutineCount := uint(n)
	taskChannel = make(chan Task)
	wg := &sync.WaitGroup{}
	wg.Add(n + 1)
	go func() {
		defer wg.Done()
		publisher(tasks, maxErrCount, &errCount)
	}()
	for i := uint(0); i < goroutineCount; i++ {
		go func() {
			defer wg.Done()
			worker(taskChannel, &errCount)
		}()
	}
	wg.Wait()
	if errCount >= int32(m) {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func publisher(tasks []Task, setErrCount uint, errCount *int32) {
	for _, t := range tasks {
		mutex.Lock()
		if *errCount == int32(setErrCount) {
			mutex.Unlock()
			close(taskChannel)
			return
		}
		mutex.Unlock()
		taskChannel <- t
	}
	close(taskChannel)
}

func worker(taskChannel chan Task, errCount *int32) {
	for {
		task, ok := <-taskChannel
		if !ok {
			return
		}
		if task() != nil {
			mutex.Lock()
			atomic.AddInt32(errCount, 1)
			mutex.Unlock()
		}
	}
}
