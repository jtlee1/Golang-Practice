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
	s.Color = 0
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

func checkRight(col, row, goal int, ch chan int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if goal+col > 7 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goal; i++ {
		con += m[col+i][row].Color
	}
	switch con {
	case -goal:
		//fmt.Println("RIGHT -")
		ch <- -1
	case goal:
		//fmt.Println("RIGHT +")
		ch <- 1
	default:
		ch <- 0
	}
}

func checkTop(col, row, goal int, ch chan int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if goal+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goal; i++ {
		con += m[col][row+i].Color
	}
	switch con {
	case -goal:
		//fmt.Println("TOP -")
		ch <- -1
	case goal:
		//fmt.Println("TOP +")
		ch <- 1
	default:
		ch <- 0
	}
}

func checkTR(col, row, goal int, ch chan int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if goal+col > 7 {
		ch <- 0
		return
	}
	if goal+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goal; i++ {
		con += m[col+i][row+i].Color
	}
	switch con {
	case -goal:
		//fmt.Println("TR -")
		ch <- -1
	case goal:
		//fmt.Println("TR +")
		ch <- 1
	default:
		ch <- 0
	}
}
func checkTL(col, row, goal int, ch chan int, m [7][6]Slot, w *sync.WaitGroup) {
	defer w.Done()
	if -goal+col+1 < 0 {
		ch <- 0
		return
	}
	if goal+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goal; i++ {
		con += m[col-i][row+i].Color
	}
	switch con {
	case -goal:
		//fmt.Println("TL -")
		ch <- -1
	case goal:
		//fmt.Println("TL +")
		ch <- 1
	default:
		ch <- 0
	}
}

func checkWinSub(m [7][6]Slot) int {
	ch := make(chan int, 69)
	wg := new(sync.WaitGroup)
	wg.Add(69) //21+24+12+12
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			go checkTop(i, j, 4, ch, m, wg)
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			go checkRight(i, j, 4, ch, m, wg)

		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			go checkTR(i, j, 4, ch, m, wg)
			go checkTL(i+3, j, 4, ch, m, wg)
		}
	}
	wg.Wait()
	f := 0
	for i := 0; i < 69; i++ {
		f += <-ch
	}
	close(ch)
	if f != 0 {
		f = f / int(math.Abs(float64(f)))
	}
	return f
}

func getScore(turn int, m [7][6]Slot, goal, weight1, weight2 int) int {
	ch := make(chan int, 42*4)
	wg := new(sync.WaitGroup)
	wg.Add(42 * 4)
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			go checkTop(i, j, goal, ch, m, wg)
			go checkRight(i, j, goal, ch, m, wg)
			go checkTR(i, j, goal, ch, m, wg)
			go checkTL(i, j, goal, ch, m, wg)
		}
	}
	wg.Wait()
	f := 0
	sum := 0
	for i := 0; i < 168; i++ {
		f = <-ch
		//fmt.Println(f)
		if f*turn > 0 {
			sum += weight1
		}
		if f*turn < 0 {
			sum += weight2
		}
	}
	close(ch)
	return sum
}

// output total estimation score
func getTotScore(turn int, m [7][6]Slot, top []int) int {
	spec1 := [3]int{3, 10, 10000}
	spec2 := [3]int{2, 8, 500}
	score := 0
	for i := 2; i < 5; i++ {
		score += getScore(turn, m, i, spec1[i-2], spec2[i-2])
	}
	return score
}

func getMaxScore(m [7][6]Slot, top []int, turn, iter int) (int, int) {
	if iter <= 0 {
		return 0, 0
	}
	s1 := []int{0, 0, 0, 0, 0, 0, 0}
	var matrix [7][7][6]Slot
	var matrix2 [7][7][6]Slot
	var t1 [7][]int
	var t2 [7][]int
	for i := 0; i < 7; i++ {
		t1[i] = append(t1[i], top...)
		t2[i] = append(t2[i], top...)
		//matrix[i] = setup(m)
		matrix[i] = matrixCopy(m)
		matrix2[i] = matrixCopy(m)
		matrix[i], _ = place(i, turn, matrix[i], t1[i])
		matrix2[i], _ = place(i, -turn, matrix2[i], t2[i])

		if top[i] >= 6 {
			s1[i] = -99999
		} else {
			tempMax1, _ := getMaxScore(matrix[i], t1[i], -turn, iter-1)
			//tempMax2, _ := getMaxScore(matrix2[i], t2[i], -turn, iter-1)
			fmt.Println("get1: ", getTotScore(turn, matrix[i], t1[i]))
			fmt.Println("get2: ", getTotScore(turn, matrix2[i], t2[i]))
			s1[i] = getTotScore(turn, matrix[i], t1[i]) + getTotScore(turn, matrix2[i], t2[i]) - int(math.Abs(float64(i-3))) - (tempMax1)/2
		}
	}
	max := -10000
	loc := 0
	for i, e := range s1 {
		fmt.Println(e)
		if i == 0 || e > max {
			max = e
			loc = i
		}
	}
	return max, loc
}

func matrixCopy(m [7][6]Slot) [7][6]Slot {
	var matrix [7][6]Slot
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			matrix[i][j] = m[i][j]
		}
	}
	return matrix
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
	matrix = setup(matrix)
	//turn := -1
	//matrix = setup(matrix)

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
		_, cal := getMaxScore(matrix, top, -1, 1)
		matrix, _ = place(cal, -1, matrix, top)
		//fmt.Println("\033[2J")
		PrintMatrix(matrix)
		CheckWin(matrix)
		fill += 2
		if fill == 42 {
			fmt.Println("Draw")
			break
		}
	}
	//matrix, _ = place(2, 1, matrix, top)
	//matrix, _ = place(3, 1, matrix, top)
	//sc, loc := getMaxScore(matrix, top, -1, 1)
	//fmt.Println("sc: ", sc)
	//fmt.Println("loc: ", loc)

}
