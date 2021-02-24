// 反转链表
// 难度：简单
/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

import { ListNode } from "./types";

// 递归
export function reverseList(head: ListNode | null): ListNode | null {
  if (head === null || head.next === null) {
    return head;
  }
  const next = reverseList(head.next);
  head.next.next = head;
  head.next = null;
  return next;
}

// 迭代
function reverseList2(head: ListNode | null): ListNode | null {
  let prev = null;
  let curr = head;
  while (curr !== null) {
    let nextTemp = curr.next;
    curr.next = prev;
    prev = curr;
    curr = nextTemp;
  }
  return prev;
}
