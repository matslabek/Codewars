package main

import s "strings"

func FirstNonRepeating(str string) string {
	sl := s.Split(str, "")
	STR := s.ToUpper(str)
	for i, r := range STR {
		x := s.Count(STR, string(r))
		if x == 1 {
			return sl[i]
		}
	}
	return ""
}
