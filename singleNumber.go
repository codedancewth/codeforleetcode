package main

// https://leetcode.cn/problems/single-number/submissions/543842376/?envType=study-plan-v2&envId=top-100-liked
// 只出现一次的数字
// 用map即可
func singleNumber(nums []int) int {
	a := make(map[int]int)
	for _, index := range nums {
		a[index]++
	}

	for k, index := range a {
		if index == 1 {
			return k
		}
	}
	return 0
}
