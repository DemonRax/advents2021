package main

import (
	"advents2021/util"
	"testing"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in   string
		want int
	}{
		{
			in:   `16,1,2,0,4,2,7,1,2,14`,
			want: 37,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := cheapest(util.ReadString(test.in)); got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
