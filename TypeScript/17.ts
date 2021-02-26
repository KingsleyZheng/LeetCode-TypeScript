// 电话号码的字母组合
// 中等
const map = new Map<string, string[]>([
  ['2', ['a', 'b', 'c']],
  ['3', ['d', 'e', 'f']],
  ['4', ['g', 'h', 'i']],
  ['5', ['j', 'k', 'l']],
  ['6', ['m', 'n', 'o']],
  ['7', ['p', 'q', 'r', 's']],
  ['8', ['t', 'u', 'v']],
  ['9', ['w', 'x', 'y', 'z']]
]);

function letterCombinations(digits: string): string[] {
  let ans: string[] = [];

  if (!digits.length) {
    return ans;
  }
  helpFunc(digits.charAt(0), digits.slice(1), '', ans);

  return ans;
}

function helpFunc(number: string, rest: string, prefix: string, ans: string[]) {
  const charList = map.get(number);
  charList!.forEach((char) => {
    const word = prefix + char;
    if (rest.length === 0) {
      ans.push(word);
    } else {
      helpFunc(rest.charAt(0), rest.slice(1), word, ans);
    }
  });
}

console.log(letterCombinations('563'));
