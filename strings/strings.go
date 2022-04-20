package main

import "math"

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
func main() {
	lengthOfLongestSubstring("abcabcbb")
}

/*
tag: 字符串，滑动窗口
*/
func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]bool{}
	var (
		left, right int
	)

	res := math.MinInt
	bytes := []byte(s)
	for right < len(bytes) {
		// 不满足条件，则从左边缩小窗口，找到最后一个符合条件的
		for m[bytes[right]] {
			delete(m, bytes[left])
			left++
		}
		m[bytes[right]] = true
		res = max(res, len(m))
		right++
	}

	if res == math.MinInt {
		res = 0
	}
	return res
}

/*
tag: 数组、滑动窗口
*/
func minSubArrayLen(target int, nums []int) int {
	var (
		left, right, cur int
	)

	res := math.MaxInt

	for right < len(nums) {
		cur += nums[right]
		/*
			找到满足条件的集合，从左边缩小窗口，看能不能更小
		*/
		for cur >= target {
			res = min(res, right-left+1)
			cur -= nums[left]
			left++
		}
		right++
	}

	if res == math.MaxInt {
		res = 0
	}
	return res
}
