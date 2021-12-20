package main

import "fmt"

// 构建小顶堆
type Heap struct {
	heap []int
	bool // true为小根堆，false为大根堆
}

// swap 结点交换
func (h Heap) swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

// less 结点比较
func (h Heap) less(i, j int) bool {
	if h.bool {
		return h.heap[i] < h.heap[j]
	} else {
		return h.heap[i] > h.heap[j]
	}
}

// up 上浮操作
func (h Heap) up(i int) {
	for {
		// f 为父结点
		f := (i - 1) / 2
		if i == f || h.less(f, i) {
			break
		}

		h.swap(f, i)
		i = f
	}
}

func (h *Heap) Push(x int) {
	h.heap = append(h.heap, x)
	h.up(len(h.heap) - 1)
}

// down 下沉操作
func (h *Heap) down(i int) {
	for {
		l := 2*i + 1
		if l >= len(h.heap) {
			// i 已经是叶子结点了
			break
		}

		j := l

		// 如果右子结点小于左子结点
		if r := l + 1; r < len(h.heap) && h.less(r, l) {
			// 将右子结点与左子结点中最小值赋给j
			j = r
		}
		// 比较i处结点与i的的最小子树
		if h.less(i, j) {
			break
		}
		// 交换父结点和子结点
		h.swap(i, j)
		// 继续向下比较
		i = j
	}
}

// Pop弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	n := len(h.heap) - 1
	h.swap(0, n)
	x := (h.heap)[n]
	h.heap = (h.heap)[0:n]
	h.down(0)
	return x
}

func (h Heap) Init() {
	n := len(h.heap)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

func main() {
	test := []int{3, 20, 10, 15, 7, 25, 30, 17, 19}
	tmp := &Heap{
		heap: test,
		bool: false,
	}

	tmp.Init()
	fmt.Println(tmp.heap)
	tmp.Push(6)
	fmt.Println(tmp.heap)

	p1 := tmp.Pop()
	p2 := tmp.Pop()
	p3 := tmp.Pop()
	fmt.Println(tmp.heap)

	fmt.Println(p1, p2, p3)
}
