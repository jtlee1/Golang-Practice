package main

import (
	"fmt"
	"time"
)

func proc1() {
	fmt.Println("proc1 start")
	time.Sleep(1 * time.Second)
	fmt.Println("proc1 end")
}

func proc2() {
	fmt.Println("proc2 start")
	time.Sleep(2 * time.Second)
	fmt.Println("proc2 end")
}

func proc3() {
	fmt.Println("proc3 start")
	time.Sleep(3 * time.Second)
	fmt.Println("proc3 end")
}

func Add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	proc1()
	proc2()
	proc3()

}
