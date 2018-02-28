// This package implements a number of sorting algorithms
package demosort

func merge(left []int, right []int) []int {
	var result []int
	var i,j int

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	if i < len(left) {
		result = append(result, left[i:]...)
	}

	if j < len(right) {
		result = append(result, right[j:]...)
	}

	return result
}

func MergeSort(values []int) []int {
	valuesLenght := len(values)
	if valuesLenght <= 1 {
		return values
	} else {
		idx := valuesLenght / 2

		p1 := MergeSort(values[:idx])
		p2 := MergeSort(values[idx:])

		return merge(p1, p2)
	}
}
