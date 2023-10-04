package main

import (
	"fmt"
	"time"
)

func proc1(ch chan<- string) {
	for {
		time.Sleep(1 * time.Second)
		ch <- "proc1"
	}
}

func proc2(ch chan<- string) {
	for {
		time.Sleep(2 * time.Second)
		ch <- "proc2"
	}
}

func proc3(ch chan<- string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "proc3"
	}
}

func Done(ch chan<- string, i time.Duration) {
	time.Sleep(i * time.Second)
	ch <- "Done"
}

// use main as listener
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)
	go proc1(ch1)
	go proc2(ch2)
	go proc3(ch3)
	go Done(ch4, 8)
	//use loop because for loop is an outer loop
Loop:
	for {
		select {
		case out := <-ch1:
			fmt.Println(out)
		case out := <-ch2:
			fmt.Println(out)
		case out := <-ch3:
			fmt.Println(out)
		case out := <-ch4:
			fmt.Println(out)
			break Loop
		}
	}
}
