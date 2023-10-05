package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"sync"
)

// -1 = red. 1 = black
type Slot struct {
	Taken bool
	Color int
}

func selectLevel() int {
	fmt.Println("select Level (1~5) 5 is the hardest :")
	var in string
	fmt.Scanln(&in)
	t, _ := strconv.Atoi(in)
	if t > 0 && t <= 5 {
		return t
	}
	return selectLevel()
}

func getUserMove() int {
	fmt.Println("place your move (0~6) :")
	var in string
	fmt.Scanln(&in)
	t, err := strconv.Atoi(in)
	if t >= 0 && t <= 6 && err == nil {
		return t
	}
	return getUserMove()
}

func newSlot() Slot {
	var s Slot
	s.Taken = false
	s.Color = 0
	return s
}

func setup(m [7][6]Slot) [7][6]Slot {
	wg := new(sync.WaitGroup)
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			wg.Add(1)
			go func(i, j int) {
				m[i][j] = newSlot()
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
	goalSub := goal
	if goalSub == 3 {
		goalSub = 4
	}
	if goalSub+col > 7 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goalSub; i++ {
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
	goalSub := goal
	if goalSub == 3 {
		goalSub = 4
	}
	if goalSub+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goalSub; i++ {
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
	goalSub := goal
	if goalSub == 3 {
		goalSub = 4
	}
	if goalSub+col > 7 {
		ch <- 0
		return
	}
	if goalSub+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goalSub; i++ {
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
	goalSub := goal
	if goalSub == 3 {
		goalSub = 4
	}
	if -goalSub+col+1 < 0 {
		ch <- 0
		return
	}
	if goalSub+row > 6 {
		ch <- 0
		return
	}
	con := 0
	for i := 0; i < goalSub; i++ {
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
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	count := 0
	for i := 0; i < 7; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j].Taken {
				count += 1
				wg.Add(1)
				go checkTop(i, j, 4, ch, m, wg)
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			if m[i][j].Taken {
				count += 1
				wg.Add(1)
				go checkRight(i, j, 4, ch, m, wg)
			}
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if m[i][j].Taken {
				count += 2
				wg.Add(2)
				go checkTR(i, j, 4, ch, m, wg)
				go checkTL(i+3, j, 4, ch, m, wg)
			}
		}
	}
	f := 0
	for i := 0; i < count; i++ {
		f += <-ch
	}
	if f != 0 {
		f = f / int(math.Abs(float64(f)))
	}
	wg.Wait()
	return f
}

func getScore(turn int, m [7][6]Slot, goal, weight1, weight2 int) int {
	ch := make(chan int)
	count := 0
	wg := new(sync.WaitGroup)
	for i := 0; i < 7; i++ {
		for j := 0; j < 6; j++ {
			if m[i][j].Taken {
				wg.Add(4)
				count += 4
				go checkTop(i, j, goal, ch, m, wg)
				go checkRight(i, j, goal, ch, m, wg)
				go checkTR(i, j, goal, ch, m, wg)
				go checkTL(i, j, goal, ch, m, wg)
			}
		}
	}
	f := 0
	sum := 0
	for i := 0; i < count; i++ {
		f = <-ch
		//fmt.Println(f)
		if f*turn > 0 {
			sum += weight1
		}
		if f*turn < 0 {
			sum += weight2
		}
	}
	wg.Wait()
	//close(ch)
	return sum
}

// output total estimation score
func getTotScore(turn int, m [7][6]Slot, top []int) int {
	spec1 := [3]int{3, 10, 100000}
	spec2 := [3]int{2, 8, 9999}
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
		matrix[i] = matrixCopy(m)
		matrix2[i] = matrixCopy(m)

		if top[i] >= 6 {
			s1[i] = -999999
		} else {
			matrix[i], _ = place(i, turn, matrix[i], t1[i])
			matrix2[i], _ = place(i, -turn, matrix2[i], t2[i])
			tempMax1, _ := getMaxScore(matrix[i], t1[i], -turn, iter-1)
			//fmt.Println("get1: ", getTotScore(turn, matrix[i], t1[i]))
			//fmt.Println("get2: ", getTotScore(turn, matrix2[i], t2[i]))
			checkSol1 := getTotScore(turn, matrix[i], t1[i])
			checkSol2 := getTotScore(turn, matrix2[i], t2[i])
			if checkSol1 >= 100000 {
				return 100000, i
			} else if checkSol2 >= 9999 {
				s1[i] = 90000
			} else {
				s1[i] = checkSol1 + checkSol2 - int(math.Abs(float64(i-3))) - (tempMax1)/5*4
			}
		}
	}
	max := -10000
	loc := 0
	for i, e := range s1 {
		//fmt.Println(e)
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

func printMatrix(m [7][6]Slot) {
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

func checkWin(m [7][6]Slot) {
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
	LVin := selectLevel()
	LV := LVin
	if LVin >= 4 {
		LV = 4
	}
	top := []int{0, 0, 0, 0, 0, 0, 0}
	var matrix [7][6]Slot
	matrix = setup(matrix)

	fill := 0
	for {
		ok := true
		for {
			if ok {
				t := getUserMove()
				matrix, ok = place(t, 1, matrix, top)
			} else {
				break
			}
		}
		checkWin(matrix)
		_, cal := getMaxScore(matrix, top, -1, LV)
		matrix, _ = place(cal, -1, matrix, top)
		fmt.Println("\033[2J")
		printMatrix(matrix)
		checkWin(matrix)
		fill += 2
		if slices.Max(top) == 6 && LVin == 5 {
			LV = 5
		}
		if fill == 42 {
			fmt.Println("Draw")
			break
		}
	}

}
