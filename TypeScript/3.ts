// 无重复字符的最长子串
// 中等

// 滑动窗口，用一个set 存储当前窗口的所有字符
function lengthOfLongestSubstring(s: string): number {
  const occ = new Set<string>();
  const n = s.length;
  let rk = -1,
    ans = 0;
  for (let i = 0; i < n; i++) {
    if (i !== 0) {
      occ.delete(s.charAt(i - 1));
    }
    while (rk + 1 < n && !occ.has(s.charAt(rk + 1))) {
      occ.add(s.charAt(rk + 1));
      rk++;
    }
    ans = Math.max(ans, rk - i + 1);
  }
  return ans;
}
