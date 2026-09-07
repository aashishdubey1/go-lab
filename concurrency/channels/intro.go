package channels

import (
	"sync"
)

func compute(ch chan int) {
	sum := 0
	for i := range 1000000 {
		sum += i
	}
	ch <- sum
}

func printId(id int, wg *sync.WaitGroup,ch chan int) { 
	defer wg.Done()
	ch <- id
}

func Run() {

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
