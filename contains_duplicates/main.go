// https://leetcode.com/problems/contains-duplicate/description/

package main

import (
	"fmt"
)

func containsDuplicate(nums []int) bool {
	found := make(map[int]bool)
	for _, num := range nums {
		if _, ok := found[num]; ok {
			return true
		}
		found[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	fmt.Println(containsDuplicate(nums))
}
