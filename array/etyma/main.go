/*
词根：在词根后面添加一些其他词组或另一个较长的单词-我们称这个词为继承词。
例如，词根am，跟随着单词other，可以形成新的单词another。
现在，给定一个由许多词根组成的词典dictionary和一个用空格分割单词形成的句子sentence.
你需要将句子中的所有继承词用词根替换掉。如果继承词有许多可以形成它的词根，则用最短的词根替换它。
你需要输出替换之后的句子。
示例1
输入：dictionary=["cat", "bat", "rat"], sentence="the cattle was rattled by the battery"
输出："the cat was rat by the bat"

示例2
输入：dictionary=["a", "b", "c"], sentence="aadsfasf absba bbab cadsfafs"
输出："a a b c"
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	dic := []string{"cat", "bat", "rat"}
	str := "the cattle was rattled by the battery"
	//dic := []string{"a", "b", "c"}
	//str := "aadsfas absbs bbab cadsfafs"
	fmt.Println(replaceSen(str, dic))
}

func replaceSen(str string, dic []string) string {
	ret := ""

	ones := strings.Split(str, " ")
	for i := 0; i < len(ones); i++ {
		rep := replace(ones[i], dic)
		if rep != "" {
			ret = ret + rep + " "
		} else {
			ret = ret + ones[i] + " "
		}
	}

	ret = strings.TrimRight(ret, " ")

	return ret
}

func replace(str string, dic []string) string {
	minLen := 0
	first := true
	rep := ""
	for i := 0; i < len(dic); i++ {
		idx := strings.Index(str, dic[i])
		if idx == 0 {
			if first {
				minLen = len(dic[i])
				rep = dic[i]
				first = false
			}
			if len(dic[i]) < minLen {
				rep = dic[i]
				minLen = len(dic[i])
			}
		}
	}

	return rep
}
