package main

import (
	"fmt"
	"sync"
	"time"
)

func loop(c chan int) {
	start := time.Now()

	for i := 0; i < 10; i = i + 1 {
		go func(v int) {
			time.Sleep(1 * time.Second)
			c <- v
		}(i)
	}

	for i := 0; i < 10; i = i + 1 {
		v := <-c
		fmt.Printf("%v", v)
	}
	fmt.Println()
	elapsed := time.Since(start)
	fmt.Println("go routine with channel", elapsed)

}

func loop2() {
	start := time.Now()
	var slice []int
	wg := new(sync.WaitGroup)
	wg.Add(10)

	for i := 0; i < 10; i = i + 1 {
		go func(v int) {
			time.Sleep(1 * time.Second)
			slice = append(slice, v)
			//fmt.Println(v)
			wg.Done()
		}(i)
	}
	fmt.Println(slice)
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Println("go routine with wait", elapsed)
}

func loop3() {
	start := time.Now()
	var slice []int

	for i := 0; i < 10; i = i + 1 {
		func(v int) {
			time.Sleep(1 * time.Second)
			slice = append(slice, v)
		}(i)
	}
	fmt.Println(slice)

	elapsed := time.Since(start)
	fmt.Println("normal loop: ", elapsed)
}

func main() {
	c := make(chan int)
	loop(c)
	loop2()
	loop3()

}
