// 删除链表倒数第n个结点
// 中等

import { ListNode } from './types';

// 双指针法
// first 比 second 超前n个结点，first到结尾时，second的下一个结点就是要清除的节点
export function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
  let dummy = new ListNode(0, head);
  let first = head;
  let second = dummy;
  for (let i = 0; i < n; i++) {
    first = first!.next;
  }
  while (first !== null) {
    first = first.next;
    second = second.next!;
  }
  second.next = second.next!.next;
  let ans = dummy.next;
  return ans;
}
