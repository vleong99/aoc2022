package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type directory struct {
	name   string
	files  []int
	dir    []*directory
	parent *directory
}

func main() {

	var allDirs []*directory

	var parentDir *directory

	currentDir := &directory{"/", []int{}, []*directory{}, parentDir}

	allDirs = append(allDirs, currentDir)

	file, err := os.Open("day7.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")

		fmt.Println(line)

		//contents
		if line[0] != "$" { //list dirs
			if line[0] == "dir" {
				dirName := line[1]
				newDir := directory{dirName, []int{}, []*directory{}, currentDir}
				currentDir.dir = append(currentDir.dir, &newDir)
				allDirs = append(allDirs, &newDir)
				continue
			} else { //list files
				num, _ := strconv.Atoi(line[0])
				currentDir.files = append(currentDir.files, num)
				continue
			}
		} else {
			if line[1] == "ls" {
				continue
			}
			if line[1] == "cd" {
				if line[2] == "/" {
					currentDir = allDirs[0]
				}
				if line[2] == ".." { // move up one dir
					currentDir = parentDir
					parentDir = currentDir.parent
				}
				for _, v := range currentDir.dir {
					if v.name == line[2] {
						parentDir = currentDir
						currentDir = v
					}
				}
			}
		}
	}

	//part 1
	// sum := 0
	// for _, v := range allDirs {
	// 	total := add(*v)
	// 	if total <= 100000 {
	// 		sum += total
	// 	}
	// }
	// fmt.Println(sum)

	//part 2

	var allTotals []int

	for _, v := range allDirs {
		total := add(*v)
		allTotals = append(allTotals, total)
	}

	sort.Ints(allTotals)

	spaceLeft := 70000000 - allTotals[len(allTotals)-1]

	spaceNeeded := 30000000 - spaceLeft

	fmt.Println(spaceNeeded)

	for i := 0; i < len(allTotals); i++ {
		if allTotals[i] >= spaceNeeded {
			fmt.Println(allTotals[i])
			return
		}
	}

}

func add(d directory) int {
	count := 0
	for _, v := range d.files {
		count += v
	}
	for _, v := range d.dir {
		count += add(*v)
	}
	return count
}
