package main

import (
	"fmt"
	"sync"
)

func withWaitGroup(num int) {
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go procWithWait(i, wg)
	}
	wg.Wait()
}

func procWithWait(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("my val ", num)
	//time.Sleep(1 * time.Second)
}

func withoutWaitGroup(num int) {
	for i := 0; i < num; i++ {
		go procWithoutWait(i)
	}
}

func procWithoutWait(num int) {
	fmt.Println("my val ", num)
	//time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("withWaitGroup start")
	withWaitGroup(10)
	fmt.Println("withWaitGroup end")
	//time.Sleep(1 * time.Second)
	fmt.Println("withoutWaitGroup start")
	withoutWaitGroup(10)
	fmt.Println("withoutWaitGroup end")
}
