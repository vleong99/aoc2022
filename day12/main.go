package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type node struct {
	row, col, val, pathLength int
}

func main() {

	//set up grid

	var grid [][]node

	rowNo := 0

	colNo := 0

	//part1
	// var start []int

	var end []int

	var starts [][]int

	file, err := os.Open("day12.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		var eachLine []node

		for _, v := range line {
			n := node{rowNo, colNo, int(v), 0}
			if n.val == 83 || n.val == 97 {
				starts = append(starts, []int{rowNo, colNo})
				n.val = 97
			}
			if n.val == 69 {
				end = []int{rowNo, colNo}
				n.val = 122
			}
			eachLine = append(eachLine, n)
			colNo++
		}

		grid = append(grid, eachLine)

		rowNo++

		colNo = 0

	}

	var allLengths []int

	for _, v := range starts {

		var newGrid [][]node

		for _, r := range grid {
			var line []node
			for _, c := range r {
				line = append(line, c)
			}
			newGrid = append(newGrid, line)
		}

		n := shortestPath(newGrid, v, end)
		allLengths = append(allLengths, n)
	}

	sort.Ints(allLengths)
	ans := 0
	for i := 0; i < len(allLengths); i++ {
		if allLengths[i] > 0 {
			ans = allLengths[i]
			break
		}
	}
	fmt.Println(ans)

}

func valid(grid *[][]node, a, b []int) bool {
	// if b[0] < 0 || b[1] < 0 || b[0] > len(*grid)-1 || b[1] > len((*grid)[0])-1 {
	// 	return false
	// }

	if ((*grid)[b[0]][b[1]].val - (*grid)[a[0]][a[1]].val) > 1 {
		return false
	}

	return true

}

func shortestPath(grid [][]node, start, end []int) int {
	var queue []node

	//append start node

	grid[start[0]][start[1]].pathLength = 1

	queue = append(queue, grid[start[0]][start[1]])

	//process BFS

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.row == end[0] && curr.col == end[1] {

			return curr.pathLength - 1
		}

		moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

		for _, m := range moves {

			if curr.row+(m[0]) < 0 || curr.row+(m[0]) > len(grid)-1 || curr.col+(m[1]) < 0 || curr.col+(m[1]) > len(grid[0])-1 {
				continue
			}
			next := &grid[curr.row+(m[0])][curr.col+(m[1])]
			if valid(&grid, []int{curr.row, curr.col}, []int{next.row, next.col}) && next.pathLength == 0 {
				next.pathLength = curr.pathLength + 1
				queue = append(queue, *next)
			}
		}

	}
	return -1
}
