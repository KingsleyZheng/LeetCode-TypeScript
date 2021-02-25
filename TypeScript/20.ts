// 有效的括号
// 简单

// 利用栈
export function isValid(s: string): boolean {
  const map : any = {
    ')': '(',
    ']': '[',
    '}': '{',
  }
  const stack: string[] = []
  for (let i = 0; i < s.length; i++) {
    const char = s.charAt(i);
    if (['(', '[', '{'].includes(char)) {
      stack.push(char)
    } else {
      if (stack.pop() === map[char]) {
        continue
      } else {
        return false
      }
    }
  }
  return stack.length === 0
}

console.log(isValid(")))"))