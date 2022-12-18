package main

import (
	utils "aoc2022/libs"
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {

	var allSB [][]int //all sensors and beacones
	//figs is sCol, sRow, bCol, bRow

	file, err := os.Open("day15.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// for scanner.Scan() {
	// 	line := strings.Split(scanner.Text(), ": closest beacon is at x=")
	// 	line[0] = strings.TrimPrefix(line[0], "Sensor at x=")
	// 	s := strings.Split(line[0], ", y=")
	// 	b := strings.Split(line[1], ", y=")
	// 	sCol, sRow, bCol, bRow := utils.Ints(s[0]), utils.Ints(s[1]), utils.Ints(b[0]), utils.Ints(b[1])

	// 	tDiff := int(math.Abs(float64(bRow-sRow)) + math.Abs(float64(bCol-sCol)))

	// 	topOfD := sRow - tDiff

	// 	bottomOfD := sRow + tDiff

	// 	if 2000000 >= topOfD && 2000000 <= bottomOfD { //within sensor range
	// 		rowDiff := int(math.Abs(float64(sRow - 2000000)))
	// 		width := tDiff - rowDiff
	// 		eachRange := []int{sCol - width, sCol + width}
	// 		allRanges = append(allRanges, eachRange)
	// 	}
	// }

	//part 2

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ": closest beacon is at x=")
		line[0] = strings.TrimPrefix(line[0], "Sensor at x=")
		s := strings.Split(line[0], ", y=")
		b := strings.Split(line[1], ", y=")
		sCol, sRow, bCol, bRow := utils.Ints(s[0]), utils.Ints(s[1]), utils.Ints(b[0]), utils.Ints(b[1])

		n := []int{sCol, sRow, bCol, bRow}
		allSB = append(allSB, n)
	}

	for i := 0; i <= 2000000; i++ {
		findRanges(allSB, 2000000+i)
		findRanges(allSB, 2000000-i)
	}
	// for i := 0; i <= 20; i++ {
	// 	findRanges(allSB, i)
	// }

}

func findRanges(allSB [][]int, num int) {
	var allRanges [][]int

	for _, v := range allSB {
		sCol := v[0]
		sRow := v[1]
		bCol := v[2]
		bRow := v[3]

		tDiff := int(math.Abs(float64(bRow-sRow)) + math.Abs(float64(bCol-sCol)))

		topOfD := sRow - tDiff

		bottomOfD := sRow + tDiff

		if num >= topOfD && num <= bottomOfD { //within sensor range
			rowDiff := int(math.Abs(float64(sRow - num)))
			width := tDiff - rowDiff
			eachRange := []int{sCol - width, sCol + width}
			allRanges = append(allRanges, eachRange)
		}
	}

	sort.Slice(allRanges, func(i, j int) bool {
		return allRanges[i][0] < allRanges[j][0]
	})

	// notBeacon(allRanges)
	findBeacon(allRanges, num)

}

func notBeacon(allRanges [][]int) {
	ans := 0

	newRange := []int{0, 0}

	for i := 0; i < len(allRanges); i++ {
		if i == 0 {
			newRange = []int{allRanges[i][0], allRanges[i][1]}
			continue
		}

		if i == len(allRanges)-1 {

			if allRanges[i][1] > newRange[1] {
				newRange[1] = allRanges[i][1]
			}

			ans += newRange[1] - newRange[0] + 1

			if newRange[1] >= 1749091 && newRange[0] <= 1749091 {
				ans--
			}
			break
		}
		if allRanges[i][1] <= newRange[1] {
			continue
		}
		if allRanges[i][1] > newRange[1] && allRanges[i][0] <= newRange[1] {

			newRange[1] = allRanges[i][1]

		}
		if allRanges[i][0] > newRange[1] {
			ans += newRange[1] - newRange[0] + 1

			if newRange[1] >= 1749091 && newRange[0] <= 1749091 {
				ans--
			}

			newRange = []int{0, 0}
		}
	}
	fmt.Println(ans)
}

func findBeacon(allRanges [][]int, num int) {

	newRange := []int{0, 0}

	for i := 0; i < len(allRanges); i++ {

		//set newRange
		if i == 0 {
			newRange = []int{allRanges[i][0], allRanges[i][1]}
			continue
		}
		//x boundaries
		if allRanges[i][1] < 0 || allRanges[i][0] > 4000000 {
			continue
		}

		//if current range already encompassed
		if allRanges[i][1] <= newRange[1] {
			continue
		}

		//update largest num
		if allRanges[i][1] > newRange[1] && (allRanges[i][0]-newRange[1]) <= 1 {

			newRange[1] = allRanges[i][1]
		}

		//if there is a gap

		if allRanges[i][0]-newRange[1] > 1 {
			fmt.Println(num)
			fmt.Println(newRange[1], allRanges[i][0])
			fmt.Println(num + (newRange[1]+1)*4000000)
			return
		}

		//if iterated till the end
		if i == len(allRanges)-1 {
			return
		}

	}
}
