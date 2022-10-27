/*
字母先排序，字母相等的用数字排序
字母按位比较，位缺失的，为小，比如："B" < "BA"

思路：
有不同的字母，不需要把字母比完，就得到结果了
需要把字母比完的，那就是比完的那些字母都一样。剩下没比完的2种情况，1长短不一，则短的就小；2是长度一样，都到头了，剩下数字了
*/
package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	str := []string{"D1", "A20", "AA1", "A1", "B3"}
	sort.Sort(strNum(str))
	fmt.Println(str)
}

type strNum []string

func (sn strNum) Len() int {
	return len(sn)
}

func (sn strNum) Less(i, j int) bool {
	for k := 0; k < len(sn[i]); k++ {
		si := sn[i][k]
		sj := sn[j][k]

		if si >= 'A' && si <= 'Z' && sj >= 'A' && sj <= 'Z' {
			if si < sj {
				return true
			}
			if si > sj {
				return false
			}
		}

		// 前面的都相同，一个结束了一个还有
		// 数字0-9 ASCII 48-57
		// 大写字母 65-90
		// 小写字母 97-122
		// si字母结束了，sj还有字母
		if si < 'A' && sj >= 'A' {
			return true
		}
		if si >= 'A' && sj < 'A' {
			return false
		}

		// 2个前面的字母一样，都到数字了
		if si < 'A' && sj < 'A' {
			ni, _ := strconv.Atoi(sn[i][k:])
			nj, _ := strconv.Atoi(sn[j][k:])
			if ni <= nj {
				return true
			} else {
				return false
			}
		}
	}

	return false
}

func (sn strNum) Swap(i, j int) {
	sn[i], sn[j] = sn[j], sn[i]
}
