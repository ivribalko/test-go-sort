package main

import (
	"reflect"
	"testing"
)

// Test all sorters can sort.
func TestSort(t *testing.T) {
	cases := []struct {
		in, want [3]int
	}{
		{[3]int{}, [3]int{}},
		{[3]int{3, 2, 1}, [3]int{1, 2, 3}},
		{[3]int{1, 2, 3}, [3]int{1, 2, 3}},
		{[3]int{3, 3, 1}, [3]int{1, 3, 3}},
		{[3]int{3, 3, 3}, [3]int{3, 3, 3}},
		{[3]int{0, 1, -1}, [3]int{-1, 0, 1}},
	}
	for _, s := range []iSort{
		bubble{},
		selection{},
	} {
		for _, test := range cases {
			copy := test.in
			s.sort(copy[0:])
			if !reflect.DeepEqual(copy, test.want) {
				t.Errorf("sort(%q) is %q, want %q", test.in, copy, test.want)
			}
		}
	}
}
