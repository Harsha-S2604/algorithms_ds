package sort


// NON-DECREASING ORDER
/* insertion sort algorithm
	
	for j = 2 to A.length
		key = A[j]
		i = j - 1
		while i > 0 and A[i] > key
			A[i + 1] = A[i]
			i = i - 1
		A[i + 1] = key
		
*/

// NON-INCREASING ORDER
/* algorithm:
	
	for j = 2 to A.length
		key = A[j]
		i = j - 1
		while i > 0 and A[i] < key
			A[i + 1] = A[i]
			i = i - 1
		A[i + 1] = key
	
*/

/* ALGORITHM ANALYSIS:
	time complexity:
		worst-case => O(n^2)
		avg-case => O(n^2)
		best-case => O(n)
	space complexity: O(1)
*/

func InsertionSort(arr []int) []int {
	
	arrLen := len(arr)
	
	for i := 1; i < arrLen; i++ {
		curr := arr[i]
		index := i - 1
		for index >= 0 && arr[index] > curr {
			arr[index + 1] = arr[index]
			index--
		}

		arr[index + 1] = curr

	}
	
	return arr
}
