package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	arr      []int
	expected []int
}

func TestQuickSort(t *testing.T) {
	for _, tc := range dataProvider() {
		actual := quickSort(tc.arr)
		assert.Equal(t, tc.expected, actual)
	}
}

func dataProvider() []testCase {
	return []testCase{
		{
			arr:      []int{4},
			expected: []int{4},
		},
		{
			arr:      []int{4, 1},
			expected: []int{1, 4},
		},
		{
			arr:      []int{4, 5, 1, 6, 2},
			expected: []int{1, 2, 4, 5, 6},
		},
		{
			arr:      []int{14, 55, 61, 3, 128},
			expected: []int{3, 14, 55, 61, 128},
		},
	}
}
