package main

import (
	"fmt"
	"sync/atomic"
)

// Run with: go run atomic/007_aba_versioned_pointer.go
// Demonstrates a simple versioned pointer to mitigate ABA in CAS patterns.

type node struct{ v int }

type versioned struct {
	ptr atomic.Pointer[node]
	ver atomic.Uint64
}

func (v *versioned) Load() (*node, uint64) { return v.ptr.Load(), v.ver.Load() }

func (v *versioned) CAS(oldPtr *node, oldVer uint64, newPtr *node) bool {
	if !v.ptr.CompareAndSwap(oldPtr, newPtr) { return false }
	v.ver.Add(1)
	return true
}

func main() {
	var v versioned
	v.ptr.Store(&node{v:1})
	p, ver := v.Load()
	ok := v.CAS(p, ver, &node{v:2})
	fmt.Println("cas:", ok, "new:", v.ptr.Load().v, "ver:", v.ver.Load())
}

