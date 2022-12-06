/*
现场编程题题目内容：
一个旗舰店中有若干商品，其价格记录于数组 prices 中，下标表示商品种类编号，值表示对应商品价格。

为了迎接到来的购物节，旗舰店准备了足够多的商品，并推出了若干优惠活动。promotions[i] 和 discounts[i] 分别描述了第 i 项优惠活动的组合条件和对应的减免金额。
·         promotions[i] 是一个数组，格式为 [id0,num0,id1,num1,...,idN,numN]，可以是一种或多种商品的组合，其中每种商品的 id,num 表示编号 id 的商品最少购买 num 件；
·         当满足优惠活动 promotions[i] 的优惠条件时，可获得 discounts[i] 的金额减免；
·         一种商品至多出现在一个优惠活动中，不会出现在多个优惠活动中。
现给定某顾客的购物清单 order，格式为 [id0,num0, id1,num1, ..., idM,numM] ，其中某种商品的 id,num 表示编号 id 的商品购买了 num 件。
该顾客很节俭，他会参加所有可优惠的活动，同一个优惠活动也允许参与多次，请计算该顾客需要支付多少金额。

示例 1：

输入：
prices = [10,5,8,8,6,3]
promotions = [[0,7],[4,10],[2,6,1,9],[5,2]]
discounts = [4,3,5,1]
order = [2,17,3,10,1,27,5,2,4,9]

输出: 400

解释：
优惠活动 0 表示每购买 7 个商品 0，可以获得减免金额 4。顾客未购买商品 0，不享受该优惠减免
优惠活动 1 表示每购买 10 个商品 4，可以获得减免金额 3。顾客购买商品 4 的数量达不到优惠条件
优惠活动 2 表示每购买 6 个商品 2 以及 9 个商品 1，可以获得减免金额 5。顾客购买了 17 个商品 2 和 27 个商品 1，享受两次该优惠减免
优惠活动 3 表示每购买 2 个商品 5 ，可以获得减免金额 1。顾客购买了 2 个商品 5 ，享受一次该优惠减免
总支付金额为 8*17 + 8*10 + 5*27 + 3*2 + 6*9 - 2*5 - 1*1 = 400

示例 2：

输入：
prices = [10,10,10]
promotions = [[1,2],[0,5,2,3]]
discounts = [10,10]
order = [0,10]

输出: 100

解释：
优惠活动 1 表示每购买 5 个商品 0 以及 3 个商品 2，可以获得减免金额 10。该顾客仅购买了商品 0 ，未购买商品 2，所以不能享受该优惠减免

示例 3：

输入：
prices = [10,4]
promotions = []
discounts = []
order = [0,2,1,3]

输出: 32

提示：
·         1 <= prices.length <= 1000
·         1 <= prices[i] <= 10000
·         0 <= promotions.length == discounts.length <= 100
·         2 <= promotions[i].length <= 200
·         0 <= promotions[i][2*j] < prices.length
·         1 <= promotions[i][2*j+1] <= 100
·         1 <= discounts[i] <= 商品总原价
·         2 <= order.length <= 2000
·         0 <= order[2*i] < prices.length
·         1 <= order[2*i+1] <= 100
*/
package main

func count(price []int, promotions [][]int, discounts []int, order []int) int {
	total := 0
	totalInit := countInit(price, order)
	totalDiscount := countDiscount(order, promotions, discounts)
	total = totalInit - totalDiscount

	return total
}

func countInit(price []int, order []int) int {
	total := 0
	for i := 0; i < len(order); i = i + 2 {
		total = total + price[order[i]]*order[i+1]
	}
	//fmt.Println(total)
	return total

}

/*
prices = [10,5,8,8,6,3]
promotions = [[0,7],[4,10],[2,6,1,9],[5,2]]
discounts = [4,3,5,1]
order = [2,17,3,10,1,27,5,2,4,9]
*/
func countDiscount(order []int, promotions [][]int, discounts []int) int {
	mHas := make(map[int]struct{})
	total := 0
	for i := 0; i < len(order); i = i + 2 {
		if _, ok := mHas[i]; ok {
			continue
		} else {
			mHas[i] = struct{}{}
		}
		for j := 0; j < len(promotions); j++ {
			for k := 0; k < len(promotions[j]); k = k + 2 {
				if promotions[j][k] == order[i] {
					num := order[i+1] / promotions[j][k+1]
					minNum := num
					enough := 1
					for q := 0; q < len(promotions[j]); q = q + 2 {
						if q == k {
							continue
						}
						for p := 0; p < len(order); p = p + 2 {
							if promotions[j][q] == order[p] {
								enough++
								num = order[p+1] / promotions[j][q+1]
								if num < minNum {
									minNum = num
								}
								mHas[p] = struct{}{}
							}
						}
					}
					if enough == len(promotions[j])/2 {
						total = total + minNum*discounts[j]
					}
				}
			}
		}
	}

	return total
}
