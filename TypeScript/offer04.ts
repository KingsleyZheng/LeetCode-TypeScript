// 二维数组中的查找
// 中等

// 从右上角开始查找
function findNumberIn2DArray(matrix: number[][], target: number): boolean {
  if (matrix == null || matrix.length == 0 || matrix[0].length == 0) {
    return false;
  }
  let rows = matrix.length,
    columns = matrix[0].length;
  let row = 0,
    column = columns - 1;
  while (row < rows && column >= 0) {
    let num = matrix[row][column];
    if (num == target) {
      return true;
    } else if (num > target) {
      column--;
    } else {
      row++;
    }
  }
  return false;
}
