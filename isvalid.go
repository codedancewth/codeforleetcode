package main

// 有效的括号
// https://leetcode.cn/problems/valid-parentheses/description/?envType=study-plan-v2&envId=top-100-liked
/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
/*func isValid(s string) bool {
	b := []byte(s)
	endMap := map[byte]byte{
		41:  40,
		125: 123,
		93:  91,
	} // 结尾的

	sb := []byte{}

	for index := 0; index < len(b); index++ {
		key := b[index]
		if endMap[key] > 0 {
			if len(sb) == 0 || (len(sb) > 0 && sb[len(sb)-1] != endMap[key]) { // 遇到最右边的数据，则要出栈
				return false
			}

			if len(sb) > 0 {
				sb = sb[:len(sb)-1]
			}

		} else {
			sb = append(sb, key)
		}
	}

	return len(sb) == 0
}*/

// isValid 仿照栈的写法
func isValidV2(s string) bool {
	truemap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	// 存入的字母是否存在
	arrayStack := []byte{}
	for i := range s {

		if len(arrayStack) >= 1 {
			if arrayStack[len(arrayStack)-1] == truemap[s[i]] {
				if len(arrayStack) >= 2 {
					arrayStack = arrayStack[:len(arrayStack)-1]
				} else {
					arrayStack = []byte{}
				}
			} else {
				arrayStack = append(arrayStack, s[i])
			}

		} else {
			arrayStack = append(arrayStack, s[i])
		}
	}
	return len(arrayStack) == 0
}
