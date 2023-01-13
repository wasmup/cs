package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := &ListNode{1, &ListNode{20, &ListNode{30, &ListNode{4, &ListNode{70, &ListNode{90, nil}}}}}}
	show(a)
	show(sort(a))
}

func sort(a *ListNode) *ListNode {
	if a == nil || a.Next == nil {
		return a
	}
	n := 0
	for p := a; p != nil; p = p.Next {
		n++
	}

	b := a
	for n = n/2 - 1; n > 0; n-- {
		b = b.Next
	}
	p := b
	b = b.Next
	p.Next = nil

	return merge(sort(a), sort(b))
}

func merge(a, b *ListNode) *ListNode {
	var head ListNode
	for p, v := &head, 0; a != nil || b != nil; p = p.Next {
		if b != nil && (a == nil || b.Val <= a.Val) {
			v = b.Val
			b = b.Next
		} else {
			v = a.Val
			a = a.Next
		}
		p.Next = &ListNode{v, nil}
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	if p == nil {
		return ""
	}
	return strconv.Itoa(p.Val)
}

func show(p *ListNode) {
	for s := ""; p != nil; p = p.Next {
		fmt.Print(s, p.Val)
		s = ","
	}
	fmt.Println()
}
