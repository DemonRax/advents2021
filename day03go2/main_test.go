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
			want: 230,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := oxygen(util.ReadString(test.in)); got != 23 {
				t.Fatalf("oxygen %d", got)
			}
			if got := co2(util.ReadString(test.in)); got != 10 {
				t.Fatalf("co2 %d", got)
			}
			if got := diagnose(util.ReadString(test.in)); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_LowHigh(t *testing.T) {
	for _, test := range []struct {
		desc      string
		zero, one int
		wantLow,
		wantHigh uint8
	}{
		{
			desc:     "empty",
			wantLow:  '0',
			wantHigh: '1',
		},
		{
			desc:     "more zeroes",
			zero:     10,
			one:      9,
			wantLow:  '1',
			wantHigh: '0',
		},
		{
			desc:     "more zeroes",
			zero:     10,
			one:      9,
			wantLow:  '1',
			wantHigh: '0',
		},
		{
			desc:     "more ones",
			zero:     10,
			one:      11,
			wantLow:  '0',
			wantHigh: '1',
		},
		{
			desc:     "no ones",
			zero:     10,
			one:      0,
			wantLow:  '0',
			wantHigh: '0',
		},
		{
			desc:     "no zeroes",
			zero:     0,
			one:      1,
			wantLow:  '1',
			wantHigh: '1',
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if got := low(test.zero, test.one); got != test.wantLow {
				t.Errorf("low = %d, want %d", got, test.wantLow)
			}
			if got := high(test.zero, test.one); got != test.wantHigh {
				t.Errorf("high = %d, want %d", got, test.wantHigh)
			}
		})
	}
}
