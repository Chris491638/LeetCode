package main

import (
	"fmt"
)

/*
// 328ms 7.62% chris
type Node struct {
	key   int
	value int
	next  *Node
}

type LRUCache struct {
	head *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{&Node{capacity, 0, nil}}
}

func (this *LRUCache) Get(key int) int {
	current := this.head
	for current.next != nil {
		if current.next.key == key {
			p := current.next
			current.next = p.next
			p.next = this.head.next
			this.head.next = p
			return p.value
		} else {
			current = current.next
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node := &Node{key, value, nil}
	current := this.head
	p := current

	for current.next != nil {
		if current.next.key == key {
			current.next = current.next.next
			node.next = this.head.next
			this.head.next = node
			return
		} else {
			p = current
			current = current.next
		}
	}
	if this.head.key == this.head.value {
		p.next = nil
	} else {
		this.head.value += 1
	}
	node.next = this.head.next
	this.head.next = node
}

func (this *LRUCache) PrintAll() {
	current := this.head
	fmt.Printf("cap=%d;len=%d\n", this.head.key, this.head.value)
	for current.next != nil {
		fmt.Printf("key=%d;value=%d\n", current.next.key, current.next.value)
		current = current.next
	}
	fmt.Println()
}
*/

// 316ms 7.62% chris
type Node struct {
	prev  *Node
	key   int
	value int
	next  *Node
}

type LRUCache struct {
	head *Node
}

func Constructor(capacity int) LRUCache {
	node := &Node{nil, capacity, 0, nil}
	node.prev = node
	node.next = node
	return LRUCache{node}
}

func (this *LRUCache) Get(key int) int {
	current := this.head
	for current.next != this.head {
		p := current.next
		if p.key == key {
			current.next = p.next
			p.next.prev = current
			p.next = this.head.next
			p.prev = this.head
			this.head.next.prev = p
			this.head.next = p
			return p.value
		} else {
			current = current.next
		}
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	current := this.head
	for current.next != this.head {
		p := current.next
		if p.key == key {
			p.value = value
			current.next = p.next
			p.next.prev = current
			p.next = this.head.next
			p.prev = this.head
			this.head.next.prev = p
			this.head.next = p
			return
		} else {
			current = current.next
		}
	}
	if this.head.key == this.head.value {
		p := this.head.prev
		p.key = key
		p.value = value

		p.prev.next = this.head
		this.head.prev = p.prev

		p.next = this.head.next
		this.head.next.prev = p

		p.prev = this.head
		this.head.next = p
	} else {
		node := &Node{this.head, key, value, nil}

		this.head.next.prev = node
		node.next = this.head.next
		this.head.next = node
		node.prev = this.head

		this.head.value += 1
	}
}

func (this *LRUCache) PrintAll() {
	fmt.Printf("cap=%d;len=%d\n", this.head.key, this.head.value)
	for current := this.head; current.next != this.head; current = current.next {
		fmt.Printf("key=%d;value=%d\n", current.next.key, current.next.value)
	}
	fmt.Println()
}

func main() {
	obj := Constructor(2)
	obj.Put(2, 1)
	obj.PrintAll()
	obj.Put(1, 1)
	obj.PrintAll()
	obj.Put(2, 3)
	obj.PrintAll()
	obj.Put(4, 1)
	obj.PrintAll()

	obj.Get(1)
	obj.PrintAll()
	obj.Get(2)
	obj.PrintAll()
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
