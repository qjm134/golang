/*
剑指 Offer 30. 包含min函数的栈
定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

示例:
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.min();   --> 返回 -3.
minStack.pop();
minStack.top();      --> 返回 0.
minStack.min();   --> 返回 -2.

提示：
各函数的调用总次数不超过 20000 次

思路：
出栈，如果是最小值，那需要去前面再找一个最小值，避免这样，可以把之前次小的，次次小的...依次都保存起来
入栈，遇到小的就压入保存栈，栈顶就是最小的，小的依次保存在栈中了
*/
package main

import (
	"errors"
	"fmt"
)

type stack struct {
	data []int
	min  []int
}

func NewStack() *stack {
	return &stack{data: make([]int, 0), min: make([]int, 0)}
}

func (s *stack) push(ele int) {
	s.data = append(s.data, ele)

	if len(s.min) == 0 {
		s.min = append(s.min, ele)
		return
	}

	if ele <= s.min[len(s.min)-1] {
		s.min = append(s.min, ele)
	}
}

func (s *stack) pop() (int, error) {
	if len(s.data) == 0 {
		return 0, errors.New("stack has no data")
	}

	res := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	if res == s.min[len(s.min)-1] {
		s.min = s.min[:len(s.min)-1]
	}

	return res, nil
}

func (s *stack) getMin() (int, error) {
	if len(s.min) == 0 {
		return 0, errors.New("stack has no data")
	}

	return s.min[len(s.min)-1], nil
}

func main() {
	s := NewStack()
	s.push(1)
	fmt.Println(s.getMin())

	s.push(-1)
	fmt.Println(s.getMin())

	s.push(2)
	s.push(3)
	fmt.Println(s.getMin())

	s.pop()
	fmt.Println(s.getMin())

	s.pop()
	s.pop()
	fmt.Println(s.getMin())
}
