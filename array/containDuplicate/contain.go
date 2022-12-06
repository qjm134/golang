/*
217. Contains Duplicate
Given an array of integers, find if the array contains any duplicates.

Your function should return true if any value appears at least twice in the array, and it should return false if every element is distinct.

Example 1:

Input: [1,2,3,1]
Output: true
Example 2:

Input: [1,2,3,4]
Output: false
Example 3:

Input: [1,1,1,3,3,4,3,2,4,2]
Output: true

思路一
排序（快排，再比较相邻）
时间复杂度：O(NlogN)，其中 N 为数组的长度。需要对数组进行排序。
空间复杂度：O(logN)，其中 N 为数组的长度。注意我们在这里应当考虑递归调用栈的深度。

思路二
map
时间复杂度：O(N)，其中 N 为数组的长度。
空间复杂度：O(N)，其中 N为数组的长度
*/
package containDuplicate

func IsDuplicate(nums []int) bool {
	dataToNum := make(map[int]int)
	for _, v := range nums {
		dataToNum[v]++
		if dataToNum[v] >= 2 {
			return true
		}
	}
	return false
}
