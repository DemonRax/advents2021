package main

import (
	"advents2021/util"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := util.ReadFile("./day02go1/input.txt")
	fmt.Println(position(input))
}

func position(ss []string) int {
	var x, y int
	for _, s := range ss {
		sp := strings.Split(s, " ")
		command := sp[0]
		value, _ := strconv.Atoi(sp[1])
		switch command {
		case "forward":
			x += value
		case "up":
			y -= value
		case "down":
			y += value
		}
	}
	return x * y
}
