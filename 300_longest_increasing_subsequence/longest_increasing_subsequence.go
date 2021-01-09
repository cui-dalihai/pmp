package main

import (
	"fmt"
)

func lengthOfLIS(nums []int) int {

	l := len(nums)
	if l == 0 {
		return 0
	}

	dp := make([]int, l)

	max := 1
	for j := 0; j < l; j++ {

		// 规模每增加1, 对规模内每个位置进行比较大小
		imax := 1
		for i := 0; i < j; i++ {
			// 如果比该位置大, 那么可以在使用该位置数字的序列长度加1
			if nums[j] > nums[i] {
				m := dp[i] + 1
				if imax < m {
					imax = m
				}
			}

		}
		// 计算规模内所有位置的最大值作为该位置的解
		dp[j] = imax

		if max < dp[j] {
			max = dp[j]
		}

	}

	return max
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	nums = []int{1, 3, 6, 7, 9, 4, 10, 5, 6}
	fmt.Println(lengthOfLIS(nums))
}
