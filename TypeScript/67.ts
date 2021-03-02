// 二进制求和
// 简单

// 模拟加法
function addBinary(a: string, b: string): string {
  let res = '';
  let na = a.length;
  let nb = b.length;
  let n = Math.max(na, nb);
  let carry = false; // 是否需要进位
  // 在字符串左边补全0至ab长度相等
  if (na > nb) {
    for (let i = 0; i < na - nb; i++) {
      b = '0' + b;
    }
  } else if (na < nb) {
    for (let i = 0; i < nb - na; i++) {
      a = '0' + a;
    }
  }
  //
  for (let i = n - 1; i >= 0; i--) {
    let tmp = 0;
    if (carry) {
      tmp = Number(a[i]) - 0 + (Number(b[i]) - 0) + 1;
    } else {
      tmp = Number(a[i]) - 0 + (Number(b[i]) - 0);
    }
    if (tmp == 0) {
      res = '0' + res;
      carry = false;
    } else if (tmp == 1) {
      res = '1' + res;
      carry = false;
    } else if (tmp == 2) {
      res = '0' + res;
      carry = true;
    } else if (tmp == 3) {
      res = '1' + res;
      carry = true;
    }
  }
  if (carry) res = '1' + res;
  return res;
}
