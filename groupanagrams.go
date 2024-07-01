package main

//字母异位词分组
//https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
//一开始的思路，但是超时了，del可惜，想要通过单个字符去一个一个匹配的，但是容易超时。
//func groupAnagrams(strs []string) [][]string {
//	resultMap := make(map[string]int64)
//	resultList := make([][]string, 0)
//
//	for k, index := range strs {
//		if resultMap[index] == 1 {
//			continue
//		}
//		if index == "" || len(index) == 1 {
//			// 直接找有没有空的元素,额外处理下
//			resultIndex := make([]string, 0)
//			resultIndex = append(resultIndex, index)
//			resultMap[index] = 1
//			for v, index2 := range strs {
//				if k == v {
//					continue
//				}
//
//				if index == index2 {
//					resultIndex = append(resultIndex, index2)
//				}
//			}
//			resultList = append(resultList, resultIndex)
//			continue
//		}
//
//		resultIndex := make([]string, 0)
//		arrayMap := make(map[byte]int64)
//
//		for _, index2 := range []byte(index) {
//			arrayMap[index2] = arrayMap[index2] + 1
//		}
//
//		// 处理数据
//		resultIndex = append(resultIndex, index)
//		resultMap[index] = 1
//
//		for v, index2 := range strs {
//			if k == v {
//				continue
//			}
//
//			if index2 == index {
//				resultMap[index2] = 1
//				resultIndex = append(resultIndex, index2)
//				continue
//			}
//
//			arrayMap2 := make(map[byte]int64)
//			count := 0
//			for _, index3 := range []byte(index2) {
//				if arrayMap[index3] > 0 {
//					arrayMap2[index3] = arrayMap2[index3] + 1
//					count++
//				} else {
//					break
//				}
//
//				if count == len(index2) && len(index2) == len(index) { // 正确的次数等于长度
//					cflag := true
//					for ik, indexm := range arrayMap2 {
//						if indexm != arrayMap[ik] {
//							cflag = false // 元素个数也相等
//						}
//					}
//
//					if cflag {
//						// 重复的元素
//						resultMap[index2] = 1
//						resultIndex = append(resultIndex, index2)
//					}
//
//				}
//			}
//		}
//
//		resultList = append(resultList, resultIndex)
//	}
//
//	return resultList
//}

// 真正的解决方式 ，通过单个字符的数值二进制
// https://leetcode.cn/problems/group-anagrams/description/?envType=study-plan-v2&envId=top-100-liked
// tm最后一道题g了我日
func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[int]int)
	resultList := make([][]string, 0)
	for k, index := range strs {
		if resultMap[k] == 1 { // 判断过了直接过
			continue
		}

		if index == "" {
			// 直接找有没有空的元素,额外处理下
			resultIndex := make([]string, 0)
			resultIndex = append(resultIndex, index)
			resultMap[k] = 1
			for v, index2 := range strs {
				if resultMap[v] == 1 {
					continue
				}

				if index == index2 {
					resultMap[v] = 1
					resultIndex = append(resultIndex, index2)
				}
			}
			resultList = append(resultList, resultIndex)
			continue
		}

		resultMap[k] = 1
		resultIndex := make([]string, 0)
		resultIndex = append(resultIndex, index)

		arrayCurrent := int32(index[0])

		for ks, index2 := range index {
			if ks == 0 {
				continue
			}
			arrayCurrent = arrayCurrent * index2
		}

		for k1, index2 := range strs {
			if k >= k1 || resultMap[k1] == 1 {
				continue
			}

			if len(index2) <= 0 {
				continue
			}

			arrayCurrent2 := int32(index2[0])
			for ks2, index3 := range index2 {
				if ks2 == 0 {
					continue
				}
				arrayCurrent2 = arrayCurrent2 * index3
			}
			if arrayCurrent2 == arrayCurrent {
				resultMap[k1] = 1
				resultIndex = append(resultIndex, index2)
			}

		}

		resultList = append(resultList, resultIndex)
	}

	return resultList
}

// leetcode大佬的思路
// [26]int 用来表示key,每个index对应一个字母，值是其出现次数 如 "aabc" 对应[2,1,1,0....]
// 然后根据串生成key存到map中即可，map的值是个字符串数组指针，因为我们需要改变字符串数组的值
//
// 作者：steven
// 链接：https://leetcode.cn/problems/group-anagrams/solutions/84565/golangmapbi-jiao-qiang-da-na-shu-zu-dang-key-by-se/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
//func groupAnagrams2(strs []string) [][]string {
//	dict := make(map[[26]int]*[]string)
//	for _, s := range strs {
//		k := getKey(s)
//		if v, ok := dict[k]; ok {
//			*v = append(*v, s)
//		} else {
//			dict[k] = &[]string{s}
//		}
//	}
//	var res [][]string
//	for _, v := range dict {
//		res = append(res, *v)
//	}
//	return res
//}
//
//func getKey(s string) [26]int {
//	var k [26]int
//	for i := range s {
//		k[int(s[i]-'a')]++
//	}
//	return k
//}

// 复写这种牛逼的方式，回忆对应的思路
// 这里的实现方式使用的是字符出现的次数来做对应的存储的
func groupAnagrams2(strs []string) [][]string {
	repeatMap := make(map[[26]int64]int64)
	resultMap := make(map[[26]int64][]string, 0)
	resultList := make([][]string, 0)

	for _, index := range strs {
		do := getKey(index)
		if repeatMap[do] == 1 {
			resultMap[do] = append(resultMap[do], index)
		} else {
			resultMap[do] = make([]string, 0)
			resultMap[do] = append(resultMap[do], index)
			repeatMap[do] = 1
		}
	}

	for _, index := range resultMap {
		resultList = append(resultList, index)
	}

	return resultList
}

func getKey(str string) [26]int64 {
	var k [26]int64
	for _, index := range str {
		k[index-'a']++
	}
	return k
}
