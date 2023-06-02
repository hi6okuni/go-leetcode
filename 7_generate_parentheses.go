package main

import "fmt"

func generateParenthesis(n int) []string {
	length := n * 2
	ans := []string{}
	for bit := 0; bit < (1 << length); bit++ {
		s := ""
		for i := 0; i < length; i++ {
			// 1が立ってたらsに)を追加する
			if bit&(1<<i) > 0 {
				s += ")"
			} else {
				// 0だったらsに(を追加する
				s += "("
			}
		}
		fmt.Println(s)
		// s = "((()))"
		acc := 0
		flag := true
		for i := 0; i < length; i++ {
			if s[i] == '(' {
				acc++
			} else {
				acc--
			}
			if acc < 0 {
				flag = false
				break
			}
		}
		if flag && acc == 0 {
			ans = append(ans, s)
		}
	}

	return ans
}

/*
bit : 0 0 0 1 1 1 = (bit = 7)
(i=2)
1<<i: 0 0 0 1 0 0
→     0 0 0 1 0 0
(i=3)
1<<i: 0 0 1 0 0 0
→     0 0 0 0 0 0 = 0


0: (
1: )

0 → 0 0 0 0 0 0
1 → 0 0 0 0 0 1

1<<3

0 0 0 0 0 1
↓
0 0 1 0 0 0


0 0 0 0 0 0
0 0 0 0 0 1
0 0 0 0 1 0
0 0 0 0 1 1
,,,
0 0 0 1 1 1
( ( ( ) ) )
1 1 1 1 1 1
*/
