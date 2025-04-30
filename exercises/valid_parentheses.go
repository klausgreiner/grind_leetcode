package exercises

import "fmt"

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

func ValidParentheses(s string) bool {
	stack := []rune{}
	openBrackets := map[rune]bool{'(': true, '[': true, '{': true}
	closingBrackets := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, ch := range s {
		println("Processing character:", string(ch))
		println("Current stack before append:", string(stack))

		if openBrackets[ch] {
			stack = append(stack, ch)
		}
		println("Current stack:", string(stack))
		if closingBrackets[ch] != 0 { // it is a closing bracket?
			println("Closing bracket found:", string(ch))
			if len(stack) == 0 || stack[len(stack)-1] != closingBrackets[ch] {
				println("Unmatched closing bracket:", string(ch))
				fmt.Println("Unmatched closing bracket:", string(ch))
				return false
			}
			println("Matching bracket found. Popping from stack:", string(stack[len(stack)-1]))
			stack = stack[:len(stack)-1]
			println("Current stack after pop:", string(stack))
		}

	}
	println("Finished processing string. Final stack:", string(stack))
	return len(stack) == 0

}

// Test cases for ValidParenthesis
var validParenthesisTestCases = []struct {
	input    string
	expected bool
}{
	{"(){[]}}", false},
	{"()[]{}", true},
	{"([{}])", true},
	{"(((", false},
	{")))", false},
	{"([]", false},
	{"([)]", false},
	{"{[]}", true},
	{"", true},
	{"a(b)", true},
}

// RunValidParenthesisTests executes the test cases for ValidParenthesis
func RunValidParenthesesTests() {
	fmt.Println("--- Running ValidParenthesis Tests ---")
	for _, tc := range validParenthesisTestCases {
		result := ValidParentheses(tc.input)
		fmt.Printf("Input: \"%s\", Expected: %t, Got: %t\n", tc.input, tc.expected, result)
		if result != tc.expected {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
	}
	fmt.Println("--- Finished ValidParenthesis Tests ---")
}
