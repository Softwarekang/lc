package main

func main() {
	lengthLongestPath("dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext")
}

/*
tag: 字符串、stack、深搜
*/
func lengthLongestPath(input string) int {
	cur, res, nums := 0, 0, make([]int, len(input)+1)
	for cur < len(input) {
		depth := 1
		for cur < len(input) && input[cur] == '\t' {
			depth++
			cur++
		}

		length, isFileName := 0, false
		for cur < len(input) && input[cur] != '\n' {
			if input[cur] == '.' {
				isFileName = true
			}
			length++
			cur++
		}

		cur++

		if depth > 1 {
			length = nums[depth-1] + length + 1
		}

		if isFileName {
			res = max(res, length)
		} else {
			nums[depth] = length
		}
	}

	return res
}
