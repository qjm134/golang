/*
76. Minimum Window Substring
Given two strings s and t, return the minimum window in s which will contain all the characters in t. If there is no such window in s that covers all characters in t, return the empty string "".

Note that If there is such a window, it is guaranteed that there will always be only one unique minimum window in s.



Example 1:

Input: s = "ADOBECODEBANC", t = "ABC"
Output: "BANC"
Example 2:

Input: s = "a", t = "a"
Output: "a"


Constraints:

1 <= s.length, t.length <= 105
s and t consist of English letters.


Follow up: Could you find an algorithm that runs in O(n) time?

思路：
窗，右扩/左缩

map，窗内是否有该字符，有几个
slice，窗内字符位置

入窗
开始字符，是否重复，位置队列长度==1，入队的时候看是否跟开始字符重复，若重复，则开始字符移到当前字符：入窗，开始字符出窗

出窗
开始字符，是否重复，原开始字符出窗后，新开始字符是否跟下一个字符重复，通过位置队列位置，找s，得到字符，若相同，出窗

1.遍历t，字符建map（小于23个字母），知道一共多少字符
2.遍历s，有个窗map[byte]int，key是字符，value是字符出现次数，可以跟第一步map共用
3.有个窗中不重字符数，字符在map中，若value ==0，则之前未在窗中出现过，不重数+1，map 数量+1
  若value > 0，则该字符出现过，位置入队
  需要记录第一个入窗的字符，还有一个最小窗宽，及最小窗宽开始位置 和 结束位置
  如果不重数等于了t中字符数，则计算窗宽：最后入窗的字符位置 - 窗开始的字符位置 + 1
  若窗宽小于最小窗宽，则更新最小窗宽，及窗开始及结束位置
  还要一个窗内字符位置队列
  窗开始字符出窗，即其位置出队，map 数-1，若map新数>0，则出窗，直到map数=0，不重字符数--
*/
package minWinSub

func minWin(s, t string) string {
	return ""
}
