package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	changed := true
	for changed {
		before := len(s)
		changed = false
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "{}", "", -1)
		after := len(s)
		if before != after {
			changed = true
		}
	}
	return len(s) == 0
}

func main() {
	fmt.Println(isValid("()[](}"))
}
