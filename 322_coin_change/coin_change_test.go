package main

import "testing"

var tests = []struct {
	coins   []int
	amounts int
	want    int
}{
	{[]int{1, 2, 5}, 11, 3},
	{[]int{2}, 3, -1},
	{[]int{1}, 0, 0},
	{[]int{1}, 1, 1},
	{[]int{1}, 2, 2},
	{[]int{1, 2, 5, 10}, 10, 1},
	{[]int{2, 5, 10, 1}, 27, 4},
	{[]int{186, 419, 83, 408}, 6249, 20},
}

func TestCoinChange(t *testing.T) {

	for _, test := range tests {
		if got := CoinChange(test.coins, test.amounts); got != test.want {
			t.Errorf("CoinChange(%v, %v) = %v\n", test.coins, test.amounts, got)
		}
	}
}

func TestCoinChangeBottomUp(t *testing.T) {

	for _, test := range tests {
		if got := CoinChangeBottomUp(test.coins, test.amounts); got != test.want {
			t.Errorf("CoinChangeBottomUp(%v, %v) = %v\n", test.coins, test.amounts, got)
		}
	}
}
