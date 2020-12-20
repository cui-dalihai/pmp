package main

import "fmt"

// 长度:价格 表
// 求n长度时切割(切割可以为0)的最高价格
var P = map[int]int{
	1:  1,
	2:  5,
	3:  8,
	4:  9,
	5:  10,
	6:  17,
	7:  17,
	8:  20,
	9:  24,
	10: 30,
}

// 时间: 2^n
// 空间: 1
func CutRod(n int) int {
	if n == 0 {
		return 0
	}
	max := -1

	// 长度为n的钢条, 每1的长度位置都有切和不切两种方案, 所以所有方案是2^n-1
	// 而这个递归实际上就是在遍历这些方案
	// 时间复杂度为2^n, 即n每增加1, 求解时间大约增加一倍
	for i := 1; i <= n; i++ {
		subr := P[i] + CutRod(n-i)
		if max < subr {
			max = subr
		}
	}
	return max
}

// 时间: n^2
// 空间: n
func CutRodTopDownMap(n int) int {
	var mem = map[int]int{}
	var subCutRod func(int) int

	subCutRod = func(n int) int {
		if v, ok := mem[n]; ok {
			return v
		}

		if n == 0 {
			mem[n] = 0
			return 0
		}

		max := -1
		for i := 1; i <= n; i++ {
			subp := P[i] + subCutRod(n-i)
			if max < subp {
				max = subp
			}
		}
		mem[n] = max
		return max
	}

	return subCutRod(n)
}

// 时间: n^2
// 空间: n
func CutRodTopDownArr(n int) int {

	// subCutRod在递归求解过程使用的mem
	// 要对0和n求解, 所以是n+1长度
	var mem = make([]int, n+1)
	var subCutRod func(int) int

	subCutRod = func(n int) int {

		// 如果n已经被计算过, 直接返回
		if mem[n] != 0 {
			return mem[n]
		}

		if n == 0 {
			return 0
		}

		// 求解n要递归调用n-i
		max := 0
		for i := 1; i <= n; i++ {
			subp := P[i] + subCutRod(n-i)
			if max < subp {
				max = subp
			}
		}
		mem[n] = max
		return max
	}

	return subCutRod(n)
}

// 时间: n^2
// 空间: n
func CutRodBottomUp(n int) int {

	// 使用数组下标作为问题规模
	var mem = make([]int, n+1)

	// 升序遍历每种规模的问题
	// 每次循环就是对规模为j的问题进行求解
	// 而对规模为j的求解过程实际上和上面递归求解是相同的, 都是对j的子问题进行递归求解
	// 只不过, 前面的递归求解中每次从最大规模的j开始, 再计算规模递减的子问题
	// 虽然实际计算过程中最小规模的子问题最先计算出来, 但这是由递归结构导致的, 整个求解过程最后计算的是规模最小的问题
	// 而自底向上的求解方法中先求解最小规模的j并存储结果, 当求解更大规模的j时, 直接从mem中获取结果
	// 实际上, 这种方法从小求解的过程要和mem配合使用, 如果去掉这两点, 那就是去掉外层递增问题规模j的for直接使用j=n
	// 内层mem[j-i]则要替换成递归CutRodBottomUp(j-i)的形式, 这样就回到了最初的那种递归求解形态,
	// 所以从底向上的求解过程跟递归过程并没有本质的区别, 相比递归形态, 差别仅仅在于对不同规模问题的求解顺序上
	// 即递归求解是问题规模从大到小, 而自底向上则是从小到大
	// 在两者都使用了mem进行优化的情况下, 对所有子问题都计算一次, 两者的时间复杂度没有区别
	for j := 1; j <= n; j++ {

		// 内部对于规模为j的子问题的求解过程和递归求解形态完全一致
		max := 0
		for i := 1; i <= j; i++ {
			// 由于mem的计算是升序的, 每个比j小的子问题已经被计算
			subr := P[i] + mem[j-i]
			if max < subr {
				max = subr
			}
		}

		// 计算完j后存入mem
		mem[j] = max
	}

	// 填充完后, 直接返回对应问题规模的解
	return mem[n]
}

// 返回最大收益的同时返回切割方案
func CutRodBottomUpWithSolu(n int) (int, []int) {
	var mem = make([]int, n+1)
	var solu = make(map[int][]int)
	solu[0] = []int{}

	for j := 1; j <= n; j++ {

		max := 0
		// 最多每1个长度切割一次
		subSolu := make([]int, j)
		for i := 1; i <= j; i++ {
			subr := P[i] + mem[j-i]
			if max < subr {
				// 和最大收益计算原理一致,
				// 规模为j的最优切割方案等于规模为j-i的最优切割方案加上i的长度
				subSolu = append(solu[j-i], i)
				max = subr
			}
		}
		solu[j] = subSolu
		mem[j] = max
	}
	return mem[n], solu[n]
}

func main() {
	fmt.Println(CutRodBottomUpWithSolu(15))
}
