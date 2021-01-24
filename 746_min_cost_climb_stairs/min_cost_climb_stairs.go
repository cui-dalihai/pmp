package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {

	// 解的最后一次可能是倒数第一个或倒数第二个阶梯, 如果使用这个思路往前一个状态转换的化会变成:
	// 到达最后一个阶梯的最小解，包含倒数第一个和倒数第二个的解
	// 到达倒数第二个阶梯的最小解，包含倒数第三和第四的解
	// 这就很复杂了，
	// 如果我们增加一个空的阶梯到最后，那么问题就变成跳到最后一个阶梯的最小代价，
	// 其状态转化就变为倒数第一和倒数第二的两个解，
	// m := len(cost)
	// cost = append(cost, 0)
	// 那么 opt(m) 就表示跳到下标为m的阶梯，所需要的代价, 注意, 不包括这个m阶梯上的代价, 只有在m上继续往后跳时, 才会opt(m) + cost[m]
	// opt(m) = min{ opt(m-1)+cost[m-1], opt[m-2]+cost[m-2] }
	// 所以, opt(0)表示跳到下标为0的阶梯所需要的代价, 而第0阶梯是起始阶梯，不需要耗费任何体力就可以站上去的，所以opt(0) = 0
	// 同理 opt(1) = 0
	cost = append(cost, 0)

	// 每次只要维护前一个和前两个阶梯的最小代价就可以了，那么对于一个新的阶梯计算完之后作为p1, 原来的p1作为p2, 用于下一个阶梯的计算
	// 所以p1就是当前阶梯的最优解
	p2, p1 := 0, 0

	for i, _ := range cost[2:] {

		a := p1 + cost[i-1+2] // 加 2 是为了补偿从下标为2的地方开始
		b := p2 + cost[i-2+2]
		if a < b {
			temp := p1
			p1 = a
			p2 = temp
		} else {
			temp := p1
			p1 = b
			p2 = temp
		}
	}
	return p1
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))
	fmt.Println(minCostClimbingStairs([]int{0, 0, 1, 1}))
	fmt.Println(minCostClimbingStairs([]int{0, 1, 2, 0}))
}
