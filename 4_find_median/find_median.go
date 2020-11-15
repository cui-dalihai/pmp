package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j int
	var tn []int

	// 借鉴归并算法
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			tn = append(tn, nums1[i])
			i++
		} else {
			tn = append(tn, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		tn = append(tn, nums1[i:]...)
	}

	if j < len(nums2) {
		tn = append(tn, nums2[j:]...)
	}

	l := len(tn)
	if l%2 == 0 {
		m1 := tn[(l/2)-1]
		m2 := tn[(l / 2)]
		return float64(m1+m2) / 2

	} else {
		m := tn[(l-1)/2]
		return float64(m)
	}
}

func main() {
	nums1 := []int{2}
	nums2 := []int{1, 3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
