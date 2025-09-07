package main

import "testing"

func TestSinglyReverse(t *testing.T) {
	var l LinkedListS
	for i := 1; i <= 5; i++ { l.Append(i) }
	l.Reverse()
	got := l.Traverse()
	exp := []int{5,4,3,2,1}
	if len(got) != len(exp) { t.Fatalf("len mismatch: %v vs %v", got, exp) }
	for i := range exp {
		if got[i] != exp[i] { t.Fatalf("idx %d: got %d exp %d", i, got[i], exp[i]) }
	}
}

