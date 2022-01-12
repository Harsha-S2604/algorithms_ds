package sort


/*PSEUDOCODE:
	for i = 1 to n-2
		flag = 0
		for j = 1 to n - i - 2
			if(arr[i] > arr[i+1)
				swap(arr[i], aee[i+1])
				flag = 1
		if flag == 0
			break
*/

/* algorithm analysis 
	worst: O(n^2)
	avg: O(n^2)
	best: O(n)
	space complexity: O(1)
*/
func BubbleSort(arr []int) []int {
	
	arrLen := len(arr)
	
	for j := 0 ; j < arrLen - 1; j++ {
		flag := 0
		for i := 0; i < arrLen - j - 1; i++ {
			if arr[i] > arr[i+1] {
				swap(&arr[i], &arr[i+1])
				flag = 1
			}
		}
		if flag == 0 {
			break
		}

	}
	return arr
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

