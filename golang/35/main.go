package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(nums, 0))
}

// 简单
// 35. 搜索插入位置
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 你可以假设数组中无重复元素。
func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	l, r := 0, n-1
	mid := 0
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	if nums[mid] > target {
		return mid
	}
	return mid + 1
}
