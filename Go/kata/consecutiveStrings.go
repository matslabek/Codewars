package kata

// https://www.codewars.com/kata/56a5d994ac971f1ac500003e/train/go
import (
	"fmt"
	"sort"
)

func LongestConsec(strarr []string, k int) string {
	l := len(strarr)
	//handle edge cases
	if l == 0 || k > l || k <= 0 {
		return ""
	}

	concatStrings := make([]string, 0)
	for i := 0; i <= l-k; i++ {
		cStr := strarr[i]
		for j := 1; j < k; j++ {
			cStr += strarr[i+j]
		}
		concatStrings = append(concatStrings, cStr)
	}
	fmt.Println("Concatenated strings", concatStrings)
	sort.SliceStable(concatStrings, func(i, j int) bool { return len(concatStrings[i]) > len(concatStrings[j]) })
	fmt.Println("After sort", concatStrings)
	return concatStrings[0]
}

var testCase = []string{"wlwsasphmxx", "owiaxujylentrklctozmymu", "wpgozvxxiu"}
