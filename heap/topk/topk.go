/*
347. Top K Frequent Elements
Medium

3975

237

Add to List

Share
Given a non-empty array of integers, return the k most frequent elements.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
Example 2:

Input: nums = [1], k = 1
Output: [1]
Note:

You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
It's guaranteed that the answer is unique, in other words the set of the top k frequent elements is unique.
You can return the answer in any order.

思路：
1.统计各数字出现的次数
2.出现的次数可能有一样的[1，1，3，3，4]，比如1出现2次，3也出现2次，那么2是一个待排的，频次数组里只要一个2去跟其他频次数排就行，如果2排前k里，输出的时候，并排把1和3也输出
3.转换为一个无重复数组的topk问题，快排一次n，找到一个位置pos，这个位置有3种情况，小于、大于、等于k
  等于就不用再排了，前k个就是需要的；
  大于k，则继续只排前一半就行，省去后一半不用管了
  小于k，则前一半都是，后一半还需再找到前 k-pos 的
*/

package topk

func topFrenquency(nums []int, k int) []int {
	dataToFre := make(map[int]int)
	for _, v := range nums {
		dataToFre[v]++
	}
	freToData := make(map[int][]int)
	fre := make([]int, 0, len(dataToFre))
	for key, v := range dataToFre {
		freToData[v] = append(freToData[v], key)
		if len(freToData[v]) == 1 {
			fre = append(fre, v)
		}
	}

	if k >= len(fre) {
		k = len(fre)
	} else {
		topK(fre, 0, len(fre)-1, k)
	}
	res := make([]int, 0)
	for i := 0; i < k; i++ {
		res = append(res, freToData[fre[i]]...)
	}
	return res
}

func topK(fre []int, start, end, k int) {
	if start < 0 || end >= len(fre) || start >= len(fre) || end < 0 {
		return
	}
	pos := sortOnce(fre, start, end)
	if pos+1 == k {
		return
	} else if pos+1 < k {
		topK(fre, pos+1, end, k-pos-1)
	} else {
		topK(fre, 0, pos-1, k)
	}
}

func sortOnce(a []int, start, end int) int {
	tmp := a[end]
	i := start
	j := end
	for i < j {
		for ; i < j; i++ {
			if a[i] < tmp {
				a[j] = a[i]
				j--
				break
			}
		}
		for ; i < j; j-- {
			if a[j] > tmp {
				a[i] = a[j]
				i++
				break
			}
		}
	}
	a[i] = tmp
	return i
}
