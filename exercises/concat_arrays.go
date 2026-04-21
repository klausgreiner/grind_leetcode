package exercises

import (
	"fmt"
	"reflect"
)

// Input: nums = [1,2,1]
// Output: [1,2,1,1,2,1]
func getConcatenation(nums []int) []int {
	ans := append(nums, nums...)
	return ans
}

// Test cases for Concatenation of Array
var concatTestCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2, 1}, []int{1, 2, 1, 1, 2, 1}},
	{[]int{1, 3, 2, 1}, []int{1, 3, 2, 1, 1, 3, 2, 1}},
	{[]int{1}, []int{1, 1}},
	{[]int{0, 0}, []int{0, 0, 0, 0}},
}

// RunConcatenation executes the test cases for the getConcatenation function
func RunConcatenation() {
	fmt.Println("--- Running Concatenation Tests ---")
	for _, tc := range concatTestCases {
		// We use reflect.DeepEqual to compare slices
		result := getConcatenation(tc.nums)

		fmt.Printf("Input:    %v\nExpected: %v\nGot:      %v\n",
			tc.nums, tc.expected, result)

		if !reflect.DeepEqual(result, tc.expected) {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Concatenation Tests ---")
}
