package main

import (
	"fmt"
	"github.com/Harsha-S2604/algorithms_ds/sort"
)

func main() {
	fmt.Println("Algorithms and data structure")
	arr := []int{100, 100, 98, 1, 4, 9, 8, 1, 4,5, 6, 98, 1, 4, 9, 8, 1, 4,5, 6,100, 98, 1, 4, 9, 8, 1, 4,5, 6,}
	sort.MergeSort(arr)
	fmt.Println(arr)
}
