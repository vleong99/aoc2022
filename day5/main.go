package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//general parser solution

func main() {

	var crates [][]string

	file, err := os.Open("day5.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	instructions := false

	line0 := true

	for scanner.Scan() {

		line := scanner.Text()

		if line0 == true {
			numOfStacks := (len(line) + 1) / 4
			for i := 0; i < numOfStacks; i++ {
				crates = append(crates, []string{})
			}
			line0 = false
		}

		if instructions == false {
			for i := 1; i < len(line); i += 4 {
				if string(line[i]) >= "A" && string(line[i]) <= "Z" {
					crates[i/4] = append([]string{string(line[i])}, crates[i/4]...)
				}
			}
		}

		if line == "" {
			instructions = true
			continue
		}

		if instructions == true {
			numOfCrates, _ := strconv.Atoi(strings.TrimSpace(line[5:7]))

			from, _ := strconv.Atoi(strings.TrimSpace(line[12:14]))

			to, _ := strconv.Atoi(strings.TrimSpace(line[17:]))

			crates = move(crates, numOfCrates, from-1, to-1)
		}

	}

	for i := 0; i < len(crates); i++ {
		fmt.Print(crates[i][len(crates[i])-1])
	}
}

// func main() {

// 	crates := [][]string{
// 		{"B", "G", "S", "C"},
// 		{"T", "M", "W", "H", "J", "N", "V", "G"},
// 		{"M", "Q", "S"},
// 		{"B", "S", "L", "T", "W", "N", "M"},
// 		{"J", "Z", "F", "T", "V", "G", "W", "P"},
// 		{"C", "T", "B", "G", "Q", "H", "S"},
// 		{"T", "J", "P", "B", "W"},
// 		{"G", "D", "C", "Z", "F", "T", "Q", "M"},
// 		{"N", "S", "H", "B", "P", "F"},
// 	}

// 	// crates := [][]string{
// 	// 	{"Z", "N"},
// 	// 	{"M", "C", "D"},
// 	// 	{"P"},
// 	// }

// 	file, err := os.Open("day5.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {

// 		line := scanner.Text()

// 		numOfCrates, _ := strconv.Atoi(strings.TrimSpace(line[5:7]))

// 		from, _ := strconv.Atoi(strings.TrimSpace(line[12:14]))

// 		to, _ := strconv.Atoi(strings.TrimSpace(line[17:]))

// 		crates = move(crates, numOfCrates, from-1, to-1)

// 	}

// 	// for _, v := range crates {
// 	// 	fmt.Println(v)
// 	// }

// 	for i := 0; i < len(crates); i++ {
// 		fmt.Print(crates[i][len(crates[i])-1])
// 	}
// }

func move(crates [][]string, numOfCrates, from, to int) [][]string {

	//part 1
	// for i := 0; i < numOfCrates; i++ {
	// 	elementMoved := crates[from][len(crates[from])-1]
	// 	crates[to] = append(crates[to], elementMoved)
	// 	crates[from] = crates[from][:len(crates[from])-1]
	// }
	// return crates

	// //part 2

	fromStack := crates[from]

	toStack := crates[to]

	elementsMoved := fromStack[len(fromStack)-numOfCrates:]

	toStack = append(toStack, elementsMoved...)

	fromStack = fromStack[:len(fromStack)-numOfCrates]

	crates[from] = fromStack

	crates[to] = toStack

	return crates
}
