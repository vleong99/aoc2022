package main

import (
	utils "aoc2022/libs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	blocksOnGrid := make(map[[3]int]bool)
	file, err := os.Open("day18.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		var coords [3]int
		for i := 0; i < len(line); i++ {
			coords[i] = utils.Ints(line[i])
		}
		blocksOnGrid[coords] = true
	}

	fmt.Println(part1(blocksOnGrid))

	waterSurfaceArea := part1(part2(blocksOnGrid)) - (25 * 25 * 6)
	fmt.Println(waterSurfaceArea)

}

func part1(blocksOnGrid map[[3]int]bool) int {
	ans := 0
	for k := range blocksOnGrid {
		openFaces := 6

		//check six faces
		possibleMoves := [][3]int{
			{-1, 0, 0},
			{1, 0, 0},
			{0, -1, 0},
			{0, 1, 0},
			{0, 0, -1},
			{0, 0, 1},
		}
		for _, v := range possibleMoves {
			if blocksOnGrid[[3]int{k[0] + v[0], k[1] + v[1], k[2] + v[2]}] {
				openFaces--
			}
		}
		ans += openFaces
	}
	return ans
}

func part2(blocksOnGrid map[[3]int]bool) map[[3]int]bool {

	gridSize := 23
	visited := make(map[[3]int]bool)

	possibleMoves := [][3]int{
		{-1, 0, 0},
		{1, 0, 0},
		{0, -1, 0},
		{0, 1, 0},
		{0, 0, -1},
		{0, 0, 1},
	}
	var queue [][3]int

	queue = append(queue, [3]int{-1, -1, -1}) // because there is a block with a coordinate of zero

	for len(queue) > 0 {
		currentBlock := queue[0]

		queue = queue[1:]

		for _, v := range possibleMoves {
			n := [3]int{v[0] + currentBlock[0], v[1] + currentBlock[1], v[2] + currentBlock[2]} //neighbour

			if n[0] >= -1 && n[0] <= gridSize && n[1] >= -1 && n[1] <= gridSize && n[2] >= -1 && n[2] <= gridSize && // check within bounds
				!blocksOnGrid[n] && !visited[n] { // check not a block and not already visited
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}

	return visited
}
