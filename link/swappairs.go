//链表操作
// 24 两两交换链表中的节点

package link




 //递归实现
 //思路: 两两分组 。相互交换，交换后每组的第二个的next 接受下一组的第一个。
 func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}
	temp1 := head   
	temp2 := head.Next

	temp1.Next = swapPairs(temp2.Next)  // 每组交换后的第二个 
	temp2.Next = temp1

	return temp2
}