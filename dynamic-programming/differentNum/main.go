/*
一个数组，每次随机读取一个位置的值，每个元素只能返回一次

思路：
在0-len(nums)中随机生成一个数，如果之前访问过，还需要不断重试，如果每次都产生相同的数，那时间就复杂了
可以把所有位置都放一个集合，每次从集合里随机取一个，用了将其从集合中删除
*/

package main

import "fmt"

type randNum struct {
	nums []int
	m    map[int]struct{}
}

func NewRandNum(numsIn []int) *randNum {
	newM := make(map[int]struct{})
	for i := 0; i < len(numsIn); i++ {
		newM[i] = struct{}{}
	}

	return &randNum{numsIn, newM}
}

func (r *randNum) get() int {
	for k, _ := range r.m {
		delete(r.m, k)
		return r.nums[k]
	}
	return 0
}

func main() {
	nums := []int{3, 2, 2, 5}
	r := NewRandNum(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Println(r.get())
	}
}
