// LIFO and FIFO in Golang
package main

import (
	"fmt"
)

type CallbackFunctionType func(index int, node *Node)

type QueueInterface interface {
	ForEach(QueueTraverseInterface)
}

type QueueTraverseInterface interface {
}

type QueueTraverseBackward struct {
	callback CallbackFunctionType
}

type QueueTraverseForWard struct {
	callback CallbackFunctionType
}

// ForEach traverve a queue
func (q *ArrayQueue) ForEach(i QueueTraverseInterface) {
	switch v := i.(type) {
	case QueueTraverseForWard:
		fmt.Printf("(%v, %T)\n", i, i)
		fmt.Printf("%v \n", v.callback)
		q.ForEachForward(v.callback)
		break

	case QueueTraverseBackward:
		fmt.Printf("(%v, %T)\n", i, i)
		fmt.Printf("%v \n", v.callback)
		q.ForEachBackward(v.callback)
	}
}

// ForEachForward traverve a queue from head to tail
func (q *ArrayQueue) ForEachForward(callback CallbackFunctionType) {
	node := &Node{}
	head := q.head
	counter := 0

	for counter != q.count {
		node = q.nodes[head]
		callback(head, node)

		head = (head + 1) % len(q.nodes)
		counter++
	}
}

// ForEachBackward traverve a queue from tail to head
func (q *ArrayQueue) ForEachBackward(callback CallbackFunctionType) {
	node := &Node{}
	tail := q.tail
	counter := q.count

	for counter > 0 {
		if tail == 0 {
			tail = len(q.nodes)
		}

		tail = (tail - 1) % len(q.nodes)
		node = q.nodes[tail]
		callback(tail, node)

		counter--
	}
}

// Node is a node contain value
type Node struct {
	value int
}

// ArrayQueue is a basic FIFO ArrayQueue based on a circular list that resizes as needed.
type ArrayQueue struct {
	nodes []*Node
	head  int
	tail  int
	count int
}

// NewArrayQueue returns a new ArrayQueue with the given initial size.
func NewArrayQueue(size int) *ArrayQueue {
	return &ArrayQueue{
		nodes: make([]*Node, size),
	}
}

// IsFull checks the ArrayQueue is full
func (q *ArrayQueue) IsFull() bool {
	return q.count == len(q.nodes)
}

// IsEmpty checks the ArrayQueue is empty
func (q *ArrayQueue) IsEmpty() bool {
	return q.count == 0
}

// EnArrayQueue adds a node to the ArrayQueue.
func (q *ArrayQueue) EnArrayQueue(n *Node) bool {
	if q.IsFull() {
		return false
	}

	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes) // to reuse offset if q.head != 0
	q.count++
	return true
}

// DeArrayQueue removes and returns a node from the ArrayQueue in first to last order.
func (q *ArrayQueue) DeArrayQueue() *Node {
	if q.IsEmpty() {
		return nil
	}

	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes) // to reuse offset if q.head != 0
	q.count--

	return node
}

func main() {
	q := NewArrayQueue(3)

	// data := []int{5, 3, 1, 6, 7, 4}

	// for _, value := range data {
	// 	q.EnArrayQueue(&Node{value})
	// }

	q.EnArrayQueue(&Node{3})
	q.EnArrayQueue(&Node{2})
	q.EnArrayQueue(&Node{1})
	q.DeArrayQueue()
	q.EnArrayQueue(&Node{4})

	var callback CallbackFunctionType = func(index int, node *Node) {
		fmt.Println(index)
		fmt.Println(node)
	}

	// var queueTraverseForWard = QueueTraverseForWard{callback}
	var queueTraverseBackward = QueueTraverseBackward{callback}

	// var queueTraverseInterface QueueTraverseInterface = queueTraverseForWard
	var queueTraverseInterface QueueTraverseInterface = queueTraverseBackward

	q.ForEach(queueTraverseInterface)
}
