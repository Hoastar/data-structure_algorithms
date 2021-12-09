// package linklist
package main

import "fmt"

type Node struct {
	Data interface{}
	Next *Node
}

type LList struct {
	Header *Node
	Length int
}

func NewNode(v interface{}) *Node {
	return &Node{v, nil}
}

func CreatLList() *LList {
	header := &Node{}
	return &LList{header, 0}
}

func (l *LList) AddHead(data interface{}) {
	newNode := NewNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = newNode
	} else {
		// 将当前为插入数据之前的头结点赋值给新结点
		newNode.Next = l.Header
		// 在将新结点的值赋值给头结点
		l.Header = newNode
	}
}

func (l *LList) Insert(i int, data interface{}) {
	defer func() {
		l.Length++
	}()

	if i >= l.Length {
		l.AppendTail(data)
		return
	}

	if i <= 0 {
		l.AddHead(data)
		return
	}
	newNode := NewNode(data)

	j := 1
	pre := l.Header

	for j != i && pre.Next != nil {
		pre = pre.Next
		j++
	}

	// 得出 i+1位置处的 链表元素
	after := pre.Next
	// 将 i所在位置的 Next 指向新插入的元素
	pre.Next = newNode
	// 将新加入的元素的 Next 指向之前得出 i+1位置处的链表元素
	newNode.Next = after
}

func (l *LList) Delete(i int) {

	if i >= l.Length || i < 0 {
		return
	}

	// 删除头结点，索引从0开始
	if i == 0 {
		// 把header指向第二个结点
		l.Header = l.Header.Next
		l.Length--
		return
	}

	current := l.Header
	for i > 1 {
		current = current.Next
		i--
	}
	after := current.Next
	current.Next = after.Next
	after.Next = nil
	l.Length--
}

func (l *LList) AppendTail(data interface{}) {
	node := NewNode(data)
	defer func() {
		l.Length++
	}()

	if l.Length == 0 {
		l.Header = node
		return
	}

	if l.Length > 0 {
		current := l.Header
		// 循环找到最后一个结点，循环条件：current.Next != nil，如果等于nil，即为尾结点
		for current.Next != nil {
			current = current.Next
		}
		// 将新插入的结点的地址赋值给当前最后一个结点的Next
		// 此时node为当前新链表的最后一个结点。
		current.Next = node
		return
	}
}

func (l *LList) Scan() {
	current := l.Header
	i := 0
	for current.Next != nil {
		fmt.Printf("第%d的节点值是%d\n", i, current.Data)
		current = current.Next
		i++
	}
	fmt.Printf("第%d的节点值是%d\n", i, current.Data)
}

// 链表的头结点索引从0开始
// Get 查询当前结点对应的值（数据域）
func (l *LList) Get(index int) interface{} {
	if index < 0 {
		return nil
	}
	if index == 0 {
		return l.Header.Data
	}

	if index < l.Length-1 {
		j := 0
		current := l.Header
		for j != index && current.Next != nil {
			current = current.Next
			j++
		}
		return current.Data
	}

	if index >= l.Length-1 {
		current := l.Header
		index := 0
		for current.Next != nil {
			current = current.Next
			index++
		}
		return current.Data
	}
	return nil
}

func main() {
	llist := CreatLList()
	llist.AddHead(2)
	llist.Scan()
	fmt.Printf("...\n")
	llist.Delete(1)
	llist.Scan()
	fmt.Printf("...\n")
	llist.AddHead(2)
	llist.AddHead(7)
	llist.Scan()
}
