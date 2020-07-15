// go 排序算法
package main

func main() {

}

// BubbleSort 冒泡排序
// 两两比较相邻值，交换位置。每次循环后最大值会位于尾部
func BubbleSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		// 提前结束的标识
		flag := false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

// InsertionSort 插入排序
// 类似扑克牌排序，分为有序区和无序区
func InsertionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = value
	}
}

// SelectionSort 选择排序
// 从未排序序列中找到最大或最小元素，放到起始位置，然后从剩余未排序元素中继续寻找最大元素，放到已排序序列的末尾
func SelectionSort(a []int) {
	n := len(a)
	if n <= 1 {
		return
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		a[i], a[minIndex] = a[minIndex], a[i]
	}
}

// MergeSort 归并排序
// 采用分治法，利用递归排序
func MergeSort(n []int , start, end int) {
	if start >= end {
		return
	}
	mid := (start+end)/2
	MergeSort(n, start,mid)
	MergeSort(n,mid+1, end)
	Merge(n, start, mid,end)
}

// Merge 归并排序
func Merge(n []int, start, mid, end int) {
	var temp []int
	i := start
	k := mid + 1
	j := 0

	for ;i<= mid && k <= end;j++ {
		if n[i]< n[k] {
			temp = append(temp, n[i])
			i++
		}else {
			temp = append(temp, n[k])
			k++
		}
	}
	if i > mid {
		temp = append(temp, n[k:end+1]...)
	} else {
		temp = append(temp, n[i:mid+1]...)
	}
	copy(n[start:end+1], temp)
}