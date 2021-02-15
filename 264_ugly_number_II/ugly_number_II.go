package main

import (
	"fmt"
	"math"
)

// nthUglyNumber1 递增i, 如果i能被2, 3, 5整除, 则除后的结果是丑数, 那么i也是丑数
// 超时, i是加1递增, 会考察很多非丑数, 应该想办法让i直接递增到下一个丑数
func nthUglyNumber1(n int) int {

	if n == 1 {
		return 1
	}
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

type heap struct {
	nums *[]int
}

func (h heap) heapifyMin(i int) {

	length := len(*h.nums)
	if length <= 1 {
		return
	}
	l := (i + 1) << 1
	r := (i+1)<<1 + 1

	var smallest int
	if l <= length && (*h.nums)[i] > (*h.nums)[l-1] {
		smallest = l - 1
	} else {
		smallest = i
	}

	if r <= length && (*h.nums)[smallest] > (*h.nums)[r-1] {
		smallest = r - 1
	}

	if smallest != i {
		(*h.nums)[smallest], (*h.nums)[i] = (*h.nums)[i], (*h.nums)[smallest]
		h.heapifyMin(smallest)
	}

}
func (h heap) decreaseKey(i, k int) {

	if (*h.nums)[i] < k {
		panic("key in place is smaller.")
	}
	(*h.nums)[i] = k
	p := (i+1)>>1 - 1
	for p >= 0 && (*h.nums)[p] > (*h.nums)[i] {
		(*h.nums)[p], (*h.nums)[i] = (*h.nums)[i], (*h.nums)[p]
		i = p
		p = (i+1)>>1 - 1
	}
}
func (h heap) insert(num int) {
	*h.nums = append(*h.nums, math.MaxInt64)
	h.decreaseKey(len(*h.nums)-1, num)
}
func (h heap) extractMin() int {
	r := (*h.nums)[0]
	(*h.nums)[0] = (*h.nums)[len(*h.nums)-1]
	h.heapifyMin(0)
	*h.nums = (*h.nums)[:len(*h.nums)-1]
	return r
}

func nthUglyNumber(n int) int {

	nums := make([]int, n)
	H := heap{nums: &[]int{1}}

	i := 0
	for i < n {
		p := H.extractMin()
		if i > 1 && nums[i-1] == p {
			continue
		}
		for _, v := range []int{p * 2, p * 3, p * 5} {
			H.insert(v)
		}
		nums[i] = p
		i++
	}
	return nums[n-1]
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

// nthUglyNumber2 官方解法
// 使用nums来存储从小到大的丑数
// 遍历丑数的方法:
// 从最小丑数1开始, 后面的每个丑数r都是由前面的某个丑数k乘以2，乘以3，或者乘以5得到的
// 而为了遍历所有丑数, 前面的每个丑数都要乘以2，3，5分别得到三个新的丑数, 同样这三个新的丑数再乘以2，3，5得到九个新的丑数
// 开始时把三个指针都指向1, 指针停留的位置就是用于产生下一个丑数的当前值
// 一旦某个指针乘以了自己对应的质因数, 比如i2=0时, nums[0]*2=2, 这个指针将向下移动, i2++后变为1,
// 但是这个nums[0]还停留着i3和i5两个指针, 这表示nums[0]还可以乘以3和5产生另外两个丑数,
// 为什么每次移动要找最小的丑数？
// 是为了计算前n个丑数, 如果仅仅是为了产生丑数, 比如产生丑数不考虑顺序, 那么可以先把一个指针移动结束, 再移动另一个
// 所以每次都比较当前三个指针所能产生的丑数, 使用最小的那个
// 为什么每次指针++？
// 是为了遍历所有丑数, 每个丑数都会经过三个指针
func nthUglyNumber2(n int) int {

	nums := make([]int, n)
	nums[0] = 1
	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {

		r := min([]int{nums[i2] * 2, nums[i3] * 3, nums[i5] * 5})
		nums[i] = r

		if r == nums[i2]*2 {
			i2++
		}
		if r == nums[i3]*3 {
			i3++
		}
		if r == nums[i5]*5 {
			i5++
		}
	}
	return nums[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
	//fmt.Println(nthUglyNumber(114) == nthUglyNumber1(114))
	//fmt.Println(nthUglyNumber(10) == nthUglyNumber1(10))
	//fmt.Println(nthUglyNumber(1352) == nthUglyNumber1(1352))
}
