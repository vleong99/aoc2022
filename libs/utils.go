package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func Ints(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Max(i []int) int {
	sort.Ints(i)
	return i[len(i)-1]
}

func Min(i []int) int {
	sort.Ints(i)
	return i[0]
}

func Read(s string) *bufio.Scanner {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	return scanner
}

func GridInt(row, col, num int) [][]int {
	grid := [][]int{}
	for i := 0; i < row; i++ {
		line := []int{}
		for j := 0; j < col; j++ {
			line = append(line, num)
		}
		grid = append(grid, line)
	}
	return grid
}

func gridString(row, col int, s string) [][]string {
	grid := [][]string{}
	for i := 0; i < row; i++ {
		line := []string{}
		for j := 0; j < col; j++ {
			line = append(line, s)
		}
		grid = append(grid, line)
	}
	return grid
}

func copySofS(slice [][]any) [][]any {
	return slice
}
