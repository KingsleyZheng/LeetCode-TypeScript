package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := 10
	fmt.Println(isPalindrome(s))
}

// 简单
// 9. 回文数
// 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
// 进阶:
// 你能不将整数转为字符串来解决这个问题吗？

// 自己实现：转成字符串再遍历，时间复杂度为O(log(n))
func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	for i := 0; i < (l+1)/2; i++ {
		if s[i] != s[l-1-i] {
			return false
		}
	}
	return true
}

// golang 0ms范例，原理就是把x倒过来再比较
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	sum := x
	palin := 0
	for sum > 0 {
		palin = palin*10 + sum%10
		sum /= 10
	}
	return palin == x
}

// 官方解法：反转一半数字，只反转一半数字，比较这一半数字和剩余数字即可
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber*10 + x%10
		x /= 10
	}
	// 当数字长度为奇数时，我们可以通过 revertedNumber/10 去除处于中位的数字。
	// 例如，当输入为 12321 时，在 while 循环的末尾我们可以得到 x = 12，revertedNumber = 123，
	// 由于处于中位的数字不影响回文（它总是与自己相等），所以我们可以简单地将其去除。

	return x == revertedNumber || x == revertedNumber/10
}
