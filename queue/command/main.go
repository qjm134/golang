/*
双端队列，可以往队头和队尾放数字，只能从队头出数字。
放入1，2，3，4，5
head add 1
tail add 2
remove
出必须按1，2..的顺序出，如果不是这个顺序，需要调整，求最少的调整次数
*/
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	n := 0
	fmt.Scan(&n)
	data := make([]int, 0, n)
	s := ""
	removeCount := 0
	adjustCount := 0
	for i := 0; i < 2*n; i++ {
		fmt.Scan(&s)
		cmd, val := parse(s)
		switch cmd {
		case head:
			newData := []int{val}
			data = append(newData, data...)

		case tail:
			data = append(data, val)

		case remove:
			removeCount++
			if data[0] == removeCount {
				data = data[1:]
			} else {
				adjustCount++
				sort.Sort(sort.IntSlice(data))
			}
		}
	}
	fmt.Println(adjustCount)
}

const (
	ini    = 0
	head   = 1
	tail   = 2
	remove = 3
)

func parse(s string) (int, int) {
	if strings.Contains(s, "head") {
		tmp := ""
		fmt.Scan(&tmp)
		val := 0
		fmt.Scan(&val)
		return head, val
	}
	if strings.Contains(s, "tail") {
		tmp := ""
		fmt.Scan(&tmp)
		val := 0
		fmt.Scan(&val)
		return tail, val
	}
	if strings.Contains(s, "remove") {
		return remove, 0
	}

	return ini, 0
}
