//链表操作
//83. 删除排序链表中的重复元素

package link

import (
	"openbilibili-go-common/library/net/netutil/breaker"
	"openbilibili-go-common/app/service/bbq/recsys/service/retrieve"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//暴力
func deleteDuplicates(head *ListNode) *ListNode {

	if head == nil ||  head.Next == nil{
		return head
	}

	p :=  head
	for {
		if(p.Val == p.Next.Val){
			p.Next = p.Next.Next
		}else {
			p = p.Next
		}
		if p == nil ||  p.Next == nil{
			break
		}
	}
	return head
}


//递归
func deleteDuplicates2(head *ListNode) *ListNode {

	if head == nil ||  head.Next == nil{
		return head
	}
	head.Next = deleteDuplicates2(head.Next)

	if head.Val == head.Next.Val {
		head = head.Next
	}

	return head
}
