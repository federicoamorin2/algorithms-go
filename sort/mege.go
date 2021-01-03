package sort

// MergeSort performs mergemergeSort algorith
func MergeSort(arr []int, reverse bool) []int {
	if len(arr) == 1 {
		return arr
	}
	splitPoint := len(arr) / 2
	upperSlice, lowerSlice := arr[:splitPoint], arr[splitPoint:]
	upperSlicePrime := MergeSort(upperSlice, reverse)
	lowerSlicePrime := MergeSort(lowerSlice, reverse)
	return mergePart(lowerSlicePrime, upperSlicePrime, reverse)
}

// Auxiliary function of merge sort
func mergePart(upperArr []int, lowerArr []int, reverse bool) []int {
	var pointerUpper int = 0
	var pointerLower int = 0

	mergedArr := make([]int, len(upperArr)+len(lowerArr))

	for pointerUpper+pointerLower < len(mergedArr) {
		var condition bool
		if reverse {
			condition = upperArr[pointerUpper] >= lowerArr[pointerLower]
		} else {
			condition = upperArr[pointerUpper] <= lowerArr[pointerLower]
		}
		if condition {
			mergedArr[pointerUpper+pointerLower] = upperArr[pointerUpper]
			pointerUpper++
		} else {
			mergedArr[pointerUpper+pointerLower] = lowerArr[pointerLower]
			pointerLower++
		}
		if pointerUpper == len(upperArr) {
			for i := pointerUpper + pointerLower; i < len(mergedArr); i++ {
				mergedArr[i] = lowerArr[pointerLower]
				pointerLower++
			}
			break
		}
		if pointerLower == len(lowerArr) {
			for i := pointerUpper + pointerLower; i < len(mergedArr); i++ {
				mergedArr[i] = upperArr[pointerUpper]
				pointerUpper++
			}
			break
		}
	}

	return mergedArr
}
