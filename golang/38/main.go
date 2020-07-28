package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(5))
}

// 简单
// 38. 外观数列
// 给定一个正整数 n（1 ≤ n ≤ 30），输出外观数列的第 n 项。
// 注意：整数序列中的每一项将表示为一个字符串。
// 「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。前五项如下：

// 依次迭代即可
func countAndSay(n int) string {
	ans := "1"
	for n > 1 {
		ans = say(ans)
		n--
	}
	return ans
}
func say(s string) string {
	ans := ""
	n := len(s)
	cur := 0
	for i := 0; i < n; i++ {
		if s[cur] != s[i] {
			ans += strconv.Itoa(i - cur)
			ans += string(s[cur])
			cur = i
		}
	}
	ans += strconv.Itoa(n - cur)
	ans += string(s[cur])
	return ans
}

// golang 耗时最低解法
func countAndSay1(n int) string {
	if n == 1 {
		return "1"
	}

	s := "1"
	for i := 2; i <= n; i++ {
		var temp strings.Builder
		count := 1
		current := s[0]

		for j := 1; j < len(s); j++ {
			if current == s[j] {
				count++
			} else {
				temp.WriteString(strconv.Itoa(count))
				temp.WriteByte(current)
				count = 1
				current = s[j]
			}
		}

		temp.WriteString(strconv.Itoa(count))
		temp.WriteByte(current)
		s = temp.String()
	}
	return s
}
