package main

import (
	"testing"
)

var tests = []struct {
	input int
	want  int
}{
	{0, 0},
	{1, 1},
	{6, 17},
	{7, 18},
	{9, 25},
	{10, 30},
	{15, 43},
}

func sliceEqual(a, b []int) bool {

	if (a == nil) || (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestCutRod(t *testing.T) {

	for _, test := range tests {
		if got := CutRod(test.input); got != test.want {
			t.Errorf("CutRod(%v) = %q\n", test.input, got)
		}
	}

}

func TestCutRodTopDownArr(t *testing.T) {
	for _, test := range tests {
		if got := CutRodTopDownArr(test.input); got != test.want {
			t.Errorf("CutRodTopDownArr(%v) = %q\n", test.input, got)
		}
	}

}

func TestCutRodBottomUp(t *testing.T) {
	for _, test := range tests {
		if got := CutRodBottomUp(test.input); got != test.want {
			t.Errorf("CutRodBottomUp(%v) = %q\n", test.input, got)
		}
	}
}

func TestCutRodBottomUpWithSolu(t *testing.T) {
	var tests = []struct {
		input int
		maxv  int
		solu  []int
	}{
		{0, 0, []int{}},
		{1, 1, []int{1}},
		{6, 17, []int{6}},
		{7, 18, []int{6, 1}},
		{9, 25, []int{6, 3}},
		{10, 30, []int{10}},
		{15, 43, []int{10, 3, 2}},
	}
	for _, test := range tests {
		if gotr, gots := CutRodBottomUpWithSolu(test.input); gotr != test.maxv || !sliceEqual(gots, test.solu) {
			t.Errorf("CutRodBottomUpWithSolu(%v) = %v, %v, expect: %v, %v\n", test.input, gotr, gots, test.maxv, test.solu)
		}
	}
}
