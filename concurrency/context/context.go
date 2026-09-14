package context

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopped")
			return
		case <-time.After(time.Second):
			fmt.Println("done something")
		}
	}

}

func RunContext() {

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)

	go worker(ctx, &wg)

	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()
}
