// 从尾到头打印链表

import { ListNode } from './types';

function reversePrint(head: ListNode | null): number[] {
  const stack: ListNode[] = [];
  let temp = head;
  while (temp !== null) {
    stack.push(temp);
    temp = temp.next;
  }
  let length = stack.length;
  const print: number[] = [];
  for (let i = length - 1; i >= 0; i--) {
    const element = stack[i];
    print.push(element.val);
  }
  return print;
}
