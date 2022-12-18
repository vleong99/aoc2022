package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	cals := 0

	var elves []int

	file, err := os.Open("day1-1.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, cals)
			cals = 0
		} else {
			num, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println(err)
			}
			cals += num
		}
	}

	sort.Ints(elves)

	// day1
	// fmt.Println(elves[len(elves)-1])

	//day2

	length := len(elves) - 1

	fmt.Println(elves[length] + elves[length-1] + elves[length-2])
}
