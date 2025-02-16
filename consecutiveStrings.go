package main

import (
	"fmt"
	"sort"
)

func LongestConsec(strarr []string, k int) string {
	sort.SliceStable(strarr, func(i, j int) bool { return len(strarr[i]) > len(strarr[j]) })
	fmt.Println(strarr)
	result := ""
	if len(strarr) != 0 {
		for i := 0; i < k; i++ {
			result += strarr[i]
		}
	}

	return result
}

var a = []string{"zone", "abigail", "theta", "form", "libe", "zas"}
var b = []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
var c = []string{"owiaxujylentrklctozmymu", "wlwsasphmxx", "wpgozvxxiu"}

func main() {
	fmt.Println(LongestConsec(c, 2))
}
