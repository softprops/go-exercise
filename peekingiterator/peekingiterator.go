// Peeking Iterator
//
// #medium
//
// https://leetcode.com/problems/peeking-iterator/description/
package peekingiterator

type Iterator struct {
	data []int
}

func NewIterator(data []int) *Iterator {
	return &Iterator{data}
}

func (i *Iterator) next() int {
	next := i.data[0]
	i.data = i.data[1:]
	return next
}

func (i *Iterator) hasNext() bool {
	return len(i.data) > 0
}

type PeekingIterator struct {
	it   *Iterator
	next *int
}

func NewPeekingIterator(it *Iterator) *PeekingIterator {
	var next *int = new(int)
	if it.hasNext() {
		*next = it.next()
	}
	return &PeekingIterator{it, next}
}

func (i *PeekingIterator) HasNext() bool {
	return i.next != nil
}

func (i *PeekingIterator) Next() int {
	n := *i.next
	if i.it.hasNext() {
		*i.next = i.it.next()
	}
	return n
}

func (i *PeekingIterator) Peek() int {
	return *i.next
}
