// 括号生成
// 中等

function generateParenthesis(n: number): string[] {
  let ans = new Array<string>();
  backtrack(ans, "", 0, 0, n);
  return ans;
}

function backtrack(
  ans: string[],
  cur: string,
  open: number,
  close: number,
  max: number
) {
  if (cur.length === max * 2) {
    ans.push(cur);
    return;
  }
  if (open < max) {
    cur += "(";
    backtrack(ans, cur, open + 1, close, max);
    cur = cur.slice(0, cur.length - 1);
  }
  if (close < open) {
    cur += ")";
    backtrack(ans, cur, open, close + 1, max);
    cur = cur.slice(0, cur.length - 1);
  }
}
