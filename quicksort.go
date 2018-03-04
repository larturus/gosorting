// This package implements a number of sorting algorithms
package gosort

import (
	"fmt"
)

func swap(values []int, index1 int, index2 int) {
	tmp := values[index1]
	values[index1] = values[index2]
	values[index2] = tmp
}

func partition(values []int) int {
	swapIndex := 0
	pivotIndex := len(values) - 1
	pivot := values[pivotIndex]

	for j := 0; j < pivotIndex; j++ {
		if values[j] <= pivot {
			swap(values, swapIndex, j)
			swapIndex++
		}
	}
	swap(values, swapIndex, pivotIndex)

	return swapIndex
}

func QuickSort(values []int) {
	if len(values) > 1 {
		pivotIndex := partition(values)
		QuickSort(values[:pivotIndex])
		QuickSort(values[pivotIndex:])
	}
}

func main() {
	values := []int{8, 1, 7, 9, 2}
	QuickSort(values)
	fmt.Println(values)
}
