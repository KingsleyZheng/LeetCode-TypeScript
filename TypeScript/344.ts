// 反转字符串
// 简单

function reverseString(s: string[]): void {
  const n = s.length;
  for (let left = 0, right = n - 1; left < right; ++left, --right) {
    [s[left], s[right]] = [s[right], s[left]];
  }
}
