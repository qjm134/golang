package nonOverlap

import (
	"fmt"
	"testing"
)

type Exm struct {
	arg    [][]int
	expect int
}

func TestEraseOverlap(t *testing.T) {
	exms := []Exm{
		{[][]int{{0, 2}, {1, 3}, {1, 3}, {2, 4}, {3, 5}, {3, 5}, {4, 6}}, 4},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 100}, {2, 12}, {1, 11}, {11, 22}}, 2},
	}
	for _, exm := range exms {
		res := eraseOverlap(exm.arg)
		if res != exm.expect {
			fmt.Printf("fail, expect: %v, result: %v, arg: %v\n", exm.expect, res, exm.arg)
		}
	}
}
