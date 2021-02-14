package main

import (
	"fmt"
	"math"
)

/*
堆
	a. 以下使用1开始的下标, 而不是0,
	b. 仅讨论大堆, 小堆同理
	c. 涉及到floor_int(i/2)即i除以2后向下取整的计算, 都是用i>>1来代替
	d. lgn 表示以2为底, n的对数

1. 节点关系
	当前节点i,
	父节点	: i>>1
	左子节点	: i<<1
	右子节点	: i<<1 + 1
	每一层从左至右填充

2. 堆属性
	A[parent(i)] >= A[i], 仅要求父节点大于等于子节点,

3. 对比二叉树属性
	a. 要求更低, 仅要求父节点大于等于子节点
	b. 除了最底层, 二叉堆的树是充满的, 所以高 h 的二叉堆, 至少 2^h 个元素, 最多 2^(h+1) - 1个元素, 反过来 n 个元素的堆, 树高是 lgn

4. 对于已经降序排序的数组, 就是大堆, 因为满足堆属性, 反之则不然, 因为对属性不要求两个子节点的相对大小

5. 推导属性:
	a. 对于n个元素的堆, 下标n>>1+1...n都是叶子节点, 注意叶子节点不一定都出现在最后一层, 也可能是倒数第二层

6. 最大优先队列, 基于大堆实现, 最小优先队列就是小堆
	a. maximum:	查找最大优先元素, 返回堆首元素 O(1)
	b. extract-max:	取出最大优先元素, 把堆首返回, 把堆尾放堆首, heapify堆首 O(lgn)
	c. increase-key: 增大某个元素优先, 不断与i>>2元素即父元素对比, 交换 O(lgn)
	d. insert: 插入新元素, 追加一个无穷小元素到堆尾, 再增大这个元素到key大小 O(lgn)
	总之, n个元素的堆中, 所有优先队列的操作的时间复杂度都是O(lgn)

*/

func isMaxHeap(nums []int) (bool, int) {
	l := len(nums)
	for i, v := range nums {

		// 任何一个节点:
		// 右移一位得到父节点
		// 左移一位得到左子节点
		// 左移一位加一得到右子节点
		// parent := (i + 1) >> 1
		left := (i + 1) << 1
		right := (i+1)<<1 + 1

		if left <= l && v < nums[left-1] {
			return false, nums[left-1]
		}
		if right <= l && v < nums[right-1] {
			return false, nums[right-1]
		}

	}
	return true, -1
}

// maxHeapifyRecursive 是从当前节点向下调整, 使当前节点下的节点满足堆属性
// 时间复杂度: O(lgn), 即O(h)
func maxHeapifyRecursive(nums []int, i int) {

	le := len(nums)

	var largest int
	l := (i + 1) << 1
	r := (i+1)<<1 + 1

	if l <= le && nums[i] < nums[l-1] {
		largest = l - 1
	} else {
		largest = i
	}

	if r <= le && nums[largest] < nums[r-1] {
		largest = r - 1
	}

	if largest != i {
		nums[largest], nums[i] = nums[i], nums[largest]

		// 交换时largest是l或r, 交换后i, l, r三个节点满足了堆属性, 但是largest中存储了原来i的值
		// 即, 原来的l或者r的值替换为原来i的值, 这可能导致l或r子树违反堆属性, 所以要递归继续检查largest
		maxHeapifyRecursive(nums, largest)
	}
}

// maxHeapifyLoop 上面对应的非递归版本
func maxHeapifyLoop(nums []int, i int) {
	le := len(nums)

	var largest int
	cur := i

	for {
		l := (cur + 1) << 1
		r := (cur+1)<<1 + 1

		if l <= le && nums[cur] < nums[l-1] {
			largest = l - 1
		} else {
			largest = cur
		}

		if r <= le && nums[largest] < nums[r-1] {
			largest = r - 1
		}

		if largest == cur {
			break
		} else {
			nums[largest], nums[i] = nums[i], nums[largest]
			cur = largest
		}
	}

}

// buildHeap 这是根据推导属性: n个元素的堆, n>>1+1以后的元素都是叶节点, 这些节点可以看作是没有子节点的大堆的根
// 所以从n>>1往前依次构建堆即可
// 第一次循环是构建一层堆的过程, 因为第一次的i=l>>1, 有且仅有一层孩子
// 这样, 在i从l>>1依次向0递减时, 一直维护着如下的循环不变量:
// i+1, i+2 ... n 这些节点都是一个大堆的根
// 时间复杂度: O(nlgn)
func buildHeap(nums []int) {
	l := len(nums)
	for i := l >> 1; i >= 0; i-- {
		maxHeapifyRecursive(nums, i)
	}
}

