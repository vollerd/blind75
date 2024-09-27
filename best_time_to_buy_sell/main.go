package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	lowest := prices[0]
	prof := 0
	for _, p := range prices {
		if p < lowest {
			lowest = p
		}
		prof = max(prof, p-lowest)
	}
	return prof
}

func main() {
	fmt.Println(maxProfit([]int{2, 1, 4}))
}
