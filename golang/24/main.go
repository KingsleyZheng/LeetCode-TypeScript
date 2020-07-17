package main

func main() {

}

// 中等
// 24. 两两交换链表中的节点
// 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
// 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

type ListNode struct {
	Val  int
	Next *ListNode
}

// 暴力法，先分解再组合(超出时间限制)
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	single, double := []*ListNode{}, []*ListNode{}
	cur := head
	index := 0
	for cur != nil {
		if index%2 == 0 {
			single = append(single, cur)
		} else {
			double = append(double, cur)
		}
		index++
		cur = cur.Next
	}
	ret := &ListNode{}
	cur = ret
	n := len(single) + len(double)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			cur.Next = double[i/2]
		} else {
			cur.Next = single[i/2]
		}
		cur = cur.Next
	}
	return ret.Next
}

// 官方解法一：递归。很巧妙的方法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode := head
	secondNode := head.Next
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

// 官方解法二：迭代。对指针的运用写的太好了
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	prevNode := dummy

	for head != null && head.Next != null {
		firstNode := head
		secondNode := head.Next

		prevNode.Next = secondNode
		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode

		prevNode = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}
