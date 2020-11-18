package longestIncreaseSubsequence

import (
	"fmt"
	"testing"
)

type Example struct {
	args   []int
	expect int
}

func TestLengthOfLIS(t *testing.T) {
	exms := []Example{
		{[]int{1}, 1},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{4, 10, 4, 3, 8, 9}, 3},
	}

	for _, exm := range exms {
		res := LengthOfLIS(exm.args)
		if res != exm.expect {
			fmt.Printf("failed, expect: %+v, result: %+v, arg: %+v\n", exm.expect, res, exm.args)
		}

	}
}
