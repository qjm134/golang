/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
示例 3：

输入：s = "a"
输出："a"
示例 4：

输入：s = "ac"
输出："a"


提示：

1 <= s.length <= 1000
s 仅由数字和英文字母（大写和/或小写）组成


思路：
暴力
最长回文，总以某个字符为中心，遍历所有字符，作为中心
向2遍扩展，判断对称位置是否相等，相等继续扩，不等停止，记录串长，及回文子串
中心有2种情况
1.跟前一个相等
2.跟前第二个相等
*/

package LongestPalindrome

func longestPal(s string) string {
	if len(s) <= 0 {
		return ""
	}
	max := 1
	maxStr := s[0:1]
	for k := 1; k < len(s); k++ {
		var left, right int
		if s[k] == s[k-1] {
			left = k - 2
		} else if k > 1 && s[k] == s[k-2] {
			left = k - 3
		} else {
			continue
		}

		right = k + 1
		for left >= 0 && right < len(s) {
			if s[left] == s[right] {
				left--
				right++
			} else {
				break
			}
		}
		maxTmp := right - left - 1
		if maxTmp > max {
			max = maxTmp
			maxStr = s[(left + 1):right]
		}
	}
	return maxStr
}

func LongestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	max := 1
	maxStr := s[0:1]
	rear := len(s) - 1

	for start := 0; start < rear; start++ {
		for end := rear; end-start+1 > max; end-- {
			i := start
			j := end
			for i < j {
				if s[i] == s[j] {
					i++
					j--
				} else {
					break
				}
			}

			if i >= j {
				max = end - start + 1
				maxStr = s[start : end+1]
			}
		}
	}
	return maxStr
}
