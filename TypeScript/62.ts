// 不同路径
// 中等

// 动态规划
function uniquePaths(m: number, n: number): number {
  const f: number[][] = new Array(m).fill(0).map(() => new Array(n).fill(0));
  for (let i = 0; i < m; i++) {
    f[i][0] = 1;
  }
  for (let j = 0; j < n; j++) {
    f[0][j] = 1;
  }
  for (let i = 1; i < m; i++) {
    for (let j = 1; j < n; j++) {
      f[i][j] = f[i - 1][j] + f[i][j - 1];
    }
  }
  return f[m - 1][n - 1];
}
