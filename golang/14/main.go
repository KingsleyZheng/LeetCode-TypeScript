package main

import "fmt"

func main() {
	strs := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(strs))
}

// 简单
// 14. 最长公共前缀
// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。
// 说明:所有输入只包含小写字母 a-z 。

// 自己实现：将第一个字符串作为基准值，逐个比较。要注意切片或者字符串下标越界的情况
func longestCommonPrefix(strs []string) string {
	ret := ""
	l := len(strs)
	if l == 0 {
		return ""
	}
	if l == 1 {
		return strs[0]
	}
	s0 := strs[0]
	l0 := len(s0)
	if l0 <= 0 {
		return ret
	}
label1:
	for i := 0; i < l0; i++ {
		for j := 1; j < len(strs); j++ {
			if i > len(strs[j])-1 {
				break label1
			}
			if s0[i] != strs[j][i] {
				break label1
			}
		}
		ret += string(s0[i])
	}
	return ret
}

// go范例。逐个找出最短的字符串，再去比较，就能防止数组越界
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		s := len(res)
		if s > len(strs[i]) {
			s = len(strs[i])
			res = res[:s]
		}
		for j := 0; j < s; j++ {
			if res[j] != strs[i][j] {
				res = res[:j]
				break
			}
		}
	}
	return res
}

// 官方解法一：横向扫描
// 基于该结论，可以得到一种查找字符串数组中的最长公共前缀的简单方法。依次遍历字符串数组中的每个字符串，对于每个遍历到的字符串，更新最长公共前缀，当遍历完所有的字符串以后，即可得到字符串数组中的最长公共前缀。
// 如果在尚未遍历完所有的字符串时，最长公共前缀已经是空串，则最长公共前缀一定是空串，因此不需要继续遍历剩下的字符串，直接返回空串即可。
func longestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := min(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 官方解法二：纵向扫描
func longestCommonPrefix3(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
