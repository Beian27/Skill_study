package main

import (
	"fmt"
	"sort"
)

func main() {
	unique := isUnique("dasdasdas")
	fmt.Println(unique)
}

func isUnique(astr string) bool {
	m := make(map[int32]int,len(astr))

	for k ,v := range astr {
		_, ok := m[v]
		if ok {
			return false
		}
		m[v] = k
	}
	return true
}

func CheckPermutation(s1 string, s2 string) bool {
	m := make(map[int32]int, max(len(s1), len(s2)))
	s := s1 + s2
	for _, i := range s {
		m[i]++
		if m[i] > 2 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return b
	}
	return a
}




type ListNode struct {
     Val int
     Next *ListNode
  }
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA := headA
	curB := headB
	lenA, lenB := 0, 0
	// 求A，B的长度
	for curA != nil {
		curA = curA.Next
		lenA++
	}
	for curB != nil {
		curB = curB.Next
		lenB++
	}
	var step int
	var fast, slow *ListNode
	// 请求长度差，并且让更长的链表先走相差的长度
	if lenA > lenB {
		step = lenA - lenB
		fast, slow = headA, headB
	} else {
		step = lenB - lenA
		fast, slow = headB, headA
	}
	for i:=0; i < step; i++ {
		fast = fast.Next
	}
	// 遍历两个链表遇到相同则跳出遍历
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}


func deleteNode(node *ListNode) {
	len := 0
	for node != nil {
		node = node.Next
		len++
	}
	if len%2 == 0 {
		len--
	}
	for node.Next != nil {
		count := 1

		if count == len/2 {
			node.Next = node.Next.Next
		}
		count++
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	if len(nums1)%2 == 0 {
		num1 := nums1[len(nums1)/2]
		num2 := nums1[len(nums1)/2 - 1]
		return float64((num1 + num2))/2
	} else {
		return float64(nums1[len(nums1)/2])
	}
}

