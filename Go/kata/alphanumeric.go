package kata

// https://www.codewars.com/kata/526dbd6c8c0eb53254000110/train/go
func Alphanumeric(str string) bool {
	for _, char := range str {
		if !(('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')) {
			return false
		}
	}
	return len(str) > 0
}
