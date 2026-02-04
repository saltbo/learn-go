package main

import "fmt"

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
func arrayOfProducts(nums []int) []int {
	result := make([]int, len(nums))

	leftRunningProduct := 1
	for i := 0; i < len(nums); i++ {
		result[i] = leftRunningProduct
		leftRunningProduct *= nums[i]
	}

	rightRunningProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] *= rightRunningProduct
		rightRunningProduct *= nums[i]
	}

	return result
}

func main() {
	array := []int{5, 1, 4, 2}
	result := arrayOfProducts(array)
	fmt.Println(result) // Output: [8, 40, 10, 20]
}

// Brute force solution
// r[0] = nums[1] * nums[2] * nums[3]
// r[1] = nums[0] * nums[2] * nums[3]
// r[2] = nums[0] * nums[1] * nums[3]
// r[3] = nums[0] * nums[1] * nums[2]

// Smart solution
// L[i] = nums[0]⋅nums[1]⋯nums[i−1]
// R[i] = nums[i+1]⋅nums[i+2]⋯nums[n−1]
// ans[i] = L[i] ​× R[i]​
