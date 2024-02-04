package main

// bit manipulation
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

// backtracking
func subsets2(nums []int) [][]int {
	ans := [][]int{}
	var backtrack func(idx, cap int, list []int)
	backtrack = func(idx, cap int, list []int) {
		if len(list) == cap {
			temp := make([]int, cap)
			copy(temp, list)
			ans = append(ans, temp)
			return
		}

		for i := idx; i < len(nums); i++ {
			backtrack(i+1, cap, append(list, nums[i]))
		}
	}

	for i := 0; i <= len(nums); i++ {
		backtrack(0, i, []int{})
	}

	return ans
}

// func subsets(nums []int) [][]int {
//     res := [][]int{}

//     for i := 0; i <= len(nums); i++ {
//         find(nums, 0, i, []int{}, &res)
//     }

//     return res
// }

// func find(nums []int, idx, k int, list []int, res *[][]int) {
//     if len(list) == k {
//         temp := make([]int, k)
//         copy(temp, list)
//         (*res) = append(*res, temp)
//         return
//     }

//     for i := idx; i < len(nums); i++ {
//         find(nums, i+1, k, append(list, nums[i]), res)
//     }
// }

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
