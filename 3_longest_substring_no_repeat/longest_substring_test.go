package main

import (
	"testing"
)

// 功能测试函数: 以 Test 前缀, 检测逻辑的正确性, go test 会执行这些函数
// 基准测试函数: 以 Benchmark 开头, 测试某些操作的性能,
// 示例函数: 以 Example 开头
func TestLongestSubString(t *testing.T) {

	var tests = []struct {
		input string
		want  int
	}{
		{"aaa", 1},
		{"abcccc", 3},
		{"cccab", 3},
		{"", 0},
		{"abcabcbb", 3},
		{"pwwkew", 3},
	}

	for _, test := range tests {
		if got := LongestSubString(test.input); got != test.want {
			t.Errorf("LongestSubString(%q) = %v", test.input, got)
		}
	}
}
