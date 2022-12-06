/*
Write a program that outputs all possibilities to put +，-，or nothing between the numbers 1,2,…,9 (in this order) such that the result is 100.
1+2+3-4+5+6+78+9 = 100
123+4-5+67-89 = 100
...

思路
找到所有的组合表达式，从所有的表达式中选出和为100的
对于1-9的每个数字，都有3种可能
pre + curNum
pre - curNum
pre curNum
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(find())
}

func find() []string {
	return merge("2", "1", "10", 100)
}

func merge(num string, pre string, over string, target int) []string {
	if num == over {
		//fmt.Println(pre)
		if add(pre) == target {
			return []string{pre}
		} else {
			return nil
		}
	}

	curNum, _ := strconv.Atoi(num)
	nextNum := strconv.Itoa(curNum + 1)

	curPlus := pre + "+" + num
	resPlus := merge(nextNum, curPlus, over, target)

	curMinus := pre + "-" + num
	resMinus := merge(nextNum, curMinus, over, target)

	curDirect := pre + num
	resDirect := merge(nextNum, curDirect, over, target)

	res := make([]string, 0)
	res = append(res, resPlus...)
	res = append(res, resMinus...)
	res = append(res, resDirect...)
	return res
}

func add(s string) int {
	sum := 0
	preNum := make([]byte, 0)
	preSymbol := uint8('n')
	for i := 0; i < len(s); i++ {
		if s[i] == '+' || s[i] == '-' {
			num, _ := strconv.Atoi(string(preNum))
			if preSymbol == '+' || preSymbol == 'n' {
				sum = sum + num
			}
			if preSymbol == '-' {
				sum = sum - num
			}
			preSymbol = s[i]
			preNum = []byte{}
			continue
		}
		preNum = append(preNum, s[i])
	}

	num, _ := strconv.Atoi(string(preNum))
	if preSymbol == '+' || preSymbol == 'n' {
		sum = sum + num
	}
	if preSymbol == '-' {
		sum = sum - num
	}

	return sum
}
