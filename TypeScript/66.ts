// 加一
// 简单

function plusOne(digits: number[]): number[] {
  let len = digits.length;
  let carry = false;
  while (len--) {
    let num = +digits[len];
    let _num = num;
    if (carry) {
      num++;
      carry = false;
    }
    if (num >= 10 || _num + 1 >= 10) {
      carry = true;
      digits[len] = 0;
      if (len == 0) {
        digits.unshift(1);
      }
      continue;
    }
    digits[len] = _num + 1;
    break;
  }
  return digits;
}
