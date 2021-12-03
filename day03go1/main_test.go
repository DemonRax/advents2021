package main

import (
	"testing"

	"advents2021/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int64
	}{
		{
			in: `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`,
			want: 198,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := power(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
