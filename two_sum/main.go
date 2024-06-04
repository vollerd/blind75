// https://leetcode.com/problems/two-sum/description/

package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	locs := make(map[int]int)
	for i, num := range nums {
		if j, ok := locs[target-num]; ok {
			return []int{j, i}
		}
		locs[num] = i
	}
	return nil

}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 2, 4}
	target = 6
	fmt.Println(twoSum(nums, target))

	nums = []int{3, 3}
	target = 6
	fmt.Println(twoSum(nums, target))
}
