package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items     []int
	operator  string
	operation int
	test      int
	testT     *monkey
	testF     *monkey
	inspected int
}

func main() {

	//set up monkeys

	var allM []*monkey

	for i := 0; i < 8; i++ {
		m := monkey{}
		allM = append(allM, &m)
	}

	file, err := os.Open("day11.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var currentM *monkey

	supermodulo := 1

	for scanner.Scan() {

		line := strings.Split(strings.TrimSpace(scanner.Text()), " ")

		switch line[0] {
		case "Monkey":
			n, _ := strconv.Atoi(string(line[1][0]))
			currentM = allM[n]
		case "Starting":
			for i := 2; i < len(line); i++ {
				n, _ := strconv.Atoi(line[i][0:2])
				currentM.items = append(currentM.items, n)
			}
		case "Operation:":
			n, _ := strconv.Atoi(line[len(line)-1])
			currentM.operation = n
			if line[len(line)-2] == "+" {
				currentM.operator = "add"
			} else {
				currentM.operator = "times"
			}
		case "Test:":
			n, _ := strconv.Atoi(line[len(line)-1])
			currentM.test = n
			supermodulo *= n
		case "If":
			if line[1] == "true:" {
				n, _ := strconv.Atoi(line[len(line)-1])
				currentM.testT = allM[n]
			} else {
				n, _ := strconv.Atoi(line[len(line)-1])
				currentM.testF = allM[n]
			}
		}

	}

	//run rounds

	for i := 0; i < 10000; i++ {
		//each round
		for j := 0; j < len(allM); j++ {
			cM := allM[j]
			turn(cM, supermodulo)
		}
	}

	var allSums []int
	for i := 0; i < len(allM); i++ {
		allSums = append(allSums, allM[i].inspected)

	}
	sort.Ints(allSums)
	fmt.Println(allSums)
}

func turn(m *monkey, s int) {
	for i := 0; i < len(m.items); i++ {
		m.inspected++
		num := m.items[i]
		if m.operator == "add" {
			if m.operation == 0 {
				num += num
			} else {
				num += m.operation
			}
		} else {
			if m.operation == 0 {
				num *= num
			} else {
				num *= m.operation
			}
		}
		num %= s
		if num%m.test == 0 {
			(m.testT).items = append((m.testT).items, num)
		} else {
			(m.testF).items = append((m.testF).items, num)
		}
	}
	m.items = []int{}
}
