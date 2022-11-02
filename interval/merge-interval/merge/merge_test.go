package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Example struct {
	args   [][]int
	expect [][]int
}

func TestMerge(t *testing.T) {
	examples := []Example{
		{args: [][]int{{1, 4}, {0, 4}},
			expect: [][]int{{0, 4}},
		},
		{args: [][]int{{2, 5}, {3, 7}},
			expect: [][]int{{2, 7}},
		},
	}

	for _, eg := range examples {
		result := Merge(eg.args)
		if !reflect.DeepEqual(eg.expect, result) {
			fmt.Printf("fail, expect: %+v, result: %+v", eg.expect, result)
		}
	}
}
