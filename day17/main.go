package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

//initialise basic blocks
var blocks [][][]int

var one, two, three, four, five [][]int

//initialise grid of fallen rocks
var grid map[[2]int]int

//initialise slice of states

var states [][]int

var heights []int

func main() {
	one = [][]int{{0, 0}, {1, 0}, {2, 0}, {3, 0}}
	two = [][]int{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}
	three = [][]int{{0, 0}, {1, 0}, {2, 0}, {2, 1}, {2, 2}}
	four = [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}
	five = [][]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}}

	blocks = [][][]int{one, two, three, four, five}

	file, err := os.Open("day17.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	//jetcycle
	ss, _ := ioutil.ReadAll(file)
	s := 0

	grid = make(map[[2]int]int)

	for i := 0; i < 7; i++ {
		grid[[2]int{i, 0}] = 1
	}

	// ans := 0
	//each round
	for i := 0; i <= 1000000000000; i++ {
		//check height
		h := 0
		for k := range grid {
			if k[1] > h {
				h = k[1]
			}
		}

		if i == 2022 {
			fmt.Println(h)
		}
		//check current block
		b := blocks[i%5]

		//store state - block no, jetcycle index, height

		// if i == 61400 {
		// 	fmt.Println(h)
		// 	return
		// }
		if i == 4442 {
			x := 1000000000000
			repeat := (x - 2746) / 1695
			modulo := (x - 2746) % 1695

			ans := 4203 + repeat*2610 + (states[modulo][2] - states[0][2])

			fmt.Println(ans)
			fmt.Println("hi")
			return

		}
		if i >= 2746 && i <= 4441 {
			newState := []int{i % 5, s, h}

			states = append(states, newState)

			//find height at x = 40036+1

			// repeat := (40037 - 2746) / 1695

			// ans = 4203 + repeat*2610

			// fmt.Println(ans)
			// i = 40036
			// continue

		}
		// newState := []int{i % 5, s}

		// fullFloor := true
		// for i := 0; i < 7; i++ {
		// 	if grid[[2]int{i, h}] != 1 {
		// 		fullFloor = false
		// 	}
		// }
		// if contains(newState) && fullFloor {
		// 	fmt.Println("hi")
		// 	fmt.Println(i)
		// 	fmt.Println(h)
		// 	for i := 0; i < 7; i++ {
		// 		fmt.Print(grid[[2]int{i, h - 2}])
		// 	}
		// 	fmt.Println()
		// 	for i := 0; i < 7; i++ {
		// 		fmt.Print(grid[[2]int{i, h - 1}])
		// 	}
		// 	fmt.Println()

		// 	for i := 0; i < 7; i++ {
		// 		fmt.Print(grid[[2]int{i, h}])
		// 	}
		// 	fmt.Println()

		// 	// return

		// 	// //number of times the cycle repeats
		// 	// repeat := 1000000000000 / len(states)
		// 	// modulo := 1000000000000 % len(states)
		// 	// ans := repeat*newState[3] + states[modulo-1][3]
		// 	// fmt.Println(ans)
		// 	// return
		// }

		//initialise starting position

		x, y := 2, h+4

		//simulate current rock falling
		for {

			//intialise offset left or right
			dx := 0
			if ss[s] == '>' {
				dx = 1
			} else {
				dx = -1
			}
			s = (s + 1) % len(ss)

			//check if offset valid
			valid := true
			for _, v := range b {
				if v[0]+dx+x < 0 || v[0]+dx+x >= 7 || grid[[2]int{v[0] + dx + x, v[1] + y}] == 1 {
					valid = false
				}
			}
			if valid == true {
				x += dx
			}

			//check if rock has landed on another rock

			haslanded := false
			for _, coords := range b {
				x2 := coords[0]
				y2 := coords[1]

				if grid[[2]int{x + x2, y + y2 - 1}] == 1 {
					haslanded = true
					break
				}
			}
			if haslanded == true {

				break
			}
			y -= 1

		}

		//once broken out of cycle, rock has fallen, append rocks to grid
		for _, coords := range b {
			x2 := coords[0]
			y2 := coords[1]

			grid[[2]int{x + x2, y + y2}] = 1
		}

	}

}

func contains(newState []int) bool {
	for _, v := range states {
		if newState[0] == v[0] && newState[1] == v[1] {
			return true
		}
	}
	return false
}
