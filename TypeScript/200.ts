// 岛屿数量
// 中等

// 类似于图的题目
// 直奔题解的一题
// 深度优先遍历
// 扫描整个二维网络，如果遇到1则开始深度优先搜索，搜索过程中将1变为0，防止重复搜索。
// 最终岛屿的数量就是我们进行深度优先搜索的次数
function dfs(grid: string[][], r: number, c: number) {
  let nr = grid.length;
  let nc = grid[0].length;
  if (r < 0 || c < 0 || r >= nr || c >= nc || grid[r][c] === '0') {
    return;
  }
  grid[r][c] = '0';
  dfs(grid, r - 1, c);
  dfs(grid, r + 1, c);
  dfs(grid, r, c - 1);
  dfs(grid, r, c + 1);
}

function numIslands(grid: string[][]): number {
  if (grid === null || grid.length === 0) {
    return 0;
  }
  let nr = grid.length;
  let nc = grid[0].length;
  let num_islands = 0;
  for (let r = 0; r < nr; r++) {
    for (let c = 0; c < nc; c++) {
      if (grid[r][c] === '1') {
        num_islands++;
        dfs(grid, r, c);
      }
    }
  }
  return num_islands;
}
