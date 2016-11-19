package main

import (
	"fmt"
)

func partition(a []int, s int, e int) int {
	temp := a[s]
	for e > s {
		for e > s && a[e] >= temp {
			e--
		}

		a[s] = a[e]

		for e > s && a[s] <= temp {
			s++
		}

		a[e] = a[s]
	}
	a[s] = temp
	return s
}

func fast_sort(a []int, s int, e int) {
	var i = 0
	if s < e {
		i = partition(a, s, e)
		fast_sort(a, s, i-1)
		fast_sort(a, i+1, e)
	}
}

func main() {
	a := []int{0, 3, 2, 5, 4, 7, 8}
	fmt.Println("before sort:", a)
	fast_sort(a, 0, len(a)-1)
	fmt.Println("after sort:", a)
}
