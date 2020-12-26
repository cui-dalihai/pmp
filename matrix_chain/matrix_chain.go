package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func matrixMulti(a, b [][]int) ([][]int, error) {

	if len(a[0]) != len(b) {
		return [][]int{}, errors.New("incompatible dimensions")
	}

	c := make([][]int, len(a))
	for i, _ := range c {
		c[i] = make([]int, len(b[0]))
	}

	// (m, n) * (n * j) = (m, j)
	// 共计需要 m * j * n 次标量乘法
	for i, _ := range a {
		for j := 0; j < len(b[0]); j++ {

			for k := 0; k < len(a[0]); k++ {
				c[i][j] += a[i][k] * b[k][j]
			}

		}
	}
	return c, nil
}

func matrixChainBottomUp(p []int) ([][]int, [][]int) {

	n := len(p) - 1

	m := make([][]int, n+1)
	s := make([][]int, n+1)
	for i, _ := range m {
		m[i] = make([]int, n+1)
		s[i] = make([]int, n+1)
	}

	// 使用 l 递增问题规模,
	// l:=1 的情况在初始化时已经完成
	// l:=2 计算在原矩阵链中自链长度为2时的所有标量计算代价
	// l:=3 同样计算长度为3时的代价
	// l:=n 即问题的最后解, n个矩阵链时所有标量计算代价
	for l := 2; l <= n; l++ {

		// i 和 j 用来按照固定l长度从原矩阵链开头滚动自链位置
		// 例如, 原矩阵链长度为5, 那 l:=3 时, 就要计算 m[1,2,3], m[2,3,4], m[3,4,5] 并记录最小值

		// 每个l和i的所有组合就是所有的子问题空间
		// 而每个子问题空间对k的循环就是遍历该问题空间下所有可能的选择
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt64

			// k 用来遍历对一般子链的所有分割情况
			// 例如, 如果子链此时是m[2,4], 那就是A2A3A4这三个矩阵, 那么k取值为2,3,
			// 就是遍历把m[2,3,4]分割为 m[2,2],m[3,4] 和 m[2,3],m[4,4] 两种情况,
			// 其中m[2,2]是初始化的0, 而m[3,4]的规模为1, 已经在上一轮l:=1时计算过, 所以能直接从m中取到
			// 记录这两种情况的最小值到m[2,4]中, 记录最小情况时的分割点k到s[2,4]中
			// s[2,4]=k表示, 对于子链A2A3A4, 在k处分割能取到最小值
			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]
				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	for i, _ := range m {
		fmt.Println(m[i])
	}

	fmt.Println("---")

	for i, _ := range s {
		fmt.Println(s[i])
	}
	return m, s
}

func printSolu(s [][]int, i, j int) string {
	if i == j {
		return "A" + strconv.Itoa(i)
	} else {
		return "(" + printSolu(s, i, s[i][j]) + printSolu(s, s[i][j]+1, j) + ")"
	}
}

func main() {

	//a := [][]int{
	//	{1, 2, 3},
	//	{4, 5, 6},
	//	{1, 1, 1},
	//}
	//b := [][]int{
	//	{1, 2},
	//	{3, 4},
	//	{5, 6},
	//}
	//if c, err := matrixMulti(a, b); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(c)
	//}

	p := []int{30, 35, 15, 5, 10, 20, 25}
	mm, s := matrixChainBottomUp(p)

	st := printSolu(s, 1, 6)
	fmt.Println(st)
	fmt.Println(mm)
}
