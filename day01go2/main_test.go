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
			in: `199
200
208
210
200
207
240
269
260
263`,
			want: 5,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := increased(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}
