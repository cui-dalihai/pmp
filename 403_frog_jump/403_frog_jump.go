package main

import "fmt"

// 超时
func canCross1(stones []int) bool {
	l := len(stones)

	m := make(map[int]map[int]bool)

	for i := 0; i < l; i++ {
		m[i] = map[int]bool{}
	}
	m[0][0] = true

	// j 递增问题规模
	for j := 1; j < l; j++ {

		// 从i能跳到j, i取值[0, j-1]
		for i := 0; i < j; i++ {

			// 此次使用的步为k
			k := stones[j] - stones[i]

			// 那么k是i的步的某一个元素+1,+0,-1
			for v, _ := range m[i] {
				if v == k {
					m[j][k] = true
					break
				} else if v+1 == k {
					m[j][k] = true
					break
				} else if v-1 == k {
					m[j][k] = true
					break
				}
			}
		}
	}
	if len(m[l-1]) != 0 {
		return true
	}
	return false
}

// 优化1
func canCross2(stones []int) bool {
	l := len(stones)

	m := make(map[int]map[int]bool)

	for i := 0; i < l; i++ {
		m[i] = map[int]bool{}
	}
	m[0][0] = true

	// j 递增问题规模
	for j := 1; j < l; j++ {

		// 从i能跳到j, i取值[0, j-1]
		for i := 0; i < j; i++ {

			// 如果i的步为空, 即i以前的石头无法跳到i上, 那也就无法从i跳到j上
			if len(m[i]) == 0 {
				continue
			}

			// 此次使用的步为k
			k := stones[j] - stones[i]

			// 如果i,j之间的步长大于i所能跳的最大距离(即i前面每一跳都+1)
			if k > stones[i]+i+1 {
				continue
			}

			// 检查i是否可以使用k来跳到j
			// i到j的步为k, 那么i的步中包含k-1, k, k+1步可以跳到j的
			for s := k - 1; s <= k+1; s++ {
				if _, ok := m[i][s]; ok {

					m[j][k] = true
					break
				}
			}
		}
	}
	if len(m[l-1]) != 0 {
		return true
	}
	return false
}

func main() {

	s := []int{0, 1, 3, 5, 6, 8, 12, 17}
	s = []int{0, 1, 2, 3, 4, 8, 9, 11}
	s = []int{0, 2}

	fmt.Println(canCross2(s))

}
