/*
56. Merge Intervals
https://leetcode.com/problems/merge-intervals/

Description
Given a collection of intervals, merge all overlapping intervals.

Example 1:

Input: [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
Example 2:

Input: [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
NOTE: input types have been changed on April 15, 2019. Please reset to default code definition to get new method signature.

思路：
1.遍历每一个区间，跟之前已经合并过的结果区间集合比较，从左节点开始，跟当前比较区间，有3种可能，每种可能做相应的处理；然后再比较右节点，在左节点可能的每种情况，又相应有各种情况，枚举所有的可能；时间复杂度n的平方

2.以上方法每个区间去跟已经合并的各区间比较，就是因为不知道这个区间的位置，那可以先排下序，那么后面的区间只要跟已经合并的区间的最后一个比较即可
  a.按左节点排序区间
  b.排序后，左节点只要2种可能，在当前比较区间的中间，或右边
  c.中间的话，取右节点最大的为合并区间的右节点；右边的话，不能合并，直接添加到合并区间集合里

*/

package merge

func Merge(interval [][]int) [][]int {
	sorted := sort(interval, 0, len(interval)-1)
	var res [][]int
	for _, v := range sorted {
		l := len(res)
		if l == 0 {
			res = append(res, v)
		} else {
			last := res[l-1]
			if v[0] <= last[1] {
				if v[1] > last[1] {
					last[1] = v[1]
				}
			} else {
				res = append(res, v)
			}
		}
	}

	return res
}

func sort(interval [][]int, left, right int) [][]int {
	if right <= left {
		return interval
	}
	mid := sortOnce(interval, left, right)
	sort(interval, left, mid-1)
	sort(interval, mid+1, right)
	return interval
}

func sortOnce(interval [][]int, left, right int) int {
	tmp := interval[right]
	i := left
	j := right
	for i < j {
		for ; i < j; i++ {
			if interval[i][0] > tmp[0] {
				interval[j] = interval[i]
				j--
				break
			}
		}
		for ; j > i; j-- {
			if interval[j][0] < tmp[0] {
				interval[i] = interval[j]
				i++
				break
			}
		}
	}
	interval[i] = tmp

	return i
}
