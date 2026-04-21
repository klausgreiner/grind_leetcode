package exercises

import (
	"fmt"
	"slices"
)

// Input: arr = [3,5,1]
// Output: true
func canMakeArithmeticProgression(arr []int) bool {
	slices.Sort(arr)
	var diff int
	for i := 0; i < len(arr)-1; i++ {
		currentDiff := arr[i] - arr[i+1]
		if i > 0 && diff != currentDiff {
			return false
		}
		diff = currentDiff
	}
	return true
}

// Test cases for Arithmetic Progression
var arithmeticTestCases = []struct {
	arr      []int
	expected bool
}{
	{[]int{3, 5, 1}, true},
	{[]int{1, 2, 4}, false},
	{[]int{0, 0, 0, 0}, true},
	{[]int{-13, -17, -8, -1, -20, -5, -11}, false},
	{[]int{1, 100}, true},
}

// RunArithmeticProgression executes the test cases
func RunArithmeticProgression() {
	fmt.Println("--- Running Arithmetic Progression Tests ---")
	for _, tc := range arithmeticTestCases {
		// We pass a copy of the slice to avoid side effects if the function sorts in-place
		inputCopy := make([]int, len(tc.arr))
		copy(inputCopy, tc.arr)

		result := canMakeArithmeticProgression(inputCopy)

		fmt.Printf("Input: %v\nExpected: %t\nGot:      %t\n",
			tc.arr, tc.expected, result)

		if result != tc.expected {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Arithmetic Progression Tests ---")
}
