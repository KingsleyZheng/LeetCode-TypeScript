package main

func main() {

}

// 简单
// 面试题 08.03. 魔术索引
func findMagicIndex(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		if i == nums[i] {
			return i
		}
	}
	return -1
}
