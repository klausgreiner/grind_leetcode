// 125. Valid Palindrome
// Easy
// Topics
// Companies
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:

// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
// Example 2:

// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
// Example 3:

// Input: s = " "
// Output: true
// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.

// Constraints:

// 1 <= s.length <= 2 * 105
// s consists only of printable ASCII characters.

package exercises

import (
	"fmt"
	"strings"
	"unicode"
)

/*
Problem Description: Valid Parentheses (LeetCode #20)

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:
1. Open brackets must be closed by the same type of brackets.
2. Open brackets must be closed in the correct order.
3. Every close bracket has a corresponding open bracket of the same type.

Example 1:
Input: s = "()"
Output: true

Example 2:
Input: s = "()[]{}"
Output: true

Example 3:
Input: s = "(]"
Output: false

Constraints:
* 1 <= s.length <= 10^4
* s consists of parentheses only '()[]{}'.

(Description fetched/reconstructed, add link: https://leetcode.com/problems/valid-parentheses/)
*/
func removeNonAlphanumericToLower(s string) string {
	var builder strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			builder.WriteRune(unicode.ToLower(r))
		}
	}
	return builder.String()
}

// func isPalindromeTwo(s string) bool {
// 	cleanedString := []rune(removeNonAlphanumericToLower(s))
// 	n := len(cleanedString)

//		for i := 0; i < n/2; i++ {
//			if cleanedString[n-1-i] != cleanedString[i] {
//				return false
//			}
//		}
//		return true
//	}
func isPalindrome(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		// Skip non-alphanumeric on the left
		for left < right && !unicode.IsLetter(rune(s[left])) && !unicode.IsDigit(rune(s[left])) {
			left++
		}
		// Skip non-alphanumeric on the right
		for left < right && !unicode.IsLetter(rune(s[right])) && !unicode.IsDigit(rune(s[right])) {
			right--
		}

		if unicode.ToLower(rune(s[left])) != unicode.ToLower(rune(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

// Test cases for ValidParenthesis
var validPalindromeTestCases = []struct {
	input    string
	expected bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
	{"olá", false},
}

// RunValidParenthesisTests executes the test cases for ValidParenthesis
func RunValidIsPalindromeTests() {
	fmt.Println("--- Running Palindrome Tests ---")
	for _, tc := range validPalindromeTestCases {
		result := isPalindrome(tc.input)
		fmt.Printf("Input: \"%s\", Expected: %t, Got: %t\n", tc.input, tc.expected, result)
		if result != tc.expected {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
	}
	fmt.Println("--- Finished ValidParenthesis Tests ---")
}
