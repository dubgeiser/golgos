package queue

type Node struct {
	Value any
	Next  *Node
}

func NewNode(v any) *Node {
	n := Node{v, nil}
	return &n
}

type Queue struct {
	head   *Node
	tail   *Node
	length int
}

func New() *Queue {
	q := Queue{nil, nil, 0}
	return &q
}

func (q *Queue) Enqueue(v any) {
	n := NewNode(v)
	if q.head == nil {
		q.head = n
	}
	if q.tail != nil {
		q.tail.Next = n
	}
	q.tail = n
	q.length++
}

func (q *Queue) Deque() any {
	dq := q.head
	q.head = q.head.Next
	dq.Next = nil
	q.length--
	return dq.Value
}

func (q *Queue) Peek() any {
	return q.head.Value
}
