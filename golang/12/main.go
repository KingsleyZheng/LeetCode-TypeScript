package main

import "fmt"

func main() {
	num := 58
	fmt.Println(intToRoman(num))
	fmt.Println(intToRoman2(num))
}

// 中等
// 12. 整数转罗马数字
// 给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

// 自己实现：直接建立映射，暴力破解
// 输入在 1 到 3999 内 很多边角情况不用考虑。把所有位数的情况列出来逐个字符串拼接即可。时间复杂度是 O(1)
// 缺点是不好扩展
func intToRoman(num int) string {
	thousandSlice := []string{"", "M", "MM", "MMM"}
	hundredSlice := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	tenSlice := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	oneSlice := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	ret := ""
	ret += thousandSlice[num/1000]
	num = num % 1000
	ret += hundredSlice[num/100]
	num = num % 100
	ret += tenSlice[num/10]
	num = num % 10
	ret += oneSlice[num]
	return ret
}

// 官方解法：贪心算法，每一步都找出适合余数的最大符号，直到余数为0
func intToRoman2(num int) string {
	ret := ""
	// 用map是不行的，map遍历是无序的
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	for i := 0; i < len(values) && num >= 0; i++ {
		for values[i] <= num {
			num -= values[i]
			ret += symbols[i]
		}
	}

	return ret
}
