// 83. 删除排序链表中的重复元素
// 简单

import { ListNode } from './types';

export function deleteDuplicates(head: ListNode | null): ListNode | null {
  if (!head) {
    return head;
  }

  let cur = head;
  while (cur.next) {
    if (cur.val === cur.next.val) {
      cur.next = cur.next.next;
    } else {
      cur = cur.next;
    }
  }
  return head;
}
