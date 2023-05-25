package main

// original approach
func isValid(s string) bool {
	aCount, bCount, cCount := 0, 0, 0 // (), {}, []
	accWho := make([]rune, 0)

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			aCount++
			accWho = append(accWho, 'a')
		case ')':
			aCount--
			if len(accWho) < 1 || accWho[len(accWho)-1] != 'a' {
				return false
			}
			accWho = accWho[:len(accWho)-1]
		case '{':
			bCount++
			accWho = append(accWho, 'b')
		case '}':
			bCount--
			if len(accWho) < 1 || accWho[len(accWho)-1] != 'b' {
				return false
			}
			accWho = accWho[:len(accWho)-1]
		case '[':
			cCount++
			accWho = append(accWho, 'c')
		case ']':
			cCount--
			if len(accWho) < 1 || accWho[len(accWho)-1] != 'c' {
				return false
			}
			accWho = accWho[:len(accWho)-1]
		}

		if aCount < 0 || bCount < 0 || cCount < 0 {
			return false
		}
	}

	if aCount == 0 && bCount == 0 && cCount == 0 {
		return true
	} else {
		return false
	}
}

// cool approach
func isValid2(s string) bool {
	l := make([]rune, 0)
	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			l = append(l, v)
		} else if len(l) > 0 {
			e := l[len(l)-1]
			if (v == '}' && e != '{') || (v == ']' && e != '[') || (v == ')' && e != '(') {
				return false
			}
			l = l[:len(l)-1]
		} else {
			return false
		}
	}
	return len(l) == 0
}

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"
// Output: true
// Example 2:

// Input: s = "()[]{}"
// Output: true
// Example 3:

// Input: s = "(]"
// Output: false

// Constraints:

// 1 <= s.length <= 104
// s consists of parentheses only '()[]{}'.
