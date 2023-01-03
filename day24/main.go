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

var allBStates [][]b

var allW [][2]int

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
			if eachLine[i][j] == '#' {
				allW = append(allW, [2]int{i, j})
			}
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

	for i := 0; i < LCM; i++ {
		allBStates = append(allBStates, newB(allB, i))
	}

	a := shortestPath(start, end, 0)

	b := shortestPath(end, start, a) - a

	c := shortestPath(start, end, a+b) - (a + b)

	fmt.Println(a, b, c)

	fmt.Println(a + b + c)

	// shortestPath(end, start)

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

	var seen []node

	//append start node

	queue = append(queue, node{start[0], start[1], turn})

	//process BFS

	moves := [][]int{{0, 0}, {-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		cB := &allBStates[curr.pathLength%LCM]

		for _, m := range moves {

			if curr.row+(m[0]) < 1 || curr.row+(m[0]) > rows-2 || curr.col+(m[1]) < 1 || curr.col+(m[1]) > cols-2 {
				if !((curr.row+(m[0]) == end[0] && curr.col+(m[1]) == end[1]) || (curr.row+(m[0]) == start[0] && curr.col+(m[1]) == start[1])) {
					continue
				}
			}
			next := node{curr.row + (m[0]), curr.col + (m[1]), curr.pathLength + 1}

			if next.row == end[0] && next.col == end[1] {
				return next.pathLength - 1
			}

			if !blizzard(*cB, [2]int{next.row, next.col}) && !seenNode(seen, next) {
				queue = append(queue, next)
				seen = append(seen, next)
				fmt.Println(next)
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

func seenNode(s []node, n node) bool {
	for i := 0; i < len(s); i++ {
		if s[i].row == n.row && s[i].col == n.col && s[i].pathLength == n.pathLength {
			return true
		}
	}
	return false
}

func newB(blizzards []b, turn int) []b {

	newBlizzards := []b{}

	newBlizzards = append(newBlizzards, blizzards...)

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
		(newBlizzards)[i].r = new[0]
		(newBlizzards)[i].c = new[1]
	}
	return newBlizzards
}
