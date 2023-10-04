package main

import (
	"fmt"
	"time"
)

// use close with range
func example1() {
	queue := make(chan string)
	go func(queue chan string) {
		queue <- "one"
		queue <- "two"
		queue <- "one"
		queue <- "two"
		close(queue)
	}(queue)

	for elem := range queue {
		time.Sleep(1 * time.Second)
		fmt.Println(elem)
	}
}

// cite https://gobyexample.com/closing-channels
// double channel for communication (done channel stops main from finishing)
func example2() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}

func main() {
	example2()
}
