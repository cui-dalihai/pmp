package main

import "fmt"

func massage(nums []int) int {

	l := len(nums)
	dp := make([]int, l+1)
	dp[1] = nums[0]

	for j := 2; j <= l; j++ {

		var a, b int
		if j >= 2 {
			a = dp[j-2] + nums[j-1]
		} else {
			a = nums[j-1]
		}
		if j >= 3 {
			b = dp[j-3] + nums[j-2]
		} else {
			b = nums[j-2]
		}
		if a < b {
			dp[j] = b
		} else {
			dp[j] = a
		}

	}
	return dp[l]
}

func main() {
	fmt.Println(massage([]int{1, 2, 3, 1}))
	fmt.Println(massage([]int{2, 7, 9, 3, 1}))
	fmt.Println(massage([]int{2, 1, 4, 5, 3, 1, 1, 3}))
}
