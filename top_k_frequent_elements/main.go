// https://leetcode.com/problems/top-k-frequent-elements/description/

package main

import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// Step 1: Count the frequency of each number
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	// Step 2: Create a frequency array where the index represents the frequency
	freq := make([][]int, len(nums)+1)
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	// Step 3: Gather the top k frequent elements
	res := []int{}
	for i := len(freq) - 1; i > 0 && len(res) < k; i-- {
		for _, num := range freq[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	result := topKFrequent(nums, k)
	fmt.Println(result)
}
