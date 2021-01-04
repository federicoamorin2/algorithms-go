package sort

// MaxHeapSort performs the max heap sort algorithm.
func MaxHeapSort(arr []int) []int {
	// Get lenght of array and return for trivial case.
	lenArr := len(arr)
	if lenArr <= 1 {
		return arr
	}

	// Order array into max Heap.
	arr = maxHeapify(arr)

	// Initialize empty array to return.
	sortedArr := make([]int, 0)

	// Iter through heap getting max valeu and appending to return array.
	for i := 0; i < lenArr; i++ {
		// Add max value to return array.
		sortedArr = append(sortedArr, arr[0])

		// Prepare array for next iteration.
		arr[0] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		arr = doHeapify(arr, 0)
	}
	return sortedArr
}

// Auxiliary function to fully "maxheapify" an array
func maxHeapify(arr []int) []int {
	// Define the starting point in the array
	startPoint := len(arr) / 2

	// Iter descendingly starting from the mid point of the array performing
	// heapify
	for i := startPoint; i >= 0; i-- {
		arr = doHeapify(arr, i)
	}
	return arr
}

// doHeapify corrects one heap downwords, assuming child heaps are already max heaps
func doHeapify(arr []int, index int) []int {

	// Initialize indexes to be used, correcting for zero-base indexing
	var pairIndex int = (index+1)*2 - 1
	var oddIndex int = (index + 1) * 2

	// Check if both pointer are smaller than array.
	lenArr := len(arr)
	if pairIndex < lenArr && oddIndex < lenArr {
		// If condition is met see if array is in incorrect order and rearrange
		// if needed.
		if arr[pairIndex] > arr[index] || arr[oddIndex] > arr[index] {
			if arr[pairIndex] > arr[oddIndex] {
				arr[pairIndex], arr[index] = arr[index], arr[pairIndex]
				arr = doHeapify(arr, pairIndex)
			} else {
				arr[oddIndex], arr[index] = arr[index], arr[oddIndex]
				arr = doHeapify(arr, oddIndex)
			}
		}
		// Taking care of corner cases...
	} else if pairIndex >= lenArr && oddIndex >= lenArr {

		return arr
	} else if pairIndex >= lenArr {
		if arr[oddIndex] > arr[index] {
			arr[oddIndex], arr[index] = arr[index], arr[oddIndex]
			arr = doHeapify(arr, oddIndex)
		}
	} else {
		if arr[pairIndex] > arr[index] {
			arr[pairIndex], arr[index] = arr[index], arr[pairIndex]
			arr = doHeapify(arr, pairIndex)
		}
	}
	return arr
}
