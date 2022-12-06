/*
有一组版本号编码，是一个二维数组，元素由0，1构成
数组的索引号表示版本号
元素1表示，版本i-版本j有一个版本
元素0表示，版本i-版本j没有版本

示例
[[1,1,0],[1,1,1],[0,1,1]]
表示版本0-0有1个版本，0-1有版本，0-2没有版本。  版本0-1
1-0（0-1）有版本，1-1有，1-2有。 版本0-1-2
2-0（0-2）没有，2-1（1-2）有，2-2有。 版本0-1-2
求最长的版本号
*/
package main

func findMax(nums [][]int) int {
	ms := make([]map[int]struct{}, 0)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			if nums[i][j] == 1 {
				k := 0
				for ; k < len(ms); k++ {
					if _, ok := ms[k][i]; ok {
						ms[k][j] = struct{}{}
						break
					}
					if _, ok := ms[k][j]; ok {
						ms[k][i] = struct{}{}
						break
					}
				}
				if k >= len(ms) {
					tmpM := make(map[int]struct{})
					tmpM[i] = struct{}{}
					tmpM[j] = struct{}{}
					ms = append(ms, tmpM)
				}
			}
		}
	}

	maxLen := 0
	for i := 0; i < len(ms); i++ {
		mlen := len(ms[i])
		if mlen > maxLen {
			maxLen = mlen
		}
	}
	return maxLen
}
