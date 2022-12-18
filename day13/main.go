package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var packets []string

	a := ""

	b := ""

	index := 1

	ans := 0

	file, err := os.Open("day13.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		//compare a and b
		if scanner.Text() == "" {
			if compare(a, b) == 1 {
				fmt.Printf("true at index %v\n", index)
				ans += index
			} else {
				fmt.Printf("false at index %v\n", index)
			}
			a = ""
			b = ""
			index++
			continue
		} else {
			packets = append(packets, scanner.Text())
		}
		if a == "" {
			a = scanner.Text()
		} else {
			b = scanner.Text()
		}
	}

	fmt.Println(ans)

	part2(packets)
}

func part2(packets []string) {
	var index1, index2 int

	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	sort.Slice(packets, func(i, j int) bool {
		if compare(packets[i], packets[j]) == 1 {
			return true
		} else {
			return false
		}
	})

	for i := 0; i < len(packets); i++ {
		if packets[i] == "[[2]]" {
			index1 = i + 1
		}
		if packets[i] == "[[6]]" {
			index2 = i + 1
		}
	}

	fmt.Println(index1 * index2)
}
func compare(a, b string) int {

	//check if either empty
	if a == "[]" && b != "[]" {
		return 1
	}

	if a != "[]" && b == "[]" {
		return -1
	}

	a = a[1 : len(a)-1]

	b = b[1 : len(b)-1]

	// fmt.Println(a, b)

	var aSep, bSep []string

	var aElem, bElem string

	var opens, closes int

	//parse a
	for i := 0; i < len(a); i++ {
		// append if hit end of string
		if i == len(a)-1 {
			aElem += string(a[i])
			aSep = append(aSep, aElem)
			opens = 0
			closes = 0
			break
		}
		// account for opens and closes, add char to elem
		if a[i] != ',' {
			if a[i] == '[' {
				opens++
			}
			if a[i] == ']' {
				closes++
			}
			aElem += string(a[i])
		} else {
			if opens > closes {
				aElem += string(a[i])
				continue
			}
			if opens == closes {
				aSep = append(aSep, aElem)
				aElem = ""
				opens = 0
				closes = 0
				continue
			}
		}
	}

	for i := 0; i < len(b); i++ {
		// append if hit end of string

		if i == len(b)-1 {
			bElem += string(b[i])
			bSep = append(bSep, bElem)
			opens = 0
			closes = 0
			break
		}
		// account for opens and closes, add char to elem

		if b[i] != ',' {
			if b[i] == '[' {
				opens++
			}
			if b[i] == ']' {
				closes++
			}
			bElem += string(b[i])
		} else {
			if opens > closes {
				bElem += string(b[i])
				continue
			}
			if opens == closes {
				bSep = append(bSep, bElem)
				bElem = ""
				opens = 0
				closes = 0
				continue
			}
		}
	}

	if len(a) == 1 {
		aSep = []string{a}
	}

	if len(b) == 1 {
		bSep = []string{b}
	}

	// fmt.Println("aSep")
	// for _, v := range aSep {
	// 	fmt.Println(v)
	// }
	// fmt.Println()

	// fmt.Println("bSep")
	// for _, v := range bSep {
	// 	fmt.Println(v)
	// }
	// fmt.Println()

	length := len(aSep)

	for i := 0; i < length; i++ {

		// if a longer than b
		if i > len(bSep)-1 {
			return -1
		}

		if isInt(aSep[i]) && isInt(bSep[i]) { //both int

			// fmt.Println("both int")

			// fmt.Println(aSep[i], bSep[i])

			if len(bSep[i]) > len(aSep[i]) {
				return 1
			}

			if len(bSep[i]) < len(aSep[i]) {
				return -1
			}

			if bSep[i][0] > aSep[i][0] {
				return 1
			}

			if bSep[i][0] < aSep[i][0] {
				return -1
			}

			if bSep[i][0] == aSep[i][0] && i == length-1 && len(bSep) > length {
				return 1
			}

			// continue
		}
		if isInt(aSep[i]) == false && isInt(bSep[i]) == false { //both lists
			// fmt.Println("both list")
			// fmt.Println(aSep[i], bSep[i])
			if compare(aSep[i], bSep[i]) == 1 {
				return 1
			} else if compare(aSep[i], bSep[i]) == -1 {
				return -1
			}
			if i == length-1 && len(bSep) > length {
				return 1
			}
		}
		if isInt(aSep[i]) == false && isInt(bSep[i]) { //b is int
			//convert b to list
			// fmt.Println("b int")
			bSep[i] = "[" + bSep[i] + "]"
			// fmt.Println(aSep[i], bSep[i])

			if compare(aSep[i], bSep[i]) == 1 {
				return 1
			}
			if compare(aSep[i], bSep[i]) == -1 {
				return -1
			}
			if i == length-1 && len(bSep) > length {
				return 1
			}
		}
		if isInt(aSep[i]) && isInt(bSep[i]) == false { //a is int
			//convert a to list
			// fmt.Println("a int")
			aSep[i] = "[" + aSep[i] + "]"
			// fmt.Println(aSep[i], bSep[i])

			if compare(aSep[i], bSep[i]) == 1 {
				return 1
			}
			if compare(aSep[i], bSep[i]) == -1 {
				return -1
			}
			if i == length-1 && len(bSep) > length {
				return 1
			} else {
				continue
			}
		}

	}
	return 0
}

func isInt(s string) bool {
	if s[0] != '[' {
		return true
	}
	return false
}
