package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type b struct { // blizzard
	r, c int    //coords
	dir  string //direction
}

type node struct {
	row, col, pathLength int
}

var allB []b

var allBStates map[int]map[[2]int]bool

var rows, cols, LCM int

func main() {

	input, err := ioutil.ReadFile("day24.txt")
	if err != nil {
		fmt.Println(err)
	}

	eachLine := strings.Split(string(input), "\n")

	rows = len(eachLine) //row length

	cols = len(eachLine[0]) //col length

	start, end := [2]int{0, 1}, [2]int{rows - 1, cols - 2}

	for i := 0; i < len(eachLine); i++ {
		for j := 0; j < len(eachLine[i]); j++ {
			if eachLine[i][j] == '>' || eachLine[i][j] == '<' || eachLine[i][j] == '^' || eachLine[i][j] == 'v' {
				newB := b{i, j, "NA"}
				switch eachLine[i][j] {
				case '>':
					newB.dir = "R"
				case '<':
					newB.dir = "L"
				case '^':
					newB.dir = "U"
				case 'v':
					newB.dir = "D"
				}
				allB = append(allB, newB)
			}
		}
	}

	LCM = findLCM(rows-2, cols-2) //lowest common denominator

	allBStates = make(map[int]map[[2]int]bool)

	for i := 0; i < LCM; i++ {
		allBStates[i] = newB(allB, i)
	}

	a := shortestPath(start, end, 0)
	b := shortestPath(end, start, a) - a
	c := shortestPath(start, end, a+b) - (a + b)

	fmt.Println(a, b, c)
	fmt.Println(a + b + c)
}

func findLCM(a, b int) int { //find least common multiple
	i := 1
	for {
		if i%a == 0 && i%b == 0 {
			return i
		}
		i++
	}
}

func shortestPath(start, end [2]int, turn int) int {
	var queue []node

	seen := make(map[[3]int]bool)

	queue = append(queue, node{start[0], start[1], turn}) //append start node

	moves := [][]int{{0, 0}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {

		curr := queue[0]
		queue = queue[1:]

		cB := allBStates[curr.pathLength%LCM]

		for _, m := range moves {
			next := node{curr.row + (m[0]), curr.col + (m[1]), curr.pathLength + 1}

			if next.row == end[0] && next.col == end[1] {
				return next.pathLength - 1
			}

			if next.row < 1 || next.row > rows-2 || next.col < 1 || next.col > cols-2 {
				if !(next.row == start[0] && next.col == start[1]) {
					continue
				}
			}

			if !cB[[2]int{next.row, next.col}] && !seen[[3]int{next.col, next.row, next.pathLength}] {
				queue = append(queue, next)
				seen[[3]int{next.col, next.row, next.pathLength}] = true
			}
		}
	}
	return -1
}

func blizzard(b []b, new [2]int) bool {
	for _, v := range b {
		if v.r == new[0] && v.c == new[1] {
			return true
		}
	}
	return false
}

func newB(blizzards []b, turn int) map[[2]int]bool {

	bState := make(map[[2]int]bool)

	turnRow := turn % (rows - 2)

	turnCol := turn % (cols - 2)

	for i := 0; i < len(blizzards); i++ {
		v := (blizzards)[i]
		new := [2]int{v.r, v.c}
		switch v.dir {
		case "U":
			new[0] = v.r - turnRow
			if (v.r - turnRow) < 1 {
				new[0] = (rows - 2) + (v.r - turnRow)
			}
		case "D":
			new[0] = v.r + turnRow
			if (v.r + turnRow) > rows-2 {
				new[0] = (v.r + turnRow) - (rows - 2)
			}
		case "L":
			new[1] = v.c - turnCol
			if (v.c - turnCol) < 1 {
				new[1] = (cols - 2) + (v.c - turnCol)
			}
		case "R":
			new[1] = v.c + turnCol
			if (v.c + turnCol) > cols-2 {
				new[1] = (v.c + turnCol) - (cols - 2)
			}
		}
		bState[[2]int{new[0], new[1]}] = true
	}
	return bState
}
