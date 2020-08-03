package main

func main() {

}

// 中等
// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 应该用回溯法
func permute(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	if n == 0 {
		return res
	}
}