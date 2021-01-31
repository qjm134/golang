/*
归并排序
*/

package main

import "fmt"

func mergeTwo(a, b []int) []int {
	lc := len(a) + len(b)
	c := make([]int, 0, lc)
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}
	if i == len(a) {
		c = append(c, b[j:]...)
	}
	if j == len(b) {
		c = append(c, a[i:]...)
	}
	return c
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	a := mergeSort(arr[:mid])
	b := mergeSort(arr[mid:])
	return mergeTwo(a, b)
}

func main() {
	arr := []int{3, 2, 5, 4, 1}
	fmt.Println(mergeSort(arr))
}
