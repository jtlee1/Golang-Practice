package main

import (
	"fmt"
)

func main() {
	foo := []string{"a", "b"}
	fmt.Println("foo:", foo)
	bar := foo[:1]
	foo[0] = "g"
	fmt.Println("Bar1:", bar)
	fmt.Println("foo1:", foo)
	s1 := append(bar, "c")
	fmt.Println("foo2:", foo)
	fmt.Println("s1:", s1)
	fmt.Println("Bar2:", bar)
	//fmt.Println("hello bong")
	//fmt.Println(random.String(10))
}
