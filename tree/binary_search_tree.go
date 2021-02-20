package main

import (
	"fmt"
)

/*

二叉搜索树

高度为h, 节点数为n

1. 节点关系:
对任意节点x, x.L.key <= x.Key <= x.R.key

2. 基本操作的复杂度:
search: 沿着子节点边向下 O(h)
minimum: 沿着左边向下 O(h)
maximum: 沿着右边向下 O(h)
successor: 沿着父节点向上 O(h)
predecessor: 沿着父节点向上 O(h)
insert: 小于当前节点值时沿左边下降, 大于当前节点值时沿右边下降, 直到为nil时, 新建一个节点拼接到末尾O(h)
delete: 左边为nil, 右边为nil, 左右边均不为nil时找后继节点O(h)

*/

type Stack []*Node

func (s *Stack) push(x *Node) {
	*s = append(*s, x)
}
func (s *Stack) empty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
func (s *Stack) pop() *Node {
	l := len(*s)
	if l == 0 {
		return nil
	}
	r := (*s)[l-1]
	*s = (*s)[:l-1]
	return r
}

type Node struct {
	Key int
	L   *Node
	R   *Node
	P   *Node
}

// 中序遍历, 递归版本
// 时间复杂度: O(n)
func (t *Node) inorderWalkRecursive() {
	if t != nil {
		t.L.inorderWalkRecursive()
		fmt.Println(t.Key)
		t.R.inorderWalkRecursive()
	}
}

// 中序遍历, 非递归版本
// 时间复杂度: O(n)
func (t *Node) inorderWalkLoop() {

	s := Stack{}
	cur := t

	for {
		if cur != nil {
			s.push(cur)
			cur = cur.L
		} else {
			if s.empty() {
				break
			}
			cur = s.pop()
			fmt.Println(cur.Key)
			cur = cur.R
		}
	}

}

// 辅助函数, 构建父节点关系
func (t *Node) constructParent() {
	if t != nil {
		if t.L != nil {
			t.L.P = t
		}
		t.L.constructParent()
		if t.R != nil {
			t.R.P = t
		}
		t.R.constructParent()
	}
}

// 搜索, 递归版本, 遍历路径是沿着根到目标节点中间的边向下的
// 时间复杂度: O(h)
func (t *Node) searchRecursive(k int) *Node {
	if t == nil {
		return nil
	}
	if t.Key == k {
		return t
	}
	if k < t.Key {
		return t.L.searchRecursive(k)
	}
	return t.R.searchRecursive(k)
}

// 搜索, 非递归版本
func (t *Node) searchLoop(k int) *Node {

	cur := t
	for {
		if k == cur.Key {
			return cur
		} else if k < cur.Key {
			cur = cur.L
		} else {
			cur = cur.R
		}
		if cur == nil {
			return nil
		}
	}
}

// 最小值, 一直沿着左边查找
func (t *Node) minimum() *Node {
	cur := t
	for {
		if cur.L == nil {
			return cur
		}
		cur = cur.L
	}
}
func (t *Node) minimumRecursive() *Node {
	if t.L == nil {
		return t
	}
	return t.L.minimumRecursive()
}

// 最大值, 一直沿着右边查找
func (t *Node) maximum() *Node {
	cur := t
	for {
		if cur.R == nil {
			return cur
		}
		cur = cur.R
	}
}
func (t *Node) maximumRecursive() *Node {
	if t.R == nil {
		return t
	}
	return t.R.maximumRecursive()
}

// 后继节点, 找到比t大的最小节点
// 如果当前节点是有右分支的, 那么就是右分支中最小的节点
// 如果没有右分支, 沿着父节点向上, 直到找到一个左分支, 那这个左分支关系中的父节点就是后继节点
// 原理就是, 当前节点没有右分支的话, 那么它一定是某个节点左分支的子树中最大值, 沿着父节点找到这对左分支关系即可
// 没有的话, 说明当前节点是最大节点
func (t *Node) successor() *Node {
	if t.R != nil {
		fmt.Println(t.R.minimum())
		return t.R.minimum()
	}

	cur := t
	p := t.P
	for {
		if p == nil || p.L == cur {
			return p
		}
		cur = p
		p = p.P
	}
}

// 先驱节点, 和successor对称, 找左子树上的最大值, 或者沿着父节点向上找一对右分支关系
func (t *Node) predecessor() *Node {
	if t.L != nil {
		return t.L.maximum()
	}

	cur := t
	p := t.P
	for {
		if p == nil || p.R == cur {
			return p
		}
		cur = p
		p = p.P
	}
}

