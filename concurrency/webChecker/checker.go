package webchecker

import (
	"fmt"
	"net/http"
	"sync"
)

func check(url string,wg *sync.WaitGroup) { 
	defer wg.Done()
	res,err := http.Get(url)
	if err != nil { 
	fmt.Println(url,"down")
		return
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		fmt.Println(url,"up")
	} else {
		fmt.Println(url,"down")
	}
}

func Run() { 
	var wg sync.WaitGroup

	urls := []string{
    	"https://google.com",
    	"https://github.com",
    	"https://youtube.com",
    	"https://exale.com",
	}
	for _,url := range urls {
		wg.Add(1)
		go check(url,&wg)
	}
	wg.Wait()
}