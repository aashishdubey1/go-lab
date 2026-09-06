package jobprocessor

import (
	"fmt"
	"sync"
	"time"
)

type Job struct {
	Id int
}

func Run() {
	jobs := []Job{
		{1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10},
		{11}, {12}, {13}, {14}, {15}, {16}, {17}, {18}, {19}, {20},
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	currentIndex := 0

	for workerId := 1; workerId <= 5; workerId++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				mu.Lock()
				if currentIndex > len(jobs) {
					fmt.Println("no jobs availble")
					mu.Unlock()
					return
				}

				job := jobs[currentIndex]
				currentIndex++
				mu.Unlock()

				fmt.Printf("worker %d processing job %d\n", id, job.Id)
				time.Sleep(200 * time.Millisecond)
			}
		}(workerId)
	}
	wg.Wait()
}
