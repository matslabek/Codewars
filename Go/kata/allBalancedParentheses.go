package kata

// https://www.codewars.com/kata/5426d7a2c2c7784365000783/train/go
func BalancedParens(n int) []string {
	if n == 0 {
		return []string{""}
	}

	result := []string{}
	generate("", n, 0, 0, &result)
	return result
}

func generate(current string, n, open, close int, result *[]string) {
	// Base case: when we've used all n pairs
	if len(current) == 2*n {
		*result = append(*result, current)
		return
	}

	// Add opening parenthesis if we haven't used all n
	if open < n {
		generate(current+"(", n, open+1, close, result)
	}

	// Add closing parenthesis if it won't create an invalid sequence
	if close < open {
		generate(current+")", n, open, close+1, result)
	}
}
