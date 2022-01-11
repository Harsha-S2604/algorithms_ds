package main

import (
	"fmt"
	"github.com/Harsha-S2604/algorithms_ds/sort"
)

func main() {
	fmt.Println("Algorithms and data structure")
	arr := sort.InsertionSort([]int{10, -2, 100, -10, -1, 12, 1})
	fmt.Println(arr)
}
