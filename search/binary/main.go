package main

import "fmt"

/**
Для бинарного поиска на вход необходимо передавать отсортированный массив.
Сложность O(log2(n)). Логарифм - операция, обратная возведению в степень.
*/

func main() {
	list := []int{1,2,3,4,5}

	result := binarySearch(list, 4)

	fmt.Println(result)
}

func binarySearch(list []int, item int) *int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] == item {
			return &mid
		}

		if list[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return nil
}