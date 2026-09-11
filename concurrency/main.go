package main

import (
	"github.com/aashishdubey1/go-lab/concurrency/channels"
)

// func fetchData(source string,wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	time.Sleep(1 * time.Second)
// 	fmt.Println(source, "file downloaded")
// }

func main() {

	// var wg sync.WaitGroup
	// wg.Add(3)

	// start := time.Now()

	// go fetchData("server 1 ",&wg)
	// go fetchData("server 2 ",&wg)
	// fetchData("server 3 ",&wg)

	// wg.Wait()
	// fmt.Println("time taken =", time.Since(start))

	// // var wg sync.WaitGroup
	// for i := range 5 {
	// 	wg.Add(1)
	// 	go func(id int, wg *sync.WaitGroup) {
	// 		defer wg.Done()
	// 		time.Sleep(200*time.Millisecond)
	// 		fmt.Println(id)
	// 	}(i,&wg)
	// }

	// wg.Wait()

	// webchecker.Run()
	// counter.Run()
	// jobprocessor.Run()
	// channels.Run()
	channels.RunEx2()
}
