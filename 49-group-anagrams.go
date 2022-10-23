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

package main

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
    for _, str := range strs {
		// string型を[ ]byte型に変換
        b := []byte(str) // 出力example: [97 98 99]
		// byteを照準で並び替える
        sort.Slice(b, func(i, j int) bool {
            return b[i] < b[j]
        })
		// 並び替えたbをstring型に戻してhashmapのkeyとする
        key := string(b)
		// sliceに追加
        groups[key] = append(groups[key], str)
    }
    
    ans := make([][]string, 0, len(groups))
	// rangeで回して、groupsのvalueだけをansにappendしていく
    for _, v := range groups {
        ans = append(ans, v)        
    }
    return ans
}