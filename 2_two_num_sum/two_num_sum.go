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

	//var c int
	var res = ListNode{}
	var cur *ListNode
	cur = &res
	for {
		if l1 != nil && l2 != nil {
			v1 := l1.Val
			v2 := l2.Val

			sum := v1 + v2 + cur.Val

			v := sum % 10
			c := sum / 10

			cur.Val = v
			if c != 0 {
				cur.Next = &ListNode{Val: c}
				cur = cur.Next
			}

			l1 = l1.Next
			l2 = l2.Next

			if c == 0 && (l1 != nil || l2 != nil) {
				cur.Next = &ListNode{}
				cur = cur.Next
			}

			continue
		}

		if l1 != nil {
			sum := l1.Val + cur.Val

			v := sum % 10
			c := sum / 10

			cur.Val = v
			if c != 0 {
				cur.Next = &ListNode{Val: c}
				cur = cur.Next
				l1 = l1.Next
				continue
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
