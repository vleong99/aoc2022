package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	l := strToList("[[0,5,[[],[],2,[7,9]]],[[8,8,3,[6,3,8,9,1]],[2,0,10,7,10],4,10,[9,[1,8],4,[4,0,5,10],[4,0,8,8]]],[9,10],[],[0,[[]],4,10]]")
	for _, v := range l {
		fmt.Println(v)
	}
}

func strToList(s string) []any {
	l := make([]any, 0)
	json.Unmarshal([]byte(s), &l)
	return l
}
