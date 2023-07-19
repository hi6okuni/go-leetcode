package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// dfs(stack)
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)

}

// dfs(stack)を再帰を使わずに
func hasPathSum2(root *TreeNode, targetSum int) bool {
	type Stack struct {
		targetSum int
		treenode  *TreeNode
	}

	stack := []Stack{
		{
			targetSum: targetSum,
			treenode:  root,
		},
	}

	for len(stack) > 0 {
		curNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curNode.treenode == nil {
			continue
		}
		curNode.targetSum -= curNode.treenode.Val

		if curNode.treenode.Left == nil && curNode.treenode.Right == nil {
			if curNode.targetSum == 0 {
				return true
			}
			continue
		}

		stack = append(stack, Stack{targetSum: curNode.targetSum, treenode: curNode.treenode.Left})
		stack = append(stack, Stack{targetSum: curNode.targetSum, treenode: curNode.treenode.Right})
	}

	return false
}

// bfs(queue)
func hasPathSum3(root *TreeNode, targetSum int) bool {
	type Queue struct {
		targetSum int
		treenode  *TreeNode
	}

	q := []Queue{
		{
			targetSum: targetSum,
			treenode:  root,
		},
	}

	for len(q) > 0 {
		curNode := q[0]
		q = q[1:]
		if curNode.treenode == nil {
			continue
		}
		curNode.targetSum -= curNode.treenode.Val

		if curNode.treenode.Left == nil && curNode.treenode.Right == nil {
			if curNode.targetSum == 0 {
				return true
			}
			continue
		}

		q = append(q, Queue{targetSum: curNode.targetSum, treenode: curNode.treenode.Left})
		q = append(q, Queue{targetSum: curNode.targetSum, treenode: curNode.treenode.Right})
	}

	return false
}

// Given the root of a binary tree and an integer targetSum, return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.

// A leaf is a node with no children.

// Example 1:

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.
// Example 2:

// Input: root = [1,2,3], targetSum = 5
// Output: false
// Explanation: There two root-to-leaf paths in the tree:
// (1 --> 2): The sum is 3.
// (1 --> 3): The sum is 4.
// There is no root-to-leaf path with sum = 5.
// Example 3:

// Input: root = [], targetSum = 0
// Output: false
// Explanation: Since the tree is empty, there are no root-to-leaf paths.

// Constraints:

// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
