package stack

import (
	"fmt"
	"github.com/RicardoBeltranRamirez/Algorithms/list"
)

/*
@author (original JAVA) William Fiset, william.alexandre.fiset@gmail.com
        (Golang implementation) Ricardo Ramirez ricardor936@gmail.com
*/

//Stack is a struct representing the stack data structure
type Stack struct {
	list *list.DoublyLinkedList // we use the DoublyLinkedList as the underline data structure
}

//NewStack returns a new stack data structure
func NewStack() *Stack {
	return &Stack{list: list.NewDoublyLinkedList()}
}

//Size returns the size of the stack
func (s *Stack) Size() int {
	return s.list.Size()
}

//IsEmpty verifies if this stack is empty
func (s *Stack) IsEmpty() bool {
	return s.list.IsEmpty()
}

//Push an element on the stack
func (s *Stack) Push(elem interface{}) {
	s.list.AddLast(elem)
}

// Pop an element off the stack
// Throws an error is the stack is empty
func (s *Stack) Pop() (interface{}, error) {

	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.list.RemoveLast()
}

// Peek the top of the stack without removing an element
// Throws an exception if the stack is empty
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return s.list.PeekLast()
}
