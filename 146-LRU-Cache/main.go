package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// c := Constructor(3)
	// c.Put(1, 1)
	// c.Put(2, 2)
	// c.Put(3, 3)
	// c.Put(4, 4)
	// fmt.Println(c.Get(1))
	// fmt.Println(c.Get(2))
	// fmt.Println(c.Get(3))
	// fmt.Println(c.Get(4))
	// // c.Put(4, 4)
	// c.print()

	// c := Constructor(2)
	// c.Put(1, 0)
	// c.Put(2, 2)
	// c.Get(1)
	// c.Put(3, 3)
	// c.Get(2)
	// c.Put(4, 4)
	// c.Get(1)
	// c.Get(3)
	// c.Get(4)
	// c.print()

	//	["LRUCache","put","put","put","put","get","get"]
	// [[2],[2,1],[1,1],[2,3],[4,1],[1],[2]]

	// c := Constructor(2)
	// c.Put(2, 1)
	// c.Put(1, 1)
	// c.Put(2, 3)
	// c.Put(4, 1)
	// println(c.Get(1))
	// println(c.Get(2))
	// c.print()

	// 	["LRUCache","get","put","get","put","put","get","get"]
	// [[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	c := Constructor(2)
	c.Put(2, 6)
	c.Put(1, 5)
	c.Put(1, 2)
	println(c.Get(1))
	println(c.Get(2))
	c.print()
}

type node struct {
	val      int
	key      int
	next     *node
	previous *node
}

type LRUCache struct {
	kv       map[int]*node
	first    *node
	last     *node
	len      int
	capacity int
}

func print(x any) {
	b, err := json.MarshalIndent(x, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func (c *LRUCache) print() {
	print(c.kv)
	i := 0
	for n := c.first; n != nil; n = n.next {
		if i > c.len+1 {
			fmt.Println("puk-puk")
			break
		}
		fmt.Print(n.key, ":", n.val, "->")
		i++
	}
	fmt.Println()
	i = 0
	fmt.Print("last->first")
	for n := c.last; n != nil; n = n.previous {
		if i > c.len+1 {
			fmt.Println("puk-puk")
			break
		}
		fmt.Print(n.val, "->")
		i++
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		kv:       map[int]*node{},
		len:      0,
		capacity: capacity,
	}
}

func (c *LRUCache) delete(n *node) bool {
	//begin
	if n.previous == nil {
		if n.next == nil {
			return false
		}
		n.next.previous = nil
		c.first = n.next
		return true
	}

	//end
	if n.next == nil {
		if n.previous == nil {
			return false
		}
		c.last = n.previous
		return true
	}

	//middle
	n.previous.next = n.next
	n.next.previous = n.previous

	return true
}

func (c *LRUCache) Get(key int) int {
	if n, ok := c.kv[key]; ok {
		c.delete(n)
		// insert node in the end
		n.next = nil
		c.last.next = n
		n.previous = c.last
		c.last = n
		return n.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if n, ok := c.kv[key]; ok {
		n.val = value
		c.Get(key)
		return
	}

	// case if key is the same
	n := &node{
		val: value,
		key: key,
	}

	if c.len+1 > c.capacity {
		// delete first from map
		delete(c.kv, c.first.key)
		// insert new key
		c.kv[key] = n
		// delete first
		c.delete(c.first)
		// insert new node in the end
		c.last.next = n
		n.previous = c.last
		c.last = n
		return
	}

	// empty queue
	if c.len == 0 {
		c.first = n
		c.last = n
	} else { // non empty queue and have space for ne elem
		c.last.next = n
		n.previous = c.last
		c.last = n
	}
	c.kv[key] = n
	c.len++
}
