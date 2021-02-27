// 外观数列
// 简单

function countAndSay(n: number): string {
  if (n <= 0) return "";
  let res = "1";
  while (--n) {
    let cur = "";
    let l = res.length;
    for (let i = 0; i < l; i++) {
      let count = 1;
      while (i + 1 < l && res[i] == res[i + 1]) {
        count++;
        i++;
      }
      cur += count.toString() + res[i];
    }
    res = cur;
  }
  return res;
}
