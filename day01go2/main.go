package main

import (
	"fmt"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day01go2/input.txt")
	fmt.Println(increased(input))
}

func increased(ss []string) int {
	xx := append(ss, "0", "0")
	var count int
	for i := range ss {
		if i == 0 {
			continue
		}
		if sum(xx, i) > sum(xx, i-1) {
			count++
		}
	}
	return count
}

func sum(ss []string, i int) int {
	return util.ToInt(ss[i]) + util.ToInt(ss[i+1]) + util.ToInt(ss[i+2])
}
