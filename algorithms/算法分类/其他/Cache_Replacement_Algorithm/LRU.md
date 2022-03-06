# LRU
## LRU介绍
LRU：Least recently used，最近最少使用策略。基于数据访问历史记录来执行淘汰策略，LRU是首先淘汰最长时间未被使用的页面，LRU只关注该条数据的访问时间。核心主要是：如果一个数据在最近一段时间内没被访问，则在将来一段时间内被访问的可能性也很小。

## 代码实现

```go
// LRUCache定义
type LRUCache struct {
	size       int  // 当前缓存数据大小
	capacity   int  // 当前缓存数据容量
	cache      map[int]*DoublyLinkedNode    // 缓存数据K-v存储
	head, tail *DoublyLinkedNode            // 维护双向链表，head, tail为双向链表的虚拟头尾结点
}
// 双向链表
// 淘汰最近最少使用数据时间复杂度为O(1)
type DoublyLinkedNode struct {
	kek, value int  // 缓存数据键值
	prev, next *DoublyLinkedNode // 当前缓存数据结点的前、后结点的指针
}

func initDoublyLinkedNode(key, value int) *DoublyLinkedNode {
	return &DoublyLinkedNode{
		kek:   key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		head:     initDoublyLinkedNode(0, 0),
		tail:     initDoublyLinkedNode(0, 0),
		cache:    map[int]*DoublyLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDoublyLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			node := this.removeTail()
			delete(this.cache, node.kek)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) addToHead(node *DoublyLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DoublyLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DoublyLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DoublyLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
```