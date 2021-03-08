// 颜色分类
// 中等

// 单指针，两次遍历
function sortColors(nums: number[]): void {
  let n = nums.length;
  let ptr = 0;
  for (let i = 0; i < n; ++i) {
    if (nums[i] == 0) {
      let temp = nums[i];
      nums[i] = nums[ptr];
      nums[ptr] = temp;
      ++ptr;
    }
  }
  for (let i = ptr; i < n; ++i) {
    if (nums[i] == 1) {
      let temp = nums[i];
      nums[i] = nums[ptr];
      nums[ptr] = temp;
      ++ptr;
    }
  }
}

// 双指针，一次遍历，把0换到前面，把2换到后面
function sortColors2(nums: number[]): void {
  let n = nums.length;
  let p0 = 0,
    p2 = n - 1;
  for (let i = 0; i <= p2; ++i) {
    while (i <= p2 && nums[i] == 2) {
      let temp = nums[i];
      nums[i] = nums[p2];
      nums[p2] = temp;
      --p2;
    }
    if (nums[i] == 0) {
      let temp = nums[i];
      nums[i] = nums[p0];
      nums[p0] = temp;
      ++p0;
    }
  }
}
