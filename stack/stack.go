package stack

import "github.com/linsibolinhong/datastruct/list"

type Stack struct {
	l *list.List
}

func New() *Stack {
	return &Stack{l:list.New()}
}

func (s *Stack) Top() interface{} {
	return s.l.Front().Value
}

func (s *Stack) Push(v interface{}) {
	s.l.PushFront(v)
}

func (s *Stack) Pop() interface{}  {
	if s.l.Len() == 0 {
		return nil
	}

	v := s.l.Front()
	s.l.Remove(v)
	return v.Value
}

func (s *Stack) Size() int {
	return s.l.Len()
}


