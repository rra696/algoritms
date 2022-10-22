package main

import "fmt"

/*
RSQ(range sum query)
*/
func main() {
	nums := []int{5, 3, 8, 1, 4, 6}
	prefixSum := makePrefixSum(nums)

	sum := rsq(prefixSum, 2, 5)

	fmt.Println(sum)
}

func makePrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)

	for i := 1; i < len(nums)+1; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	return prefixSum
}

func rsq(prefixSum []int, left, right int) int {
	return prefixSum[right] - prefixSum[left]
}
