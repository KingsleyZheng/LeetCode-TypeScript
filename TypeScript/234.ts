// 回文链表
// 简单

import { ListNode } from './types';

// 快慢指针

const reverseList = (head: ListNode | null) => {
  let prev = null;
  let curr = head;
  while (curr !== null) {
    let nextTemp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = nextTemp;
  }
  return prev;
};

const endOfFirstHalf = (head: ListNode) => {
  let fast = head;
  let slow = head;
  while (fast.next !== null && fast.next.next !== null) {
    fast = fast.next.next;
    slow = slow.next!;
  }
  return slow;
};

function isPalindrome(head: ListNode | null): boolean {
  if (head == null) return true;

  // 找到前半部分链表的尾节点并反转后半部分链表
  const firstHalfEnd = endOfFirstHalf(head);
  const secondHalfStart = reverseList(firstHalfEnd.next);

  // 判断是否回文
  let p1 = head;
  let p2 = secondHalfStart;
  let result = true;
  while (result && p2 != null) {
    if (p1.val != p2.val) result = false;
    p1 = p1.next!;
    p2 = p2.next;
  }

  // 还原链表并返回结果
  firstHalfEnd.next = reverseList(secondHalfStart);
  return result;
}
