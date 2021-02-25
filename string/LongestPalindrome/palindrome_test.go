package LongestPalindrome

import (
	"fmt"
	"testing"
)

type Example struct {
	arg    string
	expect string
}

func TestLongestPalindrome(t *testing.T) {
	exms := []Example{
		{"babad", "bab"},
		{"ac", "a"},
	}

	for _, exm := range exms {
		//res := LongestPalindrome(exm.arg)
		res := longestPal(exm.arg)
		if res != exm.expect {
			fmt.Printf("fail, result: %+v, expect: %+v, arg: %v", res, exm.expect, exm.arg)
		}
	}
}
