package kata

//https://www.codewars.com/kata/5264d2b162488dc400000001/train/go
import (
	"strings"
)

func SpinWords(str string) string {
	words := strings.Split(str, " ")
	outputStr := ""
	for _, w := range words {
		if len(w) >= 5 {
			w = reverseString(w)
		}
		outputStr += w + " "
	}
	return strings.Trim(outputStr, " ")
}

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
