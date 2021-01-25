package sort

// CountSort performs count sort aglorithm
func CountSort(array []int) []int {

	// Find minumum and maximum value in array.
	min := minValue(array)
	max := maxValue(array)

	// Intialize empty count mapping with zero in all keys.
	var countDict = make(map[int]int)
	for i := min; i <= max; i++ {
		countDict[i] = 0
	}

	// Update count mapping for each occurence of each number.
	for _, v := range array {
		countDict[v]++
	}

	// Fill empty array with values from count mapping.
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

// minValue finds the minumum value in an array.
func minValue(arr []int) int {

	// Iter array updating min variable to contain minimum value seen so far.
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}

// maxValue finds the maximum value in an array.
func maxValue(arr []int) int {

	// Iter array updating max variable to contain maximum value seen so far.
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
