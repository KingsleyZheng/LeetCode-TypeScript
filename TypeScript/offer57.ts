// 和为s的连续正数序列
// 简单

// 暴力法
function findContinuousSequence(target: number): number[][] {
  let ans: number[][] = [];
  for (let i = 1; i < target / 2; i++) {
    let sum = 0;
    let j = i;
    let list: number[] = [];
    while (sum < target) {
      sum += j;
      list.push(j);
      j++;
    }
    if (sum === target) {
      ans.push(list);
    }
  }
  return ans;
}

// 还可以利用求和公式，使用双指针法

console.log(findContinuousSequence(15))