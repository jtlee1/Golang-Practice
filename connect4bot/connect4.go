package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"sync"
)

// -1 = red. 1 = black
type Slot struct {
	Taken bool
	Color int
}

func GetUserMove() int {
	fmt.Println("place your move (0~6) :")
	var in string
	fmt.Scanln(&in)
	t, _ := strconv.Atoi(in)
	if t >= 0 && t <= 6 {
		return t
	}
	return GetUserMove()
}

func NewSlot() Slot {
	var s Slot
	s.Taken = false
	s.Color = -50
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

func getColor(i int) string {
	switch i {
	case -1:
		return "X"
	case 1:
		return "O"
	default:
		return " "
	}
}

func checkPlace(loc int, top []int) (bool, error) {
	if top[loc] == 6 {
		return true, errors.New("column full")
	}
	return false, nil
}

func place(loc int, turn int, m [7][6]Slot, top []int) ([7][6]Slot, bool) {
	if _, err := checkPlace(loc, top); err != nil {
		fmt.Println(err)
		return m, true
	}
	m[loc][top[loc]].Taken = true
	m[loc][top[loc]].Color = turn
	top[loc] += 1
	return m, false
}

func checkRight(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
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
	default:
	}
}

func checkLeft(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
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
	default:
	}
}

func checkTop(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
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
	default:
	}
}

func checkBot(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
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
	default:
	}
}

func checkTR(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
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
	default:
	}
}
func checkTL(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if -goal+col+1 < 0 {
		return
	}
	if goal+row > 6 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col-i][row+i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	default:
	}
}

func checkBR(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if goal+col > 7 {
		return
	}
	if -goal+row+1 < 0 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col+i][row-i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	default:
	}
}
func checkBL(hypo, add, col, row, goal int, f *int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if -goal+col+1 < 0 {
		return
	}
	if -goal+row+1 < 0 {
		return
	}
	con := hypo
	for i := 0; i < goal; i++ {
		con += m[col-i][row-i].Color
	}
	switch con {
	case -goal:
		*f = -1 - add
	case goal:
		*f = 1 + add
	default:
	}
}

func checkWinSub(m [7][6]Slot) int {
	f := 0
	wg := new(sync.WaitGroup)
	wg.Add(69) //21+24+12+12
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			go checkTop(0, 0, i, j, 4, &f, m, wg)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			go checkRight(0, 0, i, j, 4, &f, m, wg)

		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			go checkTR(0, 0, i, j, 4, &f, m, wg)
			go checkTL(0, 0, i+3, j, 4, &f, m, wg)
		}
	}
	wg.Wait()
	return f
}

func getScore(loc int, turn int, m [7][6]Slot, top []int) int {
	s1 := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}}
	s2 := [][]int{{0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0}}
	sum := 0
	//fmt.Println(top[loc])
	wg := new(sync.WaitGroup)
	wg.Add(48)
	for i := 2; i < 5; i++ {
		go checkTop(turn+50, i*i, loc, top[loc], i, &s1[i-2][0], m, wg)
		go checkBot(turn+50, i*i, loc, top[loc], i, &s1[i-2][1], m, wg)
		go checkLeft(turn+50, i*i, loc, top[loc], i, &s1[i-2][2], m, wg)
		go checkRight(turn+50, i*i, loc, top[loc], i, &s1[i-2][3], m, wg)
		go checkTR(turn+50, i*i, loc, top[loc], i, &s1[i-2][4], m, wg)
		go checkTL(turn+50, i*i, loc, top[loc], i, &s1[i-2][5], m, wg)
		go checkBR(turn+50, i*i, loc, top[loc], i, &s1[i-2][6], m, wg)
		go checkBL(turn+50, i*i, loc, top[loc], i, &s1[i-2][7], m, wg)
		go checkTop(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][0], m, wg)
		go checkBot(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][1], m, wg)
		go checkLeft(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][2], m, wg)
		go checkRight(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][3], m, wg)
		go checkTR(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][4], m, wg)
		go checkTL(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][5], m, wg)
		go checkBR(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][6], m, wg)
		go checkBL(-turn+50, i*i-1, loc, top[loc], i, &s2[i-2][7], m, wg)
	}
	wg.Wait()
	for i := 0; i < 3; i++ {
		for j := 0; j < 8; j++ {
			//fmt.Println(i, ": ", s1[i][j])
			sum += s1[i][j] * turn
			sum += s2[i][j]
		}
	}
	return sum
}

func GetMove(turn int, m [7][6]Slot, top []int) int {
	s1 := []int{0, 0, 0, 0, 0, 0, 0}
	for i := 0; i < 7; i++ {
		if top[i] >= 6 {
			s1[i] = -99999
		} else {
			s1[i] = getScore(i, turn, m, top) - int(math.Abs(float64(i-3)))
		}
	}
	max := 0
	loc := 0
	for i, e := range s1 {
		//fmt.Println(e)
		if i == 0 || e > max {
			max = e
			loc = i
		}
	}
	return loc
}

func PrintMatrix(m [7][6]Slot) {
	for row := 5; row >= 0; row-- {
		for column := 0; column < 7; column++ {
			fmt.Print(getColor(m[column][row].Color), " ")
		}
		fmt.Print("\n")
	}
	for i := 0; i < 7; i++ {
		fmt.Print(i, " ")
	}
	fmt.Print("\n")
}

func CheckWin(m [7][6]Slot) {
	con := checkWinSub(m)
	switch con {
	case 1:
		//PrintMatrix(m)
		fmt.Println("you win")
		os.Exit(3)
	case -1:
		//PrintMatrix(m)
		fmt.Println("you lose")
		os.Exit(3)
	default:
	}
}

func main() {
	top := []int{0, 0, 0, 0, 0, 0, 0}
	var matrix [7][6]Slot
	//turn := -1
	matrix = setup(matrix)
	fill := 0
	for {
		t := GetUserMove()
		ok := true
		matrix, ok = place(t, 1, matrix, top)
		for {
			if ok {
				t = GetUserMove()
				matrix, ok = place(t, 1, matrix, top)
			} else {
				break
			}
		}
		CheckWin(matrix)
		cal := GetMove(-1, matrix, top)
		matrix, _ = place(cal, -1, matrix, top)
		fmt.Println("\033[2J")
		PrintMatrix(matrix)
		CheckWin(matrix)
		fill += 2
		if fill == 42 {
			fmt.Println("Draw")
			break
		}
	}
}
