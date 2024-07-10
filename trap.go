package main

import "fmt"

// https://leetcode.cn/problems/trapping-rain-water/description/?envType=study-plan-v2&envId=top-100-liked
// 接雨水
// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
// 输出：6
// 解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//func trap(height []int) int {
//	return 0
//}

// 普及一下什么是动态规划？
// 动态规划（Dynamic Programming）是一种解决多阶段决策问题的优化方法，通过将问题划分为若干个重叠子问题，并保存子问题的解，以避免重复计算，从而有效地求解问题。
// 动态规划通常用于优化具有重叠子问题和最优子结构性质的问题。重叠子问题是指原问题的解可以通过一系列子问题的解得到，而这些子问题可以在求解原问题时重复出现。最优子结构是指原问题的最优解可以通过一系列子问题的最优解得到。
// 动态规划的基本思想是利用已解决的子问题的解来求解当前问题的解，并将子问题的解保存在一个表格（通常是数组）中。通过自底向上的计算方式，逐步求解规模更大的子问题，最终得到原问题的解。
// 动态规划通常包括以下步骤：
// 定义状态：确定问题的状态表示，即问题的子问题中需要保存的信息。
// 确定状态转移方程：根据问题的最优子结构性质，建立当前问题与子问题之间的递推关系式，即状态转移方程。通过状态转移方程可以从子问题的解推导出当前问题的解。
// 确定边界条件：确定最简单的子问题的解，即边界条件。边界条件是递归或迭代计算的终止条件。
// 确定计算顺序：确定子问题的计算顺序，通常采用自底向上的方式，先计算规模较小的子问题，再逐步计算规模更大的子问题，直到计算出原问题的解。
// 构建解：根据计算顺序，依次计算每个子问题的解，并保存在表格中。最终得到原问题的解。
// 动态规划广泛应用于各种问题，如最短路径问题、背包问题、序列比对、图像处理等。它可以大大降低问题的时间复杂度，提高算法的效率。

// 官方题解（LeetCode）
// 方法一：动态规划
// 对于下标 i，下雨后水能到达的最大高度等于下标 i 两边的最大高度的最小值，下标 i 处能接的雨水量等于下标 i 处的水能到达的最大高度减去 height[i]。
//
// 朴素的做法是对于数组 height 中的每个元素，分别向左和向右扫描并记录左边和右边的最大高度，然后计算每个下标位置能接的雨水量。假设数组 height 的长度为 n，该做法需要对每个下标位置使用 O(n) 的时间向两边扫描并得到最大高度，因此总时间复杂度是 O(n
// 2
// )。
//
// 上述做法的时间复杂度较高是因为需要对每个下标位置都向两边扫描。如果已经知道每个位置两边的最大高度，则可以在 O(n) 的时间内得到能接的雨水总量。使用动态规划的方法，可以在 O(n) 的时间内预处理得到每个位置两边的最大高度。
//
// 创建两个长度为 n 的数组 leftMax 和 rightMax。对于 0≤i<n，leftMax[i] 表示下标 i 及其左边的位置中，height 的最大高度，rightMax[i] 表示下标 i 及其右边的位置中，height 的最大高度。
//
// 显然，leftMax[0]=height[0]，rightMax[n−1]=height[n−1]。两个数组的其余元素的计算如下：
//
// 当 1≤i≤n−1 时，leftMax[i]=max(leftMax[i−1],height[i])；
//
// 当 0≤i≤n−2 时，rightMax[i]=max(rightMax[i+1],height[i])。
//
// 因此可以正向遍历数组 height 得到数组 leftMax 的每个元素值，反向遍历数组 height 得到数组 rightMax 的每个元素值。
//
// 在得到数组 leftMax 和 rightMax 的每个元素值之后，对于 0≤i<n，下标 i 处能接的雨水量等于 min(leftMax[i],rightMax[i])−height[i]。遍历每个下标位置即可得到能接的雨水总量。
//
// 动态规划做法可以由下图体现。
func trapV2(height []int) (ans int) {
	n := len(height)
	if n == 0 {
		return
	}

	leftMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	for i, h := range height {
		ans += min(leftMax[i], rightMax[i]) - h
		fmt.Println(fmt.Sprintf("leftMax[i] %d, rightMax[i] %d,height[i] %d ans %d", leftMax[i], rightMax[i], height[i], ans))
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 自己的理解，接雨水的核心就是算好每个点的左右高度的max 的min值，然后减去当前点的高度即可点每个点的雨水
func trap(height []int) (ans int) {
	if len(height) == 0 {
		return
	}

	leftMaxArray := make([]int, len(height))
	rightMaxArray := make([]int, len(height))

	leftMaxArray[0] = height[0]
	rightMaxArray[len(height)-1] = height[len(height)-1]

	// 把每个点的左右高度最大都记住
	for index := 1; index < len(height); index++ {
		leftMaxArray[index] = max(leftMaxArray[index-1], height[index]) // 预存左边的最高高度
	}

	for index := len(height) - 2; index >= 0; index-- {
		rightMaxArray[index] = max(rightMaxArray[index+1], height[index]) // 预存右边的最高高度
	}

	for index := range height { // 高度相减 在减去当前的点，即可得到答案
		ans += min(leftMaxArray[index], rightMaxArray[index]) - height[index]
	}

	return
}
