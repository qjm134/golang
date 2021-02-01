/*
415. 字符串相加
给定两个字符串形式的非负整数 num1 和num2 ，计算它们的和。

提示：

num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字 0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库， 也不能直接将输入的字符串转换为整数形式
*/

package main

import (
	"fmt"
	"strconv"
)

func add(s1, s2 string) string {
	s := ""
	flag := 0
	for i, j := len(s1)-1, len(s2)-1; i >= 0 || j >= 0 || flag > 0; i, j = i-1, j-1 {
		x, y := 0, 0
		if i >= 0 {
			x = int(s1[i] - '0')
		}
		if j >= 0 {
			y = int(s2[j] - '0')
		}
		tmp := x + y + flag
		s = strconv.Itoa(tmp%10) + s
		flag = tmp / 10
	}
	return s
}

func main() {
	s1 := "9943"
	s2 := "71"
	fmt.Println(add(s1, s2))
}
