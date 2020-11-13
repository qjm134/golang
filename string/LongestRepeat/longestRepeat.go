/*
424. Longest Repeating Character Replacement
Medium

Given a string s that consists of only uppercase English letters, you can perform at most k operations on that string.

In one operation, you can choose any character of the string and change it to any other uppercase English character.

Find the length of the longest sub-string containing all repeating letters you can get after performing the above operations.

Note:
Both the string's length and k will not exceed 104.

Example 1:

Input:
s = "ABAB", k = 2

Output:
4

Explanation:
Replace the two 'A's with two 'B's or vice versa.


Example 2:

Input:
s = "AABABBA", k = 1

Output:
4

Explanation:
Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.

思路：
暴力，找出所有可能的，然后选一个最长的

1.有问题
可能的：从哪开始，哪结束
开始，遍历每个字符串，作为开始，优化为从不同的字符再次开始
结束，跟开始的字符不同的，减k，k为0，则结束
问题：开始的字符可能会变

2.某一个长度的字符串，重复最多的字符为基准，剩下的字符如果刚好是k个，则不论重复最多的字符怎么排列，都可以讲剩下的字符变为基准字符满足要求
长度用当前最大的
怎么找出所有的：以每个字符开始，找到当前字符开始的，可以达到的最大长度；遍历每个字符作为开始
假如最大的长度出现在中间某个段，则这个长度前面不可能有符合的，因为前面以某个字符开始的，符合的长度已经都检查的过了，达不到当前大的长度才继续往后遍历的

优化：
找当前字符开始的最大的长度，从当前最大的长度maxLen开始扩，因为前面已经可以达到当前的长度了，当前字符连当前的长度都达不到，肯定不满足，则直接往后移一个；找的是最大的，可以过滤掉一些不遍历，只关注大的就行
字符数量最多的max值也保持，如果没有比这个max再大的，肯定也不能满足maxLen - max < k

窗口中字符最多的数，能引起这个数变化的有2种，一是长度扩大，进入一个新的字符；二是右移，进入一个新的，减少一个左边的字符。

如果是进入一个新的，则新的2种可能，一是原来最多的，二不是原来的。
是原来最多的，max跟进来这个比，2取大没问题。如果不是原来最多的，因为只有新进的字符有变化，只需要看这个变化的跟当前最多的比就行。

如果是右移，右边新进的同上；左边减少的，max不需要减，因为maxLen-max>k（右移），max不满足，新进的数量如果不能比这个max大，则肯定也不满足
*/
package LongestRepeat

func CharacterReplacement(s string, k int) int {
	var maxLen, maxChar int
	var start int
	charToNum := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		charToNum[s[i]]++
		maxLen++
		if charToNum[s[i]] > maxChar {
			maxChar = charToNum[s[i]]
		}
		if maxLen > (maxChar + k) {
			charToNum[s[start]]--
			start++
			maxLen--
		}
	}
	return maxLen
}

func FindLongestRepeat(s string, k int) int {
	longest := 1
	charToNum := make(map[byte]int)
	var max int
	var maxChar byte

	for i := 0; i < len(s); i++ {
		for max+k >= longest {
			longest++
			if i+longest == len(s) {
				return longest
			}
			if i+longest > len(s) {
				return longest - 1
			}
			charToNum[s[i+longest-1]]++
			if charToNum[s[i+longest-1]] > max {
				max = charToNum[s[i+longest-1]]
				maxChar = s[i+longest-1]
			}
		}
		longest--
		charToNum[s[i]]--
		if s[i] == maxChar {
			max--
		}
	}
	return 0
}
