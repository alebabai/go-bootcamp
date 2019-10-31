package word

import (
	"strings"
)

// UseCount returns a map with word as a key and its occurrences number as a value
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count words in a string
func Count(s string) int {
	return len(strings.Fields(s))
}
