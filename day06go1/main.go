package main

import (
	"advents2021/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("./day06go1/input.txt")
	fmt.Println(simulateRaw(input, 80))
}

func simulateRaw(ss []string, days int) int {
	s := strings.Split(ss[0], ",")
	ints := make([]int, len(s))
	for i, v := range s {
		ints[i], _ = strconv.Atoi(v)
	}
	return simulate(ints, days)
}

func simulate(ff []int, days int) int {
	for day := 0; day < days; day++ {
		ff = tick(ff)
	}
	return len(ff)
}

func tick(ff []int) []int {
	res := make([]int, len(ff), len(ff)*2)
	nn := make([]int, 0, len(ff))
	for i, f := range ff {
		var spawn bool
		res[i], spawn = state(f)
		if spawn {
			nn = append(nn, 8)
		}
	}
	return append(res, nn...)
}

func state(f int) (int, bool) {
	f--
	if f < 0 {
		return 6, true
	}
	return f, false
}
