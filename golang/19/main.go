package main

func main() {

}

// 中等
// 19. 删除链表的倒数第N个节点
// 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
// 说明：
// 给定的 n 保证是有效的。
// 进阶：
// 你能尝试使用一趟扫描实现吗？
// 提示：Maintain two pointers and update one with a delay of n steps.

// ListNode 链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// 看了提示实现的
// 维护双指针，第二个指针延迟N步再动
// 要考虑极限情况：1、只有一个节点；2、删除的是头结点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	left, right := head, head
	index := 0
	for right.Next != nil {
		right = right.Next
		index++
		if index > n {
			left = left.Next
		}
	}
	if index == n-1 {
		return head.Next
	}
	left.Next = left.Next.Next
	return head
}

// 官方解法一：两次遍历算法
// 特别之处，在 head 结点前面插了一个 dummy 结点，排除只有一个结点时会出错的情况
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	length := 0
	first := head
	for first != nil {
		length++
		first = first.Next
	}
	length -= n
	first = dummy
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next
	return dummy.Next
}

// 官方解法二：一次遍历算法，类似我的解法，更简洁
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	first := head
	second := head
	for i := 1; i <= n+1; i++ {
		first = first.Next
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}
