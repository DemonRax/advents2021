package main

import (
	"advents2021/util"
	"fmt"
	"strings"
)

func main() {
	input := util.ReadFile("./day08go1/input.txt")
	fmt.Println(appear(input))
}

func appear(ss []string) int {
	var count int
	for _, s := range ss {
		_, right := scan(s)
		count += findEasy(right)
	}
	return count
}

func scan(s string) ([]string, []string) {
	split := strings.Split(s, " | ")
	return strings.Split(split[0], " "), strings.Split(split[1], " ")
}

var easy = map[int]interface{}{
	2: nil,
	3: nil,
	4: nil,
	7: nil,
}

func findEasy(ss []string) int {
	var count int
	for _, s := range ss {
		_, ok := easy[len(s)]
		if ok {
			count++
		}
	}
	return count
}
