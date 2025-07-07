package array

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-100-liked
// 多数元素
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
func majorityElement(nums []int) int {
	a := make(map[int]int)
	for _, index := range nums {
		a[index]++
	}

	for k, index := range a {
		if index > len(nums)/2 {
			return k
		}
	}
	return 0
}
