package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"abc", false},
		{"abdba", true},
		{"a", true},
		{"", true},
		{"aaaaa", true},
		{"aaaa", true},
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q)=%v", test.input, got)
		}
	}

}

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"abde", "a"},
		{"babad", "bab"},
		{"abbc", "bb"},
		{"aba", "aba"},
		{"bb", "bb"},
		{"bacabab", "bacab"},
	}

	for _, test := range tests {
		if got := LongestPalindrome2(test.input); got != test.want {
			t.Errorf("LongestPalindrome(%q) = %v", test.input, got)
		}
	}

}
