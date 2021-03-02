// 最小路径和
// 中等

// 动态规划
function minPathSum(grid: number[][]): number {
  if (grid === null || grid.length === 0 || grid[0].length === 0) {
    return 0;
  }
  let rows = grid.length;
  let columns = grid[0].length;
  const dp: number[][] = new Array(rows)
    .fill(null)
    .map(() => new Array(columns));
  dp[0][0] = grid[0][0];
  for (let i = 1; i < rows; i++) {
    dp[i][0] = dp[i - 1][0] + grid[i][0];
  }
  for (let j = 1; j < columns; j++) {
    dp[0][j] = dp[0][j - 1] + grid[0][j];
  }
  for (let i = 1; i < rows; i++) {
    for (let j = 1; j < columns; j++) {
      dp[i][j] = Math.min(dp[i - 1][j], dp[i][j - 1]) + grid[i][j];
    }
  }
  return dp[rows - 1][columns - 1];
}
