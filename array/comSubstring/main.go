/*
找2个字符串的最长公共子串
first := "hzacdfg"
second := "acbacdfe"
最长公共子串："acdf"
*/
package main

import "fmt"

func main() {
	first := "hzacdfg"
	second := "acbacdfe"
	fmt.Println(find(first, second))
}

func find(first, second string) string {
	maxLen := 0
	maxStr := ""
	for i := 0; i < len(first); i++ {
		curLen := 0
		k := i
		for j := 0; j < len(second) && k < len(first); j++ {
			if first[k] == second[j] {
				curLen++
				k++
			} else {
				if curLen > maxLen {
					maxLen = curLen
					maxStr = first[i:k]
				}
				curLen = 0
				k = i
			}
		}
		if curLen > maxLen {
			maxLen = curLen
			maxStr = first[i:k]
		}
	}
	return maxStr
}
