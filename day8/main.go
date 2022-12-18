package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {

	var trees [][]int

	// visible := 0

	file, err := os.Open("day8.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		var lineOfTrees []int

		for i := 0; i < len(line); i++ {
			n, _ := strconv.Atoi(string(line[i]))
			lineOfTrees = append(lineOfTrees, n)
		}

		trees = append(trees, lineOfTrees)
	}

	width := len(trees[0])

	height := len(trees)

	//part 1
	// visible = width*2 + height*2 - 4

	// for i := 1; i < height-1; i++ {
	// 	for j := 1; j < width-1; j++ {
	// 		if checkVertical(i, j, trees) == true || checkHorizontal(i, j, trees) == true {
	// 			visible++
	// 		}
	// 	}
	// }

	// fmt.Println(visible)

	//part 2

	var scores []int

	for i := 1; i < height-1; i++ {
		for j := 1; j < width-1; j++ {
			s := calculateScore(i, j, trees)
			scores = append(scores, s)
		}
	}

	sort.Ints(scores)

	fmt.Println(scores)

}

// func checkVertical(height, width int, trees [][]int) bool {
// 	num := trees[height][width]
// 	//check top

// 	falses := 0

// 	for i := 0; i < height; i++ {
// 		if trees[i][width] >= num && i != height {
// 			falses += 1
// 			break
// 		}
// 	}

// 	//check bottom

// 	for i := height + 1; i < len(trees); i++ {
// 		if trees[i][width] >= num && i != height {
// 			falses += 1
// 			break
// 		}
// 	}

// 	if falses == 2 {
// 		return false
// 	}
// 	return true

// }

// func checkHorizontal(height, width int, trees [][]int) bool {
// 	num := trees[height][width]

// 	falses := 0

// 	//check left

// 	for i := 0; i < width; i++ {
// 		if trees[height][i] >= num && i != width {
// 			falses += 1
// 			break
// 		}
// 	}

// 	//check right

// 	for i := width + 1; i < len(trees[0]); i++ {
// 		if trees[height][i] >= num && i != width {
// 			falses += 1
// 			break
// 		}
// 	}
// 	if falses == 2 {
// 		return false
// 	}
// 	return true
// }

func calculateScore(height, width int, trees [][]int) int {
	num := trees[height][width]

	//left trees

	left := 0

	for i := width - 1; i >= 0; i-- {
		if trees[height][i] >= num {
			left++
			break
		} else {
			left++
		}
	}

	fmt.Println(left)

	//right trees

	right := 0

	for i := width + 1; i < len(trees[0]); i++ {
		if trees[height][i] >= num {
			right++
			break
		} else {
			right++
		}
	}

	fmt.Println(right)

	//top trees

	top := 0

	for i := height - 1; i >= 0; i-- {
		if trees[i][width] >= num {
			top++
			break
		} else {
			top++
		}
	}

	fmt.Println(top)

	//bottom trees

	bottom := 0

	for i := height + 1; i < len(trees); i++ {
		if trees[i][width] >= num {
			bottom++
			break
		} else {
			bottom++
		}
	}

	fmt.Println(bottom)

	fmt.Println(top * bottom * left * right)

	fmt.Println()
	return top * bottom * left * right

}
