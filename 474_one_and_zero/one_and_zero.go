package main

import "fmt"

func findMaxForm(strs []string, m int, n int) int {

	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {

		// 对于当前str, 包含z个0, o个1
		z, o := countOneAndZero(str)

		// 对于z个0和o个1的str
		// 只需要遍历[z,m]规模的0和[o,n]规模的1即可
		// 因为这样的str不会为低于这些规模的集合新增元素
		for i := m; i >= z; i-- {
			for j := n; j >= o; j-- {

				// 那么对于当前规模i,j
				// dp[i][j] = 1 + dp[i-z][j-o]
				// 考虑到其它str可能已经在dp[i][j]这个位置计算出了结果
				// 所以取最大值
				p := 1 + dp[i-z][j-o]
				if p > dp[i][j] {
					dp[i][j] = p
				}
			}
		}
	}
	return dp[m][n]
}

func countOneAndZero(str string) (m int, n int) {

	for _, b := range str {
		if b == '0' {
			m++
		} else {
			n++
		}
	}
	return
}

func main() {
	fmt.Println(findMaxForm([]string{"11111", "100", "1101", "1101", "11000"}, 5, 7))
	//fmt.Println(findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	//fmt.Println(findMaxForm([]string{"10", "0", "1"}, 1, 1))
}
