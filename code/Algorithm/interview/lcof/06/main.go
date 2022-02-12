// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.8 MB, 在所有 Go 提交中击败了96.36%的用户
// 通过测试用例：24 / 24

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	var length int
	tmp := head
	for ; tmp != nil; length++ {
		tmp = tmp.Next
	}
	res := make([]int, length)
	tmp = head
	for tail := length - 1; tail >= 0; tail-- {
		res[tail] = tmp.Val
		tmp = tmp.Next
	}
	return res
}
