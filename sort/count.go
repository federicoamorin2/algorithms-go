package sort

// CountSort performs count sort aglorithm
func CountSort(array []int) []int {
	var countDict = make(map[int]int)
	min := minValue(array)
	max := maxValue(array)
	for i := min; i <= max; i++ {
		countDict[i] = 0
	}
	for _, v := range array {
		countDict[v]++
	}
	sortedArray := make([]int, len(array))
	sortedArrayIndex := 0
	for i := min; i <= max; i++ {
		for j := 0; j < countDict[i]; j++ {
			sortedArray[sortedArrayIndex] = i
			sortedArrayIndex++
		}
	}
	return sortedArray
}

// minValue finds the minumum value in an array
func minValue(arr []int) int {
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// maxValue finds the minumum value in an array
func maxValue(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
