package main

import (
	"fmt"
)

func CoinChange(coins []int, amount int) int {

	// 初始化记忆
	mem := map[int]int{0: 0}
	for _, c := range coins {
		mem[c] = 1
	}
	var memFunc func([]int, int) int

	memFunc = func(coins []int, amount int) int {
		// 检查记忆内容
		if v, ok := mem[amount]; ok {
			return v
		}
		// 直接在初始化中解决了amount为0的情况
		//if amount == 0 {
		//	mem[amount] = 0
		//	return 0
		//}

		min := -1
		for _, c := range coins {
			if c > amount {
				continue
			}
			opt := memFunc(coins, amount-c)
			// 跳过不能兑换的情况即-1
			if opt != -1 {
				if min == -1 {
					min = opt + 1
				} else {
					if min > opt+1 {
						min = opt + 1
					}
				}
			}
		}
		mem[amount] = min
		return min
	}
	return memFunc(coins, amount)
}

func CoinChangeBottomUp(coins []int, amount int) int {
	var mem = make([]int, amount+1)

	// 对问题规模升序求解
	for j := 1; j <= amount; j++ {

		// 约定无法兑换时返回-1
		min := -1

		// 遍历每个规模为j的问题的子问题
		for _, c := range coins {

			// 跳过比规模还大的硬币
			if c > j {
				continue
			}

			// 跳过无法兑换的子问题
			if mem[j-c] == -1 {
				continue
			}

			subMin := mem[j-c] + 1
			// 首次检查子问题
			if min == -1 {
				min = subMin
			} else {

				// 保存最小硬币数
				if min > subMin {
					min = subMin
				}
			}
		}
		// 保存j规模问题的解
		mem[j] = min
	}
	return mem[amount]
}

func main() {

	cs := []int{2, 4, 5, 10}
	am := 11
	//fmt.Println(CoinChange(cs, am))
	fmt.Println(CoinChangeBottomUp(cs, am))
	//
	//fmt.Println(0/1)
	//fmt.Println(6249%419)
}
