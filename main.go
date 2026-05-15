package main

import (
	"fmt"

	day1 "github.com/eduardoaugustolb/daily-leetcode/src/day-1"
	day2 "github.com/eduardoaugustolb/daily-leetcode/src/day-2"
)

func main() {
	fmt.Println("# Day 1")
	day1.RunTwoSum()
	day1.RunValidParentheses()
	day1.RunBestTimeToBuy()

	fmt.Println("# Day 2")
	day2.RunMergeTwoSortedList()
	day2.RunMaxSubArray()
	day2.RunContainsDuplicate()
}
