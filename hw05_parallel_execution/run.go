package hw05_parallel_execution //nolint:golint,stylecheck

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in N goroutines and stops its work when receiving M errors from tasks
func Run(tasks []Task, N int, M int) error { //nolint:gocritic
	if len(tasks) == 0 || N <= 0 {
		return nil
	}

	wg := &sync.WaitGroup{}
	taskCh := make(chan Task, len(tasks))
	resultCh := make(chan error, len(tasks))
	quitCh := make(chan struct{})

	// Start N workers
	for i := 0; i < N; i++ {
		wg.Add(1) //nolint:gomnd
		go executeTask(wg, taskCh, resultCh, quitCh)
	}

	// Submit tasks to workers
	for _, val := range tasks {
		taskCh <- val
	}
	var (
		errorCounter int
		doneCounter  int
		isError      bool
	)

	for res := range resultCh {
		if res != nil {
			errorCounter++
		} else {
			doneCounter++
		}
		if errorCounter >= N {
			// Prevent overrunning
			isError = true
			break
		} else if doneCounter == len(tasks) {
			break
		}
	}

	close(taskCh)
	close(quitCh)
	wg.Wait()

	if isError {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func executeTask(wg *sync.WaitGroup, taskCh <-chan Task, resultCh chan<- error, quitCh <-chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-quitCh:
			return
		case task, ok := <-taskCh:
			{
				if ok {
					resultCh <- task()
				}
			}
		}
	}
}
