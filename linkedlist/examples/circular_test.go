package main

import "testing"

func TestCircularBasic(t *testing.T) {
	var r CircularSingly
	// empty traverse
	if got := r.Traverse(); len(got) != 0 { t.Fatalf("expected empty traverse, got %v", got) }

	r.Append(1)
	r.Append(2)
	r.Append(3)
	if got := r.Traverse(); len(got) != 3 { t.Fatalf("want 3, got %v", got) }

	r.Prepend(0)
	if got := r.Traverse(); got[0] != 0 { t.Fatalf("prepend failed: %v", got) }

	if n := r.Search(2); n == nil || n.val != 2 { t.Fatalf("search failed: %v", n) }

	if ok := r.Delete(3); !ok { t.Fatalf("delete(3) failed") }
	if n := r.Search(3); n != nil { t.Fatalf("3 still present") }
}

