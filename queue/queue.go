package queue

import (
	"errors"
)

type Queue struct {
	Items  []any
	MaxCap int
}

func New(maxCap int) Queue {
	return Queue{
		Items:  make([]any, 0, maxCap),
		MaxCap: maxCap,
	}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue) IsFull() bool {
	return len(q.Items) == q.MaxCap
}

func (q *Queue) Enqueue(v any) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}
	q.Items = append(q.Items, v)

	return nil
}

func (q *Queue) Dequeue() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}
	top := q.Items[0]
	q.Items = q.Items[1:]

	return top, nil
}

func (q *Queue) Front() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.Items[0], nil
}

func (q *Queue) Back() (any, error) {
	if q.IsEmpty() {
		return nil, errors.New("queue is empty")
	}

	return q.Items[len(q.Items)-1], nil
}
