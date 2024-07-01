package main

// 两数之和（easy）
// https://leetcode.cn/problems/two-sum/submissions/543103221/?envType=study-plan-v2&envId=top-100-liked
func twoSum(nums []int, target int) []int {
	for k, index := range nums {
		for v, index2 := range nums {
			if k == v {
				continue
			}
			if index+index2 == target {
				return []int{k, v}
			}
		}
	}

	return []int{}
}
