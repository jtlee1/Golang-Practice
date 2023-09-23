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
	for i := 0; i < 5; i = i + 1 {
		time.Sleep(1 * time.Second)
		v := <-c
		fmt.Println(v)
	}
}

func main() {
	c := make(chan int)
	go listener(c)
	sender(c)

}
