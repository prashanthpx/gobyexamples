package bench

import (
	"testing"
)

type node struct{ val int; next *node }

type singly struct{ head *node }

func (l *singly) append(v int) {
	n := &node{val: v}
	if l.head == nil { l.head = n; return }
	cur := l.head
	for cur.next != nil { cur = cur.next }
	cur.next = n
}

func (l *singly) traverse() []int {
	var out []int
	for cur := l.head; cur != nil; cur = cur.next { out = append(out, cur.val) }
	return out
}

func BenchmarkSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int
		for j := 0; j < 10000; j++ { s = append(s, j) }
		_ = s
	}
}

func BenchmarkSinglyAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		l := &singly{}
		for j := 0; j < 10000; j++ { l.append(j) }
	}
}

func BenchmarkSinglyTraverse(b *testing.B) {
	l := &singly{}
	for j := 0; j < 10000; j++ { l.append(j) }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = l.traverse()
	}
}

