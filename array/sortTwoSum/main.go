/*
排序数组，找2个数，和为给定值
[1,2,3,4,5] k=6
[1,5]
*/

package main

import "fmt"

func findTwo(nums []int, k int) []int {
	i := 0
	j := len(nums) - 1

	for i < j {
		sum := nums[i] + nums[j]
		if sum > k {
			j--
		} else if sum < k {
			i++
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 5, 6, 7}
	k := 7
	fmt.Println(findTwo(nums, k))
}
