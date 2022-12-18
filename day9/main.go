package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var grid [][]int

func main() {

	count := 0

	for i := 0; i < 100000; i++ {
		l := make([]int, 100000)

		grid = append(grid, l)
	}
	//part 1

	h := []int{50000, 50000} //col, row

	t := []int{50000, 50000} //col, row

	//part 2

	// h, one, two, three, four, five, six, seven, eight, nine := []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000}, []int{50000, 50000} //col, row

	// knots := [][]int{h, one, two, three, four, five, six, seven, eight, nine}

	file, err := os.Open("day9.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//part 1

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		reps, _ := strconv.Atoi(line[1])
		for i := 0; i < reps; i++ {
			switch line[0] {
			case "L":
				h = moveLeft(h)
			case "R":
				h = moveRight(h)
			case "U":
				h = moveUp(h)
			case "D":
				h = moveDown(h)
			}
			h, t = touching(h, t)

			if grid[t[0]][t[1]] == 0 {
				count++
				grid[t[0]][t[1]] = 1
			}
		}
	}

	//part 2
	// for scanner.Scan() {
	// 	line := strings.Split(scanner.Text(), " ")
	// 	reps, _ := strconv.Atoi(line[1])
	// 	for i := 0; i < reps; i++ {
	// 		switch line[0] {
	// 		case "L":
	// 			h = moveLeft(h)
	// 		case "R":
	// 			h = moveRight(h)
	// 		case "U":
	// 			h = moveUp(h)
	// 		case "D":
	// 			h = moveDown(h)
	// 		}
	// 		for j := 0; j < len(knots)-1; j++ {
	// 			h, t = touching(h, t)
	// 		}

	// 		if grid[t[0]][t[1]] == 0 {
	// 			count++
	// 			grid[t[0]][t[1]] = 1
	// 		}
	// 	}
	// }
	fmt.Println(count)
}

func moveLeft(h []int) []int {
	h[1] -= 1
	return h
}

func moveRight(h []int) []int {
	h[1] += 1
	return h
}
func moveUp(h []int) []int {
	h[0] -= 1
	return h
}
func moveDown(h []int) []int {
	h[0] += 1
	return h
}
func touching(h, t []int) ([]int, []int) { //x is col and y is row, 1 is move straight and 2 is move diagonal
	//same row or col

	//right

	if h[0] == t[0] && h[1]-t[1] > 1 {
		t[1]++
	}

	//left

	if h[0] == t[0] && t[1]-h[1] > 1 {
		t[1]--
	}

	//top

	if h[1] == t[1] && t[0]-h[0] > 1 {
		t[0]--
	}
	//bottom

	if h[1] == t[1] && h[0]-t[0] > 1 {
		t[0]++
	}

	//diag top right

	if (h[0] < t[0] && h[1] > t[1]) && (t[0]-h[0] > 1 || h[1]-t[1] > 1) {
		t[0]--
		t[1]++
	}

	//diag top left

	if (h[0] < t[0] && h[1] < t[1]) && (t[0]-h[0] > 1 || t[1]-h[1] > 1) {
		t[0]--
		t[1]--
	}

	//diag bottom right

	if (h[0] > t[0] && h[1] > t[1]) && (h[0]-t[0] > 1 || h[1]-t[1] > 1) {
		t[0]++
		t[1]++
	}

	//diag bottom left

	if (h[0] > t[0] && h[1] < t[1]) && (h[0]-t[0] > 1 || t[1]-h[1] > 1) {
		t[0]++
		t[1]--
	}

	return h, t
}
