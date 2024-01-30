package singlylinkedlist

import "errors"

type Node struct {
	Val  any
	Next *Node
}

type SinglyLinkedList struct {
	Head      *Node
	NumOfNode int
}

func New() SinglyLinkedList {
	return SinglyLinkedList{
		Head:      nil,
		NumOfNode: 0,
	}
}

func (sll *SinglyLinkedList) InsertFromFront(v any) {
	newNode := &Node{v, nil}
	if sll.Head == nil {
		sll.Head = newNode
	} else {
		newNode.Next = sll.Head
		sll.Head = newNode
	}
	sll.NumOfNode++
}

func (sll *SinglyLinkedList) InsertFromBack(v any) {
	newNode := &Node{v, nil}
	if sll.Head == nil {
		sll.Head = newNode
	} else {
		currNode := sll.Head
		for currNode.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = newNode
	}
	sll.NumOfNode++
}

func (sll *SinglyLinkedList) DeleteFront() error {
	if sll.Head == nil {
		return errors.New("linked list is empty")
	} else {
		sll.Head = sll.Head.Next
		sll.NumOfNode--
	}

	return nil
}

func (sll *SinglyLinkedList) DeleteBack() error {
	if sll.Head == nil {
		return errors.New("linked list is empty")
	} else {
		currNode := sll.Head
		for currNode.Next.Next != nil {
			currNode = currNode.Next
		}
		currNode.Next = nil
		sll.NumOfNode--
	}

	return nil
}
