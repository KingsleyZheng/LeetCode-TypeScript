package main

// 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 第一种：暴力破解法，遍历两个数组，时间复杂度为O(n^2)，额外空间复杂度为O(1)
func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			if nums[i] == target - nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func main() {
	
}