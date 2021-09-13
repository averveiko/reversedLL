package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

func (list *LinkedList) Add(value int) {
	n := Node{value: value}

	if list.len == 0 {
		list.head = &n
		list.len++
		return
	}

	ptr := list.head
	for i := 0; i < list.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			list.len++
			return
		}
		ptr = ptr.next
	}

}

func (list *LinkedList) IsEmpty() bool {
	if list.len == 0 {
		return true
	}

	return false
}

func (list *LinkedList) DebugPrint() {
	if list.IsEmpty() {
		fmt.Println("empty")
		return
	}

	ptr := list.head
	for i := 0; i < list.len; i++ {
		fmt.Println(ptr.value)
		ptr = ptr.next
	}
}

func (list *LinkedList) Reverse() {
	cur := list.head

	var top *Node = nil

	for cur != nil {
		temp := cur.next
		cur.next = top
		top = cur
		cur = temp
	}

	list.head = top
}

func main() {
	ll := LinkedList{}

	ll.Add(1)
	ll.Add(2)
	ll.Add(3)
	ll.Add(4)
	ll.DebugPrint()

	fmt.Println("Reversed:")
	ll.Reverse()
	ll.DebugPrint()
}
