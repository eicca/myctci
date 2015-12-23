package base

import (
	"sort"
	"strings"
)

func uniqueMap(in string) bool {
	chars := make(map[rune]bool)
	for _, ch := range in {
		if chars[ch] {
			return false
		}
		chars[ch] = true
	}
	return true
}

func uniqueSort(in string) bool {
	arr := strings.Split(in, "")
	sort.Strings(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			return false
		}
	}
	return true
}
