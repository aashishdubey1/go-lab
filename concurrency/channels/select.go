package channels

import (
	"fmt"
)

func RunSelect() {

	ch1 := make(chan int)
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

	go func() {
		val := <-ch1
		fmt.Println(val)
	}()
	// time.Sleep(time.Millisecond)
	select {
	case ch1 <- 10:
		fmt.Println("sent")
	default:
		fmt.Println("couldn't send")
	}

}
