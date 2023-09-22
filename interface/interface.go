package main

import (
	"fmt"
)

type Animal interface {
	MakeSound()
	Walking()
}

type Dog struct {
	Sound string
}

func (d *Dog) MakeSound() {
	fmt.Println(d.Sound)
}

func (d *Dog) Walking() {
	fmt.Println("dog walking")
}

type Cat struct {
	Sound string
}

func (c *Cat) MakeSound() {
	fmt.Println(c.Sound)
}

func (c *Cat) Walking() {
	fmt.Println("cat walking")
}

func Noise(a Animal) {
	a.MakeSound()
}

func Chill(a Animal) {
	a.Walking()
}

func main() {
	d := Dog{"Woof"}
	c := Cat{"Meow"}
	Noise(&d)
	Noise(&c)
}
