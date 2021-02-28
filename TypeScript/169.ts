// 多数元素
// 简单

function majorityElement(nums: number[]): number {
  let count = 0;
  let candidate: number;

  for (let num of nums) {
    if (count == 0) {
      candidate = num;
    }
    count += num === candidate! ? 1 : -1;
  }

  return candidate!;
}
