/*
已知 sqrt (2)约等于 1.414，要求不用数学库，求 sqrt (2)精确到小数点后 10

思路：
二分
约等于1.414，1.41是精确的，最后一位4可能是3，大于5进位来的，也可能是4，舍弃了后面的
总之在 1.410 和 1.419之间，取中位数 1.415，1.414的平方如果大于2，则在1.410到1.415之间
如果，两边夹数相差1，则会多出新的一位，比如1.4和1.5二分，则出现新的位1.45
最后显示，只保留10位
*/

package main

import "fmt"

func binarySqrt(x, min, max float64) float64 {
	mid := (min + max) / 2

	if max-min <= 0.0000000001 {
		mid = float64(int(mid*10000000000)) / 10000000000
		return mid
	}

	tmp := mid * mid
	if tmp > x {
		max = mid
		return binarySqrt(x, min, max)
	} else if tmp < x {
		min = mid
		return binarySqrt(x, min, max)
	} else {
		return mid
	}
}

func main() {
	x := 2.0
	min := 1.410
	max := 1.419
	fmt.Println(binarySqrt(x, min, max))
}
