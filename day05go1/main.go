package main

import (
	"fmt"
	"strconv"
	"strings"

	"advents2021/util"
)

func main() {
	input := util.ReadFile("./day05go1/input.txt")
	fmt.Println(scan(parseVectors(input)))
}

func scan(vectors []vector, size coord) int {
	field := make([][]int, size.x+1)
	for i := range field {
		field[i] = make([]int, size.y+1)
	}
	var count int
	for _, v := range vectors {
		for x := v.start.x; x <= v.end.x; x++ {
			for y := v.start.y; y <= v.end.y; y++ {
				if field[x][y] == 1 {
					count++
				}
				field[x][y]++
			}
		}
	}
	return count
}

func parseVectors(ss []string) ([]vector, coord) {
	maxCoord := coord{}
	vectors := make([]vector, 0, len(ss))
	for _, s := range ss {
		v := parseVector(s)
		if !v.straight() {
			continue
		}
		vectors = append(vectors, v.direct())
		maxCoord = v.max(maxCoord)
	}
	return vectors, maxCoord
}

func parseVector(s string) vector {
	line := strings.Split(s, " -> ")
	return vector{start: parseCoords(line[0]), end: parseCoords(line[1])}
}

func parseCoords(s string) coord {
	line := strings.Split(s, ",")
	x, _ := strconv.Atoi(line[0])
	y, _ := strconv.Atoi(line[1])
	return coord{x: x, y: y}
}

type coord struct {
	x, y int
}

func (c coord) max(o coord) coord {
	if o.x > c.x {
		c.x = o.x
	}
	if o.y > c.y {
		c.y = o.y
	}
	return c
}

type vector struct {
	start, end coord
}

func (v vector) straight() bool {
	return v.start.x == v.end.x || v.start.y == v.end.y
}

func (v vector) max(c coord) coord {
	return v.start.max(v.end).max(c)
}

func (v vector) direct() vector {
	if v.start.x > v.end.x {
		v.start.x, v.end.x = v.end.x, v.start.x
	}
	if v.start.y > v.end.y {
		v.start.y, v.end.y = v.end.y, v.start.y
	}
	return v
}
