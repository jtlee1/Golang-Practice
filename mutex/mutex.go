package main

import (
	"fmt"
	"sync"
)

func Race(j int) {
	out := 0
	w := new(sync.WaitGroup)
	w.Add(j)
	//10
	for i := 0; i < j; i++ {
		go func() {
			defer w.Done()
			out++
		}()
	}
	w.Wait()
	fmt.Println(out)
}

func Safe(j int) {
	out := 0
	mu := new(sync.Mutex)
	//10
	w := new(sync.WaitGroup)
	w.Add(j)
	for i := 0; i < j; i++ {
		go func() {
			defer w.Done()
			mu.Lock()
			out++
			mu.Unlock()
		}()
	}
	w.Wait()
	fmt.Println(out)
}

func main() {
	fmt.Println("Race: ")
	Race(100)
	fmt.Println("Mutex: ")
	Safe(100)

}
