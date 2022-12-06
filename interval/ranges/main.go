// Package ranges defines the RangeList type, which is a aggregate of ranges.
// A pair of integers define a range, for example: [1, 5). This range includes integers: 1, 2, 3 and 4.
// A range list is a aggregate of these ranges: [1, 5), [10, 11), [100, 201)

//package ranges
package main

import "fmt"

// rangeBoundarySize is the amount of integers that defines a range.
const (
	rangeBoundarySize = 2
	leftIdx           = 0
	rightIdx          = 1
)

// RangeList is a aggregate of ranges.
type RangeList struct {
	list [][rangeBoundarySize]int
}

// Add a range to the list.
// @param singleRange - Array of two integers that specify beginning and end of range.
func (rl *RangeList) Add(singleRange [rangeBoundarySize]int) {
	leftBoundary := singleRange[leftIdx]
	rightBoundary := singleRange[rightIdx]
	leftPos := rl.findPosition(leftBoundary)
	rightPos := rl.findPosition(rightBoundary)

	newList := make([][rangeBoundarySize]int, 0)
	newRange := [rangeBoundarySize]int{}
	if leftPos < 0 {
		newRange[leftIdx] = leftBoundary
	} else {
		if leftBoundary <= rl.list[leftPos][rightIdx] {
			newList = append(newList, rl.list[:leftPos]...)
			newRange[leftIdx] = rl.list[leftPos][leftIdx]
		} else {
			newList = append(newList, rl.list[:leftPos+1]...)
			newRange[leftIdx] = leftBoundary
		}
	}

	if rightPos < 0 {
		newRange[rightIdx] = rightBoundary
	} else {
		if rightBoundary <= rl.list[rightPos][rightIdx] {
			newRange[rightIdx] = rl.list[rightPos][rightIdx]
		} else {
			newRange[rightIdx] = rightBoundary
		}
	}

	newList = append(newList, newRange)
	if rightPos < len(rl.list)-1 {
		newList = append(newList, rl.list[rightPos+1:]...)
	}

	rl.list = newList
}

// Remove a range from the list.
// @param singleRange - Array of two integers that specify beginning and end of range.
func (rl *RangeList) Remove(singleRange [rangeBoundarySize]int) {
	leftBoundary := singleRange[leftIdx]
	rightBoundary := singleRange[rightIdx]
	leftPos := rl.findPosition(leftBoundary)
	rightPos := rl.findPosition(rightBoundary)

	newList := make([][rangeBoundarySize]int, 0)
	newRange := [rangeBoundarySize]int{}
	if leftPos >= 0 {
		if leftBoundary <= rl.list[leftPos][rightIdx] {
			newList = append(newList, rl.list[0:leftPos]...)
			if leftBoundary > rl.list[leftPos][leftIdx] {
				newRange[leftIdx] = rl.list[leftPos][leftIdx]
				newRange[rightIdx] = leftBoundary
				newList = append(newList, newRange)
			}
		} else {
			newList = append(newList, rl.list[0:leftPos+1]...)
		}
	}

	if rightPos >= 0 {
		if rightBoundary < rl.list[rightPos][rightIdx] {
			newRange[leftIdx] = rightBoundary
			newRange[rightIdx] = rl.list[rightPos][rightIdx]
			newList = append(newList, newRange)
		}
		if rightPos < len(rl.list)-1 {
			newList = append(newList, rl.list[rightPos+1:]...)
		}
	} else {
		return
	}

	rl.list = newList
}

// Print out the list of ranges in the range list.
func (rl RangeList) Print() {
	for i := 0; i < len(rl.list); i++ {
		fmt.Printf("%s%d, %d%s ", "[", rl.list[i][leftIdx], rl.list[i][rightIdx], ")")
	}
	fmt.Println()
}

// findPosition use binary search to find a position.
func (rl RangeList) findPosition(number int) int {
	if len(rl.list) <= 0 {
		return -1
	}

	left := 0
	right := len(rl.list) - 1
	for left <= right {
		mid := (left + right) / 2
		if number < rl.list[mid][leftIdx] {
			right = mid - 1
		}
		if number > rl.list[mid][leftIdx] {
			left = mid + 1
		}
		if number == rl.list[mid][leftIdx] {
			return mid
		}
	}

	if left < right {
		return left
	} else {
		return right
	}
}

func main() {
	rl := new(RangeList)
	rl.Add([2]int{1, 5})
	rl.Add([2]int{10, 20})
	rl.Add([2]int{20, 20})
	rl.Add([2]int{20, 21})
	rl.Add([2]int{2, 4})
	rl.Add([2]int{3, 8})
	rl.Remove([2]int{10, 10})
	rl.Remove([2]int{10, 11})
	rl.Remove([2]int{15, 17})
	//rl.Remove([2]int{3, 19})
	rl.Print()
}
