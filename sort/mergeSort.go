package sort

func MergeSort(arr []int) {
	arrLen := len(arr)

	if arrLen < 2 {
		return
	}

	mid := arrLen / 2
	leftArr, rightArr := make([]int, mid), make([]int, arrLen - mid)
	for i := 0; i < mid; i++ {
		leftArr[i] = arr[i]
	}

	for j := mid; j < arrLen; j++ {
		rightArr[j - mid] = arr[j]
	}

	MergeSort(leftArr)
	MergeSort(rightArr)
	merge(leftArr, rightArr, arr)
}

func merge(leftArr , rightArr, arr []int) {
	i, j, k := 0, 0, 0
	leftLen, rightLen := len(leftArr), len(rightArr)

	for i < leftLen && j < rightLen {
		
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i];
			i++;
		} else {
			arr[k] = rightArr[j]
			j++;
		}
		k++
	}

	for i < leftLen {
		arr[k] = leftArr[i]
		i++
		k++
	}

	for j < rightLen {
		arr[k] = rightArr[j]
		j++
		k++
	}

}