package main

import (
	"testing"

	"advents2021/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in: `forward 5
down 5
forward 8
up 3
down 8
forward 2`,
			want: 150,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := position(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
