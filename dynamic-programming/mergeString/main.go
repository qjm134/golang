/*
有一组字符数组，例如：[a,b,c]，[d,e]，[f,g,h]
将所有数组合并为一个数组，所有字符需要保持在原有数组的顺序，数组间字符顺序任意
找出满足条件的所有合并后的数组

例如
输入：
[a,b]，[d]
输出：
[d,a,b]
[a,d,b]
[a,b,d]


思路：
暴力枚举所有情况
合并的数组的第一位只能从原数组第一个字符中选，因为如果选后面的比如选b，那a就不能入合并数组了。因此遍历所有子数组的第一位，比如选d
第二位，只能从a,e,f中选一个，比如选f
...
递归，每次传入一个Map，遍历map，map中的每一位元素都可以作为合并数组的当前位元素
根据原map生成一个新map，剔除掉原map中当前被选中的元素，加入剔除元素所在数组的下一个元素
map里保存一个数组，2个元素，第一个元素记录当前字符所在数组，第二个记录所在数组的位置


找出所有字符的全排列，找的过程中排除不满足条件的
这样不行，需要保存顺序性，每个子数组，前一个元素入合并数组了，后一个元素才有被选的权限
子数组中的每个元素，都是按顺序进合并数组的
*/

package main

import "fmt"

func selectOne(in [][]string, mCandidate map[string][]int, result []string) {
	if len(mCandidate) == 0 {
		fmt.Println(result)
		return
	}

	for character, _ := range mCandidate {
		result = append(result, character)

		newCan := make(map[string][]int)
		for cha, sP := range mCandidate {
			if cha == character {
				newSp := []int{sP[0], sP[1] + 1}
				if newSp[1] < len(in[newSp[0]]) {
					newCan[in[newSp[0]][newSp[1]]] = newSp
				}
			} else {
				newCan[cha] = sP
			}
		}

		selectOne(in, newCan, result)

		result = result[:len(result)-1]
	}
}

func merge(in [][]string) {
	mCandicate := make(map[string][]int)
	for idx, sCha := range in {
		mCandicate[sCha[0]] = append(mCandicate[sCha[0]], []int{idx, 0}...)
	}

	result := make([]string, 0)
	selectOne(in, mCandicate, result)
}

func main() {
	one := []string{"a", "b", "c"}
	two := []string{"d", "e"}

	in := [][]string{one, two}

	merge(in)
}
