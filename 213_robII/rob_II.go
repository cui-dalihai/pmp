package main

import "fmt"

func robHelp(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	dp := make([]int, l+1)
	dp[1] = nums[0]
	if nums[0] > nums[1] {
		dp[2] = nums[0]
	} else {
		dp[2] = nums[1]
	}

	for j := 3; j <= l; j++ {
		rl := dp[j-2] + nums[j-1]
		nl := dp[j-1]
		if rl > nl {
			dp[j] = rl
		} else {
			dp[j] = nl
		}
	}
	return dp[l]
}

func rob(nums []int) int {

	l := len(nums)
	// 转化为两个子问题
	// 1. 抢最后一个
	// 则对nums[1:l-1]求解
	// 2. 不抢最后一个
	// 则对nums[0:]求解
	if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}

	nums1 := nums[0 : l-1]
	nums2 := nums[1:]
	r1 := robHelp(nums1)
	r2 := robHelp(nums2)

	if r1 > r2 {
		return r1
	} else {
		return r2
	}
}

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}
