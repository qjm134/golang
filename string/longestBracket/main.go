/*
32. 最长有效括号
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：
输入：s = "(()()))("
输出： "(()())"

示例 4：
输入：s = "()("
输出：

示例 5：
输入：s = "()(()"
输出：

示例 6：
输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'


思路：
有效的括号的判断：
遇到左括号压栈，遇到右括号出栈，栈里没有，一次配对的括号结束
连续有效的，可能是由多个配对括号连续排着的

2种情况：
1. 遇到左括号
   压栈
2. 遇到右括号
   a.栈里有，能出一个，就配成一对
   b.栈里没有，配对结束

需要做的事：
找连续最长
一个保持当前最大
一个临时长度
怎么算连续长度，当前配对右括号位置 到 前面未配对的左括号位置，或上一个隔开的右括号（前面缺少左括号跟其配对的右括号，栈已为空，之前的已配对，又遇到个右括号）
临时长度什么时候开始计算，什么时候终止计算
开始计算：上一次隔开的位置
        左括号，栈里没匹配上的左括号 (  ( ()(()) ),   (  ()()()
        右括号，栈已为空，之前的已配对，又遇到个右括号 ) ()(),  ) (()())
        //右括号可以出一个栈，算一对 (()
终止计算：配对成功 //隔开，不连续
        1.右括号 //栈里没有 ())
        //2.左括号，循环结束了，栈里还有 ()((), ()(())(, ()(()()(), ()(()(()
        //栈里剩几个左括号，就隔开n+1段
*/
package main

import "fmt"

func longestValidParentheses(s string) int {
	stack := make([]int, 0)
	max := 0
	preRight := -1

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				length := 0
				if len(stack) > 0 {
					length = i - stack[len(stack)-1]
				} else {
					length = i - preRight
				}
				if length > max {
					max = length
				}
			} else {
				preRight = i
			}
		}
	}
	return max
}

/*
//需要记录前面的状态，根据后面情况再做不同操作，后面的东西是黑的，这样就很不明朗，很难做，剩这种情况(()(()。其实也有解决办法，就是从后向前再遍历一遍
上一个方法，是根据现有的情况，得出一个连续值，就好做多了
func longestValidParentheses(s string) int {
	max := 0
	tmpLong := 0
	top := 0
	pre := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "(" {
			if top > 0 {
				top++
			} else { //新开启一段配对，如果新的配对成功，跟之前的合并，形成连续
				pre = pre + tmpLong
				tmpLong = 0
			}
		} else {
			if top > 0 {
				top--
				tmpLong = tmpLong + 2
			} else { //右括号隔开，连续必结束
				if pre > 0 {
					tmpLong = tmpLong + pre
					pre = 0
				}

				if tmpLong > max {
					max = tmpLong
				}
				tmpLong = 0

			}
		}
	}

	if top > 0 { //左括号隔开
		if pre > max { //前一段
			max = pre
		}
		if tmpLong-pre > max { //后一段
			max = tmpLong - pre
		}
	} else {
		if tmpLong > max {
			max = tmpLong
		}
	}

	return max
}

*/

/*
()只有这种是有效的
*/
func longestContinuous(s string) int {
	max := 0
	tmp := 0
	top := 0

	for _, v := range s {
		if v == '(' {
			if top == 0 {
				top++
			} else {
				if tmp > max {
					max = tmp
				}
				tmp = 0
			}
		} else {
			if top == 1 {
				top--
				tmp = tmp + 2
			} else {
				if tmp > max {
					max = tmp
				}
				tmp = 0
			}
		}
	}
	if tmp > max {
		max = tmp
	}
	return max
}

func main() {
	//s := "(()))((()())("
	s := "()())((()"
	fmt.Println(longestContinuous(s))
}
