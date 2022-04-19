package main

/*
tag: 字符串，双指针
*/
func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	idx := -len(s)
	for k, v := range []byte(s) {
		if v == c {
			idx = k
		}
		res[k] = k - idx
	}

	idx = 2 * len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			idx = i
		}
		res[i] = min(res[i], idx-i)
	}
	return res
}
