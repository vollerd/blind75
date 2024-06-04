// https://leetcode.com/problems/group-anagrams/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	ans := make(map[string][]string)

	for _, s := range strs {
		sortedStr := sortString(s)
		ans[sortedStr] = append(ans[sortedStr], s)
	}

	result := make([][]string, 0, len(ans))
	for _, group := range ans {
		result = append(result, group)
	}

	return result
}

func sortString(s string) string {
	strSlice := strings.Split(s, "")
	sort.Strings(strSlice)
	return strings.Join(strSlice, "")
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
