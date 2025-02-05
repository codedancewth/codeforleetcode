package main

// longestPalindrome 最长回文子串
// leetcode https://leetcode.cn/problems/longest-palindromic-substring/description/
// 题目详情
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
func longestPalindrome(s string) string {
	if len(s) < 1 {
		return ""
	}

	array := []byte(s)
	start, end := 0, 0

	for i := range array { // 求出每个字符的最大回文
		single := centerKeep(s, i, i)   // 奇数
		double := centerKeep(s, i, i+1) // 偶数

		maxLen := max(single, double)

		if maxLen > end-start {
			start = i - (maxLen-1)/2 // 假设中心点是j，那么以j为中心点，找到最小的左数
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func centerKeep(s string, left, right int) int {

	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

//

//
//func longestPalindrome(s string) string {
//	if len(s) < 1 {
//		return ""
//	}
//
//	// 将字符串转换为 rune 数组，处理 Unicode 字符
//	runes := []byte(s)
//	start, end := 0, 0
//
//	for i := 0; i < len(runes); i++ {
//		// 奇数长度的回文（以单个字符为中心）
//		len1 := expandAroundCenter(runes, i, i)
//		// 偶数长度的回文（以两个字符的中间为中心）
//		len2 := expandAroundCenter(runes, i, i+1)
//		maxLen := max(len1, len2) // 取最大的个数
//
//		if maxLen > end-start {
//			start = i - (maxLen-1)/2
//			end = i + maxLen/2
//		}
//	}
//
//	return string(runes[start : end+1])
//}
//
//// 中心扩展函数，返回回文长度
//func expandAroundCenter(runes []byte, left, right int) int {
//	for left >= 0 && right < len(runes) && runes[left] == runes[right] {
//		left--
//		right++
//	}
//	// 返回有效回文的长度（注意：right-left-1 是回文实际长度）
//	return right - left - 1
//}
