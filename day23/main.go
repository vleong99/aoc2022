package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var grid map[[2]int]string
var check map[string][][2]int = map[string][][2]int{
	"U": {{-1, -1}, {-1, 0}, {-1, 1}}, "D": {{1, -1}, {1, 0}, {1, 1}}, //NW, N, NE, SW, S, SE
	"R": {{-1, 1}, {0, 1}, {1, 1}}, "L": {{-1, -1}, {0, -1}, {1, -1}}, //NE, E, SE, NW, W, SW
}

var order []string = []string{"U", "D", "L", "R"}

func main() {
	grid = parse()
	round()
}

func parse() map[[2]int]string {

	grid := make(map[[2]int]string)
	file, err := os.Open("day23.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		for i := 0; i < len(scanner.Text()); i++ {
			if scanner.Text()[i] == '#' {
				grid[[2]int{row, i}] = "#"
			}
		}
		row++
	}
	return grid
}

func adj(x, y int) bool { //check if any elves adjacent
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {1, -1}, {1, 0}, {1, 1}, {0, 1}, {0, -1}}
	for _, v := range directions {
		if grid[[2]int{x + v[0], y + v[1]}] == "#" {
			return true
		}
	}
	return false
}

func round() {
	i := 0
	for {
		if i == 10 { //part 1 10 rounds
			emptyTiles()
		}
		steps := make(map[[2]int][2]int)
		seen := make(map[[2]int]int)
		cOrder := []string{order[(i+0)%4], order[(i+1)%4], order[(i+2)%4], order[(i+3)%4]} //order of the directions
		for k, v := range grid {                                                           //check all elf positions
			if v == "#" && adj(k[0], k[1]) { //check an elf is there
				for i := 0; i < 4; i++ { //check all 4 directions in order
					currDir := cOrder[i] // current direction
					blocked := false
					for _, d := range check[currDir] {
						if grid[[2]int{k[0] + d[0], k[1] + d[1]}] == "#" {
							blocked = true
						}
					}
					if blocked == false {
						steps[k] = [2]int{k[0] + check[currDir][1][0], k[1] + check[currDir][1][1]}
						seen[steps[k]]++
						break
					}
				}
			}
		}
		moved := false
		for k, v := range steps {
			if seen[v] == 1 {
				grid[k] = ""
				grid[v] = "#"
				moved = true
			}
		}
		i++
		if moved == false {
			fmt.Printf("part2:%v\n", i)
			return
		}
	}
}

func emptyTiles() {
	sR, sC, lR, lC := 10000000000, 10000000000, -10000000000, -10000000000
	count := 0

	for k, v := range grid {
		if v != "#" {
			continue
		}
		count++
		sR = int(math.Min(float64(k[0]), float64(sR)))
		lR = int(math.Max(float64(k[0]), float64(lR)))
		sC = int(math.Min(float64(k[1]), float64(sC)))
		lC = int(math.Max(float64(k[1]), float64(lC)))

	}
	h := lR - sR + 1
	w := lC - sC + 1
	fmt.Printf("part1:%v\n", h*w-count)
}
