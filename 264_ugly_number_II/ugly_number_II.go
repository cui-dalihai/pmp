package main

import "fmt"

func nthUglyNumber1(n int) int {

	mem := make(map[int]bool)
	mem[0] = false
	mem[1] = true

	for i := 1; ; i++ {

		mi := 0
		if i%2 == 0 {
			mi = i / 2
		} else if i%3 == 0 {
			mi = i / 3
		} else if i%5 == 0 {
			mi = i / 5
		} else {
			continue
		}

		if mem[mi] {
			mem[i] = true
			n -= 1
			if n == 1 {
				return i
			}
		} else {
			continue
		}
	}
}

func min(nums []int) int {
	m := nums[0]
	for _, v := range nums[1:] {
		if m > v {
			m = v
		}
	}
	return m
}

func nthUglyNumber(n int) int {

	nums := []int{1}
	i2, i3, i5 := 0, 0, 0

	for i := 1; i <= 1690; i++ {
		r := min([]int{nums[i2] * 2, nums[i3] * 3, nums[i5] * 5})
		nums = append(nums, r)
		if r == nums[i2]*2 {
			i2 += 1
		}
		if r == nums[i3]*3 {
			i3 += 1
		}
		if r == nums[i5]*5 {
			i5 += 1
		}
	}
	return nums[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(114) == nthUglyNumber1(114))
}
