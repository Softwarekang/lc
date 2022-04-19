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

/*
tag: 字符串，滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]bool{}
	n := len(s)
	rk, ans := 0, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk < n && !m[s[rk]] {
			// 不断地移动右指针
			m[s[rk]] = true
			rk++
		}
		ans = max(ans, rk-i)
	}
	return ans
}
