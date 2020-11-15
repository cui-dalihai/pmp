package main

import "fmt"

// findMedianSortedArrays 先合并两个数组，然后从合并后的数组中计算结果
// 合并后的数组随着两个数组的空间线性增长: O(m+n)
// 时间平均情况下每个元素都要比较一次: O(m+n)
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
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

// findMedianSortedArrays 对于给定的两个数组，中位数是固定的，合并仅是为了方便计算，也可以在不合并的情况下计算出来
// 不需要新建slice: O(1)
// 时间: O(m+n)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	tn := len(nums1) + len(nums2)
	var i, j int
	var Done bool
	if tn%2 == 0 {
		// 中间两数均值
		im1 := (tn / 2) - 1
		im2 := tn / 2

		var m1, m2 int
		for i < len(nums1) && j < len(nums2) && (!Done) {

			if nums1[i] <= nums2[j] {
				if (i + j) == im1 {
					m1 = nums1[i]
				}

				if (i + j) == im2 {
					m2 = nums1[i]
					Done = true
				}
				i++
			} else {
				if (i + j) == im1 {
					m1 = nums2[j]
				}

				if (i + j) == im2 {
					m2 = nums2[j]
					Done = true
				}
				j++
			}
		}

		for i < len(nums1) && (!Done) {
			if (i + j) == im1 {
				m1 = nums1[i]
			}

			if (i + j) == im2 {
				m2 = nums1[i]
				Done = true
			}
			i++
		}

		for j < len(nums2) && (!Done) {
			if (i + j) == im1 {
				m1 = nums2[j]
			}

			if (i + j) == im2 {
				m2 = nums2[j]
				Done = true
			}
			j++
		}

		return float64(m1+m2) / 2

	} else {
		// 中间数
		im := (tn - 1) / 2

		var m int
		for i < len(nums1) && j < len(nums2) && (!Done) {

			if nums1[i] <= nums2[j] {
				if (i + j) == im {
					m = nums1[i]
					Done = true
				}
				i++
			} else {
				if (i + j) == im {
					m = nums2[j]
					Done = true
				}
				j++
			}
		}

		for i < len(nums1) && (!Done) {
			if (i + j) == im {
				m = nums1[i]
				Done = true
			}
			i++
		}

		for j < len(nums2) && (!Done) {
			if (i + j) == im {
				m = nums2[j]
				Done = true
			}
			j++
		}
		return float64(m)
	}
}

func main() {
	nums1 := []int{1, 2, 3, 5, 5, 6}
	nums2 := []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

}
