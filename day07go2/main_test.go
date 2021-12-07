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
			want: 168,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := cheapest(util.ReadString(test.in)); got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}

func Test_diff(t *testing.T) {
	for _, test := range []struct {
		desc string
		a, b int
		want int
	}{
		{
			desc: "diff 0",
		},
		{
			desc: "diff 1",
			a:    1, b: 2,
			want: 1,
		},
		{
			desc: "diff 2",
			a:    4, b: 2,
			want: 3,
		},
		{
			desc: "diff 3",
			a:    5, b: 2,
			want: 6,
		},
		{
			desc: "diff 4",
			a:    6, b: 2,
			want: 10,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := diff(test.a, test.b); got != test.want {
				t.Errorf("diff = %d want %d", got, test.want)
			}
		})
	}
}
