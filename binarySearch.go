package main

// 二分查找
// binarySearch
func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2 // 防止溢出

		if arr[mid] == target {
			return mid // 找到目标值，返回索引
		} else if arr[mid] < target {
			low = mid + 1 // 目标值在右半部分
		} else {
			high = mid - 1 // 目标值在左半部分
		}
	}

	return -1 // 未找到目标值
}
