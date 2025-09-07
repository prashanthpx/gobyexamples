package main

import "fmt"

// NodeD is a node in a doubly linked list with prev and next pointers.
type NodeD struct {
	val  int
	prev *NodeD
	next *NodeD
}

// DoublyLinkedList keeps head and tail pointers for O(1) append/prepend.
type DoublyLinkedList struct {
	head *NodeD
	tail *NodeD
}

// Append adds a node at the end in O(1) using tail pointer.
func (l *DoublyLinkedList) Append(v int) {
	newNode := &NodeD{val: v}
	if l.tail == nil { // empty list
		l.head, l.tail = newNode, newNode
		return
	}
	newNode.prev = l.tail
	l.tail.next = newNode
	l.tail = newNode
}

// Prepend adds a node at the beginning in O(1).
func (l *DoublyLinkedList) Prepend(v int) {
	newNode := &NodeD{val: v}
	if l.head == nil {
		l.head, l.tail = newNode, newNode
		return
	}
	newNode.next = l.head
	l.head.prev = newNode
	l.head = newNode
}

// InsertAfter inserts a new node with value v immediately AFTER the first node with value target.
func (l *DoublyLinkedList) InsertAfter(target, v int) bool {
	cur := l.head
	for cur != nil && cur.val != target {
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := &NodeD{val: v}
	newNode.prev = cur
	newNode.next = cur.next
	if cur.next != nil {
		cur.next.prev = newNode
	} else {
		l.tail = newNode
	}
	cur.next = newNode
	return true
}

// Delete removes the FIRST node that matches v. Returns true if deleted.
func (l *DoublyLinkedList) Delete(v int) bool {
	cur := l.head
	for cur != nil && cur.val != v {
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	if cur.prev != nil {
		cur.prev.next = cur.next
	} else {
		l.head = cur.next
	}
	if cur.next != nil {
		cur.next.prev = cur.prev
	} else {
		l.tail = cur.prev
	}
	return true
}

// Search returns pointer to first node with value v (or nil).
func (l *DoublyLinkedList) Search(v int) *NodeD {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.val == v {
			return cur
		}
	}
	return nil
}

// TraverseForward returns slice of values from head → tail.
func (l *DoublyLinkedList) TraverseForward() []int {
	var out []int
	for cur := l.head; cur != nil; cur = cur.next {
		out = append(out, cur.val)
	}
	return out
}

// TraverseBackward returns slice of values from tail → head.
func (l *DoublyLinkedList) TraverseBackward() []int {
	var out []int
	for cur := l.tail; cur != nil; cur = cur.prev {
		out = append(out, cur.val)
	}
	return out
}

// Reverse reverses the list in-place (swap prev/next and swap head/tail)
func (l *DoublyLinkedList) Reverse() {
	for cur := l.head; cur != nil; cur = cur.prev { // after swap, next becomes prev
		cur.prev, cur.next = cur.next, cur.prev
	}
	l.head, l.tail = l.tail, l.head
}
func main() {
	fmt.Println("=== Doubly Linked List (10,20,..,80) ===")
	var d DoublyLinkedList
	for v := 10; v <= 80; v += 10 {
		d.Append(v)
	}
	fmt.Println("Forward :", d.TraverseForward())
	fmt.Println("Backward:", d.TraverseBackward())

	ok := d.InsertAfter(40, 45)
	fmt.Println("InsertAfter(40,45) ok?", ok, "=>", d.TraverseForward())

	deleted := d.Delete(20)
	fmt.Println("Delete(20) ok?", deleted, "=>", d.TraverseForward())

	if n := d.Search(60); n != nil {
		fmt.Println("Search(60): found node with val =", n.val)
	} else {

		d.Reverse()
		fmt.Println("Reverse():", d.TraverseForward())

		fmt.Println("Search(60): not found")
	}

	d.Prepend(5)
	fmt.Println("Prepend(5):", d.TraverseForward())
	fmt.Println("Backward now:", d.TraverseBackward())
}
