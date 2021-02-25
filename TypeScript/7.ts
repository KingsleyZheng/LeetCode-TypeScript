// 整数反转
// 简单

export function reverse(x: number): number {
  let rev = 0;
  const flag = x > 0 ? 1 : -1;
  x = Math.abs(x);
  while (x !== 0) {
    let pop = x % 10;
    x = Math.floor(x / 10);
    if (rev > 2 ** 31 - 1) {
      return 0;
    }
    rev = rev * 10 + pop;
  }
  return rev > 2 ** 31 - 1 ? 0 : rev * flag;
}

console.log(reverse(1534236469));