// 先从t沿着子节点边向下，x小于t时沿左边, 否则沿着右边
// 直到nil节点, 把这个nil节点替换为新节点
func (t *Node) insert(x int) {

	// t=nil时, 方法无法修改t指针的值, 所以不允许向空节点插入新节点
	if t == nil {
		panic("empty tree.")
	}

	n := &Node{x, nil, nil, nil}
	cur := t
	var pre *Node

	for {
		if cur == nil {
			break
		}
		pre = cur
		if n.Key < cur.Key {
			cur = cur.L
		} else {
			cur = cur.R
		}
	}

	n.P = pre
	if n.Key < pre.Key {
		pre.L = n
	} else {
		pre.R = n
	}
}

// 插入的递归版本, 由于t到nil位置时无法回到parent节点, 所以每次先确认t的下一节点是否为空在进行递归
func (t *Node) insertRecursive(x int) {
	if x < t.Key {
		if t.L != nil {
			t.L.insertRecursive(x)
		} else {
			n := &Node{x, nil, nil, nil}
			n.P = t
			t.L = n
		}
	} else {
		if t.R != nil {
			t.R.insertRecursive(x)
		} else {
			n := &Node{x, nil, nil, nil}
			n.P = t
			t.R = n
		}
	}
}

// 删除分为三种情况:
// 删除z, zp为z的parent节点, zs为孩子节点, zl为z的左节点, zr为z的右节点
// 如果z没有左孩子, 那么直接移植z的右孩子到z节点, 即使z的右孩子为nil
// 如果z没有右孩子, 那么直接移植z的左孩子到z节点, 即使z的左孩子为nil
// 否则, 找到z的后继y, y一定是没有左孩子的
// 		如果y不是z的右孩子, 先把y的右孩子移植到y位置, 这样就提出y节点, 把y节点拼接到z右孩子的前面, 这时y就可以直接移植到z位置
// 		如果y是z的右孩子, 直接把y移植到z
//		由于y是没有左孩子的, 移植后需要把z原来的左孩子拼接到y的左孩子
// 时间复杂度: 除了查找z的后继是O(h)的, 其它过程都是常数时间所以delete的时间复杂度是O(h)
func (t *Node) delete(z *Node) {
	if z.L == nil {
		transplant(z, z.R)
	} else if z.R == nil {
		transplant(z, z.L)
	} else {
		y := z.successor()
		if y != z {
			transplant(y, y.R)
			y.R = z.R
			z.R.P = y
		}
		transplant(z, y)
		y.L = z.L
		z.L.P = y
	}
}

func transplant(old, new *Node) {
	if old == nil {
		panic("old node is nil")
	}

	if old.P == nil {
		*old = *new
	} else if old.P.L == old {
		old.P.L = new
	} else if old.P.R == old {
		old.P.R = new
	}

	if new != nil {
		new.P = old.P
	}
}

func main() {

	tree := &Node{15,
		&Node{6,
			&Node{3,
				&Node{2, nil, nil, nil},
				&Node{4, nil, nil, nil},
				nil},
			&Node{7,
				nil,
				&Node{13,
					&Node{9, nil, nil, nil},
					nil,
					nil},
				nil},
			nil},
		&Node{18,
			&Node{17, nil, nil, nil},
			&Node{20, nil, nil, nil},
			nil},
		nil}
	tree.constructParent()
	fmt.Println("construction completed.")

	fmt.Println("walk recursive")
	tree.inorderWalkRecursive()
	fmt.Println("walk loop")
	tree.inorderWalkLoop()

	fmt.Println(tree.searchRecursive(7))
	fmt.Println(tree.searchLoop(7))
	fmt.Println(tree.maximum() == tree.maximumRecursive())
	fmt.Println(tree.minimum() == tree.minimumRecursive())

	fmt.Println("find successor and predecessor")
	n18 := tree.R
	fmt.Println(n18.successor())
	fmt.Println(n18.predecessor())
	n7 := tree.L.R
	fmt.Println(n7.successor())
	fmt.Println(n7.predecessor())
	n13 := n7.R
	fmt.Println(n13.successor())
	fmt.Println(n13.predecessor())
	n20 := tree.R.R
	fmt.Println(n20.successor())
	fmt.Println(n20.predecessor())
	n9 := n13.L
	fmt.Println(n9.successor())
	fmt.Println(n9.predecessor())

	//ntree := &Node{nil, nil, nil, nil}
	//tree.insert(23)
	//tree.insert(5)
	//tree.insert(7)
	tree.insertRecursive(23)
	tree.insertRecursive(5)
	tree.insertRecursive(7)

	//ntree := &Node{17,
	//				 &Node{1, nil, nil, nil},
	//				 &Node{19, nil, nil, nil},
	//				 nil}
	//ntree.constructParent()
	//transplant(tree.L.L,ntree)
	//fmt.Println(tree)

	tree.delete(tree.L.L.R.R) // 5
	tree.delete(tree.R.R)     // 20
	tree.delete(tree.L.R.R)   // 13
	tree.delete(tree.R)       // 18

}
