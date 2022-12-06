/*
42
给出一个数组代表围柱的高度，求能围柱的最大的水量，例如数组{ 5，2，3，2，4 }，最大水量为5

思路：
1.
每根柱上可以蓄的数量是，当前柱左边最高的柱，右边最高的柱，2者中矮的那个决定当前柱蓄水高度
蓄水高度 - 当前柱高度 = 蓄水量
当前柱左边最大的值，可以从左向右遍历一遍，遍历到当前柱时，当前保存的最大值
右边的最大值，从右向左遍历一遍

2.
左右2端往里走
左右2边比，如果右边小，右边的节点为当前节点
当前节点，右边的最大值，即是当前节点，可以蓄水的最大高度
因为，左边大的这个往里
如果有比左边大的更大，也没用，因为右边矮，高度由右边矮的决定了
如果小，那还有左边这个大的挡住，水能上蓄
右边小，右边往里移动一步

*/

package main

import "fmt"

func water(nums []int) int {
	i := 0
	j := len(nums) - 1
	maxLeft := 0
	maxRight := 0
	w := 0

	for i < j {
		if nums[i] > maxLeft {
			maxLeft = nums[i]
		}
		if nums[j] > maxRight {
			maxRight = nums[j]
		}

		if maxLeft > maxRight {
			w = w + (maxRight - nums[j])
			j--
		} else {
			w = w + (maxLeft - nums[i])
			i++
		}
	}
	return w
}

func main() {
	nums := []int{5, 2, 3, 2, 4}
	fmt.Println(water(nums))
}

func maxRain(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}

	rain := 0
	left := 0
	right := len(nums) - 1
	leftMax := nums[left]
	rightMax := nums[right]
	for left <= right {
		if leftMax <= rightMax {
			tmp := leftMax - nums[left]
			if tmp > 0 {
				rain = rain + tmp
			}

			if nums[left] > leftMax {
				leftMax = nums[left]
			}

			left++
		} else {
			tmp := rightMax - nums[right]
			if tmp > 0 {
				rain = rain + tmp
			}

			if nums[right] > rightMax {
				rightMax = nums[right]
			}

			right--
		}
	}

	return rain
}
