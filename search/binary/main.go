package main

import (
	"fmt"
)

func binary_search(a []int, f int) {
	var start, end = 0, len(a) - 1
	var mid = (start + end) / 2
	for start <= end {
		if f == a[mid] {
			fmt.Println("Finded, position: %d", mid)
			return
		} else if f < a[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
		mid = (start + end) / 2
	}

	fmt.Println("Not find!")
}

func main() {
	a := []int{0, 2, 3, 4, 5, 7, 8, 9}
	binary_search(a, 1)
}
