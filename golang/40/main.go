package main

import (
	"fmt"
	"sort"
)

func main() {
	candidates := []int{3, 1, 3, 5, 1, 1}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}

// 中等
// 40. 组合总和 II
// 给定一个数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用一次。
// 说明：
// 所有数字（包括目标数）都是正整数。
// 解集不能包含重复的组合。

// 跟39题类似，也是用回溯算法
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ans [][]int
	backtrack(candidates, target, 0, []int{}, &ans)
	return ans
}

func backtrack(candidates []int, target, index int, path []int, ans *[][]int) {
	// 退出条件
	if target == 0 {
		*ans = append(*ans, path)
		return
	}
	n := len(candidates)
	for i := index; i < n; i++ {
		// 去除多余的重复项（调整当前层的选择进度即可，下层需要继续可选）
		rawI := i
		for i < n-1 && candidates[i] == candidates[i+1] {
			i++
		}
		// 剪枝
		if candidates[i] > target {
			return
		}
		// 作出选择
		newPath := make([]int, len(path))
		copy(newPath, path)
		newPath = append(newPath, candidates[i])
		backtrack(candidates, target-candidates[i], rawI+1, newPath, ans)
		// 撤回选择

	}
}
