package main

import "sort"

// 三数之和
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums) // 排序
	resultV2 := [][]int{}

	for i := 0; i < n-2; i++ {
		// 跳过重复的 nums[i]
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum < 0 {
				left++
			} else if sum > 0 {
				right--
			} else {
				// 找到一个有效的三元组
				resultV2 = append(resultV2, []int{nums[i], nums[left], nums[right]})
				// 跳过重复的 nums[left] 和 nums[right]
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			}
		}
	}

	return resultV2
}
