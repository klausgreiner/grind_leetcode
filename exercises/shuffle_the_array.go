package exercises

import (
	"fmt"
	"reflect" // Used for comparing slices in tests
)

// Input: nums = [2,5,1,3,4,7], n = 3
// Output: [2,3,5,4,1,7]æ
func shuffle(nums []int, n int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i])
		ans = append(ans, nums[n+i])
	}
	return ans
}

// Test cases for Shuffle
var shuffleTestCases = []struct {
	nums     []int
	n        int
	expected []int
}{
	{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
	{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
	{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
}

// RunValidShuffle executes the test cases for the shuffle function
func RunValidShuffle() {
	fmt.Println("--- Running Shuffle Tests ---")
	for _, tc := range shuffleTestCases {
		// We use reflect.DeepEqual to compare slices because == doesn't work on them
		result := shuffle(tc.nums, tc.n)

		fmt.Printf("Input: %v, n: %d\nExpected: %v\nGot:      %v\n",
			tc.nums, tc.n, tc.expected, result)

		if !reflect.DeepEqual(result, tc.expected) {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
		fmt.Println("-----------------------")
	}
	fmt.Println("--- Finished Shuffle Tests ---")
}
