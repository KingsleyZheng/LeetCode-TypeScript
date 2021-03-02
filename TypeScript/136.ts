// 只出现一次的数字
// 简单

// 任何数和其自身做异或运算，结果为0
// 和0做异或运算，结果为自身
// 异或运算满足交换律
// 额外空间为常数级的解法

function singleNumber(nums: number[]): number {
  let single = 0;
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    single ^= num;
  }
  return single;
}
