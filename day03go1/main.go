package main

import (
	"fmt"
	"strconv"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day03go1/input.txt")
	fmt.Println(power(input))
}

func power(ss []string) int64 {
	data := make(map[int]map[rune]int, len(ss[0]))
	for j, s := range ss {
		for i, r := range s {
			if j == 0 {
				data[i] = make(map[rune]int, 2)
			}
			data[i][r]++
		}
	}
	var gamma, epsilon string
	for i := range ss[0] {
		g, e := split(data[i])
		gamma += g
		epsilon += e
	}
	g, _ := strconv.ParseInt(gamma, 2, 32)
	e, _ := strconv.ParseInt(epsilon, 2, 32)
	return g * e
}

func split(bits map[rune]int) (string, string) {
	if bits['0'] > bits['1'] {
		return "0", "1"
	}
	return "1", "0"
}
