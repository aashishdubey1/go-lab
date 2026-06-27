package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("No Arguments provided")
		return
	}
	for i := len(args)-1 ; i>=0; i--{
		fmt.Println(args[i])
	}
}