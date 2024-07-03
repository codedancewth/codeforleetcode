package main

// https://leetcode.cn/problems/sort-colors/?envType=study-plan-v2&envId=top-100-liked
// 颜色分类
// 给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
// 必须在不使用库内置的 sort 函数的情况下解决这个问题
func sortColors(nums []int) {
	for k, _ := range nums {
		for i := k; i < len(nums); i++ {
			if nums[k] > nums[i] {
				temp := nums[k]
				nums[k] = nums[i]
				nums[i] = temp
			}
		}
	}
}
