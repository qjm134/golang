/*
0 1 0 1 1
1 0 1 0 1
0 1 0 1 1

找有几个岛屿
1是陆地，0是海水
上下左右是可达的
*/

package main

func find(a [][]int, i, j int) {
	if i < 0 || i >= len(a) || j < 0 || j > len(a[i]) {
		return
	}

	if a[i][j] == 1 {
		a[i][j] = 2
		for m := -1; m < 2; m = m + 2 {
			if a[i+m][j] == 1 {
				find(a, i+m, j)
			}
		}
		for n := -1; n < 2; n = n + 2 {
			if a[i][j+n] == 1 {
				find(a, i, j+n)
			}
		}
	}
}
