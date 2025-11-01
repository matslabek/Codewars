package main

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

func main() {
	//fmt.Println(LongestConsec(testCase, 2))
	//fmt.Println(TwoOldestAges([]int{1, 9, 3, 10, 2, 5, 11, 4, 5, 12, 7}))
	//fmt.Println(sortNumbers([]int{1, 2, 5, 3, 10, 6, 4}))
	//fmt.Println(Decode("XL"))
	//fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
	fmt.Println(Alphanumeric(""))
}
