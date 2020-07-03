package main

// 难度：中等
// 2.两数相加
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode singly-linked
type ListNode struct {
	Val  int
	Next *ListNode
}

// 自己想的，利用了递归
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{0, nil}
	var next *ListNode
	var cur *ListNode = l3
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	add := 0
	for {
		if l1 == nil && l2 == nil {
			if add != 0 {
				cur.Next = &ListNode{add, nil}
			}
			break
		} else if l1 == nil {
			if add != 0 {
				cur.Next = addTwoNumbers(l2, &ListNode{add, nil})
			} else {
				cur.Next = l2
			}
			break
		} else if l2 == nil {
			if add != 0 {
				cur.Next = addTwoNumbers(l1, &ListNode{add, nil})
			} else {
				cur.Next = l1
			}
			break
		}
		sum := l1.Val + l2.Val + add
		add = sum / 10
		next = &ListNode{sum % 10, nil}
		cur.Next = next
		cur = next
		l1 = l1.Next
		l2 = l2.Next
	}
	return l3.Next
}

// 官方方法一：初等数学。关键在于，若已到达l1或l2尾部，用0替代
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{0, nil}
	p := l1
	q := l2
	curr := dummyHead
	carry := 0
	for {
		if p == nil && q == nil {
			break
		}
		var x, y int
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return dummyHead.Next
}

func main() {

}
