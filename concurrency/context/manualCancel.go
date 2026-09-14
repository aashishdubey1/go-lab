package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Start 3 workers.
// Each worker continuously does some work.
// After 2 seconds:
// All three workers should stop.

func worker2(workerId int, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println(workerId, "worker doin something")
		}
	}
}

func RunManualCancellation() {

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go worker2(i, &wg, ctx)
	}

	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()
}
