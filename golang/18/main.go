package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(nums, target))
}

// 中等
// 18. 四数之和
// 给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
// 注意：
// 答案中不可以包含重复的四元组。

// 三数之和的升级版
// 先定下来第一位的数字，剩下来的就是三数之和
// 排序+双指针
func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	res := make([][]int, 0)

	if n < 4 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		// 确保第一个数有变
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			// 确保第二个数有变
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 剩下的部分跟三数之和基本一致
			L := j + 1
			R := n - 1
			for L < R {
				if nums[i]+nums[j]+nums[L]+nums[R] == target {
					res = append(res, []int{nums[i], nums[j], nums[L], nums[R]})
					for L < R && nums[L] == nums[L+1] {
						L++
					}
					for L < R && nums[R] == nums[R-1] {
						R--
					}
					L++
					R--
				} else if nums[i]+nums[j]+nums[L]+nums[R] > target {
					R--
				} else {
					L++
				}
			}
		}
	}
	return res
}
