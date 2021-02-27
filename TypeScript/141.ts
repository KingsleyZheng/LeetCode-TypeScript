// 环形链表
// 简单

import { ListNode } from "./types";

// 方法一 哈希表
function hasCycle(head: ListNode | null): boolean {
  const seen = new Set<ListNode>();
  while (head !== null) {
    if (seen.has(head)) {
      return true;
    }
    head = head.next;
  }
  return false;
}

// 方法二 快慢指针
function hasCycle2(head: ListNode | null): boolean {
  if (head === null || head.next === null) {
    return false;
  }
  let slow = head;
  let fast = head.next;
  while (slow !== fast) {
    if (fast === null || fast.next === null) {
      return false;
    }
    slow = slow.next!;
    fast = fast.next.next!;
  }
  return true;
}
