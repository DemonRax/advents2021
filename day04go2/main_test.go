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
			in: `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`,
			want: 1924,
		},
		{
			in: `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

10 18 14 22  2
16  8 21 11  0
15 23 17 13 12
 9 26 24  6  3
19 20  4  5  7`,
			want: 1924,
		},
	} {
		t.Run(test.in, func(t *testing.T) {
			if got := play(parseBingo(util.ReadString(test.in))); got != test.want {
				t.Errorf("got = %d, want %d", got, test.want)
			}
		})
	}
}

func Test_scan(t *testing.T) {
	for _, test := range []struct {
		desc    string
		card    card
		nums    []int
		wantSum int
		wantWin int
	}{
		{
			desc: "sample 1 no win",
			card: card{
				{10, 16, 15, 9, 19},
				{14, 21, 17, 24, 4},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
			nums: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21},
		},
		{
			desc: "sample 2 row win",
			card: card{
				{14, 21, 17, 24, 4},
				{10, 16, 15, 9, 19},
				{18, 8, 23, 26, 20},
				{22, 11, 13, 6, 5},
				{2, 0, 12, 3, 7},
			},
			nums:    []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			wantSum: 188,
			wantWin: 24,
		},
		{
			desc: "sample 2 col win",
			card: card{
				{14, 10, 18, 22, 2},
				{21, 16, 8, 11, 0},
				{17, 15, 23, 13, 12},
				{24, 9, 26, 6, 3},
				{4, 19, 20, 5, 7},
			},
			nums:    []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			wantSum: 188,
			wantWin: 24,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			gotSum, gotWin := test.card.scan(test.nums)
			if gotSum != test.wantSum {
				t.Errorf("sum = %d want %d", gotSum, test.wantSum)
			}
			if gotWin != test.wantWin {
				t.Errorf("win = %d want %d", gotWin, test.wantWin)
			}
		})
	}
}
