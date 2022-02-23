// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.5 MB, 在所有 Go 提交中击败了100.00%的用户
// 通过测试用例：27 / 27

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	tmp := head
	var pre *ListNode
	for next := head.Next; next != nil; next = next.Next {
		tmp.Next = pre
		pre = tmp
		tmp = next
	}
	tmp.Next = pre
	return tmp
}
