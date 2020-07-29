package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(candidates, target))
}

// 中等
// 39. 组合总和
// 给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的数字可以无限制重复被选取。
// 说明：
// 所有数字（包括 target）都是正整数。
// 解集不能包含重复的组合。

// 想不出。题解使用回溯算法
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := new([][]int)
	backtrack(candidates, target, res, 0, []int{})
	return *res
}

// 每次循环必然从数的后面开始寻找，以排除重复
func backtrack(candidates []int, target int, res *[][]int, i int, tmpList []int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*res = append(*res, tmpList)
		return
	}
	for start := i; start < len(candidates); start++ {
		if target < candidates[start] {
			break
		}
		tmpList = append(tmpList, candidates[start])
		backtrack(candidates, target-candidates[start], res, start, tmpList)
		tmpList = tmpList[0 : len(tmpList)-1]
	}
}
