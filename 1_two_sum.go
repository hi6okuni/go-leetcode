package main

// keyword: hashmap

/*
comment:
map[numberのvalue]numberのposition のようなhashmapを作ることでO(n)で解決できる
mapのkeyにvalueを入れて、mapのvalueにpositionを入れるという直感とは逆のmapを用意することが鍵
*/

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)

	for i, num := range nums {
		if position, ok := hashmap[target-num]; ok {
			return []int{position, i}
		}
		hashmap[num] = i
	}
	return []int{}
}

// func twoSum(nums []int, target int) []int {
// 	for i, num := range nums {
// 		newNumsSlice := nums[i+1:]
// 		maps := map[int]int{}

// 		for ii, num := range newNumsSlice {
// 			maps[num] = ii + i + 1
// 		}

// 		if val, ok := maps[target-num]; ok {
// 			return []int{i, val}
// 		}
// 	}

// 	return nil
// }

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.


Follow-up: Can you come up with an algorithm that is less than O(n2) time complexity?
*/
