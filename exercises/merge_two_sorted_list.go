package exercises

import "fmt"

// 21. Merge Two Sorted Lists
// Easy
// Topics
// Companies
// You are given the heads of two sorted linked lists list1 and list2.

// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.

// Return the head of the merged linked list.

// Example 1:
// (1) -> (2) -> (4)
// (1) -> (3) -> (4)
// result:
//(1) -> (1) ->(2) -> (3) -> (4) -> (4)
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]
// Example 2:

// Input: list1 = [], list2 = []
// Output: []
// Example 3:

// Input: list1 = [], list2 = [0]
// Output: [0]

// Constraints:

// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

type ListNode struct {
	Val  int
	Next *ListNode
}

func addRecursively(nodeList1 *ListNode, nodeList2 *ListNode) *ListNode {
	if nodeList1 == nil && nodeList2 == nil {
		return nil
	}

	var val int
	var next1, next2 *ListNode

	if nodeList1 == nil {
		val = nodeList2.Val
		next1 = nil
		next2 = nodeList2.Next
	} else if nodeList2 == nil {
		val = nodeList1.Val
		next1 = nodeList1.Next
		next2 = nil
	} else if nodeList1.Val <= nodeList2.Val {
		val = nodeList1.Val
		next1 = nodeList1.Next
		next2 = nodeList2
	} else {
		val = nodeList2.Val
		next1 = nodeList1
		next2 = nodeList2.Next
	}

	node := &ListNode{Val: val}
	node.Next = addRecursively(next1, next2)

	fmt.Println(listToSlice(node))
	return node
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	return addRecursively(list1, list2)
}

// Helper to build linked list from slice
func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for _, v := range vals[1:] {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head
}

// Helper to convert list to slice for easier comparison/printing
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

// mergeTwoListsTestCase holds input and expected output for test cases.
type mergeTwoListsTestCase struct {
	list1    *ListNode
	list2    *ListNode
	expected *ListNode
}

// Sample test cases for mergeTwoLists
var mergeTwoListsTestCases = []mergeTwoListsTestCase{
	{
		list1:    buildList([]int{1, 2, 4}),
		list2:    buildList([]int{1, 3, 4}),
		expected: buildList([]int{1, 1, 2, 3, 4, 4}),
	},
	{
		list1:    buildList([]int{}),
		list2:    buildList([]int{}),
		expected: buildList([]int{}),
	},
	{
		list1:    buildList([]int{}),
		list2:    buildList([]int{0}),
		expected: buildList([]int{0}),
	},
}

// RunMergeTwoListsTests executes the test cases for mergeTwoLists
func RunMergeTwoListsTests() {
	fmt.Println("--- Running mergeTwoLists Tests ---")
	for i, tc := range mergeTwoListsTestCases {
		result := mergeTwoLists(tc.list1, tc.list2)
		got := listToSlice(result)
		want := listToSlice(tc.expected)
		pass := fmt.Sprintf("PASS")
		if fmt.Sprintf("%v", got) != fmt.Sprintf("%v", want) {
			pass = "FAIL"
		}
		fmt.Printf("Test %d: Got: %v, Expected: %v [%s]\n", i+1, got, want, pass)
		if pass == "FAIL" {
			fmt.Println("   ^^^ TEST FAILED ^^^ ")
		}
	}
	fmt.Println("--- Finished mergeTwoLists Tests ---")
}
