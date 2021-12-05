package main

import (
	"fmt"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day01go1/input.txt")
	fmt.Println(increased(input))
}

func increased(ss []string) int {
	var count int
	for i := range ss {
		if i == 0 {
			continue
		}
		if util.ToInt(ss[i]) > util.ToInt(ss[i-1]) {
			count++
		}
	}
	return count
}
