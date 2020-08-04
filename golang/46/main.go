package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

// 中等
// 46. 全排列
// 给定一个 没有重复 数字的序列，返回其所有可能的全排列。

// 官方解法：搜索回溯
func permute(nums []int) [][]int {
	res := new([][]int)
	backtrack(res, &nums, 0, len(nums))
	return *res
}
func backtrack(res *[][]int, output *[]int, first, n int) {
	if first == n {
		tmp := make([]int, n)
		copy(tmp, *output)
		*res = append(*res, tmp)
		return
	}
	for i := first; i < n; i++ {
		(*output)[i], (*output)[first] = (*output)[first], (*output)[i]
		backtrack(res, output, first+1, n)
		(*output)[i], (*output)[first] = (*output)[first], (*output)[i]
	}
}
