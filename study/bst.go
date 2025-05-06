package study

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// A Binary Search Tree (BST) is a type of binary tree where:

// Each node has at most two children: left and right.

// Left Subtree Rule: All values in the left subtree are less than the node’s value.

// Right Subtree Rule: All values in the right subtree are greater than the node’s value.

// No duplicate values are allowed (unless otherwise specified).

// 1. Insert(value)
// Add a new value into the BST, maintaining the ordering rules.
func insert(value int, node *TreeNode) *TreeNode {

	if node == nil {
		return &TreeNode{Val: value}

	}
	if node.Val == value {
		fmt.Printf("\nCan't insert duplicate: %d\n", value)
		return node
	} else if node.Val > value {
		node.Left = insert(value, node.Left)
	} else {
		node.Right = insert(value, node.Right)
	}
	return node
}
func VisualizeBST(node *TreeNode, space int) {
	if node == nil {
		return
	}

	// Increase space between levels
	space += 5

	// Print right subtree first (to visualize it on the top part of the tree)
	VisualizeBST(node.Right, space)

	// Print current node after proper indentation
	fmt.Println()
	for i := 5; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", node.Val)

	// Print left subtree
	VisualizeBST(node.Left, space)
}

// 2. Search(value)
// Return true if the value exists in the BST, false otherwise.
func Search(value int, node *TreeNode) bool {
	if node == nil {
		return false
	}
	if value == node.Val {
		return true
	} else if node.Val > value {
		return Search(value, node.Left)
	} else {
		return Search(value, node.Right)
	}

}

// 3. Delete(value)
// Remove a node from the BST:

// If it has no children → remove directly.

// If it has one child → replace with child.

// If it has two children → replace with the in-order successor or predecessor.

// 4. InOrderTraversal()
// Traverse the tree in Left → Node → Right order.

// Should return values in sorted order.

// 5. PreOrderTraversal()
// Traverse as Node → Left → Right.

// 6. PostOrderTraversal()
// Traverse as Left → Right → Node.

// 7. FindMin()
// Return the minimum value in the BST (leftmost node).

// 8. FindMax()
// Return the maximum value in the BST (rightmost node).

// 9. FindHeight()
// Return the height of the tree (longest path from root to a leaf).

// 10. IsValidBST()
// Verify if a given binary tree satisfies the BST rules.
