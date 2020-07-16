package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis2(3))
}

// 中等
// 22. 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 暴力法，生成之后再检查是否有效
func generateParenthesis(n int) []string {
	res := []string{}
	if n <= 0 {
		return res
	}
	res = append(res, "(")
	for i := 1; i < 2*n; i++ {
		newRes := []string{}
		for j := 0; j < len(res); j++ {
			newRes = append(newRes, res[j]+"(")
			newRes = append(newRes, res[j]+")")
		}
		res = newRes
	}
	ret := []string{}
	for i := 0; i < len(res); i++ {
		if valid(res[i]) {
			ret = append(ret, res[i])
		}
	}

	return ret
}

func valid(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			balance--
			if balance < 0 {
				return false
			}
		}
	}
	return balance == 0
}

// 官方解法2：回溯法
func generateParenthesis2(n int) []string {
	ans := new([]string)
	backtrack(ans, "", 0, 0, n)
	return *ans
}

func backtrack(ans *[]string, cur string, open, close, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}
	if open < max {
		cur = cur + "("
		backtrack(ans, cur, open+1, close, max)
		cur = cur[:len(cur)-1]
	}
	if close < open {
		cur = cur + ")"
		backtrack(ans, cur, open, close+1, max)
		cur = cur[:len(cur)-1]
	}
}
