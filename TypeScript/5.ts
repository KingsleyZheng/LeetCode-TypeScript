// 最长回文子串
// 中等

export function longestPalindrome(s: string): string {
  const n = s.length;
  const dp: boolean[][] = [];
  for (let i = 0; i < n; i++) {
    dp.push(new Array(n))    
  }
  let ans = '';
  for (let l = 0; l < n; l++) {
    for (let i = 0; i + l < n; i++) {
      let j = i + l;
      if (l === 0) {
        dp[i][j] = true;
      } else if (l === 1) {
        dp[i][j] = s.charAt(i) === s.charAt(j);
      } else {
        dp[i][j] = s.charAt(i) === s.charAt(j) && dp[i + 1][j - 1];
      }
      if (dp[i][j] && l + 1 > ans.length) {
        ans = s.substring(i, i + l + 1);
      }
    }
  }
  return ans;
}
