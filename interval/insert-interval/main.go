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
3. 未涉及区间，即比newInterval.start小的区间，和比newInterval.end大的区间，原样返回。
*/
package main

func main() {

}

func insert() {

}
