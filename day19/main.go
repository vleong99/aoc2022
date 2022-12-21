package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type blueprint struct {
	id          int
	orebot      int
	claybot     int
	obsidianbot []int
	geodebot    []int
}

func main() {

	var allBlueprints []blueprint

	file, err := os.Open("day19.txt")
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var a blueprint
		line := strings.Split(scanner.Text(), " ")
		a.id = Ints(strings.TrimSuffix(line[1], ":"))
		a.orebot = Ints(line[6])
		a.claybot = Ints(line[12])
		a.obsidianbot = []int{Ints(line[18]), Ints(line[21])}
		a.geodebot = []int{Ints(line[27]), Ints(line[30])}
		allBlueprints = append(allBlueprints, a)

	}
	ans := 1

	for i := 0; i < len(allBlueprints); i++ {
		fmt.Println(allBlueprints[i])
		a := calculate(allBlueprints[i])

		fmt.Println(a)
		ans *= a

	}
	fmt.Println(ans)
}

func Ints(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func calculate(b blueprint) int {
	var queue [][9]int

	seen := make(map[[9]int]bool)
	iS := [9]int{0, 0, 0, 0, 1, 0, 0, 0, 32} // initial state: ore, clay, obsidian, geode, orebot, claybot, obbot, geobot, time

	queue = append(queue, iS)

	best := []int{0, 32}

	for len(queue) > 0 {
		curr := queue[0]

		queue = queue[1:]
		o, c, ob, g, r1, r2, r3, r4, t := curr[0], curr[1], curr[2], curr[3], curr[4], curr[5], curr[6], curr[7], curr[8]

		if seen[curr] {
			continue
		}
		if g > best[0] || best[0] == g && best[1] <= t {
			best = []int{g, t}
		}
		if g < best[0] {
			continue
		}

		if t == 0 {
			continue
		}

		seen[curr] = true
		if b.geodebot[0] <= o && b.geodebot[1] <= ob {
			queue = append(queue, [9]int{o + r1 - b.geodebot[0], c + r2, ob + r3 - b.geodebot[1], g + r4, r1, r2, r3, r4 + 1, t - 1})
			continue
		}
		queue = append(queue, [9]int{o + r1, c + r2, ob + r3, g + r4, r1, r2, r3, r4, t - 1})
		if b.obsidianbot[0] <= o && b.obsidianbot[1] <= c {
			queue = append(queue, [9]int{o + r1 - b.obsidianbot[0], c + r2 - b.obsidianbot[1], ob + r3, g + r4, r1, r2, r3 + 1, r4, t - 1})
		}
		if b.claybot <= o {
			queue = append(queue, [9]int{o + r1 - b.claybot, c + r2, ob + r3, g + r4, r1, r2 + 1, r3, r4, t - 1})
		}
		if b.orebot <= o {
			queue = append(queue, [9]int{o + r1 - b.orebot, c + r2, ob + r3, g + r4, r1 + 1, r2, r3, r4, t - 1})
		}
	}
	return best[0]

}
