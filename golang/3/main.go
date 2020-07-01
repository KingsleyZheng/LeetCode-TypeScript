package main

// 3. 无重复字符的最长子串
// 给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

// 自己想的：暴力解，通过两层循环，遍历字符串，找出最长的子串
// 因为有 dvdf 这种情况出现，所以需要这样做
// 应该是对的，但运行超时。。。。
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	if length <= 1 {
		return length
	}
	max := 0
	curr := 0
	map1 := map[byte]int{}
	for i := 0; i < length; i++ {
		if max > length-i {
			break
		}
		curr = 0
		map1 = map[byte]int{}
		for j := i; j < length; j++ {
			if _, ok := map1[s[j]]; ok {
				map1 = map[byte]int{}
				if curr > max {
					max = curr
				}
				curr = 0
			}
			map1[s[j]] = j
			curr++
		}
		if curr > max {
			max = curr
		}
	}

	return max
}

// 官方解法：采用双指针的方法，利用rk一定越来越右的特点

// 辅助函数
func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
func lengthOfLongestSubstring1(s string) int {
	m := map[byte]int{}
	n := len(s)
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		ans = max(ans, rk-i+1)
	}
	return ans
}

// golang最快的解法
func lengthOfLongestSubstring1(s string) int {
	freq := make([]int, 128)
	res := 0
	start, end := 0, -1
	for start < len(s) {
		if end+1 < len(s) && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		} else {
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
	}
	return res
}

func main() {

}
