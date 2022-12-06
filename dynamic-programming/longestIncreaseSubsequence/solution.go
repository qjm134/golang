/*
300. Longest Increasing Subsequence
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements.
For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].



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

// 2, -1, 3, 1, 4
// 2开始的递增序列: 2，3，4  不在这个序列中的: -1, 1
// 再从-1开始找递增序列: -1, 3  或者 -1, 1, 4
// 下面算法有问题，没有考虑上面-1开头有几种的情况
func posiableStart(nums []int) int {
	max := 0

	sStart := make([]int, 0)
	sStart = append(sStart, 0)
	for len(sStart) > 0 {
		start := sStart[0]
		sStart = sStart[1:len(sStart)]
		tmpLen := 1
		pre := nums[start]
		for i := start + 1; i < len(nums); i++ {
			if nums[i] > pre {
				tmpLen++
				pre = nums[i]
				if start > 0 {
					for idx, numIdx := range sStart {
						if numIdx == i {
							sStart = append(sStart[:idx], sStart[idx+1:]...)
							break
						}
					}
				}
			} else {
				if start == 0 {
					sStart = append(sStart, i)
				}
			}
		}
		if tmpLen > max {
			max = tmpLen
		}
	}

	return max
}
