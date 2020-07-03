package main

// 难度：简单
// 1.两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。

// 官方方法一：暴力破解法。遍历两个数组，时间复杂度为O(n^2)，额外空间复杂度为O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == target-nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 官方方法二：两遍哈希表。时间复杂度为O(n)，空间复杂度为O(n)。先存一边切片的元素到map，再遍历一次切片，看是否存在对应的key在map中，利用了map查找速度快的特点
func twoSum2(nums []int, target int) []int {
	map1 := map[int]int{}
	for i := 0; i < len(nums); i++ {
		map1[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		val, ok := map1[complement]
		if !ok {
			continue
		}
		if val != i {
			return []int{i, val}
		}
	}
	return []int{}
}

// 官方方法三：一遍哈希表。时间复杂度为O(n)，空间复杂度为O(n)。在进行迭代并将元素插入到表中的同时，回过头检查表中是否已经存在当前元素所对应的目标元素。如果存在，立即将其返回
func twoSum3(nums []int, target int) []int {
	map1 := map[int]int{}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if val, ok := map1[complement]; ok {
			return []int{val, i}
		}
		map1[nums[i]] = i
	}
	return []int{}
}

func main() {

}
