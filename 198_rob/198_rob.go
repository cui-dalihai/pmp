package main

func rob(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	} else if l == 1 {
		return nums[0]
	} else if l == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	dp := make([]int, l+1)

	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}

	for j := 2; j < l; j++ {

		pu := dp[j-2] + nums[j]
		pnu := dp[j-1]
		if pu > pnu {
			dp[j] = pu
		} else {
			dp[j] = pnu
		}

	}
	return dp[l-1]
}

func main() {

}
