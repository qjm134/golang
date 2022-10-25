/*
121. Best Time to Buy and Sell Stock
https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

Description
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

思路：
遍历，每一个当做开始，从后面找一个最大的，当做结束；保存一个当前最大的，每一个开始的收益，跟当前的比，大则更新
反过来想，每一个当做结束，前面找一个最小的，这个最小的不需要找，遍历的时候保持一个当前最小的
*/

package main

import (
	"fmt"
)

func main() {
	p := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(BestTimeOptimize(p))

	p = []int{7, 6, 4, 3, 1}
	fmt.Println(BestTimeOptimize(p))

	p = []int{6, 3, 7, 1, 2}
	fmt.Println(BestTimeOptimize(p))
}

func BestTimeOptimize(prices []int) (buyDay, sellDay, profit int) {
	if len(prices) == 0 {
		return 0, 0, 0
	}
	minPrice := prices[0]
	minDay := 0
	for i, price := range prices {
		tmpProfit := price - minPrice
		if tmpProfit > profit {
			profit = tmpProfit
			sellDay = i
			buyDay = minDay
		}
		if price < minPrice {
			minPrice = price
			minDay = i
		}
	}
	return buyDay, sellDay, profit
}

func bestTime(prices []int) (buyDay, sellDay, profit int) {
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			tmpProfit := prices[j] - prices[i]
			if tmpProfit > profit {
				profit = tmpProfit
				buyDay = i
				sellDay = j
			}
		}
	}
	return buyDay, sellDay, profit
}
