package gosort

import (
	"testing"
)

func outOfPlaceQuickSort(values []int) []int {
	got := make([]int, len(values), (cap(values)+1)*2)
	copy(got, values)
	QuickSort(got)
	return got
}

func TestQuickSort(t *testing.T) {
	testSortFunction(t, outOfPlaceQuickSort, "QuickSort")
}
