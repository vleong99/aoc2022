package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	score := 0

	file, err := os.Open("day2.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		score += outcome(line[0], line[2])
	}

	fmt.Println(score)

}

// part 1

// func outcome(a, b byte) int {
// 	value := 0

//add value of shape

// 	if b == 'X' {
// 		value += 1
// 	}

// 	if b == 'Y' {
// 		value += 2
// 	}

// 	if b == 'Z' {
// 		value += 3
// 	}

// 	if a == 'A' {
// 		if b == 'Y' {
// 			value += 6
// 		} else if b == 'X' {
// 			value += 3
// 		} else {
// 			value += 0
// 		}
// 	}

// 	if a == 'B' {
// 		if b == 'Z' {
// 			value += 6
// 		} else if b == 'Y' {
// 			value += 3
// 		} else {
// 			value += 0
// 		}
// 	}

// 	if a == 'C' {
// 		if b == 'X' {
// 			value += 6
// 		} else if b == 'Z' {
// 			value += 3
// 		} else {
// 			value += 0
// 		}
// 	}

// 	return value
// }

//part 2

func outcome(a, b byte) int {
	value := 0

	switch b {
	case 'X':
		value += 0
	case 'Y':
		value += 3
	case 'Z':
		value += 6
	}

	// lose
	if b == 'X' {
		switch a {
		case 'A':
			value += 3
		case 'B':
			value += 1
		case 'C':
			value += 2
		}
	}

	//draw

	if b == 'Y' {
		switch a {
		case 'A':
			value += 1
		case 'B':
			value += 2
		case 'C':
			value += 3
		}
	}

	//win

	if b == 'Z' {
		switch a {
		case 'A':
			value += 2
		case 'B':
			value += 3
		case 'C':
			value += 1
		}
	}

	// A for Rock, B for Paper, and C for Scissors. 1 for Rock, 2 for Paper, and 3 for Scissors

	return value
}
