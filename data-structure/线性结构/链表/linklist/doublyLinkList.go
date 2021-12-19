package main

import "fmt"

type MyLinkedList struct {
	Head   *Node
	Length int
}

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		Head:   &Node{},
		Length: 0,
	}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Length-1 {
		return -1
	}

	if index == 0 {
		return this.Head.Data
	}

	if index < this.Length-1 {
		j := 0
		current := this.Head
		for index != j && current.Next != nil {
			current = current.Next
			j++
		}
		return current.Data
	}

	if index == this.Length-1 {
		current := this.Head
		for current.Next != nil {
			current = current.Next
		}
		return current.Data
	}
	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	node := &Node{
		Data: val,
		Next: nil,
		Prev: nil,
	}
	if this.Length == 0 {
		this.Head = node
	} else {
		node.Next = this.Head
		this.Head = node
	}
	this.Length++
}

func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{
		Data: val,
		Next: nil,
		Prev: nil,
	}

	if this.Length == 0 {
		this.AddAtHead(val)
		return
	}

	current := this.Head
	for current.Next != nil {
		current = current.Next
	}
	node.Prev = current
	current.Next = node
	this.Length++
}

func (this *MyLinkedList) AddAtIndex(index, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Length {
		this.AddAtTail(val)
		return
	}

	if index > this.Length {
		return
	}

	node := &Node{
		Data: val,
	}
	j := 1
	pre := this.Head
	for j != index && pre.Next != nil {
		pre = pre.Next
		j++
	}

	after := pre.Next
	after.Prev = node
	node.Next = after
	pre.Next = node
	this.Length++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Length {
		return
	}

	if index == 0 {
		this.Head = this.Head.Next
		this.Length--
		return
	}

	current := this.Head
	for index > 1 && current.Next != nil {
		current = current.Next
		index--
	}

	after := current.Next
	current.Next = after.Next
	// after.Next.Prev = current
	this.Length--
}

// func (this *MyLinkedList) Scan() {
// 	current := this.Head
// 	i := 0
// 	for current.Next != nil {
// 		fmt.Printf("第%d的节点值是%d\n", i, current.Data)
// 		current = current.Next
// 		i++
// 	}
// 	fmt.Printf("第%d的节点值是%d\n", i, current.Data)
// }

func main() {
	llist := Constructor()
	llist.AddAtTail(1)
	// llist.AddAtHead(1)
	// llist.DeleteAtIndex(0)
	fmt.Printf("get(0) %v\n", llist.Get(0))
	// llist.Scan()
	// llist.AddAtHead(1)
	// llist.AddAtTail(3)
	// llist.AddAtIndex(1, 2)
	// llist.Scan()
	// fmt.Printf("get(1) %v\n", llist.Get(1))
	// llist.DeleteAtIndex(1)
	// fmt.Printf("....\n")
	// llist.Scan()
	// fmt.Printf("get(1) %v\n", llist.Get(1))
}
