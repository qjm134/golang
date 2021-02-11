/*
215. Kth Largest Element in an Array
Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

Input: [3,2,1,5,6,4] and k = 2
Output: 5
Example 2:

Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
Note:
You may assume k is always valid, 1 ≤ k ≤ array's length.
*/
package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	return findK(nums, k, 0, len(nums)-1)
}

func findK(a []int, k int, start, end int) int {
	if k <= 0 {
		return 0
	}
	p := findPosition(a, start, end)
	if p > k-1 {
		return findK(a, k, 0, p-1)
	} else if p < k-1 {
		return findK(a, k, p+1, end)
	} else {
		return a[p]
	}
}

func findPosition(a []int, start, end int) int {
	if start < 0 || end >= len(a) {
		return -1
	}
	tmp := a[end]
	for start < end {
		for start < end {
			if a[start] >= tmp {
				start++
			} else {
				a[end] = a[start]
				end--
				break
			}
		}

		for start < end {
			if a[end] <= tmp {
				end--
			} else {
				a[start] = a[end]
				start++
				break
			}
		}

	}
	a[start] = tmp
	return start
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	fmt.Println(findKthLargest(nums, k))
}
