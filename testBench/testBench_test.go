package main

import "testing"

func TestAdd1(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add func wrong")
	}
}

func TestAdd2(t *testing.T) {
	if Add(2, 2) != 4 {
		t.Error("Add func wrong")
	}
}

func TestAdd3(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Add func wrong")
	}
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
