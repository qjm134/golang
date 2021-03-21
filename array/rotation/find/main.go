/*
33. 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同。
在传递给函数之前，nums 在预先未知的某个下标 target（0 <= target < nums.length）上进行了 旋转，使数组变为 [nums[target], nums[target+1], ..., nums[n-1], nums[0], nums[1], ..., nums[target-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的索引，否则返回 -1 。


示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例 2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4

进阶：你可以设计一个时间复杂度为 O(log n) 的解决方案吗？


思路：
有序数组的查找，二分，跟中间值比较

旋转数组，分2个递增序列，前一半都比后一半大，前一半的第一个值都比后一半的所有值大
分2种情况：
a.前面大的部分过半，中间值大于第一个值
b.后面小的部分过半，中间值小于第一个值

456789 123
1.在第一种情况下
  a.待查找值比中间值大，那一定是在右边
  b.比中间值小，2边都有可能。再跟第一个比，比第一个大，则在左边，小则在右边

789 123456
2.第二种情况
  a.比中间值大，2边都有可能。再跟第一个比，大则在左边，小则右
  b.比中间值小，一定在左边

3412
以上基于旋转数组的结论，移动指针后，如果不是旋转数组了，比如就剩递增数组了
递增数组，中间值大于第一个，以后都走第一种情况
比中间值小，肯定比第一大，在左边没问题

旋转数组还是递增数组，以上都适用
*/

package main

import "fmt"

func find(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		if nums[mid] >= nums[start] {
			if target > nums[mid] {
				start = mid + 1
			} else if target < nums[mid] {
				if target > nums[start] {
					end = mid - 1
				} else if target < nums[start] {
					start = mid + 1
				} else {
					return start
				}
			} else {
				return mid
			}
		} else {
			if target < nums[mid] {
				end = mid - 1
			} else if target > nums[mid] {
				if target > nums[start] {
					end = mid - 1
				} else if target < nums[start] {
					start = mid + 1
				} else {
					return start
				}
			} else {
				return mid
			}
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 1, 2, 3}
	target := 5
	fmt.Println(find(nums, target))
}
