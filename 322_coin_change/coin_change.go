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

func main() {

	cs := []int{83, 186, 408, 419}
	am := 6249
	fmt.Println(CoinChange(cs, am))
	//
	//fmt.Println(0/1)
	//fmt.Println(6249%419)
}
