// LIFO and FIFO in Golang
package main

import (
	"fmt"
	"strconv"
)

// Node is a node contain value
type Node struct {
	value int
}

type QueueByArrayInterface interface {
}

// QueueByArray is a basic FIFO QueueByArray based on a circular list that resizes as needed.
type QueueByArray struct {
	nodes []*Node
	head  int
	tail  int
	count int
}

// TraverseQueue traverses all type of the queue
func TraverseQueue(i interface{}) {
	switch v := i.(type) {
	case QueueByArray:
		fmt.Printf("(%v, %T)\n", i, i)
		fmt.Printf("%v \n", v)
	}
}

// NewQueueByArray returns a new QueueByArray with the given initial size.
func NewQueueByArray(size int) *QueueByArray {
	return &QueueByArray{
		nodes: make([]*Node, size),
	}
}

// TraverseFromHead prints nodes in the QueueByArray from head
func (q *QueueByArray) TraverseFromHead() string {
	node := &Node{}
	chain := ""
	head := q.head
	counter := 0

	for counter != q.count {
		node = q.nodes[head]
		chain += strconv.Itoa(node.value) + "<-"

		head = (head + 1) % len(q.nodes) // for q.head != 0
		counter++
	}

	return chain
}

// TraverseFromTail prints nodes in the QueueByArray from tail
func (q *QueueByArray) TraverseFromTail() string {
	node := &Node{}
	chain := ""
	tail := q.tail - 1
	counter := q.count

	for counter > 0 {
		node = q.nodes[tail]
		chain += strconv.Itoa(node.value) + "->"

		if tail == 0 {
			tail = tail + 1
		} else {
			tail = tail - 1
		}

		counter--
	}

	return chain
}

// GetSize gets size of the QueueByArray
func (q *QueueByArray) GetSize() int {
	return len(q.nodes)
}

// Count counts nodes in the QueueByArray
func (q *QueueByArray) Count() int {
	return q.count
}

// IsFull checks the QueueByArray is full
func (q *QueueByArray) IsFull() bool {
	return q.count == len(q.nodes)
}

// IsEmpty checks the QueueByArray is empty
func (q *QueueByArray) IsEmpty() bool {
	return q.count == 0
}

// IncreaseQueueByArray increases the QueueByArray with size
func (q *QueueByArray) IncreaseQueueByArray(size int) *QueueByArray {
	if !q.IsFull() {
		return q
	}

	nodes := make([]*Node, len(q.nodes)+size)
	copy(nodes, q.nodes[q.head:])
	copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])

	q.head = 0
	q.tail = len(q.nodes)
	q.count = len(q.nodes)
	q.nodes = nodes

	return q
}

// EnQueueByArray adds a node to the QueueByArray.
func (q *QueueByArray) EnQueueByArray(n *Node) {
	if q.IsFull() {
		return
	}

	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes) // to reuse offset if q.head != 0
	q.count++
}

// DeQueueByArray removes and returns a node from the QueueByArray in first to last order.
func (q *QueueByArray) DeQueueByArray() *Node {
	if q.IsEmpty() {
		return nil
	}

	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes) // to reuse offset if q.head != 0
	q.count--

	return node
}

func main() {
	q := NewQueueByArray(3)

	data := []int{5, 3, 1, 6, 7, 4}

	for _, value := range data {
		q.EnQueueByArray(&Node{value})

		if q.IsFull() {
			q.IncreaseQueueByArray(5)
		}
	}

	// q.EnQueueByArray(&Node{3})
	// q.EnQueueByArray(&Node{2})
	// q.EnQueueByArray(&Node{1})
	// q.DeQueueByArray()
	// q.DeQueueByArray()
	// q.EnQueueByArray(&Node{4})
	// q.EnQueueByArray(&Node{5})

	fmt.Println(q.TraverseFromHead())

	fmt.Println(q.DeQueueByArray())
	fmt.Println(q.DeQueueByArray())
	q.EnQueueByArray(&Node{8})

	fmt.Println(q.TraverseFromTail())
}
