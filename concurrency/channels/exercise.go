package channels

import (
	"errors"
	"fmt"
	"sync"
)

// Exercise 1
// Build a worker pool that processes jobs and returns either a result or an error
// Requirements:
// 3 workers
// 10 jobs
// Each worker calculates: -> Value * Value
// If Value is negative, return an error instead.

type Jobb struct {
	Id    int
	Value int
}
type Resultt struct {
	Job    Jobb
	Output int
	Error  error
}

func workerr(jobs chan Jobb, result chan Resultt, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range jobs {
		output := value.Value * value.Value
		if value.Value < 0 {
			result <- Resultt{Job: value, Output: output, Error: errors.New("Negative number error")}
			continue
		}
		result <- Resultt{Job: value, Output: output, Error: nil}
	}
}

func RunExercise() {

	workerCount := 3
	jobCount := 10

	jobCh := make(chan Jobb)
	resultCh := make(chan Resultt)

	var wg sync.WaitGroup
	wg.Add(workerCount)

	// Adding workers
	go func() {
		for range workerCount {
			go workerr(jobCh, resultCh, &wg)
		}
	}()

	// Adding jobs
	go func() {
		for i := range jobCount {
			if i == 6 {
				jobCh <- Jobb{Id: i, Value: i * 10 * -1}
				continue
			}
			jobCh <- Jobb{Id: i, Value: i * 10}
		}
		close(jobCh)
	}()

	// Waiting to close resultCh
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}

}
