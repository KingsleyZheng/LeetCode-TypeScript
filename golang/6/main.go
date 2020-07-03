package main

import "fmt"

func main() {
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, 3))
	fmt.Println(convert(s, 4))
	fmt.Println(convert(s, 5))

	fmt.Println(convert2(s, 3))
}

// 难度：中等
// 6. Z 字形变换
// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

// 官方解法一：按行排序。通过从左向右迭代字符串，将字符添加到每行分组的队列中，再组合起来
// 关键：只有当我们向上移动到最上面的行或向下移动到最下面的行时，当前方向才会发生改变。
// 看得我目瞪口呆，小学奥数题的感觉
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, min(numRows, len(s)))
	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	ret := ""
	for _, row := range rows {
		ret += row
	}
	return ret
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 官方方法二：按行访问
// 按照与逐行读取 Z 字形图案相同的顺序访问字符串。
// 对于所有整数 k，
// 行0中的字符位于索引k(2*numRows-2)处
// 行numRows - 1中的字符位于索引k(2*numRows-2)+numRows-1处
// 内部的行i中的字符位于索引k(2*numRows-2)+i 以及(k+1)(2*numRows-2)-i处
func convert2(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	ret := ""
	n := len(s)
	cycleLen := 2*numRows - 2
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < n; j += cycleLen {
			ret += string(s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < n {
				ret += string(s[j+cycleLen-i])
			}
		}
	}
	return ret
}
