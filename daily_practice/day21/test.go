package main

import (
	"container/list"
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

func main() {
	nums := []int{4,1,0,3,10}
	////fmt.Println(removeElement1(nums, 2))
	//fmt.Println(searchInsert(nums, 5))
	//fmt.Println(1 >> 1)
	//s := "()"
	//fmt.Println(s == ReverseString(s))
	fmt.Println(sorted(&nums))
	fmt.Println(431/10)
	words := reverseWords("Let's take LeetCode contest")
	fmt.Println(words)

}
func ReverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func removeElement(nums []int, val int) int {
	var count int
	for _, value := range nums {
		if value != val {
			nums[count] = value
			count++
		}
	}
	return count
}

func removeElement1(nums []int, val int) int {
	left, right := 0, len(nums)
	for left < right {
		if nums[left] == val  {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}
	return left
}


// 不断用二分法逼近查找第一个大于等于 \textit{target}target 的下标
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n - 1
	ans := n
	for left <= right {
		//确定中心
		mid := (right - left) >> 1 + left

		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

func maxSubArray(nums []int) int {
	max := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] + nums[i - 1] > nums[i] {
			nums[i] += nums[i - 1]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		ans[i] = v * v
	}
	sort.Ints(ans)
	return ans
}

func sorted(nums *[]int) []int{
	numss := *nums
	for i := 1; i < len(numss); i++ {
		if numss[i] < numss[i-1] {
			numss[i],numss[i-1] = numss[i-1],numss[i]
		}
	}

	return numss
}

func rotate(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		//牛啤！！！！
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

//如果 \textit{digits}digits 的末尾没有 99，例如 [1, 2, 3][1,2,3]，那么我们直接将末尾的数加一，得到 [1, 2, 4][1,2,4] 并返回；
//
//如果 \textit{digits}digits 的末尾有若干个 99，例如 [1, 2, 3, 9, 9][1,2,3,9,9]，那么我们只需要找出从末尾开始的第一个不为 99 的元素，即 33，将该元素加一，得到 [1, 2, 4, 9, 9][1,2,4,9,9]。随后将末尾的 99 全部置零，得到 [1, 2, 4, 0, 0][1,2,4,0,0] 并返回。
//
//如果 \textit{digits}digits 的所有元素都是 99，例如 [9, 9, 9, 9, 9][9,9,9,9,9]，那么答案为 [1, 0, 0, 0, 0, 0][1,0,0,0,0,0]。我们只需要构造一个长度比 \textit{digits}digits 多 11 的新数组，将首元素置为 11，其余元素置为 00 即可。

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}
//开辟一个新字符串。然后从头到尾遍历原字符串，直到找到空格为止，此时找到了一个单词，并能得到单词的起止位置。随后，根据单词的起止位置，可以将该单词逆序放到新字符串当中。如此循环多次，直到遍历完原字符串，就能得到翻转后的结果。
func reverseWords1(s string) string {
	length := len(s)
	ret := []byte{}
	for i := 0; i < length; {
		start := i
		for i < length && s[i] != ' ' {
			i++
		}
		for p := start; p < i; p++ {
			ret = append(ret, s[start + i - 1 - p])
		}
		for i < length && s[i] == ' ' {
			i++
			ret = append(ret, ' ')
		}
	}
	return string(ret)
}

func reverseWords(s string) string {
	split := strings.Split(s, " ")
	str := []byte{}
	for i, s2 := range split {
		if i != len(split) - 1 {
			s2 = " " + s2
		}
		str = append(str, reverseString([]byte(s2))...)
	}
	return string(str)
}

func reverseString(s []byte) []byte {
	n := len(s) - 1
	left, right := 0, n
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
	return s
}

type Stack struct {
	list *list.List
	lock *sync.RWMutex
}

func NewStack() *Stack {
	list := list.New()
	l := &sync.RWMutex{}
	return &Stack{list, l}
}

func (stack *Stack) Push(value interface{}) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

//
//class Solution {
//public ListNode middleNode(ListNode head) {
//ListNode slow = head, fast = head;
//while (fast != null && fast.next != null) {
//slow = slow.next;
//fast = fast.next.next;
//}
//return slow;
//}
//}
type ListNode struct {
	     Val int
	    Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var slow, fast ListNode
	for fast.Next != nil {
		slow = *slow.Next
		fast = *fast.Next.Next
	}
	return &slow
}


func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}



func lengthOfLongestSubstring(s string) int {
	// 哈希集合，记录每个字符是否出现过
	m := map[byte]int{}
	n := len(s)
	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			// 左指针向右移动一格，移除一个字符
			delete(m, s[i-1])
		}
		for rk + 1 < n && m[s[rk+1]] == 0 {
			// 不断地移动右指针
			m[s[rk+1]]++
			rk++
		}
		// 第 i 到 rk 个字符是一个极长的无重复字符子串
		ans = max(ans, rk - i + 1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}

func findRepeatNumber(nums []int) int {
	m := map[int]int{}

	for k,v := range nums {
		_, ok := m[v]
		if ok {
			return v
		}
		m[v] = k
	}
	return 0
}

func singleNumber(nums []int) (k1 int) {
	m := map[int]int{}
	for _,v := range nums {
		if _, ok := m[v]; ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}
	for k1,v1 := range m {
		if v1 == 1 {
			return k1
		}
	}
	return
}

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}


func generate(numRows int) [][]int {
	ans := make([][]int, numRows)
	for i := range ans {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j] + ans[i-1][j-1]
		}
	}
	return ans
}



func printNumbers(n int) []int {
	nn := make([]int, n)
	pow10 := int(math.Pow10(n))
	for i := 2; i < n * pow10; i++ {
		nn = append(nn, i)
	}
	return nn
}


