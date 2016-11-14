package main

import (
	"fmt"
)

func bubbleSort(a []int, len int) {
	var temp int
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if a[j] > a[j+1] {
				temp = a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
}

func main() {
	a := []int{3, 2, 1, 5, 2, 9, 8}
	fmt.Println("before sort: ", a)
	bubbleSort(a, len(a))
	fmt.Println("after sort: ", a)
}