// heapSort 对nums先构建一次大堆, 那么第一个一定是最大的元素
// 把这个元素跟最后的元素替换, 这时首个元素就不会满足堆属性
// 这时再对去掉最后一个元素的数组也就是去掉最大元素的数组及其首个元素进行heapify
// 这个新堆第一个元素会是去掉最后一个元素的数组的最大值
// 继续把这个元素同这个新堆的最后一个元素交换, 就是倒数第二个元素交换
// 直到i到了正数第二元素, 也就是i=1时, 最后两个元素完成交换, 得到一个正序数组
// 时间复杂度: O(nlgn)
func heapSort(nums []int) {
	buildHeap(nums)
	for i := len(nums) - 1; i >= 1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		maxHeapifyRecursive(nums[:i], 0)
	}
}

// heapExtractMax 取出heap的最大元素, 就是把第一个元素取出, 然后把最后一个元素放到首个元素位置
// 并对第一个位置heapify, 然后再把数组最后一个元素去掉
// 时间复杂度: O(lgn)
func heapExtractMax(nums *[]int) int {
	fmt.Printf("inside func: %p\n", &(*nums))
	l := len(*nums)
	if l < 1 {
		panic("empty heap")
	}
	max := (*nums)[0]
	(*nums)[0] = (*nums)[l-1]
	maxHeapifyRecursive(*nums, 0)
	n := (*nums)[:l-1]
	*nums = n
	return max
}

// heapIncreaseKey 增大某个元素的大小到key
// 时间复杂度: O(lgn)
func heapIncreaseKey(nums []int, i, key int) {
	if nums[i] > key {
		panic("key is smaller.")
	}

	nums[i] = key
	p := (i+1)>>1 - 1
	for p >= 0 && nums[p] < key {
		nums[i], nums[p] = nums[p], nums[i]
		i = p
		p = (i+1)>>1 - 1
	}
}

// heapInsert 向堆中插入一个元素, 先在堆末追加一个无穷小, 再把这个无穷小增大到key
// 时间复杂度: O(lgn)
func heapInsert(nums *[]int, key int) {
	*nums = append(*nums, math.MinInt64)
	heapIncreaseKey(*nums, len(*nums)-1, key)
}

// heapDelete 删除一个key, 把这个key增大到无穷大, 保证会浮到堆首, 然后取出最大元素
// 时间复杂度: O(lgn)
func heapDelete(nums *[]int, i int) {
	l := len(*nums)
	if i >= l || i < 0 {
		panic("invalid index.")
	}
	heapIncreaseKey(*nums, i, math.MaxInt64)
	heapExtractMax(nums)
}

// heapDelete1 删除一个key的另一个方法, 把这个key置为堆尾元素大小, 并对这个位置heapify
// 这时就相当于把最后一个元素替换到i位置进行heapify, 再删除后面那个元素即可
func heapDelete1(nums *[]int, i int) {
	l := len(*nums)
	if (*nums)[i] > (*nums)[l-1] {
		(*nums)[i] = (*nums)[l-1]
		maxHeapifyRecursive(*nums, i)
	} else {
		heapIncreaseKey(*nums, i, (*nums)[l-1])
	}
	*nums = (*nums)[:l-2]
}

func main() {

	nums := []int{23, 17, 14, 6, 13, 10, 1, 5, 7, 12}
	fmt.Println(nums)
	fmt.Println(isMaxHeap(nums))
	maxHeapifyLoop(nums, 3)
	fmt.Println(nums)
	fmt.Println(isMaxHeap(nums))

	fmt.Println("build")
	nums = []int{12, 1, 10, 6, 13, 14, 17, 5, 7, 23}
	fmt.Println(isMaxHeap(nums))
	buildHeap(nums)
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)

	fmt.Println("sort")
	heapSort(nums)
	fmt.Println(nums)

	fmt.Println("max priority queue")
	nums = []int{23, 13, 17, 7, 12, 14, 10, 5, 6, 1}
	fmt.Println(nums)
	fmt.Printf("outside func: %p\n", &nums)
	fmt.Println(heapExtractMax(&nums))
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)
	fmt.Println(heapExtractMax(&nums))
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)
	fmt.Println(heapExtractMax(&nums))
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)

	fmt.Println("increase key")
	nums = []int{13, 12, 10, 7, 5, 1, 6}
	heapIncreaseKey(nums, 5, 11)
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)

	fmt.Println("insert key")
	nums = []int{13, 12, 10, 7, 5, 1, 6}
	fmt.Println(isMaxHeap(nums))
	heapInsert(&nums, 11)
	fmt.Println(nums)
	fmt.Println(isMaxHeap(nums))

	fmt.Println("delete key")
	nums = []int{15, 7, 9, 1, 2, 3, 8}
	fmt.Println(nums)
	fmt.Println(isMaxHeap(nums))
	heapDelete1(&nums, 4)
	fmt.Println(isMaxHeap(nums))
	fmt.Println(nums)
}
