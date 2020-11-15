package main

import "fmt"

func mergeSort(alist []int) []int {
	if len(alist) > 1 {
		mid := len(alist) / 2

		ll := mergeSort(alist[:mid])
		rr := mergeSort(alist[mid:])

		var i, j, k int
		r := make([]int, len(ll)+len(rr))

		// 归并
		for i < len(ll) && j < len(rr) {
			if ll[i] < rr[j] {
				r[k] = ll[i]
				i++
			} else {
				r[k] = rr[j]
				j++
			}
			k++
		}

		for i < len(ll) {
			r[k] = ll[i]
			i++
			k++
		}

		for j < len(rr) {
			r[k] = rr[j]
			j++
			k++
		}
		return r
	}
	return alist
}

func main() {
	al := []int{7, 3, 1, 90, 34, 55, 0, 89, 5, 2}
	aa := mergeSort(al)
	fmt.Println(aa)
}
