/*
有N个小朋友站在一排，每个小朋友都有一个评分
你现在要按以下的规则给孩子们分糖果：
1.每个小朋友至少要分得一颗糖果
2.分数高的小朋友要他比旁边得分低的小朋友分得的糖果多
3.分数一样的，糖果数也要一样
你最少要分发多少颗糖果？
比如：
[1,5,3,1]
答案：7，分法[1,3,2,1]


思路：
一、
依次遍历，当前的与前一个比，3种情况
1.大，则前一个数量再加1个，即当前数量
2.相等，则当前数量等于前一个数量
2.小，前一个数量至少减1
     如果前一个数量就是1，那减了就是0了，不能满足至少1个
     那当前数量为1，往前遍历，前一个加1，再前一个如果还大，那需要加1；如果小就不需要了，退出不继续往前；相等也要加1
     如果前一个数量大于1，则当前数量等于1


二、
从前往后遍历一遍
大于前一个 = 前一个 + 1
等于 = =
小于 = 1，这样小于的就可能不满足了
然后从后往前遍历一遍，跟上面一样
最后2个值，取大的即可，2个值，一个是满足左边关系的，一个是满足右边的，取大的2遍就都能满足了
原始数：[1,4,3,2,2,1,0]
前到后：[1,2,1,1,1,1,1]
后到前：[1,5,4,3,3,2,1]
最后的：[1,5,4,3,3,2,1]
*/

package main

import "fmt"

func minOptimize(a []int) int {
	if len(a) == 0 {
		return 0
	}
	if len(a) == 1 {
		return 1
	}

	left := make([]int, len(a))
	left[0] = 1
	for i := 1; i < len(a); i++ {
		if a[i] > a[i-1] {
			left[i] = left[i-1] + 1
		} else if a[i] == a[i-1] {
			left[i] = left[i-1]
		} else {
			left[i] = 1
		}
	}

	right := 1
	min := maxcan(1, left[len(a)-1])
	for j := len(a) - 2; j >= 0; j-- {
		if a[j] > a[j+1] {
			right++
		} else if a[j] < a[j+1] {
			right = 1
		}
		right = maxcan(left[j], right)
		min = min + right
	}
	return min
}

func maxcan(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func minCandy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	if len(ratings) <= 1 {
		return 1
	}

	t := make([]int, len(ratings))
	t[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			t[i] = t[i-1] + 1
		} else if ratings[i] == ratings[i-1] {
			t[i] = t[i-1]
		} else {
			if t[i-1]-1 <= 0 {
				t[i] = 1
				t[i-1]++
				for j := i - 1; j > 0; j-- {
					if ratings[j] <= ratings[j-1] {
						t[j-1]++
					} else {
						break
					}
				}
			} else {
				t[i] = 1
			}
		}
	}
	min := 0
	for _, v := range t {
		min = min + v
	}
	return min
}

func main() {
	ratings := []int{1, 5, 3, 1}
	fmt.Println(minCandy(ratings))
	fmt.Println(minOptimize(ratings))
}
