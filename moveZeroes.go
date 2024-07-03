package main

// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
/**
示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
*/
// = =用补充的方式真的很快，当然，人家限制了不能复制数组，也就是只能对原数组的数据进行操作
// 够坏的，只判断零，真棒棒
func moveZeroes(nums []int) {
	for k := range nums {
		if nums[k] != 0 {
			continue
		}
		tempKey := k
		for v := range nums {
			if k >= v {
				continue
			}
			if nums[v] != 0 {
				temp := nums[tempKey]
				nums[tempKey] = nums[v]
				nums[v] = temp
				tempKey = v
			} else {
				continue
			}
		}
	}
}
