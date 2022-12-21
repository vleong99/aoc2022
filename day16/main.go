package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	utils "../libs"
)

var nonZeroValves []string

var edges map[string]map[string]int

var allValves map[string]valve

var allScores []int

type valve struct {
	flowRate   int
	neighbours []string
}

func main() {
	allValves = readInput("day16.txt")
	edges = floydwarshall(allValves)

	for k, v := range allValves {
		if v.flowRate > 0 {
			nonZeroValves = append(nonZeroValves, k)
		}
	}

	// part 1
	// allPaths := part1("AA", []string{}, 30)

	// for _, v := range allPaths {
	// 	allScores = append(allScores, score1(v))
	// }

	// sort.Ints(allScores)
	// fmt.Println(allScores[len(allScores)-1])

	//part2

	allPaths := part1("AA", []string{}, 26)

	part2(allPaths)

}

func readInput(s string) map[string]valve {
	allValves := make(map[string]valve)
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.Split(strings.TrimPrefix(scanner.Text(), "Valve "), " has flow rate=")
		line = append([]string{line[0]}, strings.Split(line[1], "; tunnels lead to valves ")...)
		if len(line) < 3 {
			line = append([]string{line[0]}, strings.Split(line[1], "; tunnel leads to valve ")...)
		}
		neighbours := strings.Split(strings.TrimPrefix(line[2], "s "), ", ")
		allValves[line[0]] = valve{utils.Ints(line[1]), neighbours}
	}
	return allValves
}

func floydwarshall(allValves map[string]valve) map[string]map[string]int {
	edges := make(map[string]map[string]int)

	for k, _ := range allValves {
		edges[k] = make(map[string]int)
		for k2, _ := range allValves {
			edges[k][k2] = 100000
		}
	}

	for k, v := range allValves {
		edges[k][k] = 0
		for _, n := range v.neighbours {
			edges[k][n] = 1
		}
	}

	for k := range allValves {
		for i := range allValves {
			for j := range allValves {
				if edges[i][j] > edges[i][k]+edges[k][j] {
					edges[i][j] = edges[i][k] + edges[k][j]
				}
			}
		}
	}

	return edges

}

func part1(pos string, openValves []string, timeLeft int) [][]string {

	var options [][]string
	for _, next := range nonZeroValves {
		if !contains(next, openValves) && edges[pos][next] <= timeLeft {
			openValves = append(openValves, next)
			options = append(options, part1(next, openValves, timeLeft-edges[pos][next]-1)...)
			openValves = openValves[:len(openValves)-1]
		}
	}
	options = append(options, append([]string{}, openValves...))
	return options
}

func score1(path []string) int {
	ans := 0
	timeleft := 30
	now := "AA"
	for i := 0; i < len(path); i++ {
		if timeleft <= 0 {
			return ans
		}
		timeleft -= edges[now][path[i]] + 1
		flow := allValves[path[i]].flowRate
		ans += flow * timeleft
		now = path[i]
	}
	return ans
}

func score2(path []string) int {
	ans := 0
	timeleft := 26
	now := "AA"
	for i := 0; i < len(path); i++ {
		if timeleft <= 0 {
			return ans
		}
		timeleft -= edges[now][path[i]] + 1
		flow := allValves[path[i]].flowRate
		ans += flow * timeleft
		now = path[i]
	}
	return ans
}

func contains(s string, sl []string) bool {
	for _, v := range sl {
		if s == v {
			return true
		}
	}
	return false
}

func part2(allPaths [][]string) {

	bestScores := make(map[string]int)

	for _, path := range allPaths {
		score := score2(path)

		sort.Strings(path)
		var pathString string
		for i := 0; i < len(path); i++ {
			pathString += path[i]
		}

		if bestScores[pathString] < score {
			bestScores[pathString] = score
		}
	}

	var maxScores []int

	for k, v := range bestScores {
		for k2, v2 := range bestScores {
			if overlap2(k, k2) == false {
				ans := v + v2
				maxScores = append(maxScores, ans)
			}
		}

	}

	sort.Ints(maxScores)
	fmt.Println(maxScores[len(maxScores)-1])
}

func overlap1(a, b []string) bool {
	for _, i := range a {
		for _, j := range b {
			if i == j {
				return true
			}
		}
	}
	return false
}

func overlap2(a, b string) bool {
	for i := 0; i < len(a); i += 2 {
		for j := 0; j < len(b); j += 2 {
			if a[i:i+2] == b[j:j+2] {
				return true
			}
		}
	}
	return false
}
