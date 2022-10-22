package main

import "fmt"

func main() {
	arr := []int{5, 3, 6, 2, 10}

	result := selectionSort(arr)

	fmt.Println(result)
}

/*
Сортировка выбором. Сложность O(n^2)
*/
func selectionSort(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	length := len(arr)

	for i := 0; i < length; i++ {
		smallest := findSmallest(arr)
		newArr = append(newArr, arr[smallest])

		arr = removeSliceElem(arr, smallest)
	}

	return newArr
}

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

func removeSliceElem(arr []int, elem int) []int {
	arr[elem] = arr[len(arr)-1]

	return arr[:len(arr)-1]
}
