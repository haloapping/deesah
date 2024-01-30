package queue_test

import (
	"testing"

	"github.com/haloapping/deesah/queue"
	"github.com/stretchr/testify/assert"
)

const maxCap = 10

func TestQueue(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		q := queue.New(maxCap)
		expected := queue.Queue{
			Items:  []any{},
			MaxCap: maxCap,
		}
		assert.Equal(t, expected, q)
		assert.Equal(t, []any{}, q.Items)
	})

	t.Run("Enqueue", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			q := queue.New(maxCap)
			err := q.Enqueue(10)
			assert.Nil(t, err)
			assert.Equal(t, []any{10}, q.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap/2; v++ {
				_ = q.Enqueue(v)
			}
			err := q.Enqueue(5)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5}, q.Items)
		})

		t.Run("when full", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = q.Enqueue(v)
			}
			err := q.Enqueue(10)
			assert.EqualError(t, err, "queue is full")
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, q.Items)
		})
	})

	t.Run("Dequeue", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			q := queue.New(maxCap)
			front, err := q.Dequeue()
			assert.Nil(t, front)
			assert.EqualError(t, err, "queue is empty")
			assert.Equal(t, []any{}, q.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap/2; v++ {
				_ = q.Enqueue(v)
			}
			front, err := q.Dequeue()
			assert.Equal(t, 0, front)
			assert.Nil(t, err)
			assert.Equal(t, []any{1, 2, 3, 4}, q.Items)
		})

		t.Run("when full", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = q.Enqueue(v)
			}
			front, err := q.Dequeue()
			assert.Equal(t, 0, front)
			assert.Nil(t, err)
			assert.Equal(t, []any{1, 2, 3, 4, 5, 6, 7, 8, 9}, q.Items)
		})
	})

	t.Run("Front", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			q := queue.New(maxCap)
			front, err := q.Front()
			assert.Nil(t, front)
			assert.EqualError(t, err, "queue is empty")
			assert.Equal(t, []any{}, q.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap/2; v++ {
				_ = q.Enqueue(v)
			}
			front, err := q.Front()
			assert.Equal(t, 0, front)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4}, q.Items)
		})

		t.Run("when full", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = q.Enqueue(v)
			}
			front, err := q.Front()
			assert.Equal(t, 0, front)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, q.Items)
		})
	})

	t.Run("Back", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			q := queue.New(maxCap)
			back, err := q.Back()
			assert.Nil(t, back)
			assert.EqualError(t, err, "queue is empty")
			assert.Equal(t, []any{}, q.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap/2; v++ {
				_ = q.Enqueue(v)
			}
			back, err := q.Back()
			assert.Equal(t, 4, back)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4}, q.Items)
		})

		t.Run("when full", func(t *testing.T) {
			q := queue.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = q.Enqueue(v)
			}
			back, err := q.Back()
			assert.Equal(t, 9, back)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, q.Items)
		})
	})
}
