/*
https://leetcode.com/problems/two-sum/

Description
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/

/*
遍历
2
9-2=7
查找剩下的是否有7

是否有某元素，可建立一个map，查找o(1)
建立map的同时，判断之前是否有需要的值，比如建立到7的时候，之前2已有了。2个数，遍历到第二个的时候，就能找到第一个了
区别于遍历到第一个，去找第二个
*/

package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))
}

func twoSum(nums []int, target int) []int {
	mNumToIndex := make(map[int]int, len(nums))
	for index, one := range nums {
		another := target - one
		if anotherIndex, ok := mNumToIndex[another]; ok {
			return []int{anotherIndex, index}
		} else {
			mNumToIndex[one] = index
		}
	}
	return []int{}
}
