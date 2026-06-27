package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var lang string 
	
	flag.StringVar(&name,"name","anon","enter your name")
	flag.StringVar(&lang,"lang","en","language")
	flag.Parse()
	
	if lang == "fr"{
		fmt.Printf("Bonjour %s, comment tu vas \n",name)
	} else if lang == "ger" {
		fmt.Printf("Hallo %s, wie geht es dir \n",name)
	}else {
		fmt.Printf("Hello %s, How are you \n",name)
	}
}