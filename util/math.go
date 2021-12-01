package util

import "strconv"

func ToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
