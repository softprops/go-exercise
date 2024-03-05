// LRU Cache
// https://leetcode.com/problems/lru-cache/description/
// #hashmap
package lru

type Entry struct {
	key, val   int
	prev, next *Entry
}

type LRUCache struct {
	cap        int
	data       map[int]*Entry
	head, tail *Entry
}

func New(capacity int) LRUCache {
	head := &Entry{}
	tail := &Entry{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		cap:  capacity,
		data: make(map[int]*Entry),
		head: head,
		tail: tail,
	}
}

func (c *LRUCache) Get(key int) int {
	if v, ok := c.data[key]; ok {
		c.delete(v)
		c.add(v)
		return v.val
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if v, ok := c.data[key]; ok {
		c.delete(v)
	}
	if len(c.data) == c.cap {
		c.delete(c.tail.prev)
	}

	c.add(&Entry{
		key: key,
		val: value,
	})
}

func (c *LRUCache) add(e *Entry) {
	// set e's pointers
	e.next = c.head.next
	e.prev = c.head
	// update head's pointers
	c.head.next.prev = e
	c.head.next = e

	// update map
	c.data[e.key] = e
}

func (c *LRUCache) delete(e *Entry) {
	prev, next := e.prev, e.next
	e.prev.next = next
	e.next.prev = prev
	// clear mem
	e.prev = nil
	e.next = nil

	// clear map
	delete(c.data, e.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
