package tree

import "testing"

func TestPreOrderTraversal(t *testing.T) {
	root := &BinaryTreeNode{1, nil, nil}
	root.Left = &BinaryTreeNode{2, nil, nil}
	root.Right = &BinaryTreeNode{3, nil, nil}
	root.Right.Left = &BinaryTreeNode{4, nil, nil}
	root.Right.Left.Right = &BinaryTreeNode{5, nil, nil}

	slice := make([]interface{}, 0, 5)
	valid := []int{1, 2, 3, 4, 5}

	root.PreOrderTraversal(func(value interface{}) { slice = append(slice, value) })
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Failed PreOrderTraversal. Got %v, Expected %v", valid, slice)
		}
	}

}

func TestInOrderTraversal(t *testing.T) {
	root := &BinaryTreeNode{1, nil, nil}
	root.Left = &BinaryTreeNode{2, nil, nil}
	root.Right = &BinaryTreeNode{3, nil, nil}
	root.Right.Left = &BinaryTreeNode{4, nil, nil}
	root.Right.Left.Right = &BinaryTreeNode{5, nil, nil}

	slice := make([]interface{}, 0, 5)
	valid := []int{2, 1, 4, 5, 3}

	root.InOrderTraversal(func(value interface{}) { slice = append(slice, value) })
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Failed InOrderTraversal. Got %v, Expected %v", valid, slice)
		}
	}

}

func TestPostOrderTraversal(t *testing.T) {
	root := &BinaryTreeNode{1, nil, nil}
	root.Left = &BinaryTreeNode{2, nil, nil}
	root.Right = &BinaryTreeNode{3, nil, nil}
	root.Right.Left = &BinaryTreeNode{4, nil, nil}
	root.Right.Left.Right = &BinaryTreeNode{5, nil, nil}

	slice := make([]interface{}, 0, 5)
	valid := []int{2, 5, 4, 3, 1}

	root.PostOrderTraversal(func(value interface{}) { slice = append(slice, value) })
	for i, value := range slice {
		if value != valid[i] {
			t.Errorf("Failed PostOrderTraversal. Got %v, Expected %v", valid, slice)
		}
	}

}
