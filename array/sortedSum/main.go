/*
剑指 Offer 57 - II. 和为s的连续正数序列

自然数 1，2，3，4，5，6...
给定一个正整数，找出所有和为该数的连续数字
比如
输入：9
输出：[2，3，4] 和 [4，5] 和 [9]

思路：
暴力
连续子序列，从小到大，以每个数字作为子序列的开头，连续向后找，满足和为目标值的序列，的结束数字
和小于目标值，一直向后
和等于目标值，找到，停止，因为再向后数字越来越大，和只会大于
和大于目标值，停止，该数字开头的子序列没有符合的

优化：
滑动窗
输入：11
以1开头的子序列，结尾遍历到5的时候，和为1+2+3+4+5=15，大于目标值了，以1开头的子序列结束了
以2开头的子序列，2+3+4+5这个已经在前一个以1开头的子序列算过了
   以2开始，相当前面以1开始的窗左缩一个，前面的和减去前一个值即可，不需要从2开始重新算一遍2+3+4+5

前一个结束的时候，要么等于目标值了，要么大于目标值
如果是等于目标值，当前结束的和减去前一个，一定小于目标值，结束数字不需要回退，窗右扩
如果是大于目标值，当前和减前一个，可能大于，也可能小于，也可能等于，相当于窗右缩，或者左缩，或者右扩，
   实际上，只需要左缩即可，因为1+2+3+4+5>目标值，则不可能是2+3+4=目标值，即去2头取中间一段，
   因为，中间段如果有等于目标值的，则加上前一个即1，则必大于目标值，那在4就结尾终止了，不会到5再结尾，
   因此，不会右缩，只会不断左缩

结束，窗右边的数字，大于目标值

*/

package main

import "fmt"

func findContinuous(target int) [][]int {
	res := make([][]int, 0)
	start := 1
	end := 1
	sum := start

	for start <= end {
		if sum < target {
			end++
			sum = sum + end
		} else if sum > target {
			sum = sum - start
			start++
		} else {
			res = append(res, digitToSlice(start, end))
			sum = sum - start
			start++
		}
	}

	return res
}

func digitToSlice(start, end int) []int {
	res := make([]int, 0, end-start+1)

	for i := start; i <= end; i++ {
		res = append(res, i)
	}

	return res
}

func main() {
	fmt.Println(findContinuous(9))
}
