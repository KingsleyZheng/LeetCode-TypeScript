// 移除元素
// 简单

export function removeElement(nums: number[], val: number): number {
  let ans = 0;
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== val) {
      nums[ans] = nums[i];
      ans++;
    }
  }
  return ans;
}

let a = [0, 1, 2, 2, 3, 0, 4, 2];
console.log(removeElement(a, 2));
console.log(a);
