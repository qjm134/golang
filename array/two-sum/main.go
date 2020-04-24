/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
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
