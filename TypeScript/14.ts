// 最长公共前缀
// 简单

// 暴力法
function longestCommonPrefix(strs: string[]): string {
  if (strs.length === 0) return ''
  if (strs.length === 1) return strs[0]
  const base = strs[0]
  let ans = ''
  for (let i = 0; i < base.length; i++) {
    const char = base.charAt(i)
    for (let j = 1; j < strs.length; j++) {
      const str = strs[j];
      if (!(str.charAt(i) ===  char)) {
        return ans
      }
    }
    ans += char
  }
  return ans
};