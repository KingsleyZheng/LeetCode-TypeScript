// 实现 strStr()
// 简单

function strStr(haystack: string, needle: string): number {
  let L = needle.length,
    n = haystack.length;

  for (let start = 0; start < n - L + 1; ++start) {
    if (haystack.substring(start, start + L) === needle) {
      return start;
    }
  }
  return -1;
}
