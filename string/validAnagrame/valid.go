/*
242. Valid Anagram
Given two strings s and t , write a function to determine if t is an anagram of s.

Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false
Note:
You may assume the string contains only lowercase alphabets.

Follow up:
What if the inputs contain unicode characters? How would you adapt your solution to such case?
*/

package validAnagrame

func IsAnagram(s, t string) bool {
	charToNum := make(map[rune]int)
	for _, v := range s {
		charToNum[v]++
	}

	for _, v := range t {
		if _, ok := charToNum[v]; !ok {
			return false
		}
		charToNum[v]--
		if charToNum[v] < 0 {
			return false
		}
	}

	for _, v := range charToNum {
		if v > 0 {
			return false
		}
	}

	return true
}
