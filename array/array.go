package main

import "fmt"

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

/*
tag: 二分、中位数
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 { // 假设 nums1 的⻓度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1),
		(len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1: .................. nums1[nums1Mid-1] | nums1[nums1Mid] ........................
		// nums2: .................. nums2[nums2Mid-1] | nums2[nums2Mid] ........................
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线 划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			// nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了 // 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {

		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
tag: 链表
*/
func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	fmt.Println(containsNearbyDuplicate(nums, 2))
}

/*
tag: 数组、滑动窗口
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	left, right, h := 0, 0, make(map[int]bool)
	for right < len(nums) {
		// 维持一个 k 长度的窗口
		for right-left <= k && right < len(nums) {
			// 存在则返回 true
			if h[nums[right]] {
				return true
			}

			// 不存在则直接保存当前值
			h[nums[right]] = true
			right++
		}

		// 移除左元素，缩小窗口
		delete(h, nums[left])
		left++
	}

	return false
}
