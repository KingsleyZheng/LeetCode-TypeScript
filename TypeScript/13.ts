// 罗马数字转整数
// 简单

export function romanToInt(s: string): number {
  let sum = 0;
  let preNum = getValue(s.charAt(0));
  for (let i = 1; i < s.length; i++) {
    let num = getValue(s.charAt(i));
    if (preNum < num) {
      sum -= preNum;
    } else {
      sum += preNum;
    }
    preNum = num;
  }
  sum += preNum;
  return sum;
}

function getValue(char: string): number {
  switch (char) {
    case 'I':
      return 1;
    case 'V':
      return 5;
    case 'X':
      return 10;
    case 'L':
      return 50;
    case 'C':
      return 100;
    case 'D':
      return 500;
    case 'M':
      return 1000;
    default:
      return 0;
  }
}
