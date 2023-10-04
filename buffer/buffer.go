package main

import (
	"fmt"
	"time"
)

func sender(c chan int) {
	//start := time.Now()

	for i := 0; i < 5; i = i + 1 {
		c <- i
		fmt.Println("send: ", i)
	}

}

func listener(c chan int) {
	//time.Sleep(3 * time.Second)
	for i := 0; i < 6; i = i + 1 {
		time.Sleep(1 * time.Second)
		v := <-c
		fmt.Println(v)
	}
}

// compare buffered channel with unbuffered channel
func main() {
	//c := make(chan int)
	c := make(chan int, 3)
	c <- 11
	c <- 11
	go sender(c)
	listener(c)
	c <- 11
	//sender(c)

}
