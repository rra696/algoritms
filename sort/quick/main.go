package main

import "fmt"

func main() {
	arr := []int{4, 5, 1, 6, 2}

	result := quickSort(arr)

	fmt.Println(result)
}

/**
Сложность O(n*log2(n))
 */
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	less, greater := extractLessAndGreater(arr, pivot)

	less = quickSort(less)
	greater = quickSort(greater)

	mergedLessPartWithPivot := append(less, pivot)

	return append(mergedLessPartWithPivot, greater...)
}

func extractLessAndGreater(arr []int, pivot int) ([]int, []int) {
	var less []int
	var greater []int

	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			less = append(less, arr[i])
		}

		if arr[i] > pivot {
			greater = append(greater, arr[i])
		}
	}

	return less, greater
}
