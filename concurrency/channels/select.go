package channels

import (
	"fmt"
	"sync"
	// "time"
)

func RunSelect() {

	var once sync.Once

	once.Do(func() {
		fmt.Println("Hellodfhd")
	})

	once.Do(func() {
		fmt.Println("Hellodfd")
	})

	once.Do(func() {
		fmt.Println("Helrerelo")
	})

	// goroutine A → does some work repeatedly
	// main        → waits 2 seconds
	// main        → cancels A
	// A           → notices cancellation
	// A           → stops
	// main        → exits
	//
	// work := make(chan int)
	// done := make(chan struct{})
	//
	// go func() {
	// 	for {
	// 		select {
	// 		case val := <-work:
	// 			fmt.Println(val)
	// 		case <-done:
	// 			fmt.Println("Stopped")
	// 			return
	// 		}
	// 	}
	// }()
	//
	// for i := range 10 {
	// 	work <- i
	// }
	//
	// time.Sleep(2 * time.Millisecond)
	//
	// close(done)
	//
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func() {
	// 	ch1 <- 100
	// }()

	// go func() {
	// 	ch2 <- 200
	// }()
	//

	// time.Sleep(time.Millisecond)

	// select {
	// case v1 := <-ch1:
	// 	fmt.Println("ch1 value", v1)
	// default:
	// 	fmt.Println("Nothing happened")
	// }
	//
	// go func() {
	// 	val := <-ch1
	// 	fmt.Println(val)
	// }()
	// // time.Sleep(time.Millisecond)
	// select {
	// case ch1 <- 10:
	// 	fmt.Println("sent")
	// default:
	// 	fmt.Println("couldn't send")
	// }

}
