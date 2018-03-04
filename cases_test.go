package gosort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

const testArraySize = 1e+4

var randomArray = generateRandomArray()

func sortArray(values []int) []int {
	svalues := values[:]
	sort.Ints(svalues)

	return svalues
}

func generateRandomArray() []int {
	var result = make([]int, testArraySize, testArraySize)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i, _ := range result {
		result[i] = r1.Intn(testArraySize)
	}

	return result
}

type sortFunction func([]int) []int

func testSortFunction(t *testing.T, fn sortFunction, fnName string) {
	for _, tc := range addCases {
		in := tc.in
		want := tc.want
		got := fn(in)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf(`FAIL: %s  %s(%s)  %s want %s`, tc.description, fnName, in, got, want)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(addCases), "cases.")
}

var addCases = []struct {
	description string
	in          []int
	want        []int
}{

	{"sorting an empty array",
		[]int{},
		[]int{},
	},
	{
		"sorting an array with a single element",
		[]int{1},
		[]int{1},
	},
	{
		"sorting an array with two elements",
		[]int{2, 1},
		[]int{1, 2},
	},
	{
		"sorting an array with two elements",
		[]int{5, 5},
		[]int{5, 5},
	},
	{
		"sorting an array with three elements",
		[]int{2, 3, 1},
		[]int{1, 2, 3},
	},
	{
		"sorting an array with four elements",
		[]int{4, 2, 3, 1},
		[]int{1, 2, 3, 4},
	},
	{
		"sorting arbitrary array",
		randomArray,
		sortArray(randomArray),
	},
}
