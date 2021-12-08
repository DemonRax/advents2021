package main

import (
	"advents2021/util"
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := util.ReadFile("./day08go2/input.txt")
	fmt.Println(total(input))
}

func total(ss []string) int {
	var count int
	for _, s := range ss {
		left, right := scan(s)
		count += convert(right, findKey(left))
	}
	return count
}

func convert(ss []string, key map[uint8]uint8) int {
	rev := make(map[uint8]uint8, len(key))
	for k, v := range key {
		rev[v] = k
	}
	var res int
	for i, s := range ss {
		res += parseNumber(s, rev, i)
	}
	return res
}

func parseNumber(s string, key map[uint8]uint8, div int) int {
	res := make([]string, len(s))
	for i := 0; i < len(s); i++ {
		res[i] = string(key[s[i]])
	}
	sort.Strings(res)
	value := toInt[strings.Join(res, "")] * 1000
	for i := 0; i < div; i++ {
		value /= 10
	}
	return value
}

func findKey(ss []string) map[uint8]uint8 {
	res := map[uint8]uint8{}
	bl := byLen(ss)
	cf := bl[2][0]
	res['a'] = exclude(cf, bl[3][0])[0]
	bd := exclude(cf, bl[4][0])
	res['g'] = exclude(string(res['a']), inAll(append(bl[5], bl[6]...)))[0]
	res['e'] = exclude(known(res, cf, bd), bl[6]...)[0]
	res['d'] = exclude(known(res), inAll(bl[5]))[0]
	res['b'] = exclude(known(res), once(bl[5]))[0]
	res['f'] = exclude(known(res), inAll(bl[6]))[0]
	res['c'] = exclude(known(res), inAll(bl[7]))[0]
	return res
}

func known(m map[uint8]uint8, chunks ...string) string {
	res := ""
	for _, v := range m {
		res += string(v)
	}
	for _, s := range chunks {
		res += s
	}
	return res
}

func once(ss []string) string {
	counts := map[rune]int{}
	for _, s := range ss {
		for _, r := range s {
			counts[r]++
		}
	}
	var res []rune
	for r, c := range counts {
		if c == 1 {
			res = append(res, r)
		}
	}
	return string(res)
}

func inAll(ss []string) string {
	counts := map[rune]int{}
	var res []rune
	for _, s := range ss {
		for _, r := range s {
			counts[r]++
			if counts[r] == len(ss) {
				res = append(res, r)
			}
		}
	}
	return string(res)
}

func exclude(known string, ss ...string) string {
	sss := strings.Join(ss, "")
	var res []rune
	for _, r := range sss {
		found := false
		for _, k := range known {
			if k == r {
				found = true
				break
			}
		}
		if !found {
			res = append(res, r)
		}
	}
	return string(res)
}

func byLen(ss []string) map[int][]string {
	res := make(map[int][]string, 6)
	for _, s := range ss {
		res[len(s)] = append(res[len(s)], s)
	}
	return res
}

func scan(s string) ([]string, []string) {
	split := strings.Split(s, " | ")
	return strings.Split(split[0], " "), strings.Split(split[1], " ")
}

const (
	zero  = "abcefg"
	one   = "cf"
	two   = "acdeg"
	three = "acdfg"
	four  = "bcdf"
	five  = "abdfg"
	six   = "abdefg"
	seven = "acf"
	eight = "abcdefg"
	nine  = "abcdfg"
)

var toInt = map[string]int{
	zero:  0,
	one:   1,
	two:   2,
	three: 3,
	four:  4,
	five:  5,
	six:   6,
	seven: 7,
	eight: 8,
	nine:  9,
}
