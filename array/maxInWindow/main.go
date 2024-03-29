/*
剑指 Offer 59 - I. 滑动窗口的最大值
给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

示例:
输入: nums = [1,3,-1,-3,5,3,6,7], 和 k = 3
输出: [3,3,5,5,6,7]

解释:
  滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

提示：
你可以假设 k 总是有效的，在输入数组不为空的情况下，1 ≤ k ≤ 输入数组的大小。

注意：本题与主站 239 题相同：https://leetcode-cn.com/problems/sliding-window-maximum/


思路：
左边先出去，分2种情况：出去的不是最大值；出去的是最大值
右边再进来，在之前2种情况中：
出去的不是最大值
则进来的跟最大值比较，大则更新最大值，小则还是原来的最大值

出去的是最大值
这时进来的值就需要跟前面的比了
那么前面的值需要保存一个次最大值，当最大值出去后，次最大值顶上

用一个变量保存当前最大值，右边进来的比这个大，则更新最大值
左边出去的，如果就是当前最大值，则左边的 - 当前进来值前 的这段区间的最大值没保存，导致大的时间复杂度。
如果进来的是最大值，则把最大值放最前面，后面进来的比这个小，则列在后面；
如果比尾部的大，从尾部开始弹出，直到比尾部的小，放到尾部。
最大的到次二大的之间的，在最大的弹出后，次二大就是最大的，次二大前面的都没用
Win = 3
3 1
3 1

3 1 2
3 2

3 1 2 4
4

虽然也要从队尾向前依次比较，但是，是弹出的，比较过的，下一个不会再跟他比较了，不像数组，每次都在那，后面的每个都得与其比较一次

*/
package main

import "fmt"

func max(nums []int, k int) []int {
	if k > len(nums) {
		k = len(nums)
	}

	res := make([]int, 0, len(nums)-k+1)
	tmp := make([]int, 0, len(nums))

	left := 0
	for right := 0; right < len(nums); right++ {
		for len(tmp) > 0 && nums[right] > tmp[len(tmp)-1] {
			tmp = tmp[:len(tmp)-1]
		}
		tmp = append(tmp, nums[right])

		if right+1-k >= 0 {
			res = append(res, tmp[0])

			if nums[left] == tmp[0] {
				tmp = tmp[1:]
			}
			left++
		}
	}
	return res
}

func findMax(nums []int, k int) []int {
	if len(nums) <= 0 {
		return []int{0}
	}

	que := make([]int, 0)
	res := make([]int, 0, len(nums)-k+1)

	for i, j := 0-k+1, 0; j < len(nums); i, j = i+1, j+1 {
		if i > 0 {
			if nums[i-1] == que[0] {
				que = que[1:len(que)]
			}
		}
		for len(que) > 0 {
			if nums[j] <= que[len(que)-1] {
				que = append(que, nums[j])
				break
			} else {
				que = que[:len(que)-1]
			}
		}
		if len(que) == 0 {
			que = append(que, nums[j])
		}

		if i >= 0 {
			res = append(res, que[0])
		}
	}

	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	fmt.Println(findMax(nums, k))
}
