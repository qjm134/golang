/*
23. Merge k Sorted Lists
Hard

4225

267

Add to List

Share
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
*/

/*
输入是，每个链表的头节点组成的slice
新建一个，新链表的头节点，新链表移动的指针
遍历所有头节点，开始循环次数是k，遍历的时候，如果有空节点，则将slice最后一个（k-1）头节点放入当前空节点位置，k--
用一个变量记录当前最小的数，
用一个slice记录最小的数对应的头节点在输入slice中的位置
如果有比变量小的，更新这个变量，并清空最小slice(minSlice = minSlice[:0])，将新位置append入最小slice
遍历最小slice，最小头节点链接到新链表，头节点往后移一个
*/
package main

func main() {

}
