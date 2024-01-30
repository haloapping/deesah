package stack

import (
	"errors"
)

type Stack struct {
	Items  []any
	Top    int
	MaxCap int
}

func New(maxCap int) Stack {
	return Stack{
		Items:  make([]any, 0, maxCap),
		Top:    -1,
		MaxCap: maxCap,
	}
}

func (s *Stack) IsEmpty() bool {
	return s.Top == -1
}

func (s *Stack) IsFull() bool {
	return s.Top == s.MaxCap-1
}

func (s *Stack) Push(v any) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.Items = append(s.Items, v)
	s.Top++
	return nil
}

func (s *Stack) Pop() error {
	if s.IsEmpty() {
		return errors.New("stack is empty")
	}

	s.Items = s.Items[:len(s.Items)-1]
	s.Top--

	return nil
}

func (s *Stack) Peek() (v any, ok bool) {
	if s.IsEmpty() {
		return nil, false
	}

	return s.Items[s.Top], true
}
