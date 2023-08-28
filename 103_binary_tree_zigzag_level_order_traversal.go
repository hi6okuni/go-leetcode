package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}
	currentNest := 0

	recursiveZigzag(root, currentNest, &ans)

	return ans
}

func recursiveZigzag(root *TreeNode, currentNest int, ans *[][]int) {
	if root == nil {
		return
	}

	if currentNest >= len(*ans) {
		*ans = append(*ans, []int{root.Val})
	} else {
		(*ans)[currentNest] = append((*ans)[currentNest], root.Val)
	}

	currentNest++
	recursiveZigzag(root.Left, currentNest, ans)
	recursiveZigzag(root.Right, currentNest, ans)
}

// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1:

// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]
// Example 2:

// Input: root = [1]
// Output: [[1]]
// Example 3:

// Input: root = []
// Output: []

// Constraints:

// The number of nodes in the tree is in the range [0, 2000].
// -100 <= Node.val <= 100
