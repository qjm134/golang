/*
给你输入一个数组nums和一个正整数k，请你判断nums是否能够被平分为元素和相同的k个子集。
[1,1,2,3,1,2,2]  k=3
[1,1,2] [3,1] [2,2]

思路：
求总和：12
每个组和：12/k=4

从数组中找数，能凑成4
找k次，每次都能凑成和=4，最后数据刚好用完

每次找个系数组合，1,1,1,0,0,0,0，跟数组元素乘积和刚好是4
               0,0,0,1,1,0,0
               0,0,0,0,0,1,1

竖着看，就是每个元素只能在k组中选一组

*/
package main

import "fmt"

func oneBucket(nums []int, target, curTar int, pos int, res []int) {
	if pos == len(nums) {
		if curTar == target {
			fmt.Println(res)
		}
		return
	}

	xi := []int{0, 1}
	for _, v := range xi {
		curTar = curTar + v*nums[pos]
		if curTar > target {
			return
		}
		res = append(res, v)
		oneBucket(nums, target, curTar, pos+1, res)
		curTar = curTar - v*nums[pos]
		res = res[:len(res)-1]
	}
}

func main() {
	nums := []int{3, 2, 1, 4, 6, 5}
	target := 5
	res := make([]int, 0)

	oneBucket(nums, target, 0, 0, res)
}
