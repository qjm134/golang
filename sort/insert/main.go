package main

import (
	"fmt"
)

func insert_sort(a []int) {
	var temp int
	for i := 1; i < len(a); i++ {
		temp = a[i]
		for j := i; j > 0; j-- {
			if temp < a[j-1] {
				a[j] = a[j-1]
				a[j-1] = temp //a[0]
			} else {
				a[j] = temp
				break
			}
		}
	}
}

func main() {
	a := []int{2, 1, 0, 4, 6, 3, 7, 5, 9}
	fmt.Println("before sort:", a)
	insert_sort(a)
	fmt.Println("after sort:", a)
}
