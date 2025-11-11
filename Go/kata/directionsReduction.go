package kata

//https://www.codewars.com/kata/550f22f4d758534c1100025a/train/go
func DirReduc(arr []string) []string {
	for ; ; {
		reduced := reduce(arr)
		if !areEqual(reduced, arr) {
			arr = reduced
		} else {
			return reduced
		}
	}
}
func reduce(arr []string) []string {
	dirs := map[string]string{
		"NORTH": "SOUTH",
		"WEST":  "EAST",
		"EAST":  "WEST",
		"SOUTH": "NORTH",
	}
	reducedDirs := make([]string, 0)
	for i := 0; i < len(arr); i++ {
		current := arr[i]
		if i != len(arr)-1 {
			if current != dirs[arr[i+1]] {
				reducedDirs = append(reducedDirs, current)
			} else {
				i++
			}
		} else {
			reducedDirs = append(reducedDirs, current)
		}
	}
	return reducedDirs
}
func areEqual(foo []string, bar []string) bool {
	if len(foo) != len(bar) {
		return false
	}
	for i := 0; i < len(foo); i++ {
		if foo[i] != bar[i] {
			return false
		}
	}
	return true
}
