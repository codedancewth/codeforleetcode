package main

import "strconv"

// 给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
// 回文数回文数
// 是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
func isPalindrome(x int) bool {
	if x >= 0 {
		itoa := strconv.Itoa(x)
		return centerPoint(itoa)
	}
	return false
}

func centerPoint(s string) bool {
	if len(s) <= 1 {
		return true
	}

	for i := range s {
		if s[i] == s[len(s)-i-1] {
			continue
		} else {
			return false
		}
	}

	return true
}
