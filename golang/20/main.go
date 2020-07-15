package main

import "fmt"

func main() {
	s := "{[()]}"
	fmt.Println(isValid(s))
}

// 简单
// 20. 有效的括号
// 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
// 有效字符串需满足：
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
// 注意空字符串可被认为是有效字符串。

// 利用额外切片，看了提示之后实现的。官方解法也是这个
func isValid(s string) bool {
	slice := []byte{}
	length := len(s)
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	for i := 0; i < length; i++ {
		sLength := len(slice)
		if sLength > 0 && m[slice[sLength-1]] == s[i] {
			slice = slice[0 : sLength-1]
		} else {
			slice = append(slice, s[i])
		}
	}
	return len(slice) == 0
}
