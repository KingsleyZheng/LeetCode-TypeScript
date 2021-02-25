// 删除排序数组中的重复项
// 简单
// 给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。

// 原地修改，双指针法。一个快指针，一个慢指针
export function removeDuplicates(nums: number[]): number {
  if (nums.length === 0) {
    return 0;
  }
  let i = 0;
  for (let j = 1; j < nums.length; j++) {
    if (nums[j] !== nums[i]) {
      i++;
      nums[i] = nums[j];
    }
  }
  return i + 1;
}

console.log(removeDuplicates([1, 1]));
