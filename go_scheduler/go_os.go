package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

func add(nums []int) int {
	var v int
	for _, num := range nums {
		v = v + num
	}
	return v
}

func addConcurrent(goroutines int, nums []int) int {
	tt := len(nums)
	step := tt / goroutines
	last := goroutines - 1

	var v int64
	var n sync.WaitGroup
	n.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func(g int) {

			start := g * step
			var end int
			if g == last {
				end = tt
			} else {
				end = start + step
			}

			for _, num := range nums[start:end] {
				atomic.AddInt64(&v, int64(num))
			}
			n.Done()
		}(g)
	}
	n.Wait()
	return int(v)
}

func main() {
	al := []int{4, 5, 6, 1, 8, 32, 9, 56, 23, 234, 234, 23, 4, 654, 6, 457, 65, 875, 8, 23423, 234, 2, 12, 3, 123, 24, 5, 656, 7, 78, 98, 34, 28, 28, 2, 7, 27, 2, 63, 8, 23, 6, 3, 73, 73, 7, 37, 3, 6, 2, 63748, 58, 58, 5, 85, 7, 57, 57, 65, 456, 76654, 23452, 45678, 425, 34, 5, 45, 34, 52, 435, 24, 23, 4, 234, 23, 42, 34}

	// 物理核心数6，超线程技术，共计12
	c := runtime.NumCPU()

	// 对于这个累加计算, 计算型工作负载，顺序版和并发版在
	//              单goroutine顺序版       多goroutine并发版
	// 单核心机器上,        快                     慢           并发由于会耗费大量时间处理context-switch, 会比顺序版慢一些
	// 多核心机器上,        慢                     快           并发的优势就会出现, goroutine会被核心并行执行, 计算任务被并行处理就会快一些

	// 而对于IO型负载来说，由于存在大量的等待
	//              单goroutine顺序版       多goroutine并发版
	// 单核心机器上，       慢                        快        大量的等待在单个goroutine内线性累积，会很慢，而等待为并发版本context-switch提供了机会, 看起来就像是多个goroutine被单个核心并行执行
	// 多核心机器上，       慢                        快        多核心并不会对IO型工作负载带来收益，跟单核心机器上几乎一样的速度

	add(al)
	addConcurrent(c, al)
}
