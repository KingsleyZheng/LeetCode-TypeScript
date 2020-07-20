package main

import "fmt"

func main() {
	s := "bba"
	fmt.Println(strStr(s, s))
}

// 简单
// 28. 实现 strStr()
// 实现 strStr() 函数。
// 给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

// 相当于 indexOf 函数
// 遍历haystack，如果有相等于needle的子串则返回index
// 类似于官方解法1
func strStr(haystack string, needle string) int {
	nLength := len(needle)
	if nLength == 0 {
		return 0
	}
	hLength := len(haystack)
	if hLength < nLength {
		return -1
	}
	for i := 0; i < hLength-nLength+1; i++ {
		if haystack[i:i+nLength] == needle {
			return i
		}
	}
	return -1
}

// 官方解法2：双指针法。
// 上一个方法会比较所有子串，实际上只有在子串的第一个字符跟needle字符串第一个字符相同时才需要比较
func strStr(haystack string, needle string) int {
	L := len(needle)
	n := len(haystack)
	if l == 0 {
		return0
	}
	pn := 0
	for pn < n-L+1 {
		for pn < n-L+1 && haystack[pn] != needle[0] {
			pn++
		}
		currLen := 0
		pL := 0
		for pL < L && pn < n && haystack[pn] == needle[pL] {
			pn++
			pL++
			currLen++
		}
		if currLen == L {
			return pn - L
		}
		pn = pn - currLen + 1
	}
	return -1
}

// golang示例
// 也是双指针，但是非常简洁，没有多余变量
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}

		}

		if j == len(needle) {
			return i
		}
	}

	return -1
}
