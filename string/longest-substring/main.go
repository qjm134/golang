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
	count清0
	遍历回到存在的字符位置后一个
	map清零
*/
package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcbde"))
}

func lengthOfLongestSubstring(s string) int {
	output, count := 0, 0
	charToIndex := make(map[byte]int)

	for i := 0; i < len(s); {
		if ci, ok := charToIndex[s[i]]; ok {
			if count > output {
				output = count
			}
			count = 0
			i = ci + 1
			charToIndex = make(map[byte]int)
		} else {
			charToIndex[s[i]] = i
			count++
			i++
		}
	}

	if count > output {
		output = count
	}

	return output
}
