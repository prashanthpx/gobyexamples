package main

import "fmt"

// NodeS represents one node in a singly linked list.
// Each node stores an integer value and a pointer to the next node.
type NodeS struct {
	val  int
	next *NodeS
}

// LinkedListS holds the head pointer of the singly linked list.
type LinkedListS struct {
	head *NodeS
}

// Append adds a new node to the end of the singly linked list.
func (l *LinkedListS) Append(v int) {
	newNode := &NodeS{val: v} // next is nil by default
	if l.head == nil {        // empty list
		l.head = newNode
		return
	}
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	cur.next = newNode
}

// Prepend inserts a new node at the beginning (as the new head).
func (l *LinkedListS) Prepend(v int) {
	newNode := &NodeS{val: v, next: l.head}
	l.head = newNode
}

// InsertAfter finds the first node with value 'target' and inserts a new node after it.
// Returns true if inserted, false if target not found.
func (l *LinkedListS) InsertAfter(target, v int) bool {
	cur := l.head
	for cur != nil && cur.val != target {
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := &NodeS{val: v, next: cur.next}
	cur.next = newNode
	return true
}

// Delete removes the FIRST occurrence of 'v' from the list. Returns true if deleted.
func (l *LinkedListS) Delete(v int) bool {
	if l.head == nil {
		return false
	}
	if l.head.val == v {
		l.head = l.head.next
		return true
	}
	prev := l.head
	for prev.next != nil && prev.next.val != v {
		prev = prev.next
	}
	if prev.next == nil {
		return false
	}
	prev.next = prev.next.next
	return true
}

// Search returns a pointer to the FIRST node that contains v (or nil if not found).
func (l *LinkedListS) Search(v int) *NodeS {
	for cur := l.head; cur != nil; cur = cur.next {
		if cur.val == v {
			return cur
		}
	}
	return nil
}

// Traverse returns a slice of all values in order (left→right).
func (l *LinkedListS) Traverse() []int {
	var out []int
	for cur := l.head; cur != nil; cur = cur.next {
		out = append(out, cur.val)
	}
	return out
}

// Reverse reverses the list in-place: head->... becomes ...->head
func (l *LinkedListS) Reverse() {
	var prev *NodeS
	cur := l.head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	l.head = prev
}

// InsertAt inserts a node with value v at index idx (0-based). If idx <= 0, prepend. If idx >= len, append.
func (l *LinkedListS) InsertAt(idx int, v int) {
	if idx <= 0 || l.head == nil {
		l.Prepend(v)
		return
	}
	prev := l.head
	for i := 1; i < idx && prev.next != nil; i++ {
		prev = prev.next
	}
	newNode := &NodeS{val: v, next: prev.next}
	prev.next = newNode
}

// DeleteAt removes node at index idx (0-based). Returns true if deleted.
func (l *LinkedListS) DeleteAt(idx int) bool {
	if l.head == nil || idx < 0 {
		return false
	}
	if idx == 0 {
		l.head = l.head.next
		return true
	}
	prev := l.head
	for i := 1; i < idx && prev.next != nil; i++ {
		prev = prev.next
	}
	if prev.next == nil {
		return false
	}
	prev.next = prev.next.next
	return true
}

func main() {
	fmt.Println("=== Singly Linked List (1..10) ===")
	var s LinkedListS
	for i := 1; i <= 10; i++ {
		s.Append(i)
	}
	fmt.Println("Initial:", s.Traverse())

	ok := s.InsertAfter(5, 99)
	fmt.Println("InsertAfter(5,99) ok?", ok, "=>", s.Traverse())

	s.InsertAt(2, 77)
	fmt.Println("InsertAt(2,77):", s.Traverse())

	deleted := s.Delete(3)
	fmt.Println("Delete(3) ok?", deleted, "=>", s.Traverse())

	deleted = s.DeleteAt(4)
	fmt.Println("DeleteAt(4) ok?", deleted, "=>", s.Traverse())

	if n := s.Search(7); n != nil {
		fmt.Println("Search(7): found node with val =", n.val)
	} else {
		fmt.Println("Search(7): not found")
	}

	s.Prepend(0)
	fmt.Println("Prepend(0):", s.Traverse())

	s.Reverse()
	fmt.Println("Reverse():", s.Traverse())
}
