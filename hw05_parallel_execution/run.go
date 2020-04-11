package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func executeTask(wg *sync.WaitGroup, taskCh <-chan Task, resultCh chan<- error, quitCh <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case task := <-taskCh:
			{
				resultCh <- task()
			}
		case <-quitCh:
			return
		}
	}
}

func resultListener(errorLimit int, tasksNum int, resultCh <-chan error, errorCh chan<- bool) {
	errorCounter := 0
	doneCounter := 0

	// Receive results
	for {
		select {
		case res := <-resultCh:
			if res != nil {
				errorCounter++
			}
			if res == nil {
				doneCounter++
			}
			if errorCounter >= errorLimit {
				errorCh <- true
				return
			} else if doneCounter == tasksNum {
				errorCh <- false
				return
			}
		default:
			continue
		}
	}
}

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	if len(tasks) == 0 || N == 0 {
		return nil
	}

	wg := &sync.WaitGroup{}
	taskCh := make(chan Task, len(tasks))
	resultCh := make(chan error, len(tasks))
	quitCh := make(chan struct{})
	errorCh := make(chan bool)

	// Start N workers
	for i := 0; i < N; i++ {
		wg.Add(1) //nolint:gomnd
		go executeTask(wg, taskCh, resultCh, quitCh)
	}

	// Submit tasks to workers
	for _, val := range tasks {
		taskCh <- val
	}

	// Start listener
	go resultListener(M, len(tasks), resultCh, errorCh)

	// Receive result
	isErrorExceed := <-errorCh

	// Stop all workers
	for i := 0; i < N; i++ {
		quitCh <- struct{}{}
	}

	// Wait until all goroutines done
	wg.Wait()

	// Check if we get error exceed
	if isErrorExceed {
		return ErrErrorsLimitExceeded
	}

	return nil
}
