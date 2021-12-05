package main

import (
	"fmt"
	"strconv"
	"strings"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day02go2/input.txt")
	fmt.Println(move(input))
}

func move(ss []string) int {
	var aim, x, y int
	for _, s := range ss {
		sp := strings.Split(s, " ")
		command := sp[0]
		value, _ := strconv.Atoi(sp[1])
		switch command {
		case "forward":
			x += value
			y += aim * value
		case "up":
			aim -= value
		case "down":
			aim += value
		}
	}
	return x * y
}
