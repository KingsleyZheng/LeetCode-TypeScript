// 合并区间
// 中等

// 排序
function merge(intervals: number[][]): number[][] {
  if (intervals.length === 0) {
    return [];
  }
  intervals.sort((a: number[], b: number[]) => a[0] - b[0]);

  const merged: number[][] = [];
  for (let i = 0; i < intervals.length; i++) {
    let L = intervals[i][0];
    let R = intervals[i][1];
    if (merged.length === 0 || merged[merged.length - 1][1] < L) {
      merged.push([L, R]);
    } else {
      merged[merged.length - 1][1] = Math.max(merged[merged.length - 1][1], R);
    }
  }
  return merged;
}

console.log(
  merge([
    [2, 3],
    [2, 2],
    [3, 3],
    [1, 3],
    [5, 7],
    [2, 2],
    [4, 6],
  ])
);
