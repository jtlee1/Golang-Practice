package main

import (
	"fmt"
	"time"
)

func proc(i int) {
	fmt.Println("my val is: ", i)
	time.Sleep(1 * time.Second)
}

// shows goroutine speeds up process
func main() {
	//with go
	start := time.Now()
	fmt.Println("proc with go routine start")
	for i := 0; i < 10; i++ {
		go proc(i)
	}
	elapsed := time.Since(start)
	fmt.Println("proc with go routine took ", elapsed)

	//without go
	start = time.Now()
	fmt.Println("proc without go routine start")
	for i := 0; i < 10; i++ {
		proc(i)
	}
	elapsed = time.Since(start)
	fmt.Println("proc without go routine took ", elapsed)
}
