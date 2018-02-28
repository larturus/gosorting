package demosort

import (
    "math/rand"
    "sort"
    "time"
)

const testArraySize = 1e+5

var randomArray =  generateRandomArray();

func sortArray(values []int) []int {
    svalues := values[:]
    sort.Ints(svalues)

    return svalues
}

func generateRandomArray() []int {
    var result = make([]int, testArraySize, testArraySize)
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    for i,_ := range result {
        result[i] = r1.Intn(testArraySize)
    }

    return result
}

var addCases = []struct {
	description string
	in          []int
	want        []int
}{

    {   "sorting an empty array",
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
