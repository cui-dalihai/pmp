package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {

	sort.Ints(nums)
	var res [][]int

	// 外层循环每个元素逐个检查
	for i := 0; i < len(nums); i++ {
		// 跳过检查过的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 里层循环使用两个游标搜索满足条件的另外两个值
		n := i + 1
		m := len(nums) - 1
		for n < m {
			if nums[n]+nums[m]+nums[i] == 0 {
				res = append(res, []int{nums[i], nums[n], nums[m]})

				// 找到一组后两边同时往里移，遇到相同的跳过
				for n < m && nums[n] == nums[n+1] {
					n++
				}
				n++

				for n < m && nums[m] == nums[m-1] {
					m--
				}
				m--

				// 跟外层循环选中的值相比，内层的两数之和太小了，把左游标右移以扩大
			} else if nums[n]+nums[m] < -nums[i] {
				n++

				// 内层两数之和太大了，把右游标左移以减小，
			} else if nums[n]+nums[m] > -nums[i] {
				m--
			}
		}
	}

	return res
}

func main() {
	//n := []int{-1, 0, 1, -1,2, -4}
	n := []int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}
	r := threeSum(n)
	fmt.Println(r)
}
