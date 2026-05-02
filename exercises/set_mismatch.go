package exercises

import (
	"fmt"
	"reflect"
)

func FindErrorNums(nums []int) []int {
	mapError := make(map[int]int)
	for _, v := range nums {
		if i, ok := mapError[v]; ok {
			mapError[v] = i + 1
		} else {
			mapError[v] = 1
		}
	}

	missingNumber := 0
	result := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		if v, ok := mapError[i]; ok {
			if v > 1 {
				result = append(result, i)
			}
		} else {
			missingNumber = i
		}
	}

	result = append(result, missingNumber)
	return result
}

// Test cases for Set Mismatch
var setMismatchTestCases = []struct {
	nums     []int
	expected []int
}{
	{[]int{1, 2, 2, 4}, []int{2, 3}},
	{[]int{1, 1}, []int{1, 2}},
	{[]int{2, 2}, []int{2, 1}},
	{[]int{3, 2, 3, 4, 6, 5}, []int{3, 1}},
}

// RunSetMismatch executes the test cases for the FindErrorNums function
func RunSetMismatch() {
	fmt.Println("--- Running Set Mismatch Tests ---")
	for _, tc := range setMismatchTestCases {
		result := FindErrorNums(tc.nums)

		fmt.Printf("Input:    %v\nExpected: %v\nGot:      %v\n",
			tc.nums, tc.expected, result)

		// Using reflect.DeepEqual to compare slices
		if !reflect.DeepEqual(result, tc.expected) {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Set Mismatch Tests ---")
}
