package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {

	x := len(text1)
	y := len(text2)

	dp := make([][]int, x+1)
	for i, _ := range dp {
		dp[i] = make([]int, y+1)
	}

	// i 和 j 分别用来管理text1和text2的长度, 而不是下标
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {

			// 所以这里需要-1来获取对应字符
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				l1 := dp[i-1][j]
				l2 := dp[i][j-1]
				if l1 > l2 {
					dp[i][j] = l1
				} else {
					dp[i][j] = l2
				}
			}
		}
	}
	return dp[x][y]
}

func main() {
	s := "abcbdab"
	t := "bdcaba"
	fmt.Println(longestCommonSubsequence(s, t))
}
