package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string)
	go func(queue chan string) {
		queue <- "one"
		queue <- "two"
		queue <- "one"
		queue <- "two"
		close(queue)
	}(queue)

	for elem := range queue {
		time.Sleep(2 * time.Second)
		fmt.Println(elem)
	}
}
