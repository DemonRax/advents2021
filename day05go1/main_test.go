package main

import (
	"testing"

	"advents2021/util"
)

func Test_code(t *testing.T) {
	for _, test := range []struct {
		in      string
		wantMax coord
		want    int
	}{
		{
			in: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			wantMax: coord{x: 9, y: 9},
			want:    5,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			vectors, got := parseVectors(util.ReadString(test.in))
			if got != test.wantMax {
				t.Errorf("gotMax = %v, want %v", got, test.wantMax)
			}
			if got := scan(vectors, got); got != test.want {
				t.Errorf("got %d want %d", got, test.want)
			}
		})
	}
}
