package main

func main() {

}

// 简单
// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

// ListNode 链表头结点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 遍历法，遍历比较两条链表的节点，直到其中一条为空
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}
		if l1.Val < l2.Val {
			cur.Next = &ListNode{Val: l1.Val}
			cur = cur.Next
			l1 = l1.Next
		} else {
			cur.Next = &ListNode{Val: l2.Val}
			cur = cur.Next
			l2 = l2.Next
		}
	}
	return dummy.Next
}

// 官方解法1：递归。
func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists1(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists1(l1, l2.Next)
		return l2
	}
}
