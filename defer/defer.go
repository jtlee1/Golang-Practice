package main

import (
	"fmt"
)

func main() {
	fmt.Println("function start")
	defer fmt.Println("defer start")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	defer fmt.Println("defer end")
	fmt.Println("function end")
}
