package main

import "testing"

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
		mv    int
		solu  []int
	}{
		{0, 0, []int{}},
		{1, 1, []int{1}},
		{6, 17, []int{6}},
		{7, 18, []int{1, 6}},
		{9, 25, []int{3, 6}},
		{10, 30, []int{10}},
	}
	for _, test := range tests {
		if got, _ := CutRodBottomUpWithSolu(test.input); got != test.mv {
			t.Errorf("CutRodBottomUpWithSolu(%v) = %q\n", test.input, got)
		}
	}
}
