package demosort

import (
    "reflect"
    "testing"
)

func TestMergeSort(t *testing.T) {
	for _, tc := range addCases {
        in := tc.in
        want := tc.want
        got := MergeSort(in)
        if !reflect.DeepEqual(got, want) {
			t.Fatalf(`FAIL: %s
MergeSort(%s)
   = %s
want %s`, tc.description, in, got, want)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(addCases), "cases.")
}
