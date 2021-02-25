// 用队列实现栈
// 简单

// 数字全都存在一个队列里，另一个队列仅在入栈时作为辅助
export class MyStack {
  queue1: number[] = [];
  queue2: number[] = [];

  push(x: number): void {
    this.queue2.push(x);
    while (this.queue1.length) {
      this.queue2.push(this.queue1.shift()!);
    }
    let temp = this.queue1;
    this.queue1 = this.queue2;
    this.queue2 = temp;
  }

  pop(): number {
    if (!this.queue1.length) {
      throw new Error('stack is empty');
    } else return this.queue1.shift()!;
  }

  top(): number {
    return this.queue1[0];
  }

  empty(): boolean {
    return this.queue1.length === 0;
  }
}
