//链表操作
//83. 删除排序链表中的重复元素

package link

import (
)

type ListNode struct {
	Val  int
	Next *ListNode
}

//暴力
//  指定 p 指针指向头部 head
//  当 p 和 p.next 的存在为循环结束条件，当二者有一个不存在时说明链表没有去重复的必要了
//  当 p.val 和 p.next.val 相等时说明需要去重，则将 p 的下一个指针指向下一个的下一个，这样就能达到去重复的效果
//  如果不相等则 p 移动到下一个位置继续循环
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
