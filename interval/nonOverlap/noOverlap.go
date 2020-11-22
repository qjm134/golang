/*
435. Non-overlapping Intervals
Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.



Example 1:

Input: [[1,2],[2,3],[3,4],[1,3]]
Output: 1
Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
Example 2:

Input: [[1,2],[1,2],[1,2]]
Output: 2
Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
Example 3:

Input: [[1,2],[2,3]]
Output: 0
Explanation: You don't need to remove any of the intervals since they're already non-overlapping.


Note:

You may assume the interval's end point is always bigger than its start point.
Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.

思路：
两两考虑，有3种情况
前一个右节点小于等于后一个左节点，不重合；
前一个右节点大于后一个左节点，小于等于后一个右节点，重合，移除后一个，可能会同时解除跟再后一个的重合；
前一个右节点大于后一个右节点，重合，移除前一个，可以释放更多空间，当前节点不变，继续跟再后一节点比较。



按左边的起点排序
从左到右遍历每个区间
右边点向右移动，可能重合后面一个区间，也可能重合后面>=2个区间
a.只重合后面1个区间的，有2种移除方式，移除当前区间，或者移除后面区间；如果后面区间也有跟其后面区间重合的，那移除后面区间就能一举2得，移前面的没有任何优势
b.重合后面>=2个区间的，当前区间肯定要移除，因为移除一个，就能释放>=2个，否则要消除重合得移>=2个

Input: [[1,2],[2,3],[3,4],[1,3]]
Output: 1
排序：
[1,2],[2,3],[3,4]
[1,      3]
[1,3] 3>2，有重合，移除
[1,2] 2<=2，没有重合的，不管
[2,3]、[3,4]都没有跟后面重合的，不管

b有问题，比如[1,5],[2,8],[3,7],[6,9]，按b则移除[1,5],[2,8],[6,9]
实际，移[2,8],[3,7]即可
*/
package nonOverlap

func eraseOverlap(intervals [][]int) int {
	quickSort(intervals, 0, len(intervals)-1)
	eraseNum := 0
	present := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[present][1] > intervals[i][1] {
			eraseNum++
			present = i
		} else if (intervals[present][1] > intervals[i][0]) && (intervals[present][1] <= intervals[i][1]) {
			eraseNum++
		} else {
			present = i
		}
	}
	return eraseNum
}

func quickSort(intervals [][]int, start, end int) {
	if start > end {
		return
	}
	mid := sortOnce(intervals, start, end)
	quickSort(intervals, 0, mid-1)
	quickSort(intervals, mid+1, end)
}

func sortOnce(intervals [][]int, start, end int) int {
	tmp := intervals[end]
	for start < end {
		for ; start < end; start++ {
			if intervals[start][0] > tmp[0] {
				intervals[end] = intervals[start]
				end--
				break
			}
		}

		for ; start < end; end-- {
			if intervals[end][0] < tmp[0] {
				intervals[start] = intervals[end]
				start++
				break
			}
		}
	}
	intervals[start] = tmp
	return start
}
