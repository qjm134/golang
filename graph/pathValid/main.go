/*
1971. 寻找图中是否存在路径
有一个具有 n 个顶点的 双向 图，其中每个顶点标记从 0 到 n - 1（包含 0 和 n - 1）。图中的边用一个二维整数数组 edges 表示，其中 edges[i] = [ui, vi] 表示顶点 ui 和顶点 vi 之间的双向边。 每个顶点对由 最多一条 边连接，并且没有顶点存在与自身相连的边。
请你确定是否存在从顶点 source 开始，到顶点 destination 结束的 有效路径 。
给你数组 edges 和整数 n、source 和 destination，如果从 source 到 destination 存在 有效路径 ，则返回 true，否则返回 false 。

输入：n = 3, edges = [[0,1],[1,2],[2,0]], source = 0, destination = 2
输出：true
解释：存在由顶点 0 到顶点 2 的路径:
- 0 → 1 → 2
- 0 → 2

输入：n = 6, edges = [[0,1],[0,2],[3,5],[5,4],[4,3]], source = 0, destination = 5
输出：false
解释：不存在由顶点 0 到顶点 5 的路径.


思路一：
并查集
map存集合，集合中的元素是点
一个点在map中，则这个边上的2点都属于这个集合，都存入map。
边上的2点都不在已有的map中，则新建一个map
用slice来存所有的集合

思路二：
遍历source开始的每条路径，遍历路径中的每个点，如果有destination则连通，否则不连通
2重循环，外层是以source开始的每条路径。内层是遍历路径中的每个点，可以用递归，二维数组中找一个相等的点，递归边的另一个点
遍历的时候，每次都从头开始遍历二维数组中的每个点，遍历过的标记，遇到直接跳过
*/

package main

import "fmt"

func main() {
	edges := [][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}}
	source := 5
	destination := 9
	fmt.Println(isConnect(len(edges), edges, source, destination))
}

func isConnect(n int, edges [][]int, source int, destination int) bool {
	if len(edges) == 0 && source == 0 && destination == 0 {
		return true
	}

	for i := 0; i < len(edges); i++ {
		for j := 0; j < 2; j++ {
			if edges[i][j] == source {
				found := make([]bool, len(edges))
				found[i] = true
				res := false
				if j == 0 {
					res = find(edges, edges[i][1], destination, found)
				} else {
					res = find(edges, edges[i][0], destination, found)
				}
				if res {
					return true
				}
			}
		}
	}
	return false
}

func find(edges [][]int, next int, destination int, found []bool) bool {
	if next == destination {
		return true
	}

	for i := 0; i < len(edges); i++ {
		if found[i] == true {
			continue
		}

		for j := 0; j < 2; j++ {
			if edges[i][j] == next {
				found[i] = true
				res := false
				if j == 0 {
					res = find(edges, edges[i][1], destination, found)
				} else {
					res = find(edges, edges[i][0], destination, found)
				}
				if res {
					return true
				}
			}
		}
	}

	return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	if len(edges) <= 0 {
		return true
	}
	setList := make([]map[int]struct{}, 0, len(edges))
	sourceIdx := 0
	destinationIdx := -1

	for i := 0; i < len(edges); i++ {
		if len(edges[i]) < 2 {
			return false
		}

		if sourceIdx == destinationIdx {
			return true
		}

		flag := 0
		for k := 0; k < len(setList); k++ {
			m := setList[k]
			left := edges[i][0]
			right := edges[i][1]
			if _, ok := m[left]; ok {
				m[right] = struct{}{}

				if right == source {
					sourceIdx = k
				}
				if right == destination {
					destinationIdx = k
				}

				flag = 1
				break
			}
			if _, ok := m[right]; ok {
				m[left] = struct{}{}

				if left == source {
					sourceIdx = k
				}
				if left == destination {
					destinationIdx = k
				}

				flag = 1
				break
			}
		}

		if flag == 0 {
			left := edges[i][0]
			right := edges[i][1]

			m := make(map[int]struct{})
			m[left] = struct{}{}
			m[right] = struct{}{}

			setList = append(setList, m)

			k := len(setList) - 1
			if left == source || right == source {
				sourceIdx = k
			}
			if left == destination || right == destination {
				destinationIdx = k
			}
		}
	}

	if sourceIdx == destinationIdx {
		return true
	}

	return false
}
