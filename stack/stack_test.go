package stack_test

import (
	"testing"

	"github.com/haloapping/deesah/stack"
	"github.com/stretchr/testify/assert"
)

const maxCap = 10

func TestStack(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		s := stack.New(maxCap)
		expected := stack.Stack{
			Items:  make([]any, 0, maxCap),
			Top:    -1,
			MaxCap: maxCap,
		}
		assert.Equal(t, expected, s)
		assert.Equal(t, []any{}, s.Items)
	})

	t.Run("Push", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			s := stack.New(maxCap)
			err := s.Push(10)
			assert.Nil(t, err)
			assert.Equal(t, []any{10}, s.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < s.MaxCap/2; v++ {
				_ = s.Push(v)
			}
			err := s.Push(10)
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 10}, s.Items)
		})

		t.Run("when full", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = s.Push(v)
			}
			err := s.Push(10)
			assert.EqualError(t, err, "stack is full")
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s.Items)
		})
	})

	t.Run("Pop", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			s := stack.New(maxCap)
			err := s.Pop()
			assert.EqualError(t, err, "stack is empty")
			assert.Equal(t, []any{}, s.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < s.MaxCap/2; v++ {
				_ = s.Push(v)
			}
			err := s.Pop()
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3}, s.Items)
		})

		t.Run("when full", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = s.Push(v)
			}
			err := s.Pop()
			assert.Nil(t, err)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8}, s.Items)
		})
	})

	t.Run("Peek", func(t *testing.T) {
		t.Run("when empty", func(t *testing.T) {
			s := stack.New(maxCap)
			v, ok := s.Peek()
			assert.Nil(t, v)
			assert.False(t, ok)
			assert.Equal(t, []any{}, s.Items)
		})

		t.Run("when not empty", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < s.MaxCap/2; v++ {
				_ = s.Push(v)
			}
			v, ok := s.Peek()
			assert.Equal(t, 4, v)
			assert.True(t, ok)
			assert.Equal(t, []any{0, 1, 2, 3, 4}, s.Items)
		})

		t.Run("when full", func(t *testing.T) {
			s := stack.New(maxCap)
			for v := 0; v < maxCap; v++ {
				_ = s.Push(v)
			}
			v, ok := s.Peek()
			assert.Equal(t, 9, v)
			assert.True(t, ok)
			assert.Equal(t, []any{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, s.Items)
		})
	})
}
