package main

import "testing"

func TestAreOddsOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, false},
		{[]int{1}, true},
		{[]int{4}, false},
		{[]int{-16, 8, 4, 3}, false},
		{[]int{91, 89}, true},
		{[]int{-12, 1831, -88}, false},
		{[]int{1, 3, 5, 6}, false},
		{[]int{-21, 7, 1, 19}, true},
		{[]int{0, 13, 99}, false},
	} {

		if got := AreOddsOnly(tc.s); got != tc.want {
			t.Errorf("AreOddsOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}
func TestAreEvenOnly(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want bool
	}{
		{[]int{}, true},
		{[]int{0}, true},
		{[]int{1}, false},
		{[]int{8}, true},
		{[]int{66, -1}, false},
		{[]int{-101481, 14, -16}, false},
		{[]int{290405, 17, 2022, 11}, false},
		{[]int{1, 3, 5, 9, 6}, false},
		{[]int{-100, -222, -334}, true},
		{[]int{0, 0, 11}, false},
		{[]int{0, 1112, 0}, true},
	} {
		got := AreEvenOnly(tc.s)
		if got != tc.want {
			t.Errorf("AreEvenOnly(%v) = %v , want = %v", tc.s, got, tc.want)
		}
	}
}
