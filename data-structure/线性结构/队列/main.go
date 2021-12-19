package main

// leetcode: 239.滑动窗口最大值
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	length := len(nums)
	res := make([]int, 0)
	for i := 0; i < k; i++ { // 依次将第一个窗口内的元素弹入队列
		queue.push(nums[i])
	}
	// 将第一个窗口内的最大值加入到返回结果
	res = append(res, queue.front())

	for i := k; i < length; i++ { // 窗口右移
		queue.pop(nums[i-k]) // 依次弹出窗口最左侧队列元素，移除左侧元素
		queue.push(nums[i])  // 依次将新元素加入窗口内，新增窗口右侧元素
		res = append(res, queue.front())
	}
	return res
}

// MyQueue 单调递减队列
type MyQueue struct {
	storage []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		storage: make([]int, 0),
	}
}

func (q *MyQueue) pop(val int) {
	// 每次弹出是判断当前队列是否为空且比较当前要弹出的数值是否等于队列出口元素的数值
	if !q.isEmpty() && val == q.front() {
		// 如果相等则弹出，否则不做处理
		q.storage = q.storage[1:]
	}
}

// back 返回尾部元素
func (q *MyQueue) back() int {
	return q.storage[len(q.storage)-1]
}

func (q *MyQueue) push(val int) {
	// 每次弹出时判断当前队列是否为空
	// 比较当前要入队的数值是否大于队列入口元素的数值
	// 直到push的数值小于等于队列入口元素的数值为止
	for !q.isEmpty() && val > q.back() {
		// 如果大于则将队列尾部元素弹出。
		q.storage = q.storage[:len(q.storage)-1]
	}
	q.storage = append(q.storage, val)
}

// 返回队列首部元素即该队列中最大的元素
func (q *MyQueue) front() int {
	return q.storage[0]
}

// 判断当前队列是否为空
func (q *MyQueue) isEmpty() bool {
	return len(q.storage) == 0
}
