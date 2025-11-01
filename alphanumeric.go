package main

func Alphanumeric(str string) bool {
	for _, char := range str {
		if !(('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9')) {
			return false
		}
	}
	return len(str) > 0
}
