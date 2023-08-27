package main

func subsets(nums []int) [][]int {
	var ans [][]int
	for bit := 0; bit < (1 << len(nums)); bit++ {
		subset := []int{}
		for i := 0; i < len(nums); i++ {
			if bit&(1<<i) > 0 {
				subset = append(subset, nums[i])
			}
		}
		ans = append(ans, subset)
	}
	return ans
}

// Given an integer array nums of unique elements, return all possible
// subsets
//  (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Example 1:

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
// Example 2:

// Input: nums = [0]
// Output: [[],[0]]

// Constraints:

// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
