package main

import (
	utils "aoc2022/libs"
	"bufio"
	"fmt"
	"os"
)

func main() {

	allAns := []int{}

	total := 0
	d := 1 //denominator

	file, err := os.Open("day25.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ans := 0
		line := scanner.Text()
		for i := len(line) - 1; i >= 0; i-- {
			switch line[i] {
			case '-':
				ans += -1 * d
			case '=':
				ans += -2 * d
			default:
				ans += utils.Ints(string(line[i])) * d
			}
			d *= 5
		}
		d = 1
		allAns = append(allAns, ans)
		total += ans
	}

	fmt.Println(total)
	temp := []string{}

	d = 5
	var carry bool

	options := "012=-"

	for total > 0 {
		r := total % d //remainder
		if carry == true {
			r += 1
			carry = false
		}
		temp = append(temp, string(options[r%5]))
		if r >= 3 && r <= 5 {
			carry = true
		}
		total /= d
	}
	for i := len(temp) - 1; i >= 0; i-- {
		fmt.Print(temp[i])
	}
	fmt.Println()

}
