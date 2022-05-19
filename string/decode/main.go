/*
394. 字符串解码
给定一个经过编码的字符串，返回它解码后的字符串。
编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。

示例 1：
输入：s = "3[a]2[bc]"
输出："aaabcbc"

示例 2：
输入：s = "3[a2[c]]"
输出："accaccacc"

示例 3：
输入：s = "2[abc]3[cd]ef"
输出："abcabccdcdcdef"

示例 4：
输入：s = "abc3[cd]xyz"
输出："abccdcdcdxyz"


提示：
1 <= s.length <= 30
s 由小写英文字母、数字和方括号 '[]' 组成
s 保证是一个 有效 的输入。
s 中所有整数的取值范围为 [1, 300]

*/
package main

import (
	"fmt"
)

func decodeString(s string) string {
	stack := make([]interface{}, 30)
	top := 0

	repeatStr := make([]byte, 0)
	time := 0
	num := 0

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			// 遇到首个数字，之前的字符串压栈
			if time == 0 {
				if len(repeatStr) > 0 {
					top++
					stack[top] = repeatStr
					repeatStr = []byte{}
				}
			}

			num = num*time + int(s[i]-'0')
			time = 10
			continue
		}

		if s[i] == '[' {
			// 数字结束，压栈
			time = 0
			top++
			stack[top] = num
			num = 0
			continue
		}

		if s[i] == ']' {
			// 出栈
			if top >= 1 {
				ele := stack[top]
				top--

				// 如果栈当前的是数字，重复
				if repeatTime, ok := ele.(int); ok {
					cycleStr := make([]byte, 0)
					for j := repeatTime; j > 0; j-- {
						cycleStr = append(cycleStr, repeatStr...)
					}
					repeatStr = cycleStr
				}

				// 如果栈当前是字符串，弹出拼接。如果栈中还有，必然是数字，需要再弹出重复
				if preStr, ok := ele.([]byte); ok {
					repeatStr = append(preStr, repeatStr...)
					if top >= 1 {
						ele = stack[top]
						top--
						repeatTime := ele.(int)
						cycleStr := make([]byte, 0)
						for j := repeatTime; j > 0; j-- {
							cycleStr = append(cycleStr, repeatStr...)
						}
						repeatStr = cycleStr
					}
				}
			}
			continue
		}

		repeatStr = append(repeatStr, s[i])
	}

	// 如果栈中还有，必然是没有重复的，单纯的字符串
	for i := top; i >= 1; i-- {
		ele := stack[top]
		top--
		preStr := ele.([]byte)
		repeatStr = append(preStr, repeatStr...)
	}

	return string(repeatStr)
}

/*
适用于无嵌套的
栈
遇到字母存起来，字母结束有2种
a.遇到数字，之前保存的字符串压栈
b.遇到']'，出栈

遇到'['，数字结束
*/
func decodeWithoutNest(s string) string {
	repeatStr := make([]byte, 0)
	decodeStr := make([]byte, 0)
	time := 0
	num := 0
	stack := 0 // 存放'['

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*time + int(s[i]-'0')
			time = 10
			continue
		}
		if s[i] == '[' {
			stack++
			time = 0
			continue
		}
		if s[i] == ']' {
			stack--
			for j := 0; j < num; j++ {
				decodeStr = append(decodeStr, repeatStr...)
			}
			repeatStr = []byte{}
			num = 0
			continue
		}

		if stack > 0 {
			repeatStr = append(repeatStr, s[i])
			continue
		}

		decodeStr = append(decodeStr, s[i])
	}

	return string(decodeStr)
}

func main() {
	//s := "3[a]1[b]"
	//s := "3[a2[c]]"
	s := "2[a]3[b]4[c]5[d]"
	fmt.Println(decodeString(s))
	//fmt.Println(decodeWithoutNest(s))
}
