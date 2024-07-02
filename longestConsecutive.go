package main

import "sort"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// leetcode
// https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-liked

/*
	示例 1：

	输入：nums = [100,4,200,1,3,2]
	输出：4
	解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
	示例 2：

	输入：nums = [0,3,7,2,5,8,4,6,0,1]
	输出：9
*/

// pass
func longestConsecutive(nums []int) int {
	maxConsecutive := 0
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	}) // 先排序

	for k, index := range nums {
		if maxConsecutive == len(nums) {
			break
		}

		maxCount := 1
		startIndex := index
		for v, index2 := range nums {
			if k >= v {
				continue
			}

			if startIndex+1 == index2 {
				startIndex = index2
				maxCount++
			} else if startIndex != index2 {
				break
			}
		}

		if maxCount > maxConsecutive {
			maxConsecutive = maxCount
		}
	}

	return maxConsecutive
}

// 总结，虽然通过了，但是这里使用了sort.slice的内置的快速排序的方式，不太好
// 参考一下大佬的写法先
// 这个官方题解的方式是，想通过预存数据到map中，然后通过map的特性，遍历map，看是否存在key在递减的情况下，存在找到最长的子序列情况
// 也是一个非常好的解题思路。参考参考学习
//func longestConsecutive(nums []int) int {
//	numSet := map[int]bool{}
//	for _, num := range nums {
//		numSet[num] = true
//	}
//	longestStreak := 0
//	for num := range numSet {
//		if !numSet[num-1] {
//			currentNum := num
//			currentStreak := 1
//			for numSet[currentNum+1] {
//				currentNum++
//				currentStreak++
//			}
//			if longestStreak < currentStreak {
//				longestStreak = currentStreak
//			}
//		}
//	}
//	return longestStreak
//}
