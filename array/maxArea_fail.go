package array

import (
	"math"
)

// https://leetcode.cn/problems/container-with-most-water/?envType=study-plan-v2&envId=top-100-liked
// 盛最多水的容器
/**
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。


输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
*/

// 自己的解法，果然超过时间复杂度了，这个是O(n^2)，我服了
// 已经超出复杂度读了
func maxArea2(height []int) int {
	maxHeight := 0
	for k, index := range height {
		for v, index2 := range height {
			if k >= v {
				continue
			}

			n := index

			if index > index2 {
				n = index2
			}

			heightC := (v - k) * n

			if maxHeight < heightC {
				maxHeight = heightC
			}
		}
	}

	return maxHeight
}

// 还是超时，我不日了
func maxArea3(height []int) int {
	maxHeight := 0

	for k, index := range height {
		for i := k; i < len(height); i++ {
			if k >= i {
				continue
			}
			cur := (i - k) * int(math.Min(float64(index), float64(height[i])))
			if maxHeight < cur {
				maxHeight = cur
			}
		}
	}

	return maxHeight
}

// 看了官方的思路，应该是想使用双指针的方式，总是以最小的移动，
// 草。怎么不对啊。。先挂起吧，到时候做完回过头再去思考
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		index1 := float64(height[l])
		index2 := float64(height[r] * (r - l))
		area := math.Min(index1, index2)

		ans = int(math.Max(float64(ans), area))

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}
