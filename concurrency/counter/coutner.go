package counter

import (
	"fmt"
	"sync"
)

func increse(counter *int,wg *sync.WaitGroup,mu *sync.Mutex) { 
	defer wg.Done()
	for i:=0; i<1000; i++ { 
		mu.Lock()
		*counter = *counter + i 
		mu.Unlock()
	}
}

func Run() { 
	var wg sync.WaitGroup 
	var mu sync.Mutex

	counter := 0
	for i:= 0; i<100; i++{ 
		wg.Add(1)
		go increse(&counter,&wg,&mu)
	}
	wg.Wait()
	fmt.Println(counter)
}
