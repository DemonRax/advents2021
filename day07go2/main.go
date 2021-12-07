package main

import (
	"advents2021/util"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("./day07go2/input.txt")
	fmt.Println(cheapest(input))
}

func cheapest(ss []string) int {
	s := strings.Split(ss[0], ",")
	ints := make(map[int]int, len(s))
	var min, max int
	for _, v := range s {
		p, _ := strconv.Atoi(v)
		ints[p]++
		if min > p {
			min = p
		}
		if max < p {
			max = p
		}
	}

	lowest := math.MaxInt
	for i := min; i <= max; i++ {
		fuel := 0
		for j, c := range ints {
			fuel += c * diff(i, j)
		}
		if fuel < lowest {
			lowest = fuel
		}
	}
	return lowest
}

func diff(a, b int) int {
	a = a - b
	if a < 0 {
		a = -a
	}
	return a * (a + 1) / 2
}
