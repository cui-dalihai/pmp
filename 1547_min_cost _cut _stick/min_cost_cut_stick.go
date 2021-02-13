package main

import (
	"fmt"
	"sort"
)

// 第一版
// 直接使用两个指针遍历所有区间, 计算每个区间内的最小值
// 但无法让较小规模的问题优先计算
func minCost1(n int, cuts []int) int {

	// dp[i][j]表示(i,j)区间内的最小值, 不包括两端的i, j
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i <= n-1; i++ {
		for j := i + 1; j <= n; j++ {
			if j-i == 1 {
				dp[i][j] = 0
				continue
			}

			// 检查(i,j)区间内存在的切割点, 取最小值
			min := 0
			for _, c := range cuts {
				if i < c && c < j {
					m := dp[i][c] + dp[c][j] + j - i
					if min == 0 || m < min {
						min = m
					}
				}
			}
			dp[i][j] = min
		}
	}

	return dp[0][n]
}

// 第二版
// 使用l来递增遍历所有长度规模的子问题, 使用i来遍历每个可能的切割点, (i,i+l)就是所有区间, 在每个区间内取最小值
// 超时
func minCost2(n int, cuts []int) int {

	// dp[i][j]表示(i,j)区间内的最小值, 不包括两端的i, j
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for i := 0; i <= n-l; i++ {
			j := i + l

			// 检查(i,j)区间内存在的切割点, 取最小值
			min := 0
			for _, c := range cuts {
				if i < c && c < j {
					m := dp[i][c] + dp[c][j] + l
					if min == 0 || m < min {
						min = m
					}
				}
			}
			dp[i][j] = min
		}
	}

	return dp[0][n]
}

// 第三版
// 上一版中使用i++递增遍历的方式来考察所有可能的切割点，这会考察很多不在cuts中的切割点
// 这些切割点会在i<c<j的判断中判为0, dp会被填充很多0
// 所以这一版中仅考察cuts中的切割点, 在这些切割点上考察每个长度下的子问题
// 添加一个0切割点是为了计算原规模的问题, 即切割点为0，且长度为n时就是原问题的解
// 超出内存
func minCost3(n int, cuts []int) int {

	cuts = append(cuts, 0)
	// dp[i][j]表示(i,j)区间内的最小值, 不包括两端的i, j
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for _, i := range cuts {
			j := i + l

			// 对于cuts中的切割点i, 如果考察区间超出了n, 说明当前的i切割点加上长度l超出了总长度,
			// 就没有必要考察了, 但是cuts不是顺序的, 所以后面可能会有更小的i, 所以要continue而不是break
			if j > n {
				continue
			}

			// 检查(i,j)区间内存在的切割点, 取最小值
			min := 0
			for _, c := range cuts {
				if i < c && c < j {
					m := dp[i][c] + dp[c][j] + l
					if min == 0 || m < min {
						min = m
					}
				}
			}
			dp[i][j] = min
		}
	}

	return dp[0][n]
}

// 第四版
// 上个版本中为了遍历所有(i,j)区间, 创建了n*(n+1)的矩阵, 而这些只有在i,j恰好取值cuts中的切割点时才有值, 其余都是0
// 这个版本使用cuts中的下标映射cuts中的值, dp只要len(cuts)+1行就可以,
// 比如dp[1]就表示切割点为cuts[1]时每种长度l下的最小切割成本, 考察的就是(cuts[1], j)区间内的所有长度l
// 这样n*(n+1)规模矩阵变为l*(n+1)规模
// 超时
func minCost4(n int, cuts []int) int {
	cuts = append(cuts, 0)
	l := len(cuts)
	dp := make([][]int, l+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for l := 2; l <= n; l++ {
		for ic, i := range cuts {
			j := i + l
			if j > n {
				continue
			}

			min := 0
			for jc, c := range cuts {
				if i < c && c < j {
					m := dp[ic][c] + dp[jc][j] + l
					if min == 0 || m < min {
						min = m
					}
				}
			}
			dp[ic][j] = min
		}
	}

	return dp[l-1][n]
}

// 第五版
// 基于第四版继续优化, 把横坐标同样改用cuts下标映射,
// 规模改用cuts的下标
// 通过
func minCost5(n int, cuts []int) int {
	cuts = append(cuts, 0, n)
	sort.Ints(cuts)
	l := len(cuts)
	dp := make([][]int, l+1)
	for i, _ := range dp {
		dp[i] = make([]int, l+1)
	}

	// 通过cuts下标递增区间长度
	for step := 1; step <= l-1; step++ {
		for i, _ := range cuts {
			j := i + step + 1
			if j >= l {
				break
			}

			// 考察(i, i+step+1)区间内的最小值
			min := 0
			for k := i + 1; k <= i+step; k++ {
				m := dp[i][k] + dp[k][j] + (cuts[j] - cuts[i])
				if min == 0 || m < min {
					min = m
				}
			}
			dp[i][j] = min
		}
	}
	return dp[0][l-1]
}

func main() {
	//s := time.Now()
	//fmt.Println(minCost2(2781, []int{1051,127,890,752,66,1217,962,809,1376,2659,1819,1054,964,1350,1076,2361,215,2645,852,1024,953,581,1053,617,2383,429,701,2382,963,958,98,1000,2518,1224,2285,614,2614,1779,2495,2306,1774,758,2618,2693,2158,881,712,2541,1148,438,667,895,2359,1487,2015,1417,2002,2386,1403,765,2183,535,463,673,2571,307,1645,1479,408,190,2403,265,2338,1017,2207}))
	//fmt.Println(time.Since(s))

	//t := time.Now()
	//fmt.Println(minCost3(5709, []int{5033,3175,655,3763,1378,3633,758,3306,2928,4775,218,5052,1867,4458,4548,1275,2965,870,5141,2717,1256,3789,612,4351,1331,3923,5371,5637,2834,3445,5409,1600,963,5390,3247,2059,5428,3018,3899,5076,1664,629,2119,5302,3416,1685,1097,3292,2145,1186,3188}))
	//fmt.Println(time.Since(t))
	fmt.Println(minCost5(7, []int{1, 3, 4, 5}))
	fmt.Println(minCost5(9, []int{5, 6, 1, 4, 2}))
}
