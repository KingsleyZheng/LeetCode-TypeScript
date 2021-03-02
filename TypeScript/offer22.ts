// 链表中倒数第k个节点
// 简单

import { ListNode } from './types';

function getKthFromEnd(head: ListNode | null, k: number): ListNode | null {
  let cur: ListNode | null = head;
  while (k--) {
    cur = cur!.next;
  }
  while (cur !== null) {
    cur = cur.next;
    head = head!.next;
  }
  return head;
}
