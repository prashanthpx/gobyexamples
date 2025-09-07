package main

import "testing"

func TestSinglyEdgeCases(t *testing.T) {
	var l LinkedListS
	// delete/search on empty
	if l.Delete(1) { t.Fatal("delete on empty should be false") }
	if l.DeleteAt(0) { t.Fatal("deleteAt on empty should be false") }
	if l.Search(1) != nil { t.Fatal("search on empty should be nil") }

	// single element operations
	l.Append(42)
	if !l.DeleteAt(0) { t.Fatal("deleteAt(0) on single failed") }
	l.Append(1)
	l.InsertAt(1, 2) // append via InsertAt
	got := l.Traverse()
	exp := []int{1,2}
	if len(got) != len(exp) { t.Fatalf("len mismatch: %v vs %v", got, exp) }
	for i := range exp { if got[i] != exp[i] { t.Fatalf("idx %d got %d exp %d", i, got[i], exp[i]) } }
}

