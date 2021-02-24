// 两数相加
// 中等
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

import { ListNode } from './types';

export function addTwoNumbers(
  l1: ListNode | null,
  l2: ListNode | null
): ListNode | null {
  if (l1 === null) return l2;
  if (l2 === null) return l1;
  let head = new ListNode();
  let cur: ListNode | null = head;
  let cur1: ListNode | null = l1;
  let cur2: ListNode | null = l2;
  let isEnter = false;
  while (cur1 !== null || cur2 !== null) {
    let next = new ListNode();
    cur.next = next;
    cur = next;

    const val1: number = cur1 ? cur1.val : 0;
    const val2: number = cur2 ? cur2.val : 0;
    const total: number = val1 + val2 + (isEnter ? 1 : 0);
    cur.val = total % 10;
    isEnter = total >= 10;
    if (cur1 !== null) {
      cur1 = cur1.next;
    }
    if (cur2 !== null) {
      cur2 = cur2.next;
    }
    if (cur1 === null && cur2 === null && isEnter) {
      cur.next = new ListNode(1);
    }
  }

  return head.next;
}
