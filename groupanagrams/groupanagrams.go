package groupanagrams

import (
	"sort"
	"strings"
)

func Solution(strs []string) [][]string {
	// stores storedStr -> result index
	lookup := make(map[string]int)
	result := [][]string{}
	for _, s := range strs {
		ss := sortStr(s)
		if val, ok := lookup[ss]; ok {
			result[val] = append(result[val], s)
		} else {
			lookup[ss] = len(result)
			result = append(result, []string{s})
		}
	}
	return result
}

func sortStr(str string) string {
	chars := strings.Split(str, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
