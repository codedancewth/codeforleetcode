package main

// climbStairs 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/submissions/597276624/

// 超时了
//func climbStairs(n int) int {
//	if n <= 1 {
//		return 1
//	}
//
//	if n == 2 {
//		return 2
//	}
//
//	return climbStairs(n-1) + climbStairs(n-2)
//}

func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
