package main

import (
	"fmt"
	"testing"
)

type testStruct struct {
	prices     []int
	promotions [][]int
	discount   []int
	order      []int
	expect     int
}

func TestCount(t *testing.T) {
	tc := []testStruct{
		{[]int{10, 5, 8, 8, 6, 3},
			[][]int{{0, 7}, {4, 10}, {2, 6, 1, 9}, {5, 2}},
			[]int{4, 3, 5, 1},
			[]int{2, 17, 3, 10, 1, 27, 5, 2, 4, 9},
			400,
		},
		{
			[]int{10, 10, 10},
			[][]int{{1, 2}, {0, 5, 2, 3}},
			[]int{10, 10},
			[]int{0, 10},
			100,
		},
		{
			[]int{10, 4},
			[][]int{},
			[]int{},
			[]int{0, 2, 1, 3},
			32,
		},
	}

	for _, tOne := range tc {
		res := count(tOne.prices, tOne.promotions, tOne.discount, tOne.order)
		if res != tOne.expect {
			fmt.Println("not pass")
		} else {
			fmt.Println("pass")
		}
	}
}
