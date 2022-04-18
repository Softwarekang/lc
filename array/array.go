package array

// twoSum 题目无特殊要求有对应数值，直接返回对应 index 即可
/*
tag:数组、hash
*/
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for k, v := range nums {
		num := target - v
		if index, ok := indexMap[num]; ok {
			return []int{k, index}
		}
		indexMap[v] = k
	}

	return []int{}
}
