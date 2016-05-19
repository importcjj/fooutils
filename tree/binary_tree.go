package tree

import "fmt"

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

// Traversal node value handler
type TraversalCallback func(interface{})

// 1. Display the data part of the root (or current node).
// 2. Traverse the left subtree by recursively calling the pre-order function.
// 3. Traverse the right subtree by recursively calling the pre-order function.
func (root *BinaryTreeNode) PreOrderTraversal(callback TraversalCallback) {
	callback(root.Value)
	if root.Left != nil {
		root.Left.PreOrderTraversal(callback)
	}
	if root.Right != nil {
		root.Right.PreOrderTraversal(callback)
	}
}

// 1. Traverse the left subtree by recursively calling the in-order function.
// 2. Display the data part of the root (or current node).
// 3. Traverse the right subtree by recursively calling the in-order function.
func (root *BinaryTreeNode) InOrderTraversal(callback TraversalCallback) {
	if root.Left != nil {
		root.Left.InOrderTraversal(callback)
	}
	callback(root.Value)
	if root.Right != nil {
		root.Right.InOrderTraversal(callback)
	}

}

// 1. Traverse the left subtree by recursively calling the post-order function.
// 2. Traverse the right subtree by recursively calling the post-order function.
// 3. Display the data part of the root (or current node).
func (root *BinaryTreeNode) PostOrderTraversal(callback TraversalCallback) {
	if root.Left != nil {
		root.Left.PostOrderTraversal(callback)
	}
	if root.Right != nil {
		root.Right.PostOrderTraversal(callback)
	}
	callback(root.Value)

}

func main() {

	root := &BinaryTreeNode{"Hello, root", nil, nil}
	root.Left = &BinaryTreeNode{"Hello, left", nil, nil}
	root.Right = &BinaryTreeNode{"Hello, right", nil, nil}
	root.Left.Left = &BinaryTreeNode{"Hello, left2", nil, nil}
	root.Left.Right = &BinaryTreeNode{"Hello, right2", nil, nil}

	root.PreOrderTraversal(func(value interface{}) { fmt.Println(value) })
	fmt.Println()
	root.InOrderTraversal(func(value interface{}) { fmt.Println(value) })
	fmt.Println()
	root.PostOrderTraversal(func(value interface{}) { fmt.Println(value) })

}
