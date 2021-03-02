// x的平方根
// 简单

function mySqrt(x: number): number {
  if (x == 0) {
    return 0;
  }
  let ans = Math.exp(0.5 * Math.log(x));
  return (ans + 1) * (ans + 1) <= x ? ans + 1 : ans;
}
