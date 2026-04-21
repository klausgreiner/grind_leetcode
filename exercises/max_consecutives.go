package exercises

import (
	"fmt"
)

// Input: nums = [1,1,0,1,1,1]
// Output: 3
func findMaxConsecutiveOnes(nums []int) int {
	maxCont := 0
	currentCount := 0
	for i := range nums {
		if nums[i] == 1 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCont {
			maxCont = currentCount
		}
	}

	return maxCont
}

// Test cases for Max Consecutive Ones
var maxOnesTestCases = []struct {
	nums     []int
	expected int
}{
	{[]int{1, 1, 0, 1, 1, 1}, 3},
	{[]int{1, 0, 1, 1, 0, 1}, 2},
	{[]int{0}, 0},
	{[]int{1}, 1},
	{[]int{1, 1, 1, 1}, 4},
}

// RunMaxConsecutiveOnes executes the test cases for the findMaxConsecutiveOnes function
func RunMaxConsecutiveOnes() {
	fmt.Println("--- Running Max Consecutive Ones Tests ---")
	for _, tc := range maxOnesTestCases {
		result := findMaxConsecutiveOnes(tc.nums)

		fmt.Printf("Input: %v\nExpected: %d\nGot:      %d\n",
			tc.nums, tc.expected, result)

		if result != tc.expected {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Max Consecutive Ones Tests ---")
}
