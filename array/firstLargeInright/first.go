/*
给的一个无序正整数数组，找出数组中每个元素右边第一个比它大的元素，没有输出-1
   输入: [4, 3, 4,  5, 1,  2]
   输出: [5, 4, 5, -1, 2, -1]

思路：
遍历
2种情况
1.当前值比前一个大，输出数组相应位置写当前值
2.小于等于，暂时没找着，a.大的值可能在后面某个位置出现，b.也可能没有
  先把当前值所在的位置保存起来，后面比其大的值，会在后续遍历过程中遇到，那时就能填输出数组了
  所以遍历的时候，不止要跟前一个值比较，还要跟之前没找到的值比较
  之前没找到的值，一定是一个递减数列，当前值跟之前没找到的值比较的时候，当遇到一个大的就停止
*/

package firstLargeInright

func findFirstLarge(in []int) []int {
	out := make([]int, len(in), len(in))
	tmp := make([]int, len(in), len(in))
	position := -1

	for i := 1; i < len(in); i++ {
		if in[i] > in[i-1] {
			out[i-1] = in[i]
			for position >= 0 && in[i] > in[tmp[position]] {
				out[tmp[position]] = in[i]
				position--
			}
		} else {
			position++
			tmp[position] = i - 1
		}
	}

	for i := position; i >= 0; i-- {
		out[tmp[i]] = -1
	}
	out[len(in)-1] = -1

	return out
}
