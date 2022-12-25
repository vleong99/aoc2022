package main

import (
	"bufio"
	"fmt"
	"os"
)

type row struct {
	right, left []int
}

type col struct {
	up, down []int
}

var allRows []row

var allCols []col

func main() {

	grid := make(map[[2]int]byte)

	file, err := os.Open("day24.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	rows := 0
	cols := 0

	for i := 0; i < 8; i++ {
		allCols = append(allCols, col{[]int{}, []int{}})
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		allRows = append(allRows, row{[]int{}, []int{}})
		line := scanner.Text()
		cols = len(line)
		for i := 0; i < len(line); i++ {
			grid[[2]int{rows, i}] = line[i]
			switch line[i] {
			case '>':
				allRows[rows].right = append(allRows[rows].right, i)
			case '<':
				allRows[rows].left = append(allRows[rows].left, i)
			case '^':
				allCols[i].up = append(allCols[i].up, rows)
			case 'v':
				allCols[i].down = append(allCols[i].down, rows)
			}
		}
		rows++
	}

	//track length of path
	turn := 0

	//initialise queue and append start coords
	start := [2]int{0, 1}
	var queue [][2]int
	queue = append(queue,[2]int{0, 1})

	moves := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		turn++
		curr := queue[0]
		queue = queue[1:]
		if curr == [2]int{5,6} {
			fmt.Println(turn)
			break
		}
		for _, m := range moves {
			pM := [2]int{curr[0]+m[0], curr[1]+m[1]} //potential move
			if curr[0]+m[0] < 1 || curr[0]+m[0] > 4 || curr[1]+m[1] < 1 || curr[1]+m[1] > 6 {
				if curr[0]+m[0] == 5 && curr[1]+m[1] == 6 {
					queue = append(queue, [2]int{curr[0]+m[0],curr[1]+m[1]})
				}
				continue
			}
		}

	}

	
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
	


	// grid[[2]int{0, 1}] = 'E'

	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < cols; j++ {
	// 		fmt.Print(string(grid[[2]int{i, j}]))
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println(allRows)

	// fmt.Println(allCols)

}


func contains(b[]int, curr int ) bool {
	for _, v := range b {
		if curr == v {
			return true
		}
	}
	return false
}