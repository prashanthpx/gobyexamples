# Linked Lists in Go: Singly, Doubly, and Circular (with runnable examples)

Run these examples
- Singly linked list: go run linkedlist/examples/singly.go
- Doubly linked list: go run linkedlist/examples/doubly.go
- Circular singly linked list: go run linkedlist/examples/circular.go

---

## Table of Contents
1. [What is a linked list? (mental model)](#toc-1-overview)
2. [Singly linked list: operations](#toc-2-singly)
3. [Doubly linked list: operations](#toc-3-doubly)
4. [Circular linked list: operations](#toc-4-circular)
5. [Common mistakes and tips](#toc-5-mistakes)

---

<a id="toc-1-overview"></a>

## 1) What is a linked list? (mental model)

A linked list is a chain of nodes. Each node holds a value and one or more pointers to other nodes.

- Singly: next pointer only
- Doubly: prev and next pointers
- Circular (singly): last node points back to the first node

Singly (head → ... → nil)
```
 head
  ↓
+-----+     +-----+     +-----+
|  1  | →→  |  2  | →→  |  3  | → nil
+-----+     +-----+     +-----+
```

Doubly (head ⇄ ... ⇄ tail)
```
+-----+ ⇄  +-----+ ⇄  +-----+
|  1  |    |  2  |    |  3  |
+-----+ ⇄  +-----+ ⇄  +-----+
```

Circular (tail.next → head)
```
   ┌──────────────────────────────┐
   ↓                              │
+-----+ → +-----+ → +-----+ →────┘
|  1  |   |  2  |   |  3  |
+-----+   +-----+   +-----+
```

---

<a id="toc-2-singly"></a>

## 2) Singly linked list: operations

- Append (tail insert): walk to the end and attach new node
- Prepend (head insert): point new head’s next to old head
- InsertAfter(target): find first node with value target, link new node after it
- Delete(v): remove first node with value v by relinking previous node’s next
- Search(v): walk from head until found or nil
- Traverse(): collect values in order

Key pointer steps (Append)
```
cur := head
for cur.next != nil { cur = cur.next }
cur.next = newNode
```

Key pointer steps (Delete first v)
```
if head.val == v { head = head.next; return }
prev := head
for prev.next != nil && prev.next.val != v { prev = prev.next }
if prev.next != nil { prev.next = prev.next.next }
```

Index-based operations (singly)
```
// InsertAt(idx, v)
if idx <= 0 or head == nil: Prepend(v)
else walk idx-1 steps (or until tail), splice new node after prev

// DeleteAt(idx)
if idx == 0: head = head.next
else walk idx-1 steps, bypass next node if exists
```

See runnable: linkedlist/examples/singly.go


See runnable: linkedlist/examples/singly.go

Reverse (singly) — pointer rewiring
```
prev := nil
cur := head
for cur != nil {
  next := cur.next
  cur.next = prev
  prev = cur
  cur = next
}
head = prev
```

Before → After
```
1 -> 2 -> 3 -> nil
^head

nil <- 1 <- 2 <- 3
             ^head

Reverse (doubly) — swap prev/next, then swap head/tail
```
for cur := head; cur != nil; cur = cur.prev { // note: after swap, next becomes prev
  cur.prev, cur.next = cur.next, cur.prev
}
head, tail = tail, head
```

Before
```
head → 1 ⇄ 2 ⇄ 3 ← tail
```
After
```
head → 3 ⇄ 2 ⇄ 1 ← tail
```

```


---

<a id="toc-3-doubly"></a>

## 3) Doubly linked list: operations

- Each node has prev and next
- Keep both head and tail for O(1) prepend/append
- InsertAfter(target): rewire four pointers (`new.prev`, `new.next`, neighbors)
- Delete(v): update prev.next and next.prev accordingly; fix head/tail if needed

Wire between nodes (InsertAfter)
```
new.prev = cur
new.next = cur.next
if cur.next != nil { cur.next.prev = new }
cur.next = new
```

See runnable: linkedlist/examples/doubly.go

---

<a id="toc-4-circular"></a>

## 4) Circular linked list (singly): operations

- Maintain only a tail pointer; head is `tail.next`
- Append: insert after tail and move tail to new node
- Prepend: insert after tail; keep tail same (so new node becomes head)
- Delete/Search/Traverse: loop until you arrive back at head (stop condition)

Append with tail
```
if tail == nil { tail = new; new.next = new; return }
new.next = tail.next
tail.next = new
tail = new
```

Traverse once around the ring
```
if tail == nil { return }
for cur := tail.next; ; cur = cur.next {
  visit(cur)
  if cur == tail { break } // stop when back to tail
}
```

See runnable: linkedlist/examples/circular.go

---

<a id="toc-5-mistakes"></a>

## 5) Common mistakes and tips

- Forgetting to update head/tail when deleting boundary nodes
- In circular lists, forgetting the stop condition and looping forever
- Dangling pointers: always rewire both sides in doubly linked lists
- Favor small helper functions for clarity (e.g., `isEmpty`, `head()`, `tail()`)
