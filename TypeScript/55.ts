// 跳跃游戏
// 中等

// 贪心算法：逐步更新可以到达的最远位置
function canJump(nums: number[]): boolean {
  let n = nums.length;
  let rightmost = 0;
  for (let i = 0; i < n; ++i) {
    if (i <= rightmost) {
      rightmost = Math.max(rightmost, i + nums[i]);
      if (rightmost >= n - 1) {
        return true;
      }
    }
  }
  return false;
}
