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

	// var sums []int

	printLine := ""

	X := 1

	cycle := 0

	file, err := os.Open("day10.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//part 1

	for scanner.Scan() {

		//part1
		// if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 {
		// 	product := cycle * X
		// 	sums = append(sums, product)
		// 	fmt.Println(X)
		// }

		// if cycle == 220 {
		// 	product := cycle * X
		// 	sums = append(sums, product)
		// 	fmt.Println(X)
		// 	fmt.Println(sums)
		// 	return
		// }

		line := strings.Split(scanner.Text(), " ")

		switch line[0] {
		case "noop":
			cycle++

			if len(printLine) == X-1 || len(printLine) == X || len(printLine) == X+1 {
				printLine += "#"
			} else {
				printLine += "."
			}
			if len(printLine) == 40 {
				fmt.Println(printLine)
				printLine = ""
			}
			if cycle == 10 {
				fmt.Println(X)
				fmt.Println(printLine)
			}
		case "addx":
			num, _ := strconv.Atoi(line[1])
			cycle++

			if len(printLine) == X-1 || len(printLine) == X || len(printLine) == X+1 {
				printLine += "#"
			} else {
				printLine += "."
			}
			if len(printLine) == 40 {
				fmt.Println(printLine)
				printLine = ""
			}
			if cycle == 10 {
				fmt.Println(X)
				fmt.Println(printLine)
			}

			//part1
			// if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 {
			// 	fmt.Println(X)
			// 	product := cycle * X
			// 	sums = append(sums, product)
			// }
			// if cycle == 220 {
			// 	fmt.Println(X)
			// 	product := cycle * X
			// 	sums = append(sums, product)
			// 	fmt.Println(sums)
			// 	return
			// }

			//part2

			cycle++

			if len(printLine) == X-1 || len(printLine) == X || len(printLine) == X+1 {
				printLine += "#"
			} else {
				printLine += "."
			}
			if len(printLine) == 40 {
				fmt.Println(printLine)
				printLine = ""
			}
			if cycle == 10 {
				fmt.Println(X)
				fmt.Println(printLine)
			}
			X += num
		}

	}

}
