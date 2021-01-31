package main

import "fmt"

func canPartition(nums []int) bool {

	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	hsum := sum / 2
	l := len(nums)
	dp := make([][]bool, l+1)
	for i, _ := range dp {
		dp[i] = make([]bool, hsum+1)
	}
	dp[0][0] = false

	// 转化为背包问题
	// 问题变成: 列表中是否存在部分元素可以组成sum/2
	// dp[j][i]: j表示[0,j]包括两端的元素, i表示这些元素是否可以求和的数字
	// 如果dp[j][i] = true就表示下标[0,j]的列表元素恰好可以求和为i
	for j := 1; j <= l; j++ {
		for i := 1; i <= hsum; i++ {

			// 分三种情况讨论:
			// a. 如果我们不使用nums[j]来求和i, 则[0,j]能不能拼出i就取决于[0,j-1]能不能拼出i
			if dp[j-1][i] {
				dp[j][i] = true
			} else {

				// b. 如果我们使用了nums[j]来求和i, 而且nums[j]恰好等于i, 那么dp[j][i]恰好等于true
				if nums[j-1] == i {
					dp[j][i] = true
				}

				// c. 如果我们使用了nums[j]来求和i, 而且nums[j]小于这个i, 那么[0,j]能不能拼出i就取决于[0,j-1]能不能拼出i-nums[j]
				if nums[j-1] < i {
					dp[j][i] = dp[j-1][i-nums[j-1]]
				}

			}
		}
	}
	fmt.Println(dp)
	return dp[l][hsum]
}

func main() {
	//fmt.Println(canPartition([]int{1, 5, 11, 5}))
	//fmt.Println(canPartition([]int{1, 2, 3, 5}))
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
}
