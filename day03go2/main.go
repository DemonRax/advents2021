package main

import (
	"advents2021/util"
	"fmt"
	"strconv"
)

func main() {
	input := util.ReadFile("./day03go2/input.txt")
	fmt.Println(diagnose(input))
}

func diagnose(input []string) int64 {
	return oxygen(input) * co2(input)
}

func oxygen(ss []string) int64 {
	return reduce(ss, 0, high)
}

func co2(ss []string) int64 {
	return reduce(ss, 0, low)
}

func reduce(ss []string, i int, pred predicate) int64 {
	if len(ss) == 1 {
		res, _ := strconv.ParseInt(ss[0], 2, 64)
		return res
	}
	return reduce(filter(ss, i, pred), i+1, pred)
}

func filter(ss []string, i int, pred predicate) []string {
	data := bits(ss)
	r := pred(data[i]['0'], data[i]['1'])
	res := make([]string, 0, len(ss))
	for _, s := range ss {
		if s[i] == r {
			res = append(res, s)
		}
	}
	return res
}

func bits(ss []string) map[int]map[uint8]int {
	data := make(map[int]map[uint8]int, len(ss[0]))
	for j, s := range ss {
		for i, r := range s {
			if j == 0 {
				data[i] = make(map[uint8]int, 2)
			}
			data[i][uint8(r)]++
		}
	}
	return data
}

type predicate func(zero int, one int) uint8

func low(zero, one int) uint8 {
	switch {
	case zero == one:
		return '0'
	case zero == 0:
		return '1'
	case one > zero, one == 0:
		return '0'
	default:
		return '1'
	}
}

func high(zero, one int) uint8 {
	switch {
	case zero == one:
		return '1'
	case one == 0:
		return '0'
	case one > zero, zero == 0:
		return '1'
	default:
		return '0'
	}
}
