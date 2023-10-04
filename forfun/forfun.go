package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
	foo    int
	bar    string
}

func main() {
	p := new(SyncedBuffer)
	println(p.foo)
	fmt.Println("bar:", p.bar)
	//fmt.Printf("%#v\n", p)
}
