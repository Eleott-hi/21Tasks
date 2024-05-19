package main

import "fmt"

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func countToys(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftCount := countToys(node.Left)
	rightCount := countToys(node.Right)
	toyCount := 0
	if node.HasToy {
		toyCount = 1
	}
	return leftCount + rightCount + toyCount
}

func areToysBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftCount := countToys(root.Left)
	rightCount := countToys(root.Right)
	return leftCount == rightCount
}

func main() {
	// Balanced tree
	tree1 := &TreeNode{HasToy: false,
		Left:  &TreeNode{HasToy: false, Left: &TreeNode{HasToy: false}, Right: &TreeNode{HasToy: true}},
		Right: &TreeNode{HasToy: true}}
	fmt.Println(areToysBalanced(tree1))

	// Unbalanced tree
	tree2 := &TreeNode{HasToy: true,
		Left:  &TreeNode{HasToy: true, Left: &TreeNode{HasToy: false}, Right: &TreeNode{HasToy: false, Right: &TreeNode{HasToy: true}}},
		Right: &TreeNode{HasToy: false}}
	fmt.Println(areToysBalanced(tree2))
}
