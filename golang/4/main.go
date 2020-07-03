package main

func main() {

}

// 难度：困难
// 4.寻找两个正序数组的中位数
// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2 。
// 请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。
// 你可以假设 nums1 和 nums2 不会同时为空。

// 分析
// 给出这个要求，应该用二分查找

// 自己做的。复杂度是 O(m+n) 不符合要求
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	tl := l1 + l2
	slice := []int{}
	i1, i2 := 0, 0
	i := 0
	for i < tl/2+1 {
		if i1 >= l1 {
			slice = append(slice, nums2[i2])
			i2++
			i++
			continue
		} else if i2 >= l2 {
			slice = append(slice, nums1[i1])
			i1++
			i++
			continue
		}
		x := nums1[i1]
		y := nums2[i2]
		if x < y {
			slice = append(slice, x)
			i1++
			i++
		} else {
			slice = append(slice, y)
			i2++
			i++
		}
	}
	if tl%2 == 1 {
		return float64(slice[tl/2])
	}
	return float64(slice[tl/2]+slice[tl/2-1]) / 2.0
}
