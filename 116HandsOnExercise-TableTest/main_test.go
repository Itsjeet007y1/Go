package main

import "testing"

func TestMySum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 2}, 4},
		test{[]int{1, 22}, 23},
		test{[]int{3, 52}, 55},
		test{[]int{-4, 26}, 22},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got", x)
		}
	}
}
