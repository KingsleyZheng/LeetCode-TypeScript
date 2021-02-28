// 全排列
// 中等

// 方法一 官方题解
function permute(nums: number[]): number[][] {
  const res: number[][] = [];
  const output: number[] = [];
  for (let i = 0; i < nums.length; i++) {
    const num = nums[i];
    output.push(num);
  }
  let n = nums.length;
  backtrack(n, output, res, 0);
  return res;
}

function backtrack(
  n: number,
  output: number[],
  res: number[][],
  first: number
) {
  if (first === n) {
    res.push(Array.from(output));
  }
  for (let i = first; i < n; i++) {
    [output[first], output[i]] = [output[i], output[first]];
    backtrack(n, output, res, first + 1);
    [output[first], output[i]] = [output[i], output[first]];
  }
}

// 方法二 更好理解的版本(深度优先遍历)
function permute2(nums: number[]): number[][] {
  let len = nums.length;
  let res: number[][] = [];
  if (len === 0) {
    return res;
  }
  let used: boolean[] = new Array(len);
  let path: number[] = new Array();
  dfs(nums, len, 0, path, used, res);
  return res;
}

function dfs(
  nums: number[],
  len: number,
  depth: number,
  path: number[],
  used: boolean[],
  res: number[][]
) {
  if (depth === len) {
    res.push(Array.from(path));
    return;
  }
  for (let i = 0; i < len; i++) {
    if (!used[i]) {
      path.push(nums[i]);
      used[i] = true;

      dfs(nums, len, depth + 1, path, used, res);

      used[i] = false;
      path.pop();
    }
  }
}

console.log(permute2([1, 2, 3, 4]));
