package numberToString

import (
	"fmt"
	"testing"
)

type Example struct {
	num    int
	result int
}

func TestTranslateNum(t *testing.T) {
	exms := []Example{
		{25, 2},
		{26, 1},
		{506, 1},
		{0, 1},
	}
	for _, exm := range exms {
		res := translateNum(exm.num)
		if res != exm.result {
			fmt.Printf("fail, result: %+v, expect: %+v, arg: %+v", res, exm.result, exm.num)
		}
	}
}
