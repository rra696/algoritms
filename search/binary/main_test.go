package main

import "testing"

type testCase struct {
	list []int
	item int
	expectedPosition int
}

func TestBinarySearch(t *testing.T) {
	for _, tc := range dataProvider() {
		result := binarySearch(tc.list, tc.item)

		if result == nil {
			t.Error("Item not found")
		}

		if *result != tc.expectedPosition {
			t.Errorf("Expected %d, but got %d", tc.expectedPosition, result)
		}
	}
}

func dataProvider() []testCase {
	return []testCase{
		{
			list: []int{1, 3, 4, 6, 7, 8},
			item: 3,
			expectedPosition: 1,
		},
	}
}
