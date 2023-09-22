package main

import (
	"errors"
	"fmt"
	"sync"
)

// -1 = red. 1 = black
type Slot struct {
	Taken bool
	Color int
}

func NewSlot() Slot {
	var s Slot
	s.Taken = false
	s.Color = -5
	return s
}

func setup(m [7][6]Slot) [7][6]Slot {
	wg := new(sync.WaitGroup)
	wg.Add(42)
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			go func(i, j int) {
				m[i][j] = NewSlot()
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	return m
}

func checkPlace(loc int, top []int) (bool, error) {
	if top[loc] == 6 {
		return true, errors.New("Column full")
	}
	return false, nil
}

func place(loc int, turn int, m [7][6]Slot, top []int) [7][6]Slot {
	if _, err := checkPlace(loc, top); err != nil {
		fmt.Println(err)
	}
	m[loc][top[loc]].Taken = true
	m[loc][top[loc]].Color = turn
	top[loc] += 1
	return m
}

func checkRight(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if goal+col > 7 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col+i][row].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	}
}

func checkLeft(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if -goal+col+1 < 0 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col-i][row].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	}
}

func checkTop(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if goal+row > 6 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col][row+i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	}
}

func checkBot(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if -goal+row+1 < 0 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col][row-i].Color
	}
	//fmt.Println(goal, ": ", con)
	switch con {
	case -goal:
		*f = (-1 - add)
	case goal:
		*f = 1 + add
	}
}

func checkTR(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if goal+col > 7 {
		return
	}
	if goal+row > 6 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col+i][row+i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	}
}
func checkTL(hypo, add, col, row, goal int, f *int, m [7][6]Slot) {
	if goal+col > 7 {
		return
	}
	if goal+row > 6 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col+i][row+goal-1-i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	}
}

func checkWin(m [7][6]Slot) int {
	f := 0
	wg := new(sync.WaitGroup)
	wg.Add(12)
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			go func(i, j int) {
				checkTop(0, 0, i, j, 4, &f, m)
				checkRight(0, 0, i, j, 4, &f, m)
				checkTR(0, 0, i, j, 4, &f, m)
				checkTL(0, 0, i, j, 4, &f, m)
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	return f
}

func getScore(loc int, turn int, m [7][6]Slot, top []int) int {
	s1 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	s2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sum := 0
	//fmt.Println(top[loc])
	for i := 2; i < 5; i++ {
		go checkTop(turn+5, i*i, loc, top[loc], i, &s1[0+(i-2)*4], m)
		go checkBot(turn+5, i*i, loc, top[loc], i, &s1[1+(i-2)*4], m)
		go checkLeft(turn+5, i*i, loc, top[loc], i, &s1[2+(i-2)*4], m)
		go checkRight(turn+5, i*i, loc, top[loc], i, &s1[3+(i-2)*4], m)
		go checkTop(-turn+5, i*i-1, loc, top[loc], i, &s2[0+(i-2)*4], m)
		go checkBot(-turn+5, i*i-1, loc, top[loc], i, &s2[0+(i-2)*4], m)
		go checkLeft(-turn+5, i*i-1, loc, top[loc], i, &s2[0+(i-2)*4], m)
		go checkRight(-turn+5, i*i-1, loc, top[loc], i, &s2[0+(i-2)*4], m)
	}
	for i := 0; i < 16; i++ {
		//fmt.Println(i, ": ", s[i])
		sum += s1[i] * turn
		sum += s2[i]
	}
	return sum
}

func main() {
	top := []int{0, 0, 0, 0, 0, 0, 0}
	var matrix [7][6]Slot
	turn := -1
	matrix = setup(matrix)
	//fmt.Println(matrix[1][1].Color)
	matrix = place(1, turn, matrix, top)
	matrix = place(1, turn, matrix, top)
	matrix = place(1, turn, matrix, top)
	//fmt.Println(matrix[1][1].Color)
	//checkHori(1, -1, matrix)
	//matrix = place(1, turn, matrix)
	/*
		for row := 0; row < 7; row++ {
			for column := 0; column < 6; column++ {
				fmt.Print(matrix[row][column].Color, " ")
			}
			fmt.Print("\n")
		}*/

	fmt.Println(getScore(1, turn, matrix, top))
}
