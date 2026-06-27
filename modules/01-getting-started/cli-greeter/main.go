package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	var lang string 
	
	flag.StringVar(&name, "name", "anon", "enter your name")
	flag.StringVar(&lang, "lang", "en", "language (en, fr, es)")


	flag.Parse()
	
	switch lang {
	case "en":
		fmt.Printf("Hello %s, How are you \n",name)
	case "fr":
		fmt.Printf("Bonjour %s, comment tu vas \n",name)
	case "ger":
		fmt.Printf("Hallo %s, wie geht es dir \n",name)
		default:
		fmt.Printf("Unknown language: %s\n", lang)
	}
	
}