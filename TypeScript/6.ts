// z字形变换
// 中等

// 方法一 按行排序
function convert(s: string, numRows: number): string {
  if (numRows === 1) return s;
  let rows: string[] = [];
  for (let i = 0; i < Math.min(numRows, s.length); i++) {
    rows.push("");
  }
  let curRow = 0;
  let goingDown = false;

  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    rows[curRow] += char;
    if (curRow === 0 || curRow === numRows - 1) {
      goingDown = !goingDown;
    }
    curRow += goingDown ? 1 : -1;
  }
  let ret = "";
  for (let i = 0; i < rows.length; i++) {
    const row = rows[i];
    ret += row;
  }
  return ret;
}


console.log(convert('PAYPALISHIRING', 4));