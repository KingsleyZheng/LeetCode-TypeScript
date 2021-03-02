// 子集
// 中等

// 迭代法
function subsets(nums: number[]): number[][] {
  const ans = [];
  const n = nums.length;
  for (let mask = 0; mask < 1 << n; ++mask) {
    const t = [];
    for (let i = 0; i < n; ++i) {
      if (mask & (1 << i)) {
        t.push(nums[i]);
      }
    }
    ans.push(t);
  }
  return ans;
}

// 递归
function subsets2(nums: number[]): number[][] {
  const t: number[] = [];
  const ans: number[][] = [];
  const n = nums.length;
  const dfs = (cur: number, nums: number[]) => {
    if (cur === nums.length) {
      ans.push(t.slice());
      return;
    }
    t.push(nums[cur]);
    dfs(cur + 1, nums);
    t.pop();
    dfs(cur + 1, nums);
  };
  dfs(0, nums);
  return ans;
}
