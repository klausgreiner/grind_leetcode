package study

func RunStudy() {
	var root *TreeNode // Tree is initially empty
	root = insert(7, root)
	root = insert(5, root)
	root = insert(5, root)
	root = insert(3, root)
	root = insert(6, root)
	root = insert(9, root)
	root = insert(8, root)
	root = insert(12, root)
	root = insert(11, root)
	root = insert(4, root)
	root = insert(3, root)
	root = insert(2, root)

	VisualizeBST(root, 5)

	print("\n", Search(5, root))
	print("\n", Search(13, root))
	print("\n", Search(3, root))
	print("\n", Search(1, root))
}
