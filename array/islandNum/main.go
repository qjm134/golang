/*
200. 岛屿数量
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。


示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

package main

import "fmt"

func numIslands(grid [][]byte) int {
	islandNO := byte('1')

	for i, row := range grid {
		rowLen := len(row)
		for j, _ := range row {
			if grid[i][j] == '1' {
				islandNO++
				findSameIsland(grid, i, j, islandNO, rowLen)
			}
		}
	}

	return int(islandNO-'0') - 1
}

func findSameIsland(grid [][]byte, i, j int, islandNo byte, rowLen int) {
	grid[i][j] = islandNo

	if j-1 >= 0 {
		if grid[i][j-1] == '1' {
			findSameIsland(grid, i, j-1, islandNo, rowLen)
		}
	}
	if j+1 < rowLen {
		if grid[i][j+1] == '1' {
			findSameIsland(grid, i, j+1, islandNo, rowLen)
		}
	}
	if i-1 >= 0 {
		if grid[i-1][j] == '1' {
			findSameIsland(grid, i-1, j, islandNo, rowLen)
		}
	}
	if i+1 < len(grid) {
		if grid[i+1][j] == '1' {
			findSameIsland(grid, i+1, j, islandNo, rowLen)
		}
	}
}

func main() {
	//grid := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	grid := [][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}
	fmt.Println(numIslands(grid))
}
