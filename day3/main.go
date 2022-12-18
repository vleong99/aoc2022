package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	ans := 0

	file, err := os.Open("day3.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	//part 1
	// for scanner.Scan() {
	// 	line := scanner.Text()

	// 	a := []rune(line[:len(line)/2])

	// 	b := []rune(line[len(line)/2:])

	// 	duplicate := match(a, b)

	// 	if duplicate >= 97 && duplicate <= 122 {
	// 		ans += int(duplicate) - 96
	// 	} else {
	// 		ans += int(duplicate) - 38
	// 	}

	// }

	//part 2

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	for i := 0; i < len(lines); i += 3 {
		duplicate := match(lines[i], lines[i+1], lines[i+2])

		if duplicate >= 97 && duplicate <= 122 {
			ans += int(duplicate) - 96
		} else {
			ans += int(duplicate) - 38
		}

	}
	fmt.Println(ans)
}

// func match(a, b []rune) rune {
// 	for _, x := range a {
// 		for _, y := range b {
// 			if x == y {
// 				return x
// 			}
// 		}
// 	}
// 	return 0
// }

func match(a, b, c string) rune {
	for _, x := range a {
		for _, y := range b {
			if x == y {
				for _, z := range c {
					if y == z {
						return z
					}
				}
			}
		}
	}
	return 0
}
