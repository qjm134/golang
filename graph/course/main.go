/*
207. 课程表
你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 1：
输入：numCourses = 2, prerequisites = [[1,0]]
输出：true
解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。

示例 2：
输入：numCourses = 3, prerequisites = [[1,0],[1,2],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。

提示：
1 <= numCourses <= 105
0 <= prerequisites.length <= 5000
prerequisites[i].length == 2
0 <= ai, bi < numCourses
prerequisites[i] 中的所有课程对 互不相同


思路：
找环
因课程0-nums编号，课程号即数组的索引，所以可用二维数组保存所有课程的先修课程，则先修课程形成多叉链路
每个节点，有3种状态
0：该节点未访问过
2：该节点的所有分支都已检查完毕，没有环
1：该节点正在访问中，不是所有分支都检查完了

遍历的时候，需要标记节点状态(map) + 返回结果

进到该节点，标记为访问中1

如果访问到未访问的节点，则递归访问
如果访问的节点已检查完毕了，则将是否有环的结果直接返回，不重复检查
如果访问到未完成的节点，说明有环

所有分叉都遍历完了，退出前标记该节点为2

比如：
0 -> 1 -> 3
0 <- 1

/*
以下代码只适用于一门课程只需先修一门课程，不适合于一门课程需要先修多门课程
所有有必修前置课的课程，用map保存其必修前置课
从一门课开始，访问其链式必修前课，如果形成一个环，说明无法完成课程
判断环：所有遍历过的点记录其入口课程，如果遍历到某个点，其入口课程是入口课程，则出现环了
比如：
0 -> 1 -> 0
2 -> 1
*/
package main

import "fmt"

const (
	waitCheck        = 0 // 未检查
	checking         = 1 // 部分分支检查中
	allBranchChecked = 2 // 所有分支检查完毕
)

var coursePres [][]int
var statusMap map[int]int
var hasRing bool

func canFinish(numCoursers int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	coursePres = make([][]int, numCoursers)
	statusMap = make(map[int]int)

	for _, req := range prerequisites {
		if len(req) != 2 {
			return true
		}
		coursePres[req[0]] = append(coursePres[req[0]], req[1])
	}

	for course, _ := range coursePres {
		checkCourse(course)
	}

	return !hasRing
}

func checkCourse(course int) {
	status := statusMap[course]
	switch status {
	case waitCheck:
		statusMap[course] = checking
		for _, pre := range coursePres[course] {
			checkCourse(pre)
			if hasRing == true {
				return
			}
		}
		statusMap[course] = 2

	case checking:
		hasRing = true
		return

	case allBranchChecked:
		return
	}
}

func main() {
	numCourses := 2
	prerequisites := [][]int{{0, 1}}
	//prerequisites := [][]int{{0, 1}, {1, 2}}
	//prerequisites := [][]int{{0, 1}, {1, 0}}

	//fmt.Println(canFinishOnly(numCourses, prerequisites))
	fmt.Println(canFinish(numCourses, prerequisites))

}

type courseInfo struct {
	pre        int
	visitStart int
}

func canFinishOnly(numCourses int, prerequisites [][]int) bool {
	preMap := make(map[int]*courseInfo)
	for _, req := range prerequisites {
		co := &courseInfo{pre: req[1], visitStart: -1}
		preMap[req[0]] = co
	}

	for k, v := range preMap {
		v.visitStart = k

		pre := v.pre
		preCourse, ok := preMap[pre]
		for ok {
			if preCourse.visitStart != k {
				preCourse.visitStart = k
				preCourse, ok = preMap[preCourse.pre]
			} else {
				return false
			}
		}
	}

	return true
}
