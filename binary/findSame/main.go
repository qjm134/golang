/*
有101个数
其中100个数由1-100之间的自然数组成，没有顺序，每个数只出现一次
还有1个数是重复的
找出这个重复的数

思路
异或（按位）
相同的异或为0
0跟其他异或，按位来看，0跟0异或为0，0跟1异或为1，因此结果还是原数
从0 1之间各种组合异或得出，异或满足交换律
3^2^3^2 = 3^3 ^ 2^2 = 0^0 = 0

可以再生成100个1-100的数，这201个数做一个异或，相同的都变为0
0再跟重复的那个数异或，最后剩下的数就是那个重复的数
*/
package main

import "fmt"

func findRepeat(nums []int) int {
	res := 0
	for i := 0; i < 100; i++ {
		res = res ^ nums[i]
		res = res ^ (i + 1)
	}
	res = res ^ nums[100]
	return res
}

func main() {
	nums := make([]int, 101, 101)
	for i := 0; i < 100; i++ {
		nums[i] = i + 1
	}
	nums[100] = 66

	fmt.Println(findRepeat(nums))
}
