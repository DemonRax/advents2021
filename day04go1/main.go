package main

import (
	"fmt"
	"strconv"
	"strings"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day04go1/input.txt")
	fmt.Println(play(parseBingo(input)))
}

func play(nums []int, cards []card) int {
	current := make([]int, 0, len(nums))
	for _, n := range nums {
		current = append(current, n)
		for _, card := range cards {
			sum, win := card.scan(current)
			if sum != 0 && win != 0 {
				return sum * win
			}
		}
	}
	return 0
}

type card [][]int

func (c card) scan(nums []int) (int, int) {
	sum, win := 0, -1

	colHits := make([]int, 5)
	for _, row := range c {
		rowHits := 0
		for i, col := range row {
			hit := false
			for _, num := range nums {
				hit = hit || col == num
			}
			if hit {
				rowHits++
				colHits[i]++
			} else {
				sum += col
			}
		}
		if rowHits == 5 {
			win = nums[len(nums)-1]
		}
	}
	for _, wins := range colHits {
		if wins == 5 {
			win = nums[len(nums)-1]
		}
	}

	if win < 0 {
		win, sum = 0, 0
	}
	return sum, win
}

func parseBingo(ss []string) ([]int, []card) {
	nums := strings.Split(ss[0], ",")
	numbers := make([]int, len(nums))
	for i, num := range nums {
		numbers[i], _ = strconv.Atoi(num)
	}

	cards := make([]card, 0, len(ss))
	var row int
	c := make([][]int, 5)
	for _, line := range ss[1:] {
		if line == "" {
			continue
		}

		c[row] = make([]int, 5)
		var col int
		nums := strings.Split(line, " ")
		for _, num := range nums {
			var err error
			c[row][col], err = strconv.Atoi(num)
			if err == nil {
				col++
			}
		}
		row++
		if row == 5 {
			cards = append(cards, c)
			c = make([][]int, 5)
			row = 0
		}
	}
	return numbers, cards
}
