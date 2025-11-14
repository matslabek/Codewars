package kata

// https://www.codewars.com/kata/526d84b98f428f14a60008da/train/go
func Hammer(n int) uint {
	if n == 1 {
		return 1
	}
	hamming := make([]uint, n)
	hamming[0] = 1
	// Keep track of the multiplication
	i2, i3, i5 := 0, 0, 0
	for i := 1; i < n; i++ {
		next2 := hamming[i2] * 2
		next3 := hamming[i3] * 3
		next5 := hamming[i5] * 5
		// take the minimum of the candidates
		hamming[i] = min(next2, min(next3, next5))
		//advanced pointers that produced the minimum
		if hamming[i] == next2 {
			i2++
		}
		if hamming[i] == next3 {
			i3++
		}
		if hamming[i] == next5 {
			i5++
		}
	}

	// return the nth element
	return hamming[n-1]
}

func min(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}
