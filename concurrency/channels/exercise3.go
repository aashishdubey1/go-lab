package channels

import (
	"fmt"
	"math/rand"
	"sync"
)

// Fan-out → Fan-in with Cancellation
// Requirements
// 5 workers
// 100 jobs
// Workers process jobs concurrently.
// A worker randomly fails on some jobs.
// Once a fatal error occurs:
// producer should stop sending
// workers should stop taking new jobs
// main should eventually terminate
// No goroutine should be left hanging.
// No send on closed channel.
// No deadlock.

type Job3 struct {
	Id    int
	Value int
}

type Result3 struct {
	Job    Job3
	Output int
}

func worker3(job chan Job3, result chan Result3, done chan struct{}, wg *sync.WaitGroup, once *sync.Once) {
	defer wg.Done()
	for {
		select {
		case val, ok := <-job:
			if !ok {
				return
			}

			if rand.Intn(100) == val.Id {
				once.Do(func() {
					close(done)
				})
				return
			}

			output := Result3{Job: val, Output: val.Value * 10}

			select {
			case result <- output:
				fmt.Println("job", val.Id, "done")
			case <-done:
				fmt.Println("terminate")
				return
			}
		case <-done:
			fmt.Println("terminate")
			return
		}
	}
}

func RunEx3() {

	workerCount := 5
	jobCount := 100

	var once sync.Once

	jobCh := make(chan Job3)
	resultCh := make(chan Result3)
	done := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(workerCount)

	for range workerCount {
		go worker3(jobCh, resultCh, done, &wg, &once)
	}

	go func() {
		defer close(jobCh)
		for i := range jobCount {
			select {
			case jobCh <- Job3{Id: i, Value: i * 5}:
			case <-done:
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res.Output)
	}
}
