package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 120
	fmt.Println(reverse(x))
}

// 7. 整数反转
// 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

// 先转 string 再转 []byte
func reverse(x int) int {
	isPositive := true
	if x < 0 {
		isPositive = false
		x = -x
	}

	xString := strconv.Itoa(x)
	length := len(xString)
	slice := []byte{}
	for i := length - 1; i >= 0; i-- {
		slice = append(slice, xString[i])
	}
	res, err := strconv.Atoi(string(slice))
	if err != nil {
		return 0
	}
	if !isPositive {
		res = -res
	}
	// 2^31 2147483647,-2^31=-2147483648
	if res > 2147483647 || res < -2147483648 {
		return 0
	}
	return res
}

// 官方解法：逐个把个位数推到新数字头部，额外空间复杂度为 O(1)
func reverse1(x int) int {
	var ret int32 = 0

	var _x int32 = int32(x)
	for _x != 0 {
		pop := _x % 10
		_x /= 10
		if ret > 1<<31/10 || (ret == 1<<31/10 && pop > 7) {
			return 0
		}
		if ret < -1<<31/10 || (ret == -1<<31/10 && pop < -8) {
			return 0
		}
		ret = ret*10 + pop
	}
	return int(ret)
}
