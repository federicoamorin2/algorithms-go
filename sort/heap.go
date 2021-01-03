package sort

// MaxHeapSort performs the max heap sort algorithm
func MaxHeapSort(arr []int) []int {
	lenArr := len(arr)
	if lenArr <= 1 {
		return arr
	}
	arr = maxHeapify(arr)
	sortedArr := make([]int, 0)
	for i := 0; i < lenArr; i++ {
		sortedArr = append(sortedArr, arr[0])
		arr[0] = arr[len(arr)-1]
		arr = arr[:len(arr)-1]
		arr = doHeapify(arr, 0)
	}
	return sortedArr
}

// maxHeapify auxiliary function to fully maxheapify an array
func maxHeapify(arr []int) []int {
	splitPoint := len(arr) / 2
	for i := splitPoint; i >= 0; i-- {
		arr = doHeapify(arr, i)
	}
	return arr
}

// doHeapify corrects one heap downwords, assuming child heaps are already max heaps
func doHeapify(arr []int, index int) []int {
	var pairIndex int = (index+1)*2 - 1
	var oddIndex int = (index + 1) * 2

	lenArr := len(arr)
	if pairIndex < lenArr && oddIndex < lenArr {
		if arr[pairIndex] > arr[index] || arr[oddIndex] > arr[index] {
			if arr[pairIndex] > arr[oddIndex] {
				arr[pairIndex], arr[index] = arr[index], arr[pairIndex]
				arr = doHeapify(arr, pairIndex)
			} else {
				arr[oddIndex], arr[index] = arr[index], arr[oddIndex]
				arr = doHeapify(arr, oddIndex)
			}
		}
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
