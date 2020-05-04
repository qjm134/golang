/*
3. Longest Substring Without Repeating Characters
Medium

8511

517

Add to List

Share
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

/*
遍历字符串，字符做map的key，value为字符的位置
map中是否存在
不存在，放入map，count加1
存在，count跟output比较，count大于output，更新output为count，否则不变
	（下一个子串计算开始的位置为：重复字符第一个出现的位置的后一个。新串起始位置到当前位置已经遍历过了，可以从当前位置继续）
	count减去第一个重复字符及前的数
	map清掉第一个重复字符及前的
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	output, count := 0, 0
	charToIndex := make(map[byte]int)
	begin := 0

	for i := 0; i < len(s); {
		if ci, ok := charToIndex[s[i]]; ok {
			if count > output {
				output = count
			}
			count = count - (ci - begin + 1)
			//charToIndex = make(map[byte]int)
			//内存、运行时间都会提高
			for ; begin <= ci; begin++ {
				delete(charToIndex, s[begin])
			}
		}
		charToIndex[s[i]] = i
		count++
		i++
	}

	if count > output {
		output = count
	}

	return output
}
