package main

func MoveZeros(arr []int) []int {
	zC := 0
	resultSlice := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zC += 1
		} else {
			resultSlice = append(resultSlice, arr[i])
		}
	}
	zeroSlice := make([]int, zC)
	resultSlice = append(resultSlice, zeroSlice...)
	return resultSlice
}
