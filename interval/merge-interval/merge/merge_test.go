package merge

import (
	"fmt"
	"testing"
)

func TestMerge(t *testing.T) {
	interval := [][]int{[]int{1, 4}, []int{0, 4}}
	fmt.Println(Merge(interval))
}
