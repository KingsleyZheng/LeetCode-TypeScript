package main

import "fmt"

func main() {
	digits := "234"
	fmt.Println(letterCombinations1(digits))
}

// 中等
// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 输入："23"
// 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
// 说明:
// 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

// 首先想到就是暴力破解。列出每个数字的对应字母数组，再进行遍历
func letterCombinations(digits string) []string {
	m := map[int][]string{
		2: []string{"a", "b", "c"},
		3: []string{"d", "e", "f"},
		4: []string{"g", "h", "i"},
		5: []string{"j", "k", "l"},
		6: []string{"m", "n", "o"},
		7: []string{"p", "q", "r", "s"},
		8: []string{"t", "u", "v"},
		9: []string{"w", "x", "y", "z"},
	}
	// n := len(digits)
	res := []string{}

	for _, c := range digits {
		intc := int(c) - 48
		// fmt.Println(intc)
		if c == '0' || c == '1' {
			continue
		}
		if len(res) == 0 {
			for _, first := range m[intc] {
				res = append(res, first)
			}
		} else {
			newRes := []string{}
			for j := 0; j < len(res); j++ {
				for k := 0; k < len(m[intc]); k++ {
					newRes = append(newRes, res[j]+m[intc][k])
				}
			}
			res = newRes
		}
	}
	return res
}

// 官方解法：回溯。回溯是一种通过穷举所有可能情况来找到所有解的算法。
// 好复杂，感觉把简单问题写得不容易理解了
var phone map[string][]string = map[string][]string{
	"2": []string{"a", "b", "c"},
	"3": []string{"d", "e", "f"},
	"4": []string{"g", "h", "i"},
	"5": []string{"j", "k", "l"},
	"6": []string{"m", "n", "o"},
	"7": []string{"p", "q", "r", "s"},
	"8": []string{"t", "u", "v"},
	"9": []string{"w", "x", "y", "z"},
}
var output []string = []string{}

func backtrack(combination string, nextDigits string) {
	if len(nextDigits) == 0 {
		output = append(output, combination)
	} else {
		digit := nextDigits[0:1]
		letters := phone[digit]
		for i := 0; i < len(letters); i++ {
			letter := phone[digit][i]
			backtrack(combination+letter, nextDigits[1:])
		}
	}
}

func letterCombinations1(digits string) []string {
	if len(digits) != 0 {
		backtrack("", digits)
	}
	return output
}
