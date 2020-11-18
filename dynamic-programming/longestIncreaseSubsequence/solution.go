/*
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].



Example 1:

Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Example 2:

Input: nums = [0,1,0,3,2,3]
Output: 4
Example 3:

Input: nums = [7,7,7,7,7,7,7]
Output: 1


Constraints:

1 <= nums.length <= 2500
-104 <= nums[i] <= 104


Follow up:

Could you come up with the O(n2) solution?
Could you improve it to O(n log(n)) time complexity?

思路：
以数组中的每个元素开头，遍历，选这n个开头中的最长的
以当前元素开头，后面可能的再开头的，就是比当前元素大的作为开头，遍历比当前元素大的开头，选最长的 + 当前元素1，得到以当前元素开头最大的
*/
package longestIncreaseSubsequence

func LengthOfLIS(nums []int) int {
	for k, _ := range startToLength {
		delete(startToLength, k)
	}
	max := 0
	for i := 0; i < len(nums); i++ {
		res := maxLength(nums, i)
		if res > max {
			max = res
		}
	}
	return max
}

var startToLength = make(map[int]int)

func maxLength(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}
	maxTemp := 0
	for j := start + 1; j < len(nums); j++ {
		if nums[j] > nums[start] {
			var res int
			var ok bool
			if res, ok = startToLength[j]; !ok {
				res = maxLength(nums, j)
				startToLength[j] = res
			}
			if res > maxTemp {
				maxTemp = res
			}
		}
	}

	return maxTemp + 1
}
