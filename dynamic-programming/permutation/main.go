/*
"abcde"

找出所有全排列

*/

package main

import "fmt"

func quan(s string) {
	mStr := make(map[byte]bool)
	for _, v := range s {
		mStr[byte(v)] = false
	}

	ret := make([]byte, 0, len(s))
	one(mStr, ret)
}

func one(m map[byte]bool, ret []byte) {
	if len(m) == 0 {
		fmt.Println(string(ret))
		return
	}

	for k, _ := range m {
		ret = append(ret, k)
		nextM := make(map[byte]bool)
		for kk, v := range m {
			if kk == k {
				continue
			}
			nextM[kk] = v
		}
		one(nextM, ret)
		ret = ret[0:(len(ret) - 1)]
	}
}

func main() {
	s := "abc"
	quan(s)
	/*
		m1 := make(map[int]int)
		m1[1] = 1
		m1[2] = 2

		m2 := make(map[int]int)
		for k, v := range m1 {
			m2[k] = v
		}
		//m2 := m1
		m2[1] = 11
		fmt.Println(m1, m2)

	*/
}
