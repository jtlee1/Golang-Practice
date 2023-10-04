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
	i := 0
	for {
		time.Sleep(100 * time.Millisecond)
		ch <- i
		i++
	}
}

func example1() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go listener(ch1)
	go sender(ch2)
	fmt.Println("both event will occur")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		select {
		case a := <-ch2:
			fmt.Println("get ", a, " from sender")
		default:
		}
		ch1 <- i
	}

	time.Sleep(1 * time.Second)
}

func example2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go listener(ch1)
	go sender(ch2)
	fmt.Println("only one event will occur")
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		select {
		case a := <-ch2:
			fmt.Println("get ", a, " from sender")
		case ch1 <- i:
		}
	}
	time.Sleep(1 * time.Second)
}

func main() {
	example2()

}
