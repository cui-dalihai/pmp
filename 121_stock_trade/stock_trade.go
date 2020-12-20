package main

import (
	"fmt"
)

func maxProfit(prices []int) int {

	l := len(prices)
	if l == 0 {
		return 0
	}

	max := 0
	for buy := range prices[:l-1] {
		for sale := buy + 1; sale < l; sale++ {
			p := prices[sale] - prices[buy]
			if p < 0 {
				continue
			}
			if p > max {
				max = p
			}
		}
	}
	return max
}

func maxProfitDp(prices []int) int {

	l := len(prices)
	if l <= 1 {
		return 0
	}

	lp := prices[0]
	mp := 0

	for j := 1; j < l; j++ {

		// j天的最大收益=max(j-i天内最大收益, j天的价格减去j-1天内的最低价格)
		if prices[j]-lp > mp {
			mp = prices[j] - lp
		}
		if lp > prices[j] {
			lp = prices[j]
		}
	}
	return mp
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfitDp(prices))
	//prices = []int{}
	//fmt.Println(maxProfit(prices))
}
