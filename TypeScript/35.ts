// 搜索插入位置
// 简单

function searchInsert(nums: number[], target: number): number {
  if (nums.length === 0) return 0;
  let mid = Math.floor((nums.length - 1) / 2);
  let left = 0;
  let right = nums.length - 1;
  while (left <= right) {
    if (nums[mid] === target) {
      return mid;
    } else if (nums[mid] < target) {
      left = mid + 1;
      mid = Math.floor((left + right) / 2);
    } else {
      right = mid - 1;
      mid = Math.floor((left + right) / 2);
    }
  }
  return mid + 1;
}

console.log(searchInsert([1, 3, 5, 6], 0));
