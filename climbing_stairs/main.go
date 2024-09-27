package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n < 4 {
		return n
	}
	no_ways := make([]int, n+1)

	no_ways[2] = 2
	no_ways[3] = 3

	for i := 4; i <= n; i++ {
		no_ways[i] = no_ways[i-1] + no_ways[i-2]
	}
	return no_ways[n]
}

func main() {
	fmt.Println(climbStairs(5))
}
