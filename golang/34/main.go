package main

import "fmt"

func main() {
	nums := []int{1, 4}
	fmt.Println(searchRange(nums, 4))
}

// 中等
// 34. 在排序数组中查找元素的第一个和最后一个位置
// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

// 先找到这个数，再从这个数出发向其左右遍历
func searchRange(nums []int, target int) []int {
	n := len(nums)
	if n == 0 {
		return []int{-1, -1}
	}
	if n == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			begin, end := mid, mid
			for begin > 0 {
				if nums[begin-1] == target {
					begin--
				} else {
					break
				}
			}
			for end < n-1 {
				if nums[end+1] == target {
					end++
				} else {
					break
				}
			}
			return []int{begin, end}
		}
	}
	return []int{-1, -1}
}
