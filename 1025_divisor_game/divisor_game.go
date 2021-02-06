package main

import "fmt"

func divisorGame(N int) bool {

	dp := make([]bool, N+1)
	dp[1] = false

	// 递增问题规模
	for j := 1; j <= N; j++ {

		// 每个规模下遍历所有选择
		for i := 1; i < j; i++ {

			// 筛选满足条件的选择
			if j%i == 0 {

				// 如果对手没赢
				if !(dp[j-i]) {

					// 那么A获胜
					dp[j] = true
					break
				}
			}
		}
	}
	return dp[N]
}

func main() {
	fmt.Println(divisorGame(3))
}
