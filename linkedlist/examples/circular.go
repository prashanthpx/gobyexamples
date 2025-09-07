package main

import "fmt"

// NodeC for circular singly linked list (next wraps around)
type NodeC struct {
	val  int
	next *NodeC
}

// CircularSingly maintains only a tail pointer; head is tail.next
type CircularSingly struct { tail *NodeC }

// Append inserts a node after tail and moves tail to the new node.
func (c *CircularSingly) Append(v int) {
	newNode := &NodeC{val: v}
	if c.tail == nil {
		c.tail = newNode
		newNode.next = newNode // single node points to itself
		return
	}
	newNode.next = c.tail.next // head
	c.tail.next = newNode
	c.tail = newNode
}

// Prepend inserts a node after tail, keeping tail same (so new node becomes head).
func (c *CircularSingly) Prepend(v int) {
	newNode := &NodeC{val: v}
	if c.tail == nil {
		c.tail = newNode
		newNode.next = newNode
		return
	}
	newNode.next = c.tail.next
	c.tail.next = newNode
}

// Delete removes first node with value v. Returns true if deleted.
func (c *CircularSingly) Delete(v int) bool {
	if c.tail == nil { return false }
	prev := c.tail
	cur := c.tail.next // head
	for {
		if cur.val == v {
			if cur == prev { // single element
				c.tail = nil
				return true
			}
			prev.next = cur.next
			if cur == c.tail { c.tail = prev }
			return true
		}
		prev, cur = cur, cur.next
		if cur == c.tail.next { break } // back to head
	}
	return false
}

// Search returns pointer to first node with v (or nil).
func (c *CircularSingly) Search(v int) *NodeC {
	if c.tail == nil { return nil }
	for cur := c.tail.next; ; cur = cur.next {
		if cur.val == v { return cur }
		if cur == c.tail { break }
	}
	return nil
}

// Traverse returns values once around the ring.
func (c *CircularSingly) Traverse() []int {
	var out []int
	if c.tail == nil { return out }
	for cur := c.tail.next; ; cur = cur.next {
		out = append(out, cur.val)
		if cur == c.tail { break }
	}
	return out
}

func main() {
	fmt.Println("=== Circular Singly Linked List (ring) ===")
	var r CircularSingly
	for i := 1; i <= 6; i++ { r.Append(i) }
	fmt.Println("Once around:", r.Traverse())

	r.Prepend(0)
	fmt.Println("Prepend(0):", r.Traverse())

	deleted := r.Delete(3)
	fmt.Println("Delete(3)", deleted, "=>", r.Traverse())

	if n := r.Search(5); n != nil { fmt.Println("Search(5): found", n.val) }
}

