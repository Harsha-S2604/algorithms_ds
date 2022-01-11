package sort

import "fmt"

func InsertionSort(arr []int) []int {
	
	arrLen := len(arr)
	
	for i := 1; i < arrLen; i++ {
		curr := arr[i]
		index := i - 1
		for index >= 0 && arr[index] > curr {
			arr[index + 1] = arr[index]
			index--
			arr[index + 1] = curr
		}

	}
	
	return arr
}
