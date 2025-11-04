package main

import (
	"fmt"
	"github.com/matslabek/Codewars/Go/kata"
)

func main() {
	//	var testCase = []string{"wlwsasphmxx", "owiaxujylentrklctozmymu", "wpgozvxxiu"}
	//fmt.Println(LongestConsec(testCase, 2))
	//fmt.Println(TwoOldestAges([]int{1, 9, 3, 10, 2, 5, 11, 4, 5, 12, 7}))
	//fmt.Println(sortNumbers([]int{1, 2, 5, 3, 10, 6, 4}))
	//fmt.Println(Decode("XL"))
	//fmt.Println(MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
	//fmt.Println(Alphanumeric(""))
	//fmt.Println(Cakes(map[string]int{"apples": 500, "sugar": 200, "eggs": 1}, map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200}))
	//fmt.Println(kata.FirstNonRepeating("bba"))
	fmt.Println(kata.BalancedParens(4))
	sol := []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}
	fmt.Println(findDifference(sol, kata.BalancedParens(4)))
}

func findDifference(a, b []string) []string {
	bSet := make(map[string]bool)

	// Add all elements from slice B to the map
	for _, str := range b {
		bSet[str] = true
	}

	// Find elements in A that are not in B
	result := make([]string, 0)
	for _, str := range a {
		if !bSet[str] {
			result = append(result, str)
		}
	}

	return result
}
