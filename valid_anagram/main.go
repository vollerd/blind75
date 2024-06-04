// https://leetcode.com/problems/valid-anagram/description/

package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	runeMap := make(map[rune]int)
	for _, r := range s {
		runeMap[r]++
	}
	for _, r := range t {
		runeMap[r]--
	}
	for _, v := range runeMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))

}
