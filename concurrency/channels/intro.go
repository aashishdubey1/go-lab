package channels

import (
	"fmt"
	"sync"
)

func compute(ch chan int) {
	sum := 0
	for i := range 1000000 {
		sum += i
	}
	ch <- sum
}

func printId(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- id
}

type Job struct {
	Id int
}
type Result struct {
	job    Job
	status string
}

func worker(job chan Job, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range job {
		result <- Result{job: val, status: "done"}
	}
}

func Run() {

	var wg sync.WaitGroup

	jobCh := make(chan Job)
	resultCh := make(chan Result)

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go worker(jobCh, resultCh, &wg)
	}

	go func() {
		for i := range 6 {
			jobCh <- Job{Id: i}
		}
		close(jobCh)
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for v := range resultCh {
		fmt.Println(v)
	}

	// question 4. Close + range

	// var wg sync.WaitGroup
	// ch := make(chan int)
	// wg.Add(5)

	// for i := range 5 {
	// 	go printId(i,&wg,ch)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(ch)
	// }()

	// for v:= range ch {
	// 	fmt.Println(v)
	// }

	// Question 3 Multiple sends, one receiver

	// ch := make(chan int, 5)

	// go func(id int, ch chan int) {
	// 	ch <- id
	// }(1, ch)
	// go func(id int, ch chan int) {
	// 	ch <- id
	// }(2, ch)
	// go func(id int, ch chan int) {
	// 	ch <- id
	// }(3, ch)
	// go func(id int, ch chan int) {
	// 	ch <- id
	// }(4, ch)
	// go func(id int, ch chan int) {
	// 	ch <- id
	// }(5, ch)

	// v1 := <-ch
	// v2 := <-ch
	// v3 := <-ch
	// v4 := <-ch
	// v5 := <-ch

	// fmt.Println(v1, v2, v3, v4, v5)

}
