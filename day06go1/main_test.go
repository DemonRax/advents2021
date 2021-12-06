package main

import (
	"advents2021/util"
	"testing"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		days int
		want int
	}{
		{
			in:   `3,4,3,1,2`,
			days: 18,
			want: 26,
		},
		{
			in:   `3,4,3,1,2`,
			days: 80,
			want: 5934,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := simulateRaw(util.ReadString(test.in), test.days); got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
