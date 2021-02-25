package firstLargeInright

import (
	"fmt"
	"testing"
)

type sample struct {
	arg    []int
	expect []int
}

func TestFindFirstLarge(t *testing.T) {
	samples := []sample{
		{[]int{4, 3, 4, 5, 1, 2}, []int{5, 4, 5, -1, 2, -1}},
	}

	for _, s := range samples {
		res := findFirstLarge(s.arg)
		for idx, v := range res {
			if v != s.expect[idx] {
				fmt.Printf("fail, result: %v, expect: %v, in: %v\n", res, s.expect, s.arg)
				return
			}
		}
	}
}
