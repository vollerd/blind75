// https://leetcode.com/problems/valid-palindrome/

package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	ss := []rune(strings.ToLower(s))
	i, j := 0, len(ss)-1
	for i < j {
		if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz0123456789", ss[i]) {
			i++
		} else if !strings.ContainsRune("abcdefghijklmnopqrstuvwxyz0123456789", ss[j]) {
			j--
		} else if ss[i] != ss[j] {
			return false
		} else {
			i++
			j--
		}
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
