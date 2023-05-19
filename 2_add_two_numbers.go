package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Int, l2Int := 0, 0
	digit := 1
	for l1 != nil {
		l1Int += l1.Val * digit
		digit *= 10
		l1 = l1.Next
	}
	digit = 1
	for l2 != nil {
		l2Int += l2.Val * digit
		digit *= 10
		l2 = l2.Next
	}

	AnsInt := l1Int + l2Int
	Ans := &ListNode{
		Val: AnsInt % 10,
	}
	AnsInt /= 10

	current := Ans
	for AnsInt != 0 {
		nodeDigit := AnsInt % 10
		NextNode := &ListNode{
			Val: nodeDigit,
		}
		current.Next = NextNode
		current = current.Next
		AnsInt /= 10
	}

	return Ans
}

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/
