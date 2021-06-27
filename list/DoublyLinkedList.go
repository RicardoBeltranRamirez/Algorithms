package list

import "fmt"

/*
@author (original JAVA) William Fiset, william.alexandre.fiset@gmail.com
        (Golang implementation) Ricardo Ramirez ricardor936@gmail.com
*/

//DoublyLinkedList is a struct representing the linked list
type DoublyLinkedList struct {
	size int
	head *node
	tail *node
}

type node struct {
	data       interface{}
	prev, next *node
}

//NewDoublyLinkedList creates a new instance of the list
func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

//Clear empty the linked list, O(n)
func (l *DoublyLinkedList) Clear() {
	trav := l.head

	for trav != nil {
		next := trav.next
		trav.prev = nil
		trav.next = nil
		trav = next
	}
	l.head, l.tail, trav = nil, nil, nil
	l.size = 0
}

//Size returns the size of the linked list
func (l *DoublyLinkedList) Size() int {
	return l.size
}

//IsEmpty verifies if this linked list is empty
func (l *DoublyLinkedList) IsEmpty() bool {
	return l.size == 0
}

//AddLast add a node to the tail of the linked list, O(1)
func (l *DoublyLinkedList) AddLast(elem interface{}) {

	if l.IsEmpty() {
		l.head = &node{elem, nil, nil}
		l.tail = l.head
	} else {
		l.tail.next = &node{elem, l.tail, nil}
		l.tail = l.tail.next
	}
	l.size++
}

//Add a node o the tail of the linked list, O(1)
func (l *DoublyLinkedList) Add(elem interface{}) {
	l.AddLast(elem)
}

//AddFirst add a node to the head of the linked list, O(1)
func (l *DoublyLinkedList) AddFirst(elem interface{}) {

	if l.IsEmpty() {
		l.head = &node{elem, nil, nil}
	} else {
		l.head.prev = &node{elem, nil, l.head}
		l.head = l.head.prev
	}
	l.size++
}

//PeekFirst check the value of the first node if it exists, O(1)
func (l *DoublyLinkedList) PeekFirst() (elem interface{}, err error) {

	if l.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}
	return l.head.data, nil
}

//PeekLast check the last node if it exists, O(1)
func (l *DoublyLinkedList) PeekLast() (elem interface{}, err error) {

	if l.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}
	return l.tail.data, nil
}

//RemoveFirst removes the value at the head of the linked list, O(1)
func (l *DoublyLinkedList) RemoveFirst() (elem interface{}, err error) {

	//Can't remove data from an empty list
	if l.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	//Extract the data at the head and move the head pointer forwards one node
	data := l.head.data
	l.head = l.head.next
	l.size--

	//If the list is empty set tail as null as well
	if l.IsEmpty() {
		l.tail = nil
	} else {
		//Do a memory clean of the previous node
		l.head.prev = nil
	}

	return data, nil
}

//RemoveLast removes the value at the tail of the linked list, O(1)
func (l *DoublyLinkedList) RemoveLast() (elem interface{}, err error) {

	//Can't remove data from an empty list
	if l.IsEmpty() {
		return nil, fmt.Errorf("linked list is empty")
	}

	//Extract the data at the tail and move the tail pointer backwards one node
	data := l.tail.data
	l.tail = l.tail.prev
	l.size--

	//If the list is empty set head as null ass well
	if l.IsEmpty() {
		l.head = nil
	} else {
		//Do a memory clean of the node that was just removed
		l.tail.next = nil
	}

	return data, nil
}

//remove a specific node from the linked list
func (l *DoublyLinkedList) remove(node *node) (elem interface{}, err error) {

	// If the node to remove is somewhere either a the head or the tail, handle
	// those independently
	if node.prev == nil {
		return l.RemoveFirst()
	}

	if node.next == nil {
		return l.RemoveLast()
	}

	//Make the pointers of adjacent nodes skip over "node"
	node.next.prev = node.prev
	node.prev.next = node.next

	//Temporary store the data that we want to return
	data := node.data

	//Memory cleanup
	node.data = nil
	node.prev, node.next, node = nil, nil, nil

	l.size--

	return data, nil
}

//RemoveAt remove a node at a particular index, O(n)
func (l *DoublyLinkedList) RemoveAt(index int) (elem interface{}, err error) {

	//Make sure the index provided is valid
	if index < 0 || index >= l.size {
		return nil, fmt.Errorf("illegal index")
	}

	var i int
	var trav *node

	//Search from the front of the list
	if index < l.size/2 {
		for i, trav = 0, l.head; i != index; i++ {
			trav = trav.next
		}
	} else {
		//Search from the back of the list
		for i, trav = l.size-1, l.tail; i != index; i++ {
			trav = trav.prev
		}
	}

	return l.remove(trav)
}

//Remove a particular value in the linked list, O(n)
func (l *DoublyLinkedList) Remove(elem interface{}) (bool, error) {

	trav := l.head

	//Support searching for nil
	if elem == nil {
		for trav = l.head; trav != nil; trav = trav.next {
			if trav.data == nil {
				_, err := l.remove(trav)
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
	}

	//Search for non nil elem
	if elem != nil {
		for trav = l.head; trav != nil; trav = trav.next {
			if elem == trav.data {
				_, err := l.remove(trav)
				if err != nil {
					return false, err
				}
				return true, nil
			}
		}
	}
	return false, nil
}

//IndexOf find the index of a particular value in the linked list, O(n)
func (l *DoublyLinkedList) IndexOf(elem interface{}) int {

	var index int
	var trav *node

	//Support searching for nil
	if elem == nil {
		for trav = l.head; trav != nil; trav, index = trav.next, index+1 {
			if trav.data == nil {
				return index
			}
		}
	}

	//Search for non nil elem
	if elem != nil {
		for trav = l.head; trav != nil; trav, index = trav.next, index+1 {
			if elem == trav.data {
				return index
			}
		}
	}
	return -1
}

//Contains check if value is contained within the linked list
func (l *DoublyLinkedList) Contains(elem interface{}) bool {
	return l.IndexOf(elem) != -1
}
