package main

import (
	utils "aoc2022/libs"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type monkey struct {
	num int
	p   []string //parents
	op  string   //operator
}

func main() {

	allM := make(map[string]monkey) // all monkeys

	var names []string //monkey names
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		name := line[0][:4]
		names = append(names, name)
		if len(line) == 2 { //only num
			n := utils.Ints(line[1])
			allM[name] = monkey{n, []string{}, "NA"}
		}
		if len(line) == 4 {
			allM[name] = monkey{-10000000, []string{line[1], line[3]}, line[2]}
		}
	}

	allYelled := false

	for allYelled == false {
		allYelled = true
		for i := 0; i < len(names); i++ {
			v := allM[names[i]]
			if v.num == -10000000 && len(v.p) == 2 && allM[v.p[0]].num != -10000000 && allM[v.p[1]].num != -10000000 { //can calc
				n1 := allM[v.p[0]].num
				n2 := allM[v.p[1]].num
				switch v.op {
				case "+":
					v.num = n1 + n2
				case "-":
					v.num = n1 - n2
				case "*":
					v.num = n1 * n2
				case "/":
					v.num = n1 / n2
				}
				allM[names[i]] = v
			}
			if allM[names[i]].num == -10000000 {
				allYelled = false
			}
		}
	}

	currentMonkey := "humn"

	trace := []string{"humn"}
	for currentMonkey != "root" {
		for i := 0; i < len(names); i++ {
			v := allM[names[i]]
			if len(v.p) == 2 && (v.p[0] == currentMonkey || v.p[1] == currentMonkey) {
				trace = append(trace, names[i])
				currentMonkey = names[i]
			}
		}
	}

	var ans int
	if trace[len(trace)-2] == allM["root"].p[0] { // ans is 1
		ans = allM[allM["root"].p[1]].num
	} else {
		ans = allM[allM["root"].p[0]].num
	}

	for i := len(trace) - 2; i > 0; i-- {
		curr := allM[trace[i]]
		prev := trace[i-1]
		if curr.p[0] == prev {
			fixed := allM[curr.p[1]].num
			switch curr.op {
			case "+":
				ans -= fixed
			case "-":
				ans += fixed
			case "*":
				ans /= fixed
			case "/":
				ans *= fixed
			}

		}
		if curr.p[1] == prev {
			fixed := allM[curr.p[0]].num
			switch curr.op {
			case "+":
				ans -= fixed
			case "-":
				ans = fixed - ans
			case "*":
				ans /= fixed
			case "/":
				ans = fixed / ans
			}
		}

	}

	fmt.Println(ans)

}
