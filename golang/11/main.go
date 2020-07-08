package main

func main() {

}

// 中等
// 11. 盛最多水的容器
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
// 说明：你不能倾斜容器，且 n 的值至少为 2。
// 输入：[1,8,6,2,5,4,8,3,7]
// 输出：49

// 官方解法：双指针法。第一个想到的就是双指针法，类似最长不重复子串那题，但是想不清楚原理
// 原理就是每次移动对应数字较小的指针，因为这样移动之后的面积必定大于等于原来的面积。遍历过后给出面积的历史最大值
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	ans := 0
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		ans = max(ans, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// golang 4ms 范例：就是简化了一下代码
func maxArea1(height []int) int {
	maxVol := 0
	i, j := 0, len(height)-1
	for i < j {
		if height[i] > height[j] {
			maxVol = max(maxVol, (j-i)*height[j])
			j--
		} else {
			maxVol = max(maxVol, (j-i)*height[i])
			i++
		}
	}
	return maxVol
}
