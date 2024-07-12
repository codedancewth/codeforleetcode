package main

// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
/**
输入: nums = [1,3,5,6], target = 5
输出: 2
*/
// https://leetcode.cn/problems/search-insert-position/description/?envType=study-plan-v2&envId=top-100-liked
// 感悟。我tm服了，这道题easy到需要技巧的
func searchInsert(nums []int, target int) int {
	if nums[len(nums)-1] < target { // 如果大于最大值，往后插入
		return len(nums)
	}

	left, right := 0, len(nums)-1 // 双指针方式
	for left < right {            // 如果指针还是右边大于左边的关系

		// 二分的计算方式，右边减掉左边/2
		// 那么问题来了，为什么这么计算呢？
		// 画一下图就可以知道数据需要二分直接处理
		// 但是左边的值永远都是小的，所以可以知道左指针一直移动的位置就是差值关系
		c := left + (right-left)/2

		// 对于顺序的数组
		if nums[c] < target { // 如果当前值小于目标的值
			left = c + 1 // 那么肯定就是需要往右移动，继续向右判断
		} else {
			right = c // 如果当前值大于目标的值， 向左移动
		}
	}
	return left
}
