package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// The worker should:
// print "working" every ~500ms
// stop automatically after 2 seconds
// print "timeout" when it stops
// use context.WithTimeout
// use select
// use WaitGroup
// don't use time.Sleep for the worker loop
//

func worker1(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("working")
		}
	}
}

func RunTimeoutEx() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg.Add(1)
	go worker1(ctx, &wg)

	wg.Wait()
}
