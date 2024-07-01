package codeforleetcode

// 字母异位词分组
// https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
// 一开始的思路，但是超时了，del可惜，想要通过单个字符去一个一个匹配的，但是容易超时。
func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string]int64)
	resultList := make([][]string, 0)

	for k, index := range strs {
		if resultMap[index] == 1 {
			continue
		}
		if index == "" || len(index) == 1 {
			// 直接找有没有空的元素,额外处理下
			resultIndex := make([]string, 0)
			resultIndex = append(resultIndex, index)
			resultMap[index] = 1
			for v, index2 := range strs {
				if k == v {
					continue
				}

				if index == index2 {
					resultIndex = append(resultIndex, index2)
				}
			}
			resultList = append(resultList, resultIndex)
			continue
		}

		resultIndex := make([]string, 0)
		arrayMap := make(map[byte]int64)

		for _, index2 := range []byte(index) {
			arrayMap[index2] = arrayMap[index2] + 1
		}

		// 处理数据
		resultIndex = append(resultIndex, index)
		resultMap[index] = 1

		for v, index2 := range strs {
			if k == v {
				continue
			}

			if index2 == index {
				resultMap[index2] = 1
				resultIndex = append(resultIndex, index2)
				continue
			}

			arrayMap2 := make(map[byte]int64)
			count := 0
			for _, index3 := range []byte(index2) {
				if arrayMap[index3] > 0 {
					arrayMap2[index3] = arrayMap2[index3] + 1
					count++
				} else {
					break
				}

				if count == len(index2) && len(index2) == len(index) { // 正确的次数等于长度
					cflag := true
					for ik, indexm := range arrayMap2 {
						if indexm != arrayMap[ik] {
							cflag = false // 元素个数也相等
						}
					}

					if cflag {
						// 重复的元素
						resultMap[index2] = 1
						resultIndex = append(resultIndex, index2)
					}

				}
			}
		}

		resultList = append(resultList, resultIndex)
	}

	return resultList
}
