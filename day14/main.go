package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var smallCol, largeCol, largeRow int

func main() {

	smallCol, largeCol, largeRow = 100000, 0, 0

	var cave [][]string

	for i := 0; i < 1000; i++ {
		l := []string{}
		for j := 0; j < 1000; j++ {
			l = append(l, ".")
		}
		cave = append(cave, l)
	}

	file, err := os.Open("day14.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		var curr, prev []int

		line := strings.Split(scanner.Text(), " -> ")

		for i := 0; i < len(line); i++ {
			coordinates := strings.Split(line[i], ",")
			col, _ := strconv.Atoi(coordinates[0])
			row, _ := strconv.Atoi(coordinates[1])
			if i == 0 {
				prev = []int{col, row}
			}
			curr = []int{col, row}
			if col < smallCol {
				smallCol = col
			}
			if col > largeCol {
				largeCol = col
			}
			if row > largeRow {
				largeRow = row
			}
			//if currRow > prevRow

			if curr[1] > prev[1] {
				for i := prev[1]; i <= curr[1]; i++ {
					cave[i][curr[0]] = "#"
				}
			}

			//if currRow < prevRow

			if curr[1] < prev[1] {
				for i := curr[1]; i <= prev[1]; i++ {
					cave[i][curr[0]] = "#"
				}
			}

			// if currCol > preCol

			if curr[0] > prev[0] {
				for i := prev[0]; i <= curr[0]; i++ {
					cave[curr[1]][i] = "#"
				}
			}

			// if currCol < preCol

			if curr[0] < prev[0] {
				for i := curr[0]; i <= prev[0]; i++ {
					cave[curr[1]][i] = "#"
				}
			}

			prev = []int{col, row}
		}

	}

	part1(cave)
	// part2(cave)

}

func blocked(cave *[][]string, col, row int) bool {
	s := (*cave)[col][row]
	return (s == "#" || s == "o")
}

func moveSand(cave *[][]string, curr []int) ([]int, bool) { // curr is col,row

	//part1
	if curr[1]+1 > largeCol || curr[1]-1 < smallCol || curr[0]+1 > largeRow {
		return curr, true
	}

	//part2
	// if curr[0] >= len(*cave)-1 {
	// 	return curr, false
	// }

	if !blocked(cave, curr[0]+1, curr[1]) {
		return []int{curr[0] + 1, curr[1]}, false
	}
	if !blocked(cave, curr[0]+1, curr[1]-1) {
		return []int{curr[0] + 1, curr[1] - 1}, false
	}
	if !blocked(cave, curr[0]+1, curr[1]+1) {
		return []int{curr[0] + 1, curr[1] + 1}, false
	}
	return curr, false
}

func part1(cave [][]string) {
	sand := 0

	//simulate sand

Sim:
	for {
		sandPos := []int{0, 500}
		for {
			newSandPos, oob := moveSand(&cave, sandPos)
			if oob {
				break Sim
			}
			if newSandPos[0] == sandPos[0] && newSandPos[1] == sandPos[1] {
				cave[newSandPos[0]][newSandPos[1]] = "o"
				sand++
				for i := 0; i <= largeRow; i++ {
					fmt.Println(cave[i][smallCol : largeCol+1])

				}
				break
			}
			sandPos = newSandPos
		}
	}

	fmt.Println(sand)
}

func part2(c [][]string) {

	sand := 0

	cave := c[:largeRow+2]

	for i := 0; i < len(cave); i++ {
		fmt.Println(cave[i][smallCol-5 : largeCol+1+5])

	}

Sim:
	for {
		sandPos := []int{0, 500}
		for {
			newSandPos, _ := moveSand(&cave, sandPos)
			if newSandPos[0] == 0 && newSandPos[1] == 500 {
				sand++
				break Sim
			}
			if newSandPos[0] == sandPos[0] && newSandPos[1] == sandPos[1] {
				cave[newSandPos[0]][newSandPos[1]] = "o"
				sand++
				// for i := 0; i < len(cave); i++ {
				// 	fmt.Println(cave[i][smallCol-10 : largeCol+1+10])

				// }
				break
			}

			sandPos = newSandPos
		}
	}

	fmt.Println(sand)
}
