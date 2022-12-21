package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	var nums []int

	var order [][]int //original index, number

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		integar := Ints(scanner.Text()) * 811589153
		order = append(order, []int{len(nums), integar})
		nums = append(nums, integar)
	}

	length := len(nums)

	for a := 0; a < 10; a++ {
		for i := 0; i < len(nums); i++ {
			for j := 0; j < len(order); j++ {
				if order[j][0] != i {
					continue
				}
				curr := []int{order[j][0], order[j][1]}
				currNum := order[j][1]
				if currNum == 0 {
					break
				}
				currIndex := j

				//find new index

				newIndex := mod(currIndex+currNum, length-1)

				order = append(order[:j], order[j+1:]...)
				order = append(order[:newIndex+1], order[newIndex:]...)
				order[newIndex] = curr

				break
			}
		}
	}

	zeroIndex := 0

	for i := 0; i < len(order); i++ {
		if order[i][1] == 0 {
			zeroIndex = i
		}
	}
	index1000 := ((1000 % length) + zeroIndex) % length
	index2000 := ((2000 % length) + zeroIndex) % length
	index3000 := ((3000 % length) + zeroIndex) % length

	fmt.Printf("ans:%v\n", order[index1000][1]+order[index2000][1]+order[index3000][1])

}

func Ints(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func mod(a, b int) int {
	return (a%b + b) % b
}

// if newIndex > length {
// 	newIndex %= (length - 1)
// } else if newIndex <= 0 {
// 	newIndex = length - ((-newIndex) % length) - 1
// } else {
// 	newIndex %= length - 1
// }
