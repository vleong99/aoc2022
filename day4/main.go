package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	counter := 0

	file, err := os.Open("day4.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		eachRange := strings.Split(line, ",")

		var eachNum [][]string

		eachNum = append(eachNum, strings.Split(eachRange[0], "-"))

		eachNum = append(eachNum, strings.Split(eachRange[1], "-"))

		//part 1
		// if contained(eachNum) {
		// 	counter++
		// }

		if overlap(eachNum) {
			counter++
		}
	}

	fmt.Println(counter)

}

//part 1

// func contained(a [][]string) bool {
// 	var nums []int
// 	for _, b := range a {
// 		for _, v := range b {
// 			x, _ := strconv.Atoi(v)
// 			nums = append(nums, x)
// 		}
// 	}
// 	if (nums[0] <= nums[2] && nums[1] >= nums[3]) || (nums[0] >= nums[2] && nums[1] <= nums[3]) {
// 		return true
// 	}
// 	return false
// }

//part 2

func overlap(a [][]string) bool {
	var nums []int
	for _, b := range a {
		for _, v := range b {
			x, _ := strconv.Atoi(v)
			nums = append(nums, x)
		}
	}

	if nums[0] > nums[2] {
		nums[0], nums[1], nums[2], nums[3] = nums[2], nums[3], nums[0], nums[1]
	}

	fmt.Println(nums)

	if nums[2] <= nums[1] {
		return true
	}

	return false
}
