package main

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
func TestSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{12, 12}, []int{12, 12}},
		{[]int{1, 2, 7, 4, 10}, []int{1, 2, 4, 7, 10}},
		{[]int{2, -1}, []int{-1, 2}},
		{[]int{-1412, 8, -12, 132}, []int{-1412, -12, 8, 132}},
		{[]int{0, 14, -2}, []int{-2, 0, 14}},
	} {

		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort1(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort1(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
