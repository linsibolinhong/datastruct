package queue

import "github.com/linsibolinhong/datastruct/list"

type Queue struct {
	l *list.List
}

func New() *Queue {
	return &Queue{l:list.New()}
}

func (q *Queue) Push(v interface{}) {
	q.l.PushBack(v)
}

func (q *Queue) Pop(v interface{}) interface{}{
	front :=  q.l.Front()
	if front == nil {
		return nil
	}

	q.l.Remove(front)
	return front.Value
}

func (q *Queue) Front() interface{} {
	front :=  q.l.Front()
	if front == nil {
		return nil
	}

	return front.Value
}

func (q *Queue) Size() int {
	return q.l.Len()
}

