package deque_test

import (
	"testing"

	"github.com/haloapping/deesah/deque"
	"github.com/stretchr/testify/assert"
)

const maxCap = 5

func TestDequeEmpty(t *testing.T) {
	dq := deque.New(maxCap)
	expected := &deque.Deque{
		Items:  make([]any, 0),
		MaxCap: maxCap,
	}
	assert.Equal(t, expected, dq)
	assert.Equal(t, []any{}, dq.Items)
	assert.Equal(t, maxCap, dq.MaxCap)
	assert.True(t, dq.IsEmpty())
	assert.False(t, dq.IsFull())
}

func TestPushFront(t *testing.T) {
	t.Run("when not full", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap-1; v++ {
			_ = dq.PushFront(v)
		}
		err := dq.PushFront(4)
		assert.Nil(t, err)
		assert.Equal(t, []any{4, 3, 2, 1, 0}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})

	t.Run("when full", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushFront(v)
		}
		err := dq.PushFront(5)
		assert.EqualError(t, err, "queue is full")
		assert.Equal(t, []any{4, 3, 2, 1, 0}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})
}

func TestPushBack(t *testing.T) {
	t.Run("when not full", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap-1; v++ {
			_ = dq.PushBack(v)
		}
		err := dq.PushBack(4)
		assert.Nil(t, err)
		assert.Equal(t, []any{0, 1, 2, 3, 4}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})

	t.Run("when full", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushBack(v)
		}
		err := dq.PushBack(5)
		assert.EqualError(t, err, "queue is full")
		assert.Equal(t, []any{0, 1, 2, 3, 4}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})
}

func TestPopFront(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		front, err := dq.PopFront()
		assert.Nil(t, front)
		assert.EqualError(t, err, "queue is empty")
		assert.Equal(t, []any{}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.True(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})

	t.Run("when not empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushBack(v)
		}
		front, err := dq.PopFront()
		assert.Equal(t, 0, front)
		assert.Nil(t, err)
		assert.Equal(t, []any{1, 2, 3, 4}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})
}

func TestPopBack(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		back, err := dq.PopBack()
		assert.Nil(t, back)
		assert.EqualError(t, err, "queue is empty")
		assert.Equal(t, []any{}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.True(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})

	t.Run("when not empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushBack(v)
		}
		back, err := dq.PopBack()
		assert.Equal(t, 4, back)
		assert.Nil(t, err)
		assert.Equal(t, []any{0, 1, 2, 3}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})
}

func TestFront(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		front, err := dq.Front()
		assert.Nil(t, front)
		assert.EqualError(t, err, "queue is empty")
		assert.Equal(t, []any{}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.True(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})

	t.Run("when not empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushBack(v)
		}
		front, err := dq.Front()
		assert.Equal(t, 0, front)
		assert.Nil(t, err)
		assert.Equal(t, []any{0, 1, 2, 3, 4}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})
}

func TestBack(t *testing.T) {
	t.Run("when empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		back, err := dq.Back()
		assert.Nil(t, back)
		assert.EqualError(t, err, "queue is empty")
		assert.Equal(t, []any{}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.True(t, dq.IsEmpty())
		assert.False(t, dq.IsFull())
	})

	t.Run("when not empty", func(t *testing.T) {
		dq := deque.New(maxCap)
		for v := 0; v < maxCap; v++ {
			_ = dq.PushBack(v)
		}
		back, err := dq.Back()
		assert.Equal(t, 4, back)
		assert.Nil(t, err)
		assert.Equal(t, []any{0, 1, 2, 3, 4}, dq.Items)
		assert.Equal(t, maxCap, dq.MaxCap)
		assert.False(t, dq.IsEmpty())
		assert.True(t, dq.IsFull())
	})
}
