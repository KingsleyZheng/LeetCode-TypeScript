package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(-2147483648, 1))
}

// 中等
// 29. 两数相除
// 给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
// 返回被除数 dividend 除以除数 divisor 得到的商。
// 整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

// 不让用乘除法，就是用减法
func divide(dividend int, divisor int) int {
	biggerThanZero := true
	if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
		biggerThanZero = false
	}
	dividendAbs, divisorAbs := abs(dividend), abs(divisor)
	ans := 0
	for dividendAbs >= divisorAbs {
		if divisorAbs == 1 {
			ans = dividendAbs
			break
		}
		dividendAbs -= divisorAbs
		ans++
	}
	if !biggerThanZero {
		ans = -ans
	}
	if ans > math.MaxInt32 || ans < math.MinInt32 {
		return math.MaxInt32
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 移位也算是乘法，试用二分法优化时间
func divide2(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	a := dividend
	b := divisor
	sign := 1
	if (a > 0 && b < 0) || (a < 0 && b > 0) {
		sign = -1
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res := div(a, b)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
		return res
	}
	return -res
}

func div(a, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tb := b
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	return count + div(a-tb, b)
}
