/*
Description
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note: You may assume that you have an infinite number of each kind of coin.

提示：
1 <= coins.length <= 12
1 <= coins[i] <= 2^31 - 1
0 <= amount <= 10^4

思路：
1. 尽可能多用大的数，排序，从大的开始，总数除大的，余数再从剩下的凑，11/5=2余1，用2个5，剩下的1再除2，不行再除1；如果最小的一个不能整除，则弃掉上一个大的。这样也不行，比如 [1,3,4,5]，amout=7，则5+2*1不如3+4
2. a1*n1 + a2*n2 + a3*n3 + .... = amout，其实就是求系数，可以求出所有等于的系数，然后找一个a1+a2+a3+...最小的
   每个系数的取值：0,1,2...amount/n1, 0-amout/n2, ...，有多少种硬币循环几重，最多amount^len(coins)
*/

package main

import (
	"fmt"
)

func main() {
	coins := []int64{1, 2, 5}
	amount := int64(11)
	fmt.Println(fewestNumber(coins, amount))

	coins = []int64{2}
	amount = int64(3)
	fmt.Println(fewestNumber(coins, amount))

	coins = []int64{1}
	amount = int64(0)
	fmt.Println(fewestNumber(coins, amount))

	coins = []int64{1, 3, 4, 5}
	amount = int64(7)
	fmt.Println(fewestNumber(coins, amount))
}

func fewestNumber(coins []int64, amount int64) int64 {
	return everyCoin(0, coins, amount)
}

// 如果知道元素个数, 比如 coins = [1, 2, 5], amount = 11
// for i:=0; i<amount/coins[0]; i++ {
// 		for j:=0; j<amount/coins[1]; j++ {
//			for k:=0; k<amount/coins[2]; k++ {
//				if i*coins[0] + j*coins[1] + k*coins[2] == amount { //确定是个有效的系数组合
//					if i+j+k < min {
//						min = i+j+k //找最小的系数组合
//					}
//				}
//			}
//		}
// }
//
// 每种硬币，[0, amount/coin]种系数中，选一个最小的
// 前后关联的状态，总数减去当前这种状态，剩下的对下一种硬币的影响
// 对于每种硬币，[0, amount/coin]系数中总有一种是有效的，否则整个无效-1； 有效就是使得，递归到最后一种硬币的时候，amount可以减到 0
func everyCoin(index int, coins []int64, amount int64) (fewestNum int64) {
	if amount == 0 { //最后减到0，有效
		return 0
	}
	if index >= len(coins) { //到最后没减到0，无效
		return -1
	}

	var curFewest int64 = 10000
	coin := coins[index]
	// 对于每种硬币，[0, amount/coin]系数中选一个最小的；总要有一种是有效的，否则整个无效-1
	for i := int64(0); i <= amount/coin; i++ {
		preNum := everyCoin(index+1, coins, amount-i*coin)
		if preNum != -1 { //有效的中，找一个最小的
			if preNum+i < curFewest {
				curFewest = preNum + i
			}
		}
	}

	if curFewest == 10000 { // 系数中有有效的，总能得到一个最小值；没有有效的，则循环中没有找到有效的改变最小值
		return -1
	}

	return curFewest
}
