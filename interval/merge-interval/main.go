package main

import (
	"./merge"
	"fmt"
)

func main() {
	interval := [][]int{[]int{1, 4}, []int{0, 4}}
	fmt.Println(merge.Merge(interval))
}
