// 爬楼梯
// 简单

function climbStairs(n: number): number {
  let p: number = 0,
    q: number = 0,
    r: number = 1;
  for (let i = 1; i <= n; ++i) {
    p = q;
    q = r;
    r = p + q;
  }
  return r;
}
