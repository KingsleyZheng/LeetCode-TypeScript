package main

import "fmt"

func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s))
}

// 简单
// 13. 罗马数字转整数
// 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。

// 看了题解的思路，再实现的代码
// 遍历一遍，如果数字比右边的大或者相等，则相加，否则减去
// 时间复杂度O(n), 空间复杂度O(1)
func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ret := 0
	l := len(s)
	for i := 0; i < l; i++ {
		if i == l-1 {
			ret += m[s[i]]
		} else {
			if m[s[i]] < m[s[i+1]] {
				ret -= m[s[i]]
			} else {
				ret += m[s[i]]
			}
		}
	}
	return ret
}
