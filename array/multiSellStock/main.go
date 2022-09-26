/*
122. 买卖股票的最佳时机 II
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。

在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。

返回 你能获得的 最大 利润 。


股票可以买卖多次, 但是你必须要在买股票之前把之前买的股票卖掉。
示例 1:

输入: [7,1,5,3,6,4]
输出: 7
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 3 天（股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     随后，在第 4 天（股票价格 = 3）的时候买入，在第 5 天（股票价格 = 6）的时候卖出, 这笔交易所能获得


输入: [1,2,3,4,5]
输出: 4
解释: 在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
     注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
     因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。


思路：
找极小值，买，往后递增到最大值，卖；单调递增，中间依次卖，跟最小最大卖，钱一样
[1,2,3] 1买2卖，2买3卖 -> 跟1买3卖，钱一样    从极小到极大，单调增

如果不单调增，则中间有折叠，多次卖就比一次卖，要多，因为掉下去，折叠重复就多出来了
[1,4,2,6]   [1-4] [2-6]：则[1-2 2-4]和[2-4 4-6]，中间[2-4]就折叠多出来一段了，比直接[1-6]

所以，找出单调递增段，即为出售次数

[1,3,5,2,1,4,7]
[1,3,5]  [1,4,7]

实现：当前跟前一个比，大则卖，算一个利润，小则继续看下一个，不累加利润
*/

package main

import "fmt"

func main() {
	//nums := []int{7, 1, 5, 3, 6, 4}
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	fmt.Println(extremum(nums))
}

func maxProfit(a []int) int {
	if len(a) <= 1 {
		return 0
	}

	profit := 0
	for i := 1; i < len(a); i++ {
		tmp := a[i] - a[i-1]
		if tmp > 0 {
			profit = profit + tmp
		}
	}
	return profit
}

const (
	unknown = iota
	up
	down
)

// 极值法
func extremum(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	preFlag := unknown
	onlyTwoUp := 0
	minLocal := nums[0]
	max := 0
	for i := 1; i < len(nums); i++ {
		if preFlag == unknown {
			if nums[i] > nums[i-1] {
				preFlag = up

				if i == len(nums)-1 { // 才开始循环就结束了，进不到下面有计算利润的分支，因此需要在这计算利润
					max = nums[i] - nums[i-1]
					onlyTwoUp = 1
				}
			}
			if nums[i] < nums[i-1] {
				preFlag = down
			}
		} else {
			if preFlag == up && nums[i] < nums[i-1] { // 找到极大值
				max = max + (nums[i-1] - minLocal)
				preFlag = down
			}

			if preFlag == down && nums[i] > nums[i-1] { // 找到极小值
				minLocal = nums[i-1]
				preFlag = up
			}
		}
	}

	if onlyTwoUp == 0 && preFlag == up { // 递增结束
		max = max + (nums[len(nums)-1] - minLocal)
	}

	return max
}
