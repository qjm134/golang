/*
一个有序整数数组，数组中的数可以是正数、负数、零，请实现一个函数，这个函数返回一个整数：返回这个数组所有数的平方值中有多少种不同的取值。
举例：
nums = {-1,1,1,1},
那么你应该返回的是：1。因为这个数组所有数的平方取值都是1，只有一种取值

nums = {-1,0,1,2,3}
你应该返回4，因为nums数组所有元素的平方值一共4种取值：1,0,4,9

思路：
1.存在相同的情况：
a.相邻数字相同
b.整数跟负数
相邻数字可通过保存前一个数，当前数跟前一个比解决
整数跟负数，保存负数map，整数跟map比

2.首尾游标，双向对移动，可以找出正负数相同的
相邻相同的，还是通过保存前一个，对比当前解决
第一次遇到就算一个，后面的都不算
拿来当基准比较的，就是第一个，要先算一个
*/

package main

import (
	"fmt"
	"math"
)

func different(a []int) int {
	if len(a) == 0 {
		return 0
	}
	num := 1 //最后一个
	i := 0
	j := len(a) - 1
	prei := a[0]
	prej := a[j]
	for i < j {
		for i < j {
			if math.Abs(float64(a[i])) > math.Abs(float64(a[j])) {
				if i == 0 {
					num++
					continue
				}
				if a[i] > prei {
					num++
					prei = a[i]
				}
				i++
			} else if math.Abs(float64(a[i])) == math.Abs(float64(a[j])) {
				i++
			} else {
				num++
				prei = a[i]
				j--
				break
			}
		}

		for i < j {
			if math.Abs(float64(a[j])) > math.Abs(float64(a[i])) {
				if a[j] < prej {
					num++
					prej = a[j]
				}
				j--
			} else if math.Abs(float64(a[j])) == math.Abs(float64(a[i])) {
				j--
			} else {
				num++
				prej = a[j]
				i++
				break
			}
		}
	}
	return num
}

func differentNum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	num := 1
	m := make(map[int]struct{})
	m[int(math.Abs(float64(a[0])))] = struct{}{}
	pre := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] == pre {
			continue
		}
		if a[i] < 0 {
			m[int(math.Abs(float64(a[i])))] = struct{}{}
			num++
			pre = a[i]
		} else {
			if _, ok := m[a[i]]; !ok {
				num++
				pre = a[i]
			}
		}
	}
	return num
}

func main() {
	a := []int{-2, -2, -1, 0, 1, 2, 2, 4}
	fmt.Println(differentNum(a))
	fmt.Println(different(a))
}
