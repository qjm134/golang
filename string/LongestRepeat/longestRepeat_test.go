package LongestRepeat

import (
	"fmt"
	"testing"
)

type Example struct {
	args   Arg
	expect int
}

type Arg struct {
	str string
	k   int
}

func TestFindLongestRepeat(t *testing.T) {
	testExms := []Example{
		{Arg{"ABAB", 2}, 4},
		{Arg{"AABABBA", 1}, 4},
	}
	for _, testExm := range testExms {
		res := FindLongestRepeat(testExm.args.str, testExm.args.k)
		if res != testExm.expect {
			fmt.Printf("fail, args: %+v, expect: %+v, result: %+v", testExm.args, testExm.expect, res)
		}
	}
}

func TestCharacterReplacement(t *testing.T) {
	testExms := []Example{
		{Arg{"AABCDB", 1}, 3},
		{Arg{"ABAB", 2}, 4},
		{Arg{"AABABBA", 1}, 4},
	}
	for _, exm := range testExms {
		res := CharacterReplacement(exm.args.str, exm.args.k)
		if res != exm.expect {
			fmt.Printf("fail, args: %+v, expect: %+v, result: %+v", exm.args, exm.expect, res)
		}
	}
}
