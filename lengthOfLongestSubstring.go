package main

// lengthOfLongestSubstring 无重复字符的最短字串
// leetcode: https://leetcode.cn/problems/longest-substring-without-repeating-characters/description/
/**
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/

// 虽然a了，但是内存的复杂度还是很高
func lengthOfLongestSubstring(s string) int {
	maxlen := 0

	for i := range s {
		j := i + 1
		repMap := make(map[byte]int)
		repMap[s[i]] = 1
		for j < len(s) {
			if repMap[s[j]] != 1 { // 说明存在
				repMap[s[j]] = 1
				j++
			} else {
				break
			}
		}

		if maxlen < len(repMap) {
			maxlen = len(repMap)
		}
	}
	return maxlen
}
