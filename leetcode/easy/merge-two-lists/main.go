package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
      Val int
      Next *ListNode
 }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var head *ListNode
    var current *ListNode

    //case where one list is empty and the other is not
    if list1 == nil && list2 != nil {
        return list2
    } else if list2 == nil && list1 != nil {
        return list1
    }

    for list1 != nil && list2 != nil {
        var target *ListNode

        if list1.Val <  list2.Val {
            target =  &ListNode{list1.Val, nil}
            list1 = list1.Next
        } else {
            target =  &ListNode{list2.Val, nil}
            list2 = list2.Next 
        }

        if head == nil {
            head = target
            current = head
        } else {
            previous := current
            current = target
            previous.Next = current
        }

    }

    if list1 == nil && list2 != nil {
        current.Next = list2
    } else if list2 == nil && list1 != nil {
        current.Next = list1
    }

    return head
}


func main(){
	fmt.Println("case one([1,2,4], [1,3,4] -> expect [1,1,2,3,4,4]): ", mergeTwoLists(&ListNode{1, &ListNode{2, &ListNode{4, nil}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}))
	fmt.Println("case two([], [] -> expect []): ", mergeTwoLists(nil, nil))
	fmt.Println("case three([], [0] -> expect [0]): ", mergeTwoLists(nil, &ListNode{0, nil}))
}