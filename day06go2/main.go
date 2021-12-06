package main

import (
	"advents2021/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("./day06go2/input.txt")
	fmt.Println(simulateRaw(input, 256))
}

func simulateRaw(ss []string, days int) int {
	s := strings.Split(ss[0], ",")
	var ints [9]int
	for _, v := range s {
		i, _ := strconv.Atoi(v)
		ints[i]++
	}
	return simulate(ints, days)
}

func simulate(ff [9]int, days int) int {
	for day := 0; day < days; day++ {
		ff = tick(ff)
	}
	var sum int
	for _, v := range ff {
		sum += v
	}
	return sum
}

func tick(ff [9]int) [9]int {
	n := ff[0]
	for i := 0; i < 8; i++ {
		ff[i] = ff[i+1]
	}
	ff[6] += n
	ff[8] = n
	return ff
}
