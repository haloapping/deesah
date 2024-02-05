package deque

import (
	"errors"
	"slices"
)

type Deque struct {
	Items  []any
	MaxCap int
}

func New(maxCap int) *Deque {
	return &Deque{
		Items:  make([]any, 0),
		MaxCap: maxCap,
	}
}

func (dq *Deque) IsEmpty() bool {
	return len(dq.Items) == 0
}

func (dq *Deque) IsFull() bool {
	return len(dq.Items) == dq.MaxCap
}

func (dq *Deque) PushFront(v any) error {
	if dq.IsFull() {
		return errors.New("queue is full")
	}

	dq.Items = slices.Insert(dq.Items, 0, v)
	return nil
}

func (dq *Deque) PushBack(v any) error {
	if dq.IsFull() {
		return errors.New("queue is full")
	}

	dq.Items = slices.Insert(dq.Items, len(dq.Items), v)
	return nil
}

func (dq *Deque) PopFront() (any, error) {
	if dq.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	front := dq.Items[0]
	dq.Items = slices.Delete(dq.Items, 0, 1)
	return front, nil
}

func (dq *Deque) PopBack() (any, error) {
	if dq.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	back := dq.Items[len(dq.Items)-1]
	dq.Items = slices.Delete(dq.Items, len(dq.Items)-1, len(dq.Items))
	return back, nil
}

func (dq *Deque) Front() (any, error) {
	if dq.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return dq.Items[0], nil
}

func (dq *Deque) Back() (any, error) {
	if dq.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return dq.Items[len(dq.Items)-1], nil
}
