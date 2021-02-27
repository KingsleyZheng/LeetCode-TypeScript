// 全排列
// 中等

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
