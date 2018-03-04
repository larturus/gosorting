package gosort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	testSortFunction(t, MergeSort, "MergeSort")
}
