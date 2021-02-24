// 两数之和
// 难度：简单
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 的那 两个 整数，并返回它们的数组下标。

// 解法：利用hash表
export function twoSum(nums: number[], target: number): number[] {
  let map = new Map<number, number>();
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    const findIndex = map.get(target - num);
    if (findIndex !== undefined) {
      return [findIndex, i];
    }
    map.set(num, i);
  }
  return [];
}
