package main

// original approach
func lengthOfLongestSubstring(s string) int {
	longest := 0
	size := len(s)

	for i := 0; i < size; i++ {
		length := 0
		letters := make(map[byte]struct{})

		for j := i; j < size; j++ {
			length = j - i + 1
			c := s[j]
			if _, ok := letters[c]; ok {
				length--
				break
			}
			letters[c] = struct{}{}
		}

		if longest > length {
			// longestより小さい場合はそのまま抜ける
			continue
		} else {
			// longestより大きい場合はlongestを更新する
			longest = length
		}
	}

	return longest
}

// cool approach

/*
Given a string s, find the length of the longest
substring
 without repeating characters.



Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.


Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
