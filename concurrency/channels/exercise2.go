package channels

import (
	"fmt"
	"sync"
)

type Job2 struct {
	Id    int
	Value int
}

type Proccess struct {
	Job    Job2
	Output int
}

type Result2 struct {
	Job    Proccess
	Output int
}

func workerStage1(job chan Job2, proccess chan Proccess, wg *sync.WaitGroup) {

	defer wg.Done()

	for val := range job {
		output := Proccess{Job: val, Output: val.Value * 2}
		proccess <- output
	}
}

func workerStgage2(process chan Proccess, result chan Result2, wg *sync.WaitGroup) {

	defer wg.Done()

	for val := range process {
		output := Result2{Job: val, Output: val.Output + 10}
		result <- output
	}

}

func RunEx2() {

	jobCh := make(chan Job2)
	resultCh := make(chan Result2)
	processedCh := make(chan Proccess)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup

	wg1.Add(3)
	wg2.Add(2)

	for range 3 {
		go workerStage1(jobCh, processedCh, &wg1)
	}

	go func() {
		for i := range 10 {
			jobCh <- Job2{Id: i, Value: i * 10}
		}
		close(jobCh)
	}()

	go func() {
		wg1.Wait()
		close(processedCh)
	}()

	for range 2 {
		go workerStgage2(processedCh, resultCh, &wg2)
	}

	go func() {
		wg2.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}

}
