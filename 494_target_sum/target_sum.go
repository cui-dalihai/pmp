package main

import "fmt"

func sum(nums []int) int {
	s := 0
	for _, v := range nums {
		s += v
	}
	return s
}

func findTargetSumWays(nums []int, S int) int {

	// 仅有一个元素时
	l := len(nums)
	if l == 1 {
		if -nums[0] == S || nums[0] == S {
			return 1
		}
		return 0
	}

	// S超出所能组合的边界
	sum := sum(nums)
	if S < -sum {
		return 0
	} else if S > sum {
		return 0
	}

	dp := make([][]int, l)

	// 所有可能和
	for i, _ := range dp {
		dp[i] = make([]int, 2*sum+1)
	}

	// adjust用来调整下标和目标和之间的对应关系
	// 比如数组和为7, 那么所能表达的和的范围是[-7, 7], 那么对应到dp的横坐标就是[-7+7, 7+7]
	adjust := sum

	// 处理第一行结果, 这时只有nums[0]参与拼接, 所以只能做出-nums[0]和nums[0]
	// 而如果nums[0]恰好等于0, 那么正负号对nums[0]效果相同但计为两种方式
	// 所以使用+=可以兼容这种情况, 而如果是=, 这种情况就会被少计算一次
	dp[0][-nums[0]+adjust] += 1
	dp[0][nums[0]+adjust] += 1

	for i := 1; i < l; i++ {
		for s := -sum; s <= sum; s++ {
			si := s + adjust
			// 超出左边界, 只要计算右分支
			if si-nums[i] < 0 {
				dp[i][si] = dp[i-1][si+nums[i]]
				// 超出右边界, 只要计算左分支
			} else if si+nums[i] > 2*sum {
				dp[i][si] = dp[i-1][si-nums[i]]
				// 否则使用前i个元素拼s时就是i-1个元素拼s-nums[i]和i-i个元素拼s+nums[i]的方式和
			} else {
				dp[i][si] = dp[i-1][si-nums[i]] + dp[i-1][si+nums[i]]
			}
		}
	}

	return dp[l-1][S+adjust]
}

func main() {
	fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
