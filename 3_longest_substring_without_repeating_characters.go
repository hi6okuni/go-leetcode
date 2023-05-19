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
func lengthOfLongestSubstring2(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	l, r, max := 0, 0, 0
	dict := make(map[byte]bool)
	for r < len(s) {
		// rが右方向に進んでdict[s[r]]が重複が発生したら、
		// lをfalseで右方向に順番に塗りつぶしていき
		// lが重複した文字まで辿り着くとdict[s[r]]が再びfalseに変わる
		if dict[s[r]] == true {
			dict[s[l]] = false
			l++
		} else {
			dict[s[r]] = true
			if (r-l)+1 > max {
				max = (r - l) + 1
			}
			r++
		}
	}
	return max
}

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
