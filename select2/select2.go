package main

import (
	"fmt"
	"time"
)

func listener(ch chan int) {
	for {
		fmt.Println(<-ch, "from listener")
	}
}

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Second)
		ch <- i
	}
	close(ch)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go listener(ch1)
	go sender(ch2)
	i := 0
	for {
		time.Sleep(1 * time.Second)
		select {
		case a, ok := <-ch2:
			if !ok {
				return
			}
			fmt.Println("get ", a, " from sender")
		default:
		}
		ch1 <- i
		i++
	}
}
