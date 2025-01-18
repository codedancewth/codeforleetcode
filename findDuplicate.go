package main

// leetcode 287 寻找重复数
// 使用map的情况进行去重 （好像性能不太行）
func findDuplicate(nums []int) int {
	m := map[int]int{}
	for _, index := range nums {
		if m[index] == 0 {
			m[index] = 1
		} else {
			return index
		}
	}

	return 0
}

// https://leetcode.cn/problems/find-the-duplicate-number/
// 官方题解
// 快慢指针
func findDuplicateV2(nums []int) int {
	slow, fast := 0, 0
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		if slow == fast {
			fast = 0
			for {
				if slow == fast {
					return slow
				}
				slow = nums[slow]
				fast = nums[fast]
			}
		}
	}
}
