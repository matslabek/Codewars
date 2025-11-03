package kata

// https://www.codewars.com/kata/511f11d355fe575d2c000001/train/go
// Partition function
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low // Index keeping track where smaller elements are
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func TwoOldestAges(ages []int) [2]int {
	sortedArr := quickSortStart(ages)
	return [2]int{sortedArr[len(ages)-2], sortedArr[len(ages)-1]}
}
