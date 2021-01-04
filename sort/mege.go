package sort

// MergeSort performs mergemergeSort algorith
func MergeSort(arr []int, reverse bool) []int {
	// return for trivial cases
	if len(arr) <= 1 {
		return arr
	}
	// Find middle point of array and slpit in two
	splitPoint := len(arr) / 2
	upperSlice, lowerSlice := arr[:splitPoint], arr[splitPoint:]

	// Apply merge sort recursively to sub arrays
	upperSlicePrime := MergeSort(upperSlice, reverse)
	lowerSlicePrime := MergeSort(lowerSlice, reverse)

	// Merge sub-arrays and return
	return mergePart(lowerSlicePrime, upperSlicePrime, reverse)
}

// Auxiliary function of merge sort algorithm that merges two n-sized sorted
// arrays into one 2n-sized array.
func mergePart(upperArr []int, lowerArr []int, reverse bool) []int {
	// Initialize pointers for upper and lower arrays and empty result array.
	var pointerUpper int = 0
	var pointerLower int = 0
	mergedArr := make([]int, len(upperArr)+len(lowerArr))

	// Iter over both pointers while the sum of both is smaller than the total
	// size of array
	for pointerUpper+pointerLower < len(mergedArr) {
		// Create variable that controles order that array is sorted in
		var condition bool
		if reverse {
			condition = upperArr[pointerUpper] >= lowerArr[pointerLower]
		} else {
			condition = upperArr[pointerUpper] <= lowerArr[pointerLower]
		}

		// If condition is met add element of upper array to result array.
		if condition {
			mergedArr[pointerUpper+pointerLower] = upperArr[pointerUpper]
			pointerUpper++
			// If is not met do the oposite and add element of lower array
		} else {
			mergedArr[pointerUpper+pointerLower] = lowerArr[pointerLower]
			pointerLower++
		}

		// Handling corner cases when one pointer has reached the end but the
		// other has not
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
