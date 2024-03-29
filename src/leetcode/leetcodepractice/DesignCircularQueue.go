package leetcodepractice

type MyCircularQueue struct {
	Q        []int
	capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Q: make([]int, 0, k), capacity: k}
}

func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.IsFull() {
		return false
	}
	q.Q = append(q.Q, value)
	return true
}

func (q *MyCircularQueue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}
	q.Q = q.Q[1:]
	return true
}

func (q *MyCircularQueue) Front() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Q[0]
}

func (q *MyCircularQueue) Rear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.Q[len(q.Q)-1]
}

func (q *MyCircularQueue) IsEmpty() bool {
	return len(q.Q) == 0
}

func (q *MyCircularQueue) IsFull() bool {
	return len(q.Q) == q.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
