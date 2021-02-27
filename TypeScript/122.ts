// 买卖股票的最佳时机 II
// 简单

// 贪心算法
function maxProfit(prices: number[]): number {
  let ans = 0;
  let n = prices.length;
  for (let i = 1; i < n; ++i) {
    ans += Math.max(0, prices[i] - prices[i - 1]);
  }
  return ans;
}
