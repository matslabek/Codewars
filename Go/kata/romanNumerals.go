package kata

//https://www.codewars.com/kata/51b6249c4612257ac0000005/train/go

import (
	"fmt"
	"strings"
)

// Roman Numerals Decoder
func Decode(roman string) int {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	chars := strings.Split(roman, "")
	sum := 0
	for i := 0; i < len(chars); i++ {
		if i != len(chars)-1 {
			// Check if the next symbol is greater than the current symbol
			// If so, we're dealing with "compound" number where two symbols represent one value
			if romanNumerals[chars[i]] < romanNumerals[chars[i+1]] {
				switch chars[i+1] {
				case "V":
					sum += 4
					break
				case "X":
					sum += 9
					break
				case "L":
					sum += 40
					break
				case "C":
					sum += 90
					break
				case "D":
					sum += 400
					break
				case "M":
					sum += 900
					break
				}
				i++ // Increment the counter as we read two letters
			} else {
				// For regular numbers just read the value of the symbol
				sum += romanNumerals[chars[i]]
			}
		} else {
			sum += romanNumerals[chars[i]]
		}
	}
	fmt.Printf("The roman numeral %s equals %d \n", roman, sum)
	return sum
}
