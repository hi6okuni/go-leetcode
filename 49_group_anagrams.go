package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	var ans [][]string
	m := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortStr(str)
		m[sortedStr] = append(m[sortedStr], str)
	}

	for _, v := range m {
		ans = append(ans, v)
	}

	return ans
}

func sortStr(s string) string {
	b := []byte(s)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(b)
}

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

// Example 1:

// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
// Example 2:

// Input: strs = [""]
// Output: [[""]]
// Example 3:

// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:

// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
