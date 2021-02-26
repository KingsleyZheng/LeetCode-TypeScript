// 移动零
// 简单

/**
 Do not return anything, modify nums in-place instead.
 */
// 双指针
function moveZeroes(nums: number[]): void {
  let n = nums.length,
    left = 0,
    right = 0;
  while (right < n) {
    if (nums[right] !== 0) {
      [nums[left], nums[right]] = [nums[right], nums[left]];
      left++;
    }
    right++;
  }
}
