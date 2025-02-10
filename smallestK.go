package main

import "sort"

// smallestK 最小的k个数
// 设计一个算法，找出数组中最小的k个数。以任意顺序返回这k个数均可。
// https://leetcode.cn/problems/smallest-k-lcci/description/
func smallestK(arr []int, k int) []int {
	sort.Slice(arr, func(i, j int) bool { // 使用对应的数据结构
		return arr[i] < arr[j]
	})

	return arr[:k]
}

func smallestKv2(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}
