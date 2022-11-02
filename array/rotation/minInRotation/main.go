/*
剑指 Offer 11. 旋转数组的最小数字
把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
给你一个可能存在 重复 元素值的数组 numbers ，它原来是一个升序排列的数组，并按上述情形进行了一次旋转。
输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
例如，数组[3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。
        [4,5,1,2,3]


思路：
遍历一遍数组可以找到最小值，时间复杂度n
排序数组，找某个值，用二分可以达到logn的复杂度
这也是排序的，可以借鉴二分的思路，每次可以排除一些不用遍历
中间值跟谁比较呢，跟2端的比，比较无非大于，小于，等于3种情况，每种情况是什么意思呢，可以怎么决定最小值在哪半呢
满足什么的条件，就是要找的最小值呢

旋转，其实就是把递增数组，在中间某个地方切一刀，分2段，把后半部分大的那段放到小的那段的前面
分2段，前一段必定是大的，后一段是小的，后一段的所有值都比大的那段的最小值，就是数组的第一个值还小
比如，3,4,5最小值3都比后半段1,2都大
非旋转的正常的，则第一个值比最后一个值小，这是结束的条件

如果中间值比第一个值大，说明中间值肯定不在小的那段，因为小的那段全部都比大的那段的最小值即第一个值都小，就2段只能是在大的那段
因此，最小值在右半边

如果中间值比第一个小，说明中间值肯定不在大的那段，因为大的那段是递增的，后面的值肯定都比前面的值大，不会出现小的情况，那就是在小的那段
因此，最小值在左半边
最小值，比前一个小
第一个值的前一个值是最后一个值

中间值跟第一个相等


移动的时候，需要保证仍然是旋转数组，因此不能把中间点给移走
一开始需要排除没有旋转的数组，因为后面的判断都是基于旋转数组得出的理论，没有旋转的数组是不满足那些判断条件的
*/

package main

import "fmt"

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	//nums := []int{11, 13, 15, 17}
	nums := []int{0, 0, 1, 2, 0}
	fmt.Println(minArray(nums))
}

func minArray(nums []int) int {
	start := 0
	end := len(nums) - 1
	for start < end {
		if nums[start] < nums[end] {
			return nums[start]
		}

		mid := (start + end) / 2
		if nums[mid] > nums[start] {
			start = mid + 1
			continue
		}

		if nums[mid] < nums[start] {
			end = mid
			continue
		}

		if nums[mid] == nums[start] {
			start++
			continue
		}
	}

	return nums[start]
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	start := 0
	end := len(nums) - 1
	for start < end {
		mid := (start + end) / 2
		if nums[mid] < nums[(mid-1+len(nums))%len(nums)] {
			return nums[mid]
		}
		if nums[mid] > nums[(mid+1)%len(nums)] {
			return nums[(mid+1)%len(nums)]
		}

		if nums[mid] < nums[start] {
			end = mid
		} else {
			start = mid
		}
	}
	return nums[start]
}
