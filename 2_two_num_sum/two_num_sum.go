package main

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	var buf bytes.Buffer
	var cur = l
	for cur != nil {
		if buf.Len() > 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", cur.Val)
		cur = cur.Next
	}
	return buf.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	// 创建链表的头指针和游标
	var res = ListNode{}
	var cur = &res
	for {
		// 两个链表当前节点都不为空, 两者都未结束
		if l1 != nil && l2 != nil {
			v1 := l1.Val
			v2 := l2.Val

			// 当前位的值是两个链表对应节点的值加上进到当前位的值
			sum := v1 + v2 + cur.Val
			v := sum % 10
			c := sum / 10
			cur.Val = v

			// 当前位计算完需要进位时, 提前创建下一节点, 下一节点的Val就是进位值, 并把游标移到下一节点
			if c != 0 {
				cur.Next = &ListNode{Val: c}
				cur = cur.Next
			}

			// 滚动两个链表
			l1 = l1.Next
			l2 = l2.Next

			// 当前位计算完不需要进位时, 两个链表下一个节点只要有一个不为空, 就创建下一个空节点, 并把游标移到下一个节点
			// 空节点的Val为0
			if c == 0 && (l1 != nil || l2 != nil) {
				cur.Next = &ListNode{}
				cur = cur.Next
			}

			continue
		}

		// l2结束了
		if l1 != nil {

			// 只要l1的当前的位值和进位值参与计算就好了
			sum := l1.Val + cur.Val
			v := sum % 10
			c := sum / 10
			cur.Val = v

			// 计算完当前节点, 需要进位, 同样, 创建带有进位值的下一节点, 同时, 滚动l1和游标
			if c != 0 {
				cur.Next = &ListNode{Val: c}
				cur = cur.Next
				l1 = l1.Next
				continue

				// 计算完当前节点, 不需要进位, 如果当前不用进位, 那么后面都不会再进位, 继续循环只是把l1的剩余节点复制到游标后面
				// 这个位置没有使用复制, 而是修改指针, 直接从这个位置把l1后面的节点对接到游标
			} else {
				cur.Next = l1.Next
				l1 = nil
				break
			}
		}

		if l2 != nil {
			sum := l2.Val + cur.Val

			v := sum % 10
			c := sum / 10

			cur.Val = v
			if c != 0 {
				cur.Next = &ListNode{Val: c}
				cur = cur.Next
				l2 = l2.Next
				continue
			} else {
				cur.Next = l2.Next
				l2 = nil
				break
			}
		}

		break
	}

	return &res
}

func main() {
	var b = &ListNode{9, nil}
	var a = &ListNode{9, &ListNode{9, &ListNode{9, nil}}}

	var res = addTwoNumbers(b, a)
	fmt.Println(res)
}
