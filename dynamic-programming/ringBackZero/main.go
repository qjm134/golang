/*
圆环上有10个点，编号为0~9。从0出发，每次可以逆时针和顺时针走一步，问走n步回到0共有多少种走法。

输入: 2
输出: 2
解释：有2种方案。分别是0->1->0和0->9->0

思路：
1.
走到1又返回到0，走2步可到；走到2又返回，走4步；走到3又返回，走6步...
奇数步不能回到0
走2步，还可以走到9，再回到0，所以2步有2种走法
4步，可以走0-1-0-9-0
6步，0-1-0-9-8-9-0


2.
走n到0 = 走n-1步到1 + 走n-1到9
f(n,i) = f(n-1, (i+1)%length) + f(n-1, (i-1+length)%length)
f(n,0) = f(n-1, 1) + f(n-1, 9)
f(2,0) = f(1,1) + f(1,9) = 1+1
f(4,0) = f(3,1) + f(3,9) = f(2,2)+f(2,0) + f(2,0)+f(2,8) = 1+2 + 2+1 = 6
f(3,0) = 0
f(6,0) = f(5,1) + f(5,9) = f(4,2)+f(4,0) + f(4,0)+f(4,8) = f(3,3)+f(3,1)+6 + 6+f(3,9)+f(3,7) = 1+3+6 + 6+3+1

f(i,i)=1
f(10-i,i)=1
f(奇数,0)=0

用2维数组，当前层等于上一层，左右2边值相加
所以初始化第一层，f(1,i)或f(0,i)，后面任意层f(n,i)都能算出来
  0 1 2 3 4 5 6 7 8 9
0 1 0 0 0 0 0 0 0 0 0
1   1               1
2 2   1           1
3   3   1       1   3
4
5
6
7
8
9
*/

package main

import "fmt"

func ways(num, n int) int {
	return toPoint(num, n, 0)
}

func toPoint(num, n, i int) int {
	if n < 0 || i < 0 {
		return 0
	}
	if n == 0 {
		if i == 0 {
			return 1
		} else {
			return 0
		}
	}
	return toPoint(num, n-1, (i-1+num)%num) + toPoint(num, n-1, (i+1)%num)
}

//n 代表点的个数,k代表bushu
func GetSteps(n int, k int) int {
	if n == 0 {
		return 1
	}
	if n == 2 {
		if n%2 == 0 {
			return 1
		} else {
			return 0
		}
	}
	var dp [100][100]int

	dp[0][0] = 1
	for i := 1; i < n; i++ {
		dp[0][i] = 0
	}
	for j := 1; j <= k; j++ {
		for i := 0; i < n; i++ {
			dp[j][i] = dp[j-1][(i-1+n)%n] + dp[j-1][(i+1)%n]
		}
	}
	return dp[k][0]
}

func main() {
	n := 10
	k := 8
	fmt.Println(GetSteps(n, k))
	fmt.Println(ways(n, k))
}
