package singlylinkedlist_test

import (
	"testing"

	"github.com/haloapping/deesah/singlylinkedlist"
	"github.com/stretchr/testify/assert"
)

func TestSinglyLinkedList(t *testing.T) {
	t.Run("New", func(t *testing.T) {
		sll := singlylinkedlist.New()
		expected := singlylinkedlist.SinglyLinkedList{
			Head: nil,
		}
		assert.Equal(t, expected, sll)
		assert.Equal(t, 0, sll.NumOfNode)
	})

	t.Run("Insert From Front", func(t *testing.T) {
		t.Run("when head empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromFront(1)
			assert.Equal(t, &singlylinkedlist.Node{1, nil}, sll.Head)
			assert.Equal(t, 1, sll.NumOfNode)
		})

		t.Run("when head is not empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromFront(1)
			sll.InsertFromFront(2)
			sll.InsertFromFront(3)
			assert.Equal(t, &singlylinkedlist.Node{3, &singlylinkedlist.Node{2, &singlylinkedlist.Node{1, nil}}}, sll.Head)
			assert.Equal(t, 3, sll.NumOfNode)
		})
	})

	t.Run("Insert From Back", func(t *testing.T) {
		t.Run("when head empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromBack(1)
			assert.Equal(t, &singlylinkedlist.Node{1, nil}, sll.Head)
			assert.Equal(t, 1, sll.NumOfNode)
		})

		t.Run("when head is not empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromBack(1)
			sll.InsertFromBack(2)
			sll.InsertFromBack(3)
			assert.Equal(t, &singlylinkedlist.Node{1, &singlylinkedlist.Node{2, &singlylinkedlist.Node{3, nil}}}, sll.Head)
			assert.Equal(t, 3, sll.NumOfNode)
		})
	})

	t.Run("Delete Front", func(t *testing.T) {
		t.Run("when head is empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			err := sll.DeleteFront()
			assert.EqualError(t, err, "linked list is empty")
			assert.Equal(t, 0, sll.NumOfNode)
		})

		t.Run("when head is not empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromFront(1)
			sll.InsertFromFront(2)
			sll.InsertFromFront(3)
			sll.InsertFromFront(4)
			err := sll.DeleteFront()
			assert.Nil(t, err)
			assert.Equal(t, 3, sll.NumOfNode)
			assert.Equal(t, &singlylinkedlist.Node{3, &singlylinkedlist.Node{2, &singlylinkedlist.Node{1, nil}}}, sll.Head)
		})
	})

	t.Run("Delete Back", func(t *testing.T) {
		t.Run("when head is empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			err := sll.DeleteBack()
			assert.EqualError(t, err, "linked list is empty")
			assert.Equal(t, 0, sll.NumOfNode)
		})

		t.Run("when head is not empty", func(t *testing.T) {
			sll := singlylinkedlist.New()
			sll.InsertFromBack(1)
			sll.InsertFromBack(2)
			sll.InsertFromBack(3)
			sll.InsertFromBack(4)
			err := sll.DeleteBack()
			assert.Nil(t, err)
			assert.Equal(t, 3, sll.NumOfNode)
			assert.Equal(t, &singlylinkedlist.Node{1, &singlylinkedlist.Node{2, &singlylinkedlist.Node{3, nil}}}, sll.Head)
		})
	})
}
