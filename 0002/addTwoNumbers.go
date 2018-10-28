package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}

	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	print(addTwoNumbers(l1, l2))
}

func print(list *ListNode) {
	for list != nil {
		fmt.Printf("%d ", list.Val)
		list = list.Next
	}
}

// 20ms 100% chris
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var arr []int
	var v1, v2, t int
	var q *ListNode
	p := &ListNode{-1, nil}
	for l1 != nil || l2 != nil {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}
		arr = append(arr, v1+v2)
	}
	for _, v := range arr {
		v1 = (v + t) % 10
		t = (v + t) / 10
		if p.Val == -1 {
			p.Val = v1
			q = p
		} else {
			q.Next = &ListNode{v1, nil}
			q = q.Next
		}
	}
	if t > 0 {
		q.Next = &ListNode{1, nil}
	}
	return p
}

/* 20ms 100% others 代码更简洁
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    temp := head
     var carry int
     for l1 != nil || l2 != nil {
         var x, y int
         if l1 != nil {
             x = l1.Val
             l1 = l1.Next
         }

         if l2 != nil {
             y = l2.Val
             l2 = l2.Next
         }
        sum := x + y + carry
        carry = sum / 10
        temp.Next = &ListNode{Val: sum % 10}
        temp = temp.Next
    }

    if carry != 0 {
        temp.Next = &ListNode{Val: carry}
    }
    return head.Next
}
*/
