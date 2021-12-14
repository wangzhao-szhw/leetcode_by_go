package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l := 0
	start := 0
	dic := map[byte]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		if _, ok := dic[s[i]]; !ok {
			dic[s[i]] = i
		} else {
			if dic[s[i]]+1 > start {
				start = dic[s[i]] + 1
			}
			dic[s[i]] = i
		}
		if l < i-start+1 {
			l = i - start + 1
		}
	}
	return l
}
func main() {
	s := "abcabcbb"
	res := lengthOfLongestSubstring(s)
	fmt.Print(res)
}
