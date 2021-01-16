package main

import "fmt"

func comp(n, j int) int {
	var lower, bigger int
	if j <= n-j {
		lower = j
		bigger = n - j
	} else {
		lower = n - j
		bigger = j
	}

	var sum, sum1 uint64
	sum = 1
	for i := bigger + 1; i <= n; i++ {
		sum *= uint64(i)
	}

	sum1 = 1
	for j := 1; j <= lower; j++ {
		sum1 *= uint64(j)
	}
	return int(sum / sum1)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	// 还可以转化为对数字1和2的组合问题
	// max2 表示最多使用数字2的个数
	var max2, sum int
	if (n % 2) == 0 {
		sum = 2
		max2 = n/2 - 1
	} else {
		sum = 1
		max2 = n / 2
	}

	// 按2的使用个数来划分组合情况
	// 求和所有情况
	for j := 1; j <= max2; j++ {

		// 对n-j个位置, 向其中放入j个2, 剩余n-j-j个位置放1
		// 共有多少种组合
		// 组合涉及阶乘计算, 较大的测试用例就会出现溢出
		s := comp(n-j, j)
		sum += s
	}
	return sum
}

func climbStairsDp(n int) int {
	p2, p1 := 0, 1

	// 只需要维护前一个台阶和前两个台阶的即可
	for i := 1; i <= n; i++ {
		tmp := p2
		p2 = p1
		p1 = tmp + p1
	}
	return p1
}

func main() {
	//fmt.Println(climbStairs(2))
	//fmt.Println(climbStairs(3))

	//for i := 0; i != 100; i ++ {
	//	if climbStairsDp(i) != climbStairs(i) {
	//		fmt.Println(i)
	//	}
	//}
	fmt.Println(climbStairs(44))
	fmt.Println(climbStairsDp(44))
}
