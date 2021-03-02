// 数组中重复的数字
// 简单

function findRepeatNumber(nums: number[]): number {
  let map = new Map<number, boolean>();
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    if (!map.get(num)) {
      map.set(num, true);
    } else {
      return num;
    }
  }
  return -1;
}
