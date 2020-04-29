/*
57. Insert Interval
https://leetcode.com/problems/insert-interval/

Difficulty:
Hard

Description
Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).

You may assume that the intervals were initially sorted according to their start times.

Example 1:

Input: intervals = [[1,3],[6,9]], newInterval = [2,5]
Output: [[1,5],[6,9]]
Example 2:

Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
Output: [[1,2],[3,10],[12,16]]
Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].
*/

/*
1. 从头往后，找到比newInterval.start大的数。
	此时start所在的位置，有2种：区间里、区间外。
	此时需要对涉及区间的intervals.start和newInterval.start进行二选一，作为output.start。
	第一，区间里，即比newInterval.start大的为interval.end，则涉及区间的intervals.start为output.start，newInterval.start抛弃。
	第二，区间外，即比newInterval.start大的为interval.start，则newInterval.start为为output.start，涉及区间的intervals.start抛弃。
2. 从尾向头，找到比newInterval.end小的数。
	此时end所在的位置，有2种：区间里、区间外。
	第一，区间里，即比newInterval.end小的那个数为涉及区间的start，则涉及区间的intervals.end为output.end。
	第二，区间外，即比newInterval.end小的那个数为涉及区间的end，则newInterval.end为为output.end。

	(以上2点总结：保外舍内。区间里，取区间；区间外，取newInterval)

3. 未涉及区间，即比newInterval.start小的区间，和比newInterval.end大的区间，原样返回。
*/
package main

import "fmt"

func main() {
	intervals := [][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}
	newInterval := []int{4, 8}
	fmt.Println(insert(intervals, newInterval))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(newInterval) < 2 {
		return intervals
	}
	if len(intervals) <= 0 {
		return [][]int{newInterval}
	}
	output := make([][]int, 0, len(intervals))
	start := newInterval[0]
	end := newInterval[1]
	tmpInterval := make([]int, 2)

	//从头往后
	for _, interval := range intervals {
		if len(interval) < 2 {
			continue
		}

		if interval[0] < start {
			if interval[1] < start {
				output = append(output, interval)
			} else {
				//区间里
				tmpInterval[0] = interval[0]
				break
			}
		} else {
			//区间外
			tmpInterval[0] = start
			break
		}
	}

	//所有区间都比newInterval小
	if len(intervals[0]) == 2 {
		if tmpInterval[0] == 0 && (start > intervals[0][1]) {
			tmpInterval[0] = start
		}
	} else {
		return [][]int{newInterval}
	}

	//从尾向头
	for j := len(intervals) - 1; j >= 0; j-- {
		interval := intervals[j]
		if interval[1] > end {
			if interval[0] <= end {
				//区间里
				tmpInterval[1] = interval[1]
				output = append(output, tmpInterval)
				output = append(output, intervals[j+1:]...)
				return output
			}
		} else {
			//区间外
			tmpInterval[1] = end
			output = append(output, tmpInterval)
			output = append(output, intervals[j+1:]...)
			return output
		}
	}

	if tmpInterval[1] == 0 && intervals[len(intervals)-1][1] != 0 {
		tmpInterval[1] = end
		output = append(output, tmpInterval)
		output = append(output, intervals...)
		return output
	}

	return output
}
