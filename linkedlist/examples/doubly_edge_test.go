package main

import "testing"

func TestDoublyEdgeCases(t *testing.T) {
	var d DoublyLinkedList
	// empty operations
	if d.Delete(1) { t.Fatal("delete on empty should be false") }
	if d.Search(1) != nil { t.Fatal("search on empty should be nil") }

	// single element
	d.Append(10)
	if !d.Delete(10) { t.Fatal("delete single element failed") }

	// head/tail adjustments
	for i := 1; i <= 3; i++ { d.Append(i) }
	if !d.Delete(1) { t.Fatal("delete head failed") }
	if !d.Delete(3) { t.Fatal("delete tail failed") }
	got := d.TraverseForward()
	exp := []int{2}
	if len(got) != len(exp) || got[0] != 2 { t.Fatalf("unexpected: %v", got) }
}

